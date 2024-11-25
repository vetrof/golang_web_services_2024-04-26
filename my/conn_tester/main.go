package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// получаем переменные из .env
	step, max, url := getFromEnv()

	//объявляем канал
	outChan := make(chan int64)

	for connects := 0; connects < max; connects += step {

		for range connects {
			go getData(url[0], outChan)
		}

		var allTime int64
		//читаем из каналов
		for range connects {
			t := <-outChan
			allTime += t
		}

		var midTime int64
		if connects > 0 {
			midTime = allTime / int64(connects)
		}

		fmt.Println(connects, float64(midTime)/1000, "s")

	}
}

func getData(url string, outChan chan<- int64) {
	start := time.Now() // Засекаем время начала
	resp, err := http.Get(url)

	if err != nil {
		panic(fmt.Sprintf("Ошибка при запросе %s", err))
	}
	defer resp.Body.Close()

	status := resp.StatusCode

	if status != 200 {
		panic(fmt.Sprintf("Получен статус - %d", status))
	}

	duration := time.Since(start).Milliseconds()
	outChan <- duration
}

func getFromEnv() (step int, max int, url []string) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	step, err = strconv.Atoi(os.Getenv("STEP"))
	if err != nil || step <= 0 {
		log.Fatalf("Некорректное значение step в .env: %v", err)
	}

	max, err = strconv.Atoi(os.Getenv("MAX"))
	if err != nil || max <= 0 {
		log.Fatalf("Некорректное значение max в .env: %v", err)
	}

	url = strings.Split(os.Getenv("URL"), ",")
	if len(url) == 0 || url[0] == "" {
		log.Fatalf("URLS не заданы в .env")
	}
	return step, max, url
}
