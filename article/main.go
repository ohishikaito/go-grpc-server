package main

import (
	"app/article/client"
	"app/article/pb"
	"context"
	"fmt"
	"io"
	"log"
)

// gRPCサーバーの動作確認用
func main() {
	// clientを生成
	// c, _ := client.NewClient("localhost:50051")
	c, _ := client.NewClient("go-grpc-server_app_1:50051")
	// c, _ := client.NewClient("go-grpc-server_app_1")

	create(c)
	read(c)
	update(c)
	delete(c)
	list(c)
}

func create(c *client.Client) {
	// 記事をCREATE
	input := &pb.ArticleInput{
		Author:  "gopher",
		Title:   "gRPC",
		Content: "gRPC is so cool!",
	}
	res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to CreateArticle: %v\n", err)
	}
	fmt.Printf("CreateArticle Response: %v\n", res)
}

func read(c *client.Client) {
	// 記事をREAD
	var id int64 = 2
	res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("Failed to ReadArticle: %v\n", err)
	}
	fmt.Printf("ReadArticle Response: %v\n", res)
}

func update(c *client.Client) {
	// 記事をUPDATE
	var id int64 = 2
	input := &pb.ArticleInput{
		Author:  "GraphQL master",
		Title:   "GraphQL",
		Content: "GraphQL is very smart!",
	}
	res, err := c.Service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to UpdateArticle: %v\n", err)
	}
	fmt.Printf("UpdateArticle Response: %v\n", res)
}

func delete(c *client.Client) {
	// 記事をDELETE
	var id int64 = 1
	res, err := c.Service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("Failed to UpdateArticle: %v\n", err)
	}
	fmt.Printf("The article has been deleted (%v)\n", res)
}

func list(c *client.Client) {

	// streamレスポンスを受け取る
	stream, err := c.Service.ListArticle(context.Background(), &pb.ListArticleRequest{})
	if err != nil {
		log.Fatalf("Failed to ListArticle: %v\n", err)
	}

	// Server Streamingで渡されたレスポンスを１つ１つ出力
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to Server Streaming: %v\n", err)
		}
		fmt.Println(res)
	}
}
