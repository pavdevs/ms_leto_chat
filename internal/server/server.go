package server

import (
	"MsLetoChat/internal/channel"
	"MsLetoChat/internal/services/tokenservice"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"io"
	"sync"
)

var (
	mutex sync.Mutex
)

type WSServer struct {
	config      Config
	logger      *logrus.Logger
	connections map[int]*channel.Channel
	secretKey   []byte
}

func NewServer(config Config, logger *logrus.Logger, secretKey []byte) *WSServer {
	return &WSServer{
		config:      config,
		logger:      logger,
		connections: make(map[int]*channel.Channel),
		secretKey:   secretKey,
	}
}

func (s *WSServer) HandleWS(ws *websocket.Conn) {

	tokenString := ws.Config().Protocol[0]

	claims, err := tokenservice.ParseAccessToken(tokenString)

	if err != nil {
		s.logger.Error(err)
		return
	}

	s.logger.Info(fmt.Sprintf("new connection from client: %s, id: %d, first_name: %s, last_name: %s, email: %s", ws.RemoteAddr(), claims.Id, claims.First, claims.Last, claims.Email))

	if _, ok := s.connections[claims.Id]; ok {
		s.logger.Error("connection already exists")
		return
	}

	user := channel.NewUser(claims.Id, claims.First, claims.Last, claims.Email)
	ch := channel.NewChannel(ws, user)

	s.connections[user.Id] = ch
	s.readLoop(ch)
}

func (s *WSServer) readLoop(c *channel.Channel) {
	defer func() {
		mutex.Lock()
		delete(s.connections, c.User.Id) // Удаляем соединение после завершения
		mutex.Unlock()
		c.Conn.Close() // Закрываем соединение
	}()
	for {
		msg := make([]byte, 1024)
		n, err := c.Conn.Read(msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.logger.Error(err)
			break
		}
		s.logger.Info(string(msg[:n]))
		s.broadcast(msg[:n], c.User.Id)
	}
}

func (s *WSServer) broadcast(b []byte, userId int) {
	for _, c := range s.connections {
		go func(ws *websocket.Conn) {
			mutex.Lock()         // Здесь заблокируйте мьютекс
			defer mutex.Unlock() // И сразу разблокируйте его перед выходом из функции
			if _, err := ws.Write(b); err != nil {
				s.logger.Error(err)
				ws.Close()
			}
		}(c.Conn)
	}
}
