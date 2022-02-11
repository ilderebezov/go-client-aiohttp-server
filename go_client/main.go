package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var path string
var fullPath string
var statusCode int

func main() {
	var number, choice, url string
	var address, port, path, fileName string

	for {
		fmt.Println("\n")
		fmt.Printf("\nPress 1 to read config.cfg file\n" +
			"Press 2 to put config data manualy\n" +
			"Press 3 to STOP client\n")
		fmt.Scanf("%s", &choice)

		if choice == string('3') {
			break
		} else if choice == string('1') {
			fmt.Printf("read config.cfg\n")
			url, path = readConfig()
		} else if choice == string('2') {
			fmt.Printf("Put server data manualy\n")
			fmt.Printf("Put server adress:\n")
			fmt.Scanf("%s", &address)
			fmt.Printf("Put server port:\n")
			fmt.Scanf("%s", &port)
			fmt.Printf("Put full path to save file:\n")
			fmt.Scanf("%s", &path)

			url = address + ":" + port
		}

		fmt.Printf("\nPress 1 to GET list of files on server\n" +
			"Press 2 to GET file from server\n" +
			"Press 3 to PUT upload file to server\n" +
			"Press 4 to POST update file on server\n" +
			"Press 5 to DELETE file from server\n" +
			"Press 6 to STOP client\n")
		fmt.Scanf("%s", &number)

		if number == string('6') {
			break
		} else if number == string('5') {
			fmt.Printf("\nMethod DELETE, put filename:\n")
			fmt.Scanf("%s", &fileName)
			delFile(fileName, url)
		} else if number == string('4') {
			fmt.Printf("\nMethod POST, put filename:\n")
			fmt.Scanf("%s", &fileName)
			postFile(fileName, url)
		} else if number == string('3') {
			fmt.Printf("\nMethod PUT, put filename:\n")
			fmt.Scanf("%s", &fileName)
			putFile(fileName, url)
		} else if number == string('2') {
			fmt.Printf("\nMethod GET, put filename:\n")
			fmt.Scanf("%s", &fileName)
			getFile(fileName, url, path)
		} else if number == string('1') {
			fmt.Printf("\nGET list of files and their hash value\n")
			getList(url)
		}
	}
}

func getList(url string) {
	url = url + "/list"
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Printf("\n")
}

func getFile(fileName string, url string, path string) {
	url = url + "/file?name="
	fullPath = url + fileName
	client := http.Client{}
	resp, err := client.Get(fullPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	statusCode = resp.StatusCode
	if statusCode == 227 {
		fmt.Printf("\nNot Ok\n")
	}
	if statusCode == 228 {
		fmt.Printf("\nOk\n")
		file, err := os.Create(filepath.Join(path, filepath.Base(fileName)))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString(string(body))
	}
}

func putFile(fileName string, url string) {

	url = url + "/upload/file"

	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	// Закрываем файл по завершению
	defer file.Close()

	// Буфер для хранения нашего тела запроса в виде байтов
	var requestBody bytes.Buffer

	// Создаем писателя
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Инициализируем поле
	fileWriter, err := multiPartWriter.CreateFormFile("file", fileName)

	// Скопируем содержимое файла в поле
	_, err = io.Copy(fileWriter, file)

	// Заполняем остальные поля
	fieldWriter, err := multiPartWriter.CreateFormField("data")

	_, err = fieldWriter.Write([]byte("Value"))

	// Закрываем запись данных
	multiPartWriter.Close()

	// Создаем объект реквеста
	req, err := http.NewRequest("PUT", url, &requestBody)

	// Получаем и устанавливаем тип контента
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Отправляем запрос
	client := &http.Client{}
	response, err := client.Do(req)

	statusCode := response.StatusCode
	if statusCode == 228 {
		fmt.Println("Not Ok")
	}
	if statusCode == 229 {
		fmt.Println("Ok")
	}
}

func postFile(fileName string, url string) {
	url = url + "/update/file"
	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	// Закрываем файл по завершению
	defer file.Close()

	// Буфер для хранения нашего тела запроса в виде байтов
	var requestBody bytes.Buffer

	// Создаем писателя
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Инициализируем поле
	fileWriter, err := multiPartWriter.CreateFormFile("data", fileName)

	// Скопируем содержимое файла в поле
	_, err = io.Copy(fileWriter, file)

	// Заполняем остальные поля
	fieldWriter, err := multiPartWriter.CreateFormField("data")

	_, err = fieldWriter.Write([]byte("Value"))

	// Закрываем запись данных
	multiPartWriter.Close()

	// Создаем объект реквеста
	req, err := http.NewRequest("POST", url, &requestBody)

	// Получаем и устанавливаем тип контента
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Отправляем запрос
	client := &http.Client{}
	response, err := client.Do(req)

	statusCode := response.StatusCode
	if statusCode == 227 {
		fmt.Println("Not Ok")
	}
	if statusCode == 228 {
		fmt.Println("Not Ok")
	}
	if statusCode == 229 {
		fmt.Println("Ok")
	}
}

func delFile(fileName string, url string) {
	url = url + "/delete/file"
	// Create client
	client := &http.Client{}
	fileNameByte := strings.NewReader(fileName)
	// Create request
	req, err := http.NewRequest("DELETE", url, fileNameByte)
	if err != nil {
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode == 228 {
		fmt.Println("Not Ok")
	}
	if statusCode == 229 {
		fmt.Println("Ok")
	}
}

func readConfig() (string, string) {
	f, err := os.Open("config.cfg")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0
	address := ""
	port := ""
	path := ""
	for scanner.Scan() {
		if i == 0 {
			address += scanner.Text()
		}
		if i == 1 {
			port += scanner.Text()
		}
		if i == 2 {
			path += scanner.Text()
		}
		i += 1
	}
	url := address + ":" + port
	return url, path
}
