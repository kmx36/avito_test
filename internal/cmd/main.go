package main

import (
    "database/sql"
    "log"
    "net/http"
    "avito_test/internal/handlers"
    "avito_test/internal/middleware"
    "avito_test/internal/repository"
    "avito_test/internal/service"
    _ "github.com/lib/pq"
    "github.com/go-chi/chi/v5"
)

func main() {
    // Параметры подключения к PostgreSQL
    connStr := "user=avito_test2.0 password=7049 dbname=avito_test2.0 sslmode=disable host=localhost port=5432"

    // Подключение к базе данных
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Проверка подключения
    err = db.Ping()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    log.Println("Successfully connected to the database")

    // Инициализация репозиториев
    userRepo := repository.NewUserRepository(db)
    itemRepo := repository.NewItemRepository(db)
    transactionRepo := repository.NewTransactionRepository(db)

    // Инициализация сервисов
    authService := service.NewAuthService(userRepo, "your_jwt_secret")
    userService := service.NewUserService(userRepo)
    itemService := service.NewItemService(itemRepo)
    transactionService := service.NewTransactionService(transactionRepo, userRepo)

    // Инициализация обработчиков
    authHandler := handlers.NewAuthHandler(authService)
    infoHandler := handlers.NewInfoHandler(userService, itemService, transactionService)
    sendCoinHandler := handlers.NewSendCoinHandler(transactionService)
    buyHandler := handlers.NewBuyHandler(itemService, userService, transactionService)

    // Создание роутера с использованием chi
    r := chi.NewRouter()

    // Маршрутизация
    r.Post("/api/auth", authHandler.Authenticate)
    r.With(middleware.AuthMiddleware("your_jwt_secret")).Get("/api/info", infoHandler.GetUserInfo)
    r.With(middleware.AuthMiddleware("your_jwt_secret")).Post("/api/sendCoin", sendCoinHandler.SendCoins)
    r.With(middleware.AuthMiddleware("your_jwt_secret")).Post("/api/buy/{item}", buyHandler.BuyItem)

    // Запуск сервера
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}