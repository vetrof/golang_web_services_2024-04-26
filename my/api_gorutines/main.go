package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	repeat := 10000
	urls := []string{"https://arbuz.kz/ru/almaty/discount-catalog/225443-skidki/225164#/", "https://arbuz.kz/ru/almaty/discount-catalog/225443-skidki/225253#/"}
	//urls := []string{"https://galmart.kz/api/v1/catalog/shops"}

	outChan := make(chan int)

	for i := 0; i < repeat; i++ {
		for _, url := range urls {
			go getData(url, outChan)
		}

	}

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
