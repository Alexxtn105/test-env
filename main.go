package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// ВНИМАНИЕ! Не забываем добавлять в .gitignore ненужные файлы!!!
	// загружаем .env-файлы
	// в данном проекте исползуются следующие .env-файлы:
	// .env - для локальной разработки
	// .env.staging -
	// .env.prod 	-
	godotenv.Load()
	fmt.Println(os.Getenv("TEST_ENV_VAR"))
}
