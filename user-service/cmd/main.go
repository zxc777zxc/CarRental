package main

import (
	"database/sql"
	"log"
	"net"

	"CarRental/user-service/infrastructure/repository"
	handler "CarRental/user-service/internal/delivery/grpc"
	"CarRental/user-service/internal/usecase"
	pb "CarRental/user-service/proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/userdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, handler.NewUserHandler(uc))

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User Service started on :50052")
	s.Serve(lis)
}
