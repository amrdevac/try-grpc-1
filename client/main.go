package main

import (
	"context"
	"fmt"
	blog "grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func getByTitle(blogClient blog.MainServiceClient, title string) {
	connection, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	structImplement := blog.Blog{Title: title}
	getBlog, _ := blogClient.FindOneByTitle(connection, &structImplement)
	fmt.Println(getBlog)
}

func getByBody(blogClient blog.MainServiceClient, Bdoy string) {
	connection, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	structImplement := blog.Blog{Bdoy: Bdoy, Title: "mantap"}
	getBlog, _ := blogClient.FindOneByBody(connection, &structImplement)
	fmt.Println(getBlog)
}

func main() {
	var setupConnection []grpc.DialOption
	setupConnection = append(setupConnection, grpc.WithInsecure())
	setupConnection = append(setupConnection, grpc.WithBlock())
	connecting, err := grpc.Dial("127.0.0.1:1300", setupConnection...)
	if err != nil {
		log.Fatalln("failed to dial")
	}
	defer connecting.Close()

	client := blog.NewMainServiceClient(connecting)

	getByTitle(client, "fahmi")
	// getByBody(client, "sahrul@gmail.com")

}
