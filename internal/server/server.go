package server

import (
	docs "MsLetoChat/docs"
	"MsLetoChat/internal/api/chats"
	"MsLetoChat/internal/api/messages"
	"MsLetoChat/internal/server/client"
	"MsLetoChat/internal/support/eventsparser"
	"MsLetoChat/internal/support/tokenservice"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

var router *gin.Engine

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	config      Config
	logger      *logrus.Logger
	cl          map[int]*client.Client
	chatsApi    *chats.ChatsAPI
	messagesApi *messages.MessagesApi
	parser      *eventsparser.EventsParser
}

func init() {
	router = gin.Default()
}

func NewServer(config Config, logger *logrus.Logger, chatsApi *chats.ChatsAPI, messagesApi *messages.MessagesApi, parser *eventsparser.EventsParser) *Server {

	return &Server{
		config:      config,
		logger:      logger,
		cl:          make(map[int]*client.Client),
		chatsApi:    chatsApi,
		messagesApi: messagesApi,
		parser:      parser,
	}
}

func (s *Server) Start() error {

	router.Handle("GET", "/ws", s.echo)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", gin.WrapH(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition
	)))

	s.chatsApi.RegisterRoutes(router)

	return router.Run(s.config.Host + ":" + s.config.Port)
}

func (s *Server) Stop() error {

	return nil
}

func (s *Server) echo(ctx *gin.Context) {

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	defer func() {
		s.logger.Infof("Клиент отключился. Общее количество клиентов: %+v", s.cl)
		c.Close()
	}()

	if err != nil {
		s.logger.Error(err)
		return
	}

	cl, err := s.handleConnection(ctx.Request, c)

	if err != nil {
		s.logger.Error(err)
		return
	}

	userId := cl.GetUser().ID

	if _, exists := s.cl[userId]; exists {
		s.logger.Errorf("Клиент с таким ID = %d уже подключен", userId)
		return
	}

	s.cl[userId] = cl

	s.logger.Infof("Добавлен новый клиент.\nОбщее количество клиентов: %+v", s.cl)

	for {
		mt, message, err := c.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}

		go s.messageHandler(userId, message)
	}

	delete(s.cl, cl.GetUser().ID)
}

func (s *Server) handleConnection(r *http.Request, c *websocket.Conn) (*client.Client, error) {
	tokenString := r.URL.Query().Get("token")

	if tokenString == "" {
		return nil, fmt.Errorf("no token provided")
	}

	claims, err := tokenservice.ParseAccessToken(tokenString)

	if err != nil {
		return nil, err
	}

	user := client.NewUser(
		claims.Id,
		claims.First,
		claims.Last,
		claims.Email,
	)

	return client.NewClient(user, c), nil
}

func (s *Server) messageHandler(from int, message []byte) {
	s.logger.Infof("Сообщение от пользователя %d: %s", from, message)

	if err := s.parser.ParseEvent(message); err != nil {
		s.logger.Error(err)
	}

	for _, cl := range s.cl {

		if cl.GetUser().ID != from {
			s.send(cl, message)
		}
	}
}

func (s *Server) send(cl *client.Client, message []byte) {
	if err := cl.Send(message); err != nil {
		s.logger.Error(err)
	}
}
