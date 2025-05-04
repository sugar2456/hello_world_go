package repositories

import (
    "context"
    "fmt"
    "hello_world_go/ent"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_NAME"),
)

func GetUserByID(id int) (*ent.User, error) {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	user, err := client.User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	return user, nil
}

func CreateUser(name string, age int) (*ent.User, error) {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	user, err := client.User.Create().
		SetName(name).
		SetAge(age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	return user, nil
}

func UpdateUser(id int, name string, age int) (*ent.User, error) {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	user, err := client.User.UpdateOneID(id).
		SetName(name).
		SetAge(age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %v", err)
	}
	return user, nil
}

func DeleteUser(id int) error {
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting user: %v", err)
	}
	return nil
}