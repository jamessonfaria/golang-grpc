package main
import (
	"github.com/jamessonfaria/golang-grpc/services"
	"github.com/jamessonfaria/golang-grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
