package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tasks := make([]string, 0)

	for {
		fmt.Print("Enter a task (or 'exit' to quit): ")
		task, _ := reader.ReadString('\n')
		task = task[:len(task)-1]

		if task == "exit" {
			break
		}

		tasks = append(tasks, task)
	}
	fmt.Println("Your to-do list:")

	for i, task := range tasks {
		fmt.Printf("%d. %s", i+1, task)
	}

}
