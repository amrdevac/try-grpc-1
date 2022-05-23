package main

import (
	"context"
	"encoding/json"
	"fmt"
	blog "grpc/proto"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// Implement Function

func (blog *DataStructBlog) FindOneByTitle(c context.Context, request *blog.Blog) (*blog.Blog, error) {
	for _, value := range blog.blog {
		if value.Title == request.Title {
			return value, nil
		}
	}
	return nil, nil
}

func (blog *DataStructBlog) FindOneByBody(c context.Context, request *blog.Blog) (*blog.Blog, error) {
	for _, value := range blog.blog {
		if value.Bdoy == request.Bdoy {
			return value, nil
		}
	}
	return nil, nil
}

// Setup

func main() {
	fmt.Println("start")
	listen, _ := net.Listen("tcp", "127.0.0.1:1300")
	grpcServer := grpc.NewServer()
	blog.RegisterMainServiceServer(grpcServer, startServer())
	grpcServer.Serve(listen)

}

func startServer() *DataStructBlog {
	dataBlog := DataStructBlog{}
	dataBlog.loadData()
	return &dataBlog
}

func (dataBlog *DataStructBlog) loadData() {
	contohDataBlog, err := ioutil.ReadFile("data/json.json")
	if err != nil {
		log.Fatalf("error baca data json")
	}

	dataErr := json.Unmarshal(contohDataBlog, &dataBlog.blog)
	if dataErr != nil {
		log.Fatalf("error unmarshal data json")
	}

}

type DataStructBlog struct {
	blog.UnimplementedMainServiceServer
	mu   sync.Mutex
	blog []*blog.Blog
}
