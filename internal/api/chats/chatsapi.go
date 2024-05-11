package chats

import (
	"MsLetoChat/internal/api/authmiddleware"
	chatsdto "MsLetoChat/internal/api/chats/dto"
	"MsLetoChat/internal/api/support"
	"MsLetoChat/internal/services/chats"
	chatsservicedto "MsLetoChat/internal/services/chats/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ChatsAPI struct {
	logger *logrus.Logger
	cs     *chats.ChatsService
}

func NewChatsAPI(logger *logrus.Logger, cs *chats.ChatsService) *ChatsAPI {

	return &ChatsAPI{
		logger: logger,
		cs:     cs,
	}
}

func (c *ChatsAPI) RegisterRoutes(eng *gin.Engine) {
	chatsGroup := eng.Group("/chats")
	chatsGroup.Use(authmiddleware.JWTAuthMiddleware())
	{
		chatsGroup.POST("", c.createChat)
		chatsGroup.PATCH("/:chatID", c.editChat)
		chatsGroup.DELETE("/:chatID", c.deleteChat)
		chatsGroup.GET("/:chatID", c.getChat)
		chatsGroup.GET("", c.getChatsList)
	}
}

// @Summary Create chat
// @Security ApiKeyAuth
// @Tags Chats
// @Description Данный роут создает чат и возвращает модель созданного чата
// @ID create_chat
// @Accept json
// @Produce json
// @Param {object} body chatsdto.CreateChatRequest true "Параметры создания чата"
// @Success 200 {object} chatsdto.ChatDTO "Чат успешно создан"
// @Router /chats [post]
func (c *ChatsAPI) createChat(ctx *gin.Context) {
	var req chatsdto.CreateChatRequest

	if err := ctx.BindJSON(&req); err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request or JSON format",
		})
		return
	}

	ownerID, err := support.GetUserIDFromContext(ctx)

	if err != nil {
		c.logger.Error(fmt.Errorf("invalid user_id in header: %w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("invalid user_id in header: %w", err),
		})
		return
	}

	chatReqDTO := chatsservicedto.ChatDTO{
		Title:   req.Title,
		OwnerID: ownerID,
	}

	chatResponseDTO, err := c.cs.CreateChat(chatReqDTO)

	if err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("failed to create chat: %w", err),
		})
		return
	}

	response := chatsdto.ChatDTO{
		ID:        chatResponseDTO.ChatID,
		Title:     chatResponseDTO.Title,
		OwnerID:   chatResponseDTO.OwnerID,
		CreatedAt: chatResponseDTO.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Delete chat by id
// @Security ApiKeyAuth
// @Tags Chats
// @Description Данный роут позволяет удалить чат по его ID
// @ID delete_chat_by_id
// @Accept json
// @Produce json
// @Param chatID path int true "ID чата"
// @Success 200 {object} chatsdto.DeleteChatResponse "Чат успешно удален"
// @Router /chats/{chatID} [delete]
func (c *ChatsAPI) deleteChat(ctx *gin.Context) {
	ctxChatID := ctx.Param("chatID")

	chatID, err := strconv.ParseInt(ctxChatID, 10, 64)

	if err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request or JSON format",
		})
		return
	}

	if err := c.cs.DeleteChat(chatID); err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request or JSON format",
		})
		return
	}

	response := chatsdto.NewDeleteChatResponse()

	ctx.JSON(http.StatusOK, response)
}

// @Summary Edit chat by id
// @Security ApiKeyAuth
// @Tags Chats
// @Description Данный роут позволяет удалить чат по его ID
// @ID edit_chat_by_id
// @Accept json
// @Produce json
// @Param chatID path int true "ID чата"
// @Param {object} body chatsdto.EditChatRequest true "Параметры редактирования чата"
// @Success 200 {object} chatsdto.ChatDTO "Чат успешно изменен"
// @Router /chats/{chatID} [patch]
func (c *ChatsAPI) editChat(ctx *gin.Context) {

	chatID, err := strconv.ParseInt(ctx.Param("chatID"), 10, 64)

	if err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid query parameter",
		})
		return
	}

	var req chatsdto.EditChatRequest

	if err := ctx.BindJSON(&req); err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid body format",
		})
		return
	}

	chat, err := c.cs.UpdateChat(req.Title, chatID)

	if err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request or JSON format",
		})
		return
	}

	response := chatsdto.ChatDTO{
		ID:        chat.ChatID,
		Title:     chat.Title,
		OwnerID:   chat.OwnerID,
		CreatedAt: chat.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *ChatsAPI) getChat(ctx *gin.Context) {

}

// @Summary Get chats list
// @Security ApiKeyAuth
// @Tags Chats
// @Description Данный роут пощволяет получить список чатов пользоватедя который из запрашивает
// @ID get_chats_list
// @Accept json
// @Produce json
// @Success 200 {object} chatsdto.GetChatsResponse "Список чатов успешно получен"
// @Router /chats [get]
func (c *ChatsAPI) getChatsList(ctx *gin.Context) {
	ownerID, err := support.GetUserIDFromContext(ctx)

	if err != nil {
		c.logger.Error(fmt.Errorf("invalid user_id in header: %w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("invalid user_id in header: %w", err),
		})
		return
	}

	chatServiceResponseDTO, err := c.cs.GetChatsList(ownerID)

	if err != nil {
		c.logger.Error(fmt.Errorf("failed to get chats list: %w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("failed to get chats list: %w", err),
		})
		return
	}

	var items []chatsdto.ChatDTO

	for item := range chatServiceResponseDTO.Chats {

		items = append(items, chatsdto.ChatDTO{
			ID:        chatServiceResponseDTO.Chats[item].ChatID,
			Title:     chatServiceResponseDTO.Chats[item].Title,
			OwnerID:   chatServiceResponseDTO.Chats[item].OwnerID,
			CreatedAt: chatServiceResponseDTO.Chats[item].CreatedAt,
		})
	}

	response := chatsdto.NewGetChatsResponse(
		items,
	)

	ctx.JSON(http.StatusOK, response)
}
