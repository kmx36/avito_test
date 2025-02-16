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

    // Маршрутизация
    http.HandleFunc("/api/auth", authHandler.Authenticate)
    http.Handle("/api/info", middleware.AuthMiddleware("your_jwt_secret")(http.HandlerFunc(infoHandler.GetUserInfo)))
    http.Handle("/api/sendCoin", middleware.AuthMiddleware("your_jwt_secret")(http.HandlerFunc(sendCoinHandler.SendCoins)))
    http.Handle("/api/buy/{item}", middleware.AuthMiddleware("your_jwt_secret")(http.HandlerFunc(buyHandler.BuyItem)))

    // Запуск сервера
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}