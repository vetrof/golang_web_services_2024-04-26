package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	// получаем переменные из .env
	repeat, urls := getFromEnv()

	// объявляем канал
	outChan := make(chan int)

	// создаем горутины в цикле
	for i := 0; i < repeat; i++ {
		for _, url := range urls {
			go getData(url, outChan)
		}
	}

	// читаем из каналов
	allGorutines := repeat*len(urls) - 1
	for i := 0; i <= allGorutines; i++ {
		x := <-outChan
		fmt.Println(i, x)
	}
}

func getData(url string, outChan chan<- int) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Ошибка при запросе %s: %v\n", url, err)
		outChan <- 0
		return
	}
	defer resp.Body.Close()

	status := resp.StatusCode

	outChan <- status

}

func getFromEnv() (repeat int, urls []string) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	repeat, err = strconv.Atoi(os.Getenv("REPEAT"))
	if err != nil || repeat <= 0 {
		log.Fatalf("Некорректное значение REPEAT в .env: %v", err)
	}

	urls = strings.Split(os.Getenv("URLS"), ",")
	if len(urls) == 0 || urls[0] == "" {
		log.Fatalf("URLS не заданы в .env")
	}
	return repeat, urls
}

// getInputInt получает целое число от пользователя
func getInputInt() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil || number <= 0 {
			fmt.Println("Введите корректное положительное число.")
			continue
		}
		return number
	}
}

// getInputURLs получает список URL от пользователя
func getInputURLs() []string {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}
		input = strings.TrimSpace(input)
		urls := strings.Split(input, ",")
		if len(urls) == 0 || urls[0] == "" {
			fmt.Println("Введите хотя бы один URL.")
			continue
		}
		return urls
	}
}
