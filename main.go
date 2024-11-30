package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// ВНИМАНИЕ! Не забываем добавлять в .gitignore ненужные файлы!!!
	// загружаем .env-файлы
	// в данном проекте используются следующие .env-файлы:
	// .env - для локальной разработки
	// .env.staging -
	// .env.prod 	-

	//godotenv.Load()  // по умолчанию будет загружен файл .env
	godotenv.Load("./.env.production") // указываем конкретный файл для загрузки

	fmt.Println(os.Getenv("TEST_ENV_VAR"))
}
