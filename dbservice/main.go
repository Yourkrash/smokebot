package main

import (
	"context"
	// "errors"
	"flag"
	"fmt"
	"log"
	"net"
	// "time"

	pb "smokebot/dbservice/proto"
	// "github.com/google/uuid"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

type User struct {
	ID             int64 `gorm:"primarykey"`
	FirstName      string
	SecondName     string
	UserSubscriber []UserSubscriber `gorm:"foreignKey:ID"`
}

type UserSubscriber struct {
	ID    string `gorm:"primarykey"`
	SubID string `gorm:"primarykey"`
}

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "pass1234"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(User{})
	DB.AutoMigrate(UserSubscriber{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func (*server) RegUser(ctx context.Context, req *pb.RegUserRequest) (*pb.ErrorResponse, error) {
	user := req.GetUser()
	DB.Create(&User{
		ID:         user.GetIdUser(),
		FirstName:  user.GetName(),
		SecondName: user.GetName(),
	})
	return &pb.ErrorResponse{Error: ""}, nil
}

func (*server) IsRegUser(ctx context.Context, req *pb.UserID) (*pb.BoolResponse, error) {
	return &pb.BoolResponse{Isreg: true}, nil // TODO
}

func main() {
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMovieServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
