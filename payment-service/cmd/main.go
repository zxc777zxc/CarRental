package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"CarRental/payment-service/config"
	"CarRental/payment-service/infrastructure/repository"
	grp "CarRental/payment-service/internal/delivery/grpc"
	"CarRental/payment-service/internal/usecase"
	pb "CarRental/payment-service/proto"
	_ "github.com/lib/pq"

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

	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	repo := repository.NewPaymentRepo(db)
	uc := usecase.NewPaymentUsecase(repo)
	handler := grp.NewPaymentHandler(uc)

	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, handler)
	log.Println("Payment Service is running on :50057")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
