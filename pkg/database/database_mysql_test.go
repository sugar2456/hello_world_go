package database

import (
	"context"
	"fmt"
	"hello_world_go/ent"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ExampleUser() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// テストユーザーを作成
	u, err := client.User.Create().
		SetName("テストユーザー").
		SetAge(30).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	// 作成したユーザーを取得
	retrievedUser, err := client.User.Get(ctx, u.ID)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	fmt.Printf("ユーザー名: %s, 年齢: %d\n", retrievedUser.Name, retrievedUser.Age)
	// Output:
	// ユーザー名: テストユーザー, 年齢: 30
}

func ExampleVideos() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// テストビデオデータを作成
	video, err := client.Videos.Create().
		SetTitle("テストビデオ").
		SetDescription("これはテストビデオです").
		SetURL("http://example.com/test_video").
		SetThumbnail("http://example.com/test_thumbnail").
		SetCategory("テストカテゴリ").
		SetTags("テストタグ").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating video: %v", err)
	}
	// 作成したビデオを取得
	retrievedVideo, err := client.Videos.Get(ctx, video.ID)
	if err != nil {
		log.Fatalf("failed querying video: %v", err)
	}
	fmt.Printf("ビデオタイトル: %s, 説明: %s, URL: %s\n", retrievedVideo.Title, retrievedVideo.Description, retrievedVideo.URL)
	// Output:
	// ビデオタイトル: テストビデオ, 説明: これはテストビデオです, URL: http://example.com/test_video
}
