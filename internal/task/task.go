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

	task := Task{
		ID:          2,
		Description: args[0],
		Status:      "incomplete",
		CreatedAt:   "2021-07-01T00:00:00Z",
		UpdateAt:    "2021-07-01T00:00:00Z",
	}

	fmt.Printf("Added \"%s\" to your task list.\n", task.Description)

	err := AddNewTask(task)
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
