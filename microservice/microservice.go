package main

import (
	"CustomGRPC/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var (
	serverType = "tcp"
	port       = ":50051"
	counter    = 0
)

// Creamos una estructura que implemente
// la interfaz de la que nos provee proto.
type NotificationServer struct {
	proto.UnimplementedNotificationServiceServer
}

// Creamos un metodo de la interfaz anterior mencionada
// que use el metodo (en este caso, pueden ser varios y
// personalizados) que se necesita para implementar
// la interfaz.
func (s *NotificationServer) New(ctx context.Context, req *proto.NewNotifReq) (*proto.NewNotifResp, error) {
	counter++
	fmt.Printf("Nueva notification %v\n", counter)
	feedback := fmt.Sprintf("Este es el Feedback %v", counter)
	return &proto.NewNotifResp{
		Feedback: feedback,
	}, nil
}

// En la funcion main (en este caso), establecemos
// el servidor que escuchará en el puerto 50051.
func main() {
	// Creamos un listener con el puerto correspondiente
	listener, err := net.Listen(serverType, port)
	if err != nil {
		panic("No se pudo crear el servidor" + err.Error())
	}

	// Creamos una intancia de servidor de grpc
	server := grpc.NewServer()

	// Registramos el servicio
	proto.RegisterNotificationServiceServer(server, &NotificationServer{})

	// Iniciamos el servicio
	fmt.Printf("Server is running in localhost%v\n", port)
	fmt.Printf("- Server Type: %v\n", serverType)
	fmt.Printf("- Server Port: %v\n\n", port)
	err = server.Serve(listener)
	if err != nil {
		panic("Error ejecutando el servidor")
	}
}
