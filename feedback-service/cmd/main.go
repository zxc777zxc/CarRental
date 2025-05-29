package main

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"CarRental/feedback-service/config"
	"CarRental/feedback-service/infrastructure/repository"
	grp "CarRental/feedback-service/internal/delivery/grpc"
	"CarRental/feedback-service/internal/usecase"
	pb "CarRental/feedback-service/proto"
	_ "github.com/lib/pq"

	"github.com/prometheus/client_golang/prometheus"
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

	repo := repository.NewFeedbackRepo(db)
	uc := usecase.NewFeedbackUsecase(repo)
	handler := grp.NewFeedbackHandler(uc)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFeedbackServiceServer(grpcServer, handler)
	log.Println("Feedback Service is running on :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
