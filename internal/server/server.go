package server

import (
	"MsLetoChat/internal/server/client"
	"MsLetoChat/internal/support/tokenservice"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
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

	defer c.Close()

	if err != nil {
		s.logger.Error(err)
		return
	}

	cl, err := s.handleConnection(r, c)

	s.cl[cl.GetUser().ID] = cl

	if err != nil {
		s.logger.Error(err)
		return
	}

	for {
		mt, message, err := c.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			s.logger.Error(err)
			break
		}

		if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
			s.logger.Error(err)
			break
		}

		go s.messageHandler(message)

		s.logger.Info(string(message))
	}

	delete(s.cl, cl.GetUser().ID)
}

func (s *Server) handleConnection(r *http.Request, c *websocket.Conn) (*client.Client, error) {
	tokenString := r.Header.Get("Authorization")
	splitToken := strings.Split(tokenString, "Bearer ")

	if len(splitToken) < 2 {
		return nil, fmt.Errorf("invalid token string")
	}

	claims, err := tokenservice.ParseAccessToken(splitToken[1])

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

func (s *Server) messageHandler(message []byte) {
	fmt.Println(string(message))
	for _, cl := range s.cl {
		if err := cl.Send(message); err != nil {
			s.logger.Error(err)
		}
	}
}
