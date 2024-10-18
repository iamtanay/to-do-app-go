package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Watch Go tutorial", "Read a book"}

func main() {
	fmt.Println("Welcome to our TO-DO APP!")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Welcome to our TO-DO APP!"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for index, task := range taskItems {
		fmt.Fprintf(writer, "%d. %s\n", index+1, task)
	}
}
