package main

import (
	"CustomGRPC/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serviceClient := connection()
	for i := 1; i < 10000; i++ {
		callNew(serviceClient)
	}
}

// Creamos la conexion con el servidor que ya
// debería estar corriendo en el puerto 50051.
func connection() proto.NotificationServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se puede conectar con el servidor")
	}

	// Si la conexion es exitosa lo pasamos como argumento
	// a la estructura que trabajará como cliente.
	serviceClient := proto.NewNotificationServiceClient(conn)
	fmt.Printf("Service Client: %v; Type: %T", serviceClient, serviceClient)
	return serviceClient
}

func callNew(serviceClient proto.NotificationServiceClient) {
	// Realizamos una llamada inmediata a nuetro microservicio
	// usando el metodo preestablecido en el proto (en este caso New).
	resp, err := serviceClient.New(context.Background(), &proto.NewNotifReq{
		Msj: "Este es un mensaje",
	})

	// Verificamos un posible error
	if err != nil {
		panic("Error en peticion al microservicio")
	}

	// Si no existe un error
	// Imprimimos el resultado
	fmt.Println(resp.Feedback)
}
