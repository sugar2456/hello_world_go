package repositories

import (
	"database/sql"
	"fmt"
	"testing"
	"os"
	"context"
	"hello_world_go/ent"
)

func TestMain(m *testing.M) {
	setupTestDB()
	// テストの実行
	code := m.Run()
	os.Exit(code)
}

// テストDBの設定とテーブルクリア
func setupTestDB() {
    client, err := ent.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("テストDB接続エラー: %v\n", err)
        os.Exit(1)
    }
    defer client.Close()
    
    ctx := context.Background()
    
    // Userテーブルのデータをクリア
    _, err = client.User.Delete().Exec(ctx)
    if err != nil {
        fmt.Printf("Userテーブルクリアエラー: %v\n", err)
        os.Exit(1)
    }
    
    // AUTO_INCREMENTをリセットするために標準DBドライバーを使用
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("DB接続エラー: %v\n", err)
        os.Exit(1)
    }
    defer db.Close()
    
    // AUTO_INCREMENTをリセット
    _, err = db.ExecContext(ctx, "ALTER TABLE users AUTO_INCREMENT = 1")
    if err != nil {
        fmt.Printf("AUTO_INCREMENTリセットエラー: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("テストDB準備完了 - userテーブルクリア済み") 

	initialTestData()
}

func initialTestData() {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("テストDB接続エラー: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	ctx := context.Background()

	// テストデータの作成
	_, err = client.User.Create().
		SetName("テストユーザー").
		SetAge(30).
		Save(ctx)
	if err != nil {
		fmt.Printf("テストデータ作成エラー: %v\n", err)
		os.Exit(1)
	}
}

// TestGetUserByID に変更（通常のテスト関数として実装）
func TestGetUserByID(t *testing.T) {

    user, err := GetUserByID(1)
    if err != nil {
        t.Fatalf("failed querying user: %v", err)
    }
    fmt.Printf("ユーザー名: %s, 年齢: %d\n", user.Name, user.Age)
	// Output:
	// ユーザー名: テストユーザー, 年齢: 30
}

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("テストユーザー", 31)
	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}
	fmt.Printf("ユーザー名: %s, 年齢: %d\n", user.Name, user.Age)
	// Output:
	// ユーザー名: テストユーザー, 年齢: 31
}

func TestUpdateUser(t *testing.T) {
	user, err := UpdateUser(1, "更新されたユーザー", 32)
	if err != nil {
		t.Fatalf("failed updating user: %v", err)
	}
	fmt.Printf("ユーザー名: %s, 年齢: %d\n", user.Name, user.Age)
	// Output:
	// ユーザー名: 更新されたユーザー, 年齢: 32
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(1)
	if err != nil {
		t.Fatalf("failed deleting user: %v", err)
	}
	fmt.Println("ユーザー削除成功")
	// Output:
	// ユーザー削除成功
}