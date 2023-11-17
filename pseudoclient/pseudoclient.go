package main

import (
	"CustomGRPC/proto"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serviceClient := connection()
	for i := 0; i < 10000; i++ {
		go callNew(serviceClient, i)
	}
	// Esperamos por la respuesta
	// sin que finalize nuestra funcion main
	time.Sleep(5 * time.Second)
}

// Creamos la conexion con el servidor que ya
// debería estar corriendo en el puerto 50051.
func connection() proto.NotificationServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No se puede conectar con el servidor")
	}

	// Si la conexion es exitosa lo pasamos como argumento
	// a la estructura que trabajará como cliente.
	serviceClient := proto.NewNotificationServiceClient(conn)
	fmt.Printf("Service Client Created | Type: %T", serviceClient)
	return serviceClient
}

func callNew(serviceClient proto.NotificationServiceClient, index int) {
	// Realizamos una llamada inmediata a nuetro microservicio
	// usando el metodo preestablecido en el proto (en este caso New).
	message := fmt.Sprintf("Este es el mensaje %v", index)

	resp, err := serviceClient.New(context.Background(), &proto.NewNotifReq{
		Msj: message,
	})

	// Verificamos un posible error
	if err != nil {
		panic("Error en peticion al microservicio")
	}

	// Si no existe un error
	// Imprimimos el resultado
	fmt.Println(resp.Feedback)
}
