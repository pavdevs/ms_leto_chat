package server

import (
	"MsLetoChat/internal/server/client"
	"MsLetoChat/internal/support/tokenservice"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	config Config
	logger *logrus.Logger
	cl     map[int]*client.Client
}

func NewServer(config Config, logger *logrus.Logger) *Server {

	return &Server{
		config: config,
		logger: logger,
		cl:     make(map[int]*client.Client),
	}
}

func (s *Server) Start() error {

	http.HandleFunc("/", s.echo)
	return http.ListenAndServe(s.config.Host+":"+s.config.Port, nil)
}

func (s *Server) Stop() error {

	return nil
}

func (s *Server) echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	defer func() {
		s.logger.Infof("Клиент отключился. Общее количество клиентов: %+v", s.cl)
		c.Close()
	}()

	if err != nil {
		s.logger.Error(err)
		return
	}

	cl, err := s.handleConnection(r, c)

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
