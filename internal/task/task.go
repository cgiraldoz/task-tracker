package task

import "fmt"

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}

func AddTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify a task to add")
		return
	}

	description := Description(args[0])

	err := AddNewTask(description)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func ListTasks() {
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("You have the following tasks:")
	for _, task := range tasks {
		fmt.Printf("%d: %s\n - Status: %s\n - Created At: %s\n - Updated At: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdateAt)
	}
}
