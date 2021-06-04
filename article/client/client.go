package client

import (
	"app/article/pb"

	"google.golang.org/grpc"
)

// Clientからサービスを呼び出せるようにする
type Client struct {
	conn    *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	// client connectionを生成
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// articleサービスのclientを生成
	c := pb.NewArticleServiceClient(conn)

	// articleサービスのclientを返す
	return &Client{conn, c}, nil
}
