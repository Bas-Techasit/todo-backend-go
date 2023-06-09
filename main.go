package main

import (
	"fmt"
	"strings"
	"time"
	"todo-backend/handlers"
	"todo-backend/repositories"
	"todo-backend/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberjwt "github.com/gofiber/jwt/v3"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()
	db := initDB()

	todoRepository := repositories.NewCustomerRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	// userRepository := repositories.NewUserRepository(db)
	userRepositoryMock := repositories.NewUserRepositoryMock()
	userService := services.NewUserService(userRepositoryMock)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()

	app.Use(logger.New())
	app.Use("/todos", fiberjwt.New(fiberjwt.Config{
		SigningMethod:  "HS256",
		SigningKey:     []byte(viper.GetString("app.secret")),
		SuccessHandler: func(c *fiber.Ctx) error { return c.Next() },
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))

	app.Get("/todos/username/:username", todoHandler.GetTodos)
	app.Post("/todos/username/:username", todoHandler.NewTodo)
	app.Put("/todos/:todoID/username/:username", todoHandler.UpdateTodo)
	app.Delete("/todos/:todoID/username/:username", todoHandler.DeleteTodo)

	app.Post("/signup", userHandler.SignUp)
	app.Post("/login", userHandler.Login)

	app.Listen(":" + viper.GetString("app.port"))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func initDB() *sqlx.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
