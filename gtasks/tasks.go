package gtasks

import (
	"fmt"
	"log"

	"google.golang.org/api/tasks/v1"
)

var Service *tasks.Service

func createTask(task string, taskList string) {}
func getTasks(tasklist string) {}
func GetTaskLists() {
	r, err := Service.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}

	fmt.Println("Task Lists:")
	if len(r.Items) > 0 {
		for _, i := range r.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
		}
	} else {
		fmt.Print("No task lists found.")

	}
}
func CreateTaskList(listName string){}

