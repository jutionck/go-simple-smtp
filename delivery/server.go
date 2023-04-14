package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/go-simple-smtp/config"
	"github.com/jutionck/go-simple-smtp/delivery/controller"
	"github.com/jutionck/go-simple-smtp/manager"
	"github.com/jutionck/go-simple-smtp/utils"
)

type Server struct {
	serviceManager manager.ServiceManager
	engine         *gin.Engine
	host           string
}

func (s *Server) initController() {
	controller.NewCustomerController(s.engine, s.serviceManager.RegisterCustomerUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("failed to serve config :%s", err)
	}
	emailService := utils.NewEmailService(cfg)
	serviceManager := manager.NewServiceManager(emailService)
	r := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		serviceManager: serviceManager,
		engine:         r,
		host:           host,
	}
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		log.Printf("failed to run server :%s", err)
	}
}
