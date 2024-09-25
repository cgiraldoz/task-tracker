package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

const filePermission = 0644

func getTaskFilePath() string {
	currentWorkingDirectory := getCurrentWorkingDirectory()

	return path.Join(currentWorkingDirectory, "tasks.json")
}

func getCurrentWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory")
		os.Exit(1)
	}

	return cwd
}

func GetTasks() ([]Task, error) {
	filepath := getTaskFilePath()

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return createEmptyTaskFile(filepath)
	}

	var tasks []Task

	err := readFromFile(filepath, &tasks)

	if err != nil {
		return nil, fmt.Errorf("error reading from task file: %w", err)
	}
	fmt.Println("Cristian:")
	return tasks, nil
}

func readFromFile(filepath string, tasks *[]Task) error {
	file, fileOpeningError := os.Open(filepath)

	if fileOpeningError != nil {
		return fmt.Errorf("error opening task file: %w", fileOpeningError)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing task file")
		}
	}(file)

	if decodingError := json.NewDecoder(file).Decode(tasks); decodingError != nil {
		return fmt.Errorf("error decoding task file: %w", decodingError)
	}

	return nil
}

func createEmptyTaskFile(filepath string) ([]Task, error) {
	err := os.WriteFile(filepath, []byte("[]"), filePermission)

	if err != nil {
		return nil, fmt.Errorf("error creating and writing to task file: %w", err)
	}

	return []Task{}, nil
}

func AddNewTask(description Description) error {
	tasks, err := GetTasks()

	newTask := Task{
		ID:          getLatestTaskID(tasks) + 1,
		Description: string(description),
		Status:      "incomplete",
		CreatedAt:   getCurrentTime(),
		UpdateAt:    getCurrentTime(),
	}

	if err != nil {
		return fmt.Errorf("error getting tasks: %w", err)
	}

	tasks = append(tasks, newTask)

	newTaskStyle := NewTaskStyle(WarningColor).Render(fmt.Sprintf("ID: %d - Description: %s", newTask.ID, newTask.Description))

	fmt.Printf("Adding new task: %s\n", newTaskStyle)

	return writeTasksToFile(getTaskFilePath(), tasks)
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

func getLatestTaskID(tasks []Task) int64 {
	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].ID
}

func writeTasksToFile(filepath string, tasks []Task) error {
	file, fileOpeningError := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, filePermission)

	if fileOpeningError != nil {
		return fmt.Errorf("error opening task file: %w", fileOpeningError)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing task file")
		}
	}(file)

	encoder := json.NewEncoder(file)

	if encodingError := encoder.Encode(tasks); encodingError != nil {
		return fmt.Errorf("error encoding tasks to file: %w", encodingError)
	}

	return nil
}
