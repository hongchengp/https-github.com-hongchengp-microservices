package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hongchengp/microservices-proto/golang/order"
	"github.com/hongchengp/microservices/order/internal/ports"
	"google.golang.org/grpc"
)

// 实现adapter，也就是 grpc依赖接口的实现
// so，要实现 port的接口，而且要嵌入grpc服务器，then重写它，来真正完成功能
// 它是输入Port， 也就是给别人用的，依赖core逻辑
// db，数据库，是给core用的

type Adapter struct {
	api ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api: api,
		port: port,
	}
}


func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	fmt.Println("yyj ai hcp")
	return &order.CreateOrderResponse{}, nil
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	fmt.Println("kai shi")
	if err != nil {
		log.Fatalf("fail on listen on port %v, error %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc on port")
	}
}

