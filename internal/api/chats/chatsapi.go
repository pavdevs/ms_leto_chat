package chats

import (
	"MsLetoChat/internal/api/authmiddleware"
	chatsdto "MsLetoChat/internal/api/chats/dto"
	"MsLetoChat/internal/services/chats"
	chatsservicedto "MsLetoChat/internal/services/chats/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	{
		chatsGroup.POST("", c.createChat)
		chatsGroup.PUT("/:chatID", c.editChat)
		chatsGroup.DELETE("/:chatID", c.deleteChat)
		chatsGroup.GET("/:chatID", c.getChat)
		chatsGroup.GET("", c.getChats)
	}
	chatsGroup.Use(authmiddleware.JWTAuthMiddleware())
}

// @Summary Create chat
// @Security ApiKeyAuth
// @Tags Chats
// @Description Данный роут создает чат и возвращает модель созданного чата
// @ID create_chat
// @Accept json
// @Produce json
// @Param {object} body chatsdto.CreateChatRequest true "Параметры создания чата"
// @Success 200 {object} chatsdto.CreateChatResponse "Чат успешно создан"
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

	chatReqDTO := chatsservicedto.NewChatDTO(
		req.Title,
		req.OwnerID,
	)

	chatResponseDTO, err := c.cs.CreateChat(chatReqDTO)

	if err != nil {
		c.logger.Errorf("Invalid request or JSON format: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("failed to create chat: %w", err),
		})
		return
	}

	response := chatsdto.NewCreateChatResponse(

		chatResponseDTO.ChatID,
		chatResponseDTO.Title,
		chatResponseDTO.OwnerID,
		chatResponseDTO.CreatedAt,
	)

	ctx.JSON(http.StatusOK, response)
}

func (c *ChatsAPI) deleteChat(ctx *gin.Context) {

}

func (c *ChatsAPI) editChat(ctx *gin.Context) {

}

func (c *ChatsAPI) getChat(ctx *gin.Context) {

}

func (c *ChatsAPI) getChats(ctx *gin.Context) {

}
