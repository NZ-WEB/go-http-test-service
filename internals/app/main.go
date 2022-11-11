package app

import (
	"context"
	"github.com/NZ-WEB/go-http-test-service.git/internals/cfg"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *Server) Serve() {
	log.Println("Starting server")
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDbString())
	if err != nil {
		log.Fatalln(err)
	}

	//carsStorage := db3.NewCarStorage(server.db)
	//userStorage := db3.NewUsersStorage(server.db)

	//carsProcessor := processors.NewCarsProcessor(carsStorage)
	//usersProcessor := processors.NewUserProcessor(userStorage)

	//routes := api.CreateRoutes(userHandler, carsHandler)
	//routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr: ":" + server.config.Port,
		//Handler: routes
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (server *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shotdown Failed: #{err}")
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
