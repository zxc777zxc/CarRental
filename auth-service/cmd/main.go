package main

import (
	"database/sql"
	"log"
	"net"

	"CarRental/auth-service/infrastructure/repository"
	handler "CarRental/auth-service/internal/delivery/grpc"
	"CarRental/auth-service/internal/usecase"
	pb "CarRental/auth-service/proto"

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

	db, err := sql.Open("postgres", "postgres://postgres:123@postgres:5432/car_rental?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewAuthRepository(db)
	uc := usecase.NewAuthUsecase(repo)
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, handler.NewAuthHandler(uc))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Auth Service running on :50051")
	s.Serve(lis)
}
