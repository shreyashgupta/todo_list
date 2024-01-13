package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
}

func NewTask(content string, id int) *Task {

	task := Task{
		Id:        id,
		Content:   content,
		CreatedAt: time.Now().Local().String(),
	}
	return &task
}

func (task *Task) PrintTask() {
	w := 50
	content := fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(task.Content))/2, task.Content))
	fmt.Printf("%5d | %s | %20s\n", task.Id, content, task.CreatedAt)
}

func printColumns() {
	w := 50
	content := fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len("Task"))/2, "Task"))
	fmt.Printf("%5s | %s | %20s\n", "Id", content, "Created At")
	for i := 0; i < 100; i++ {
		fmt.Print("=")
	}
	fmt.Println("")
}

type TodoList struct {
	ActiveTasks    []Task `json:"active_tasks"`
	CompletedTasks []Task `json:"completed_tasks"`
	FilePath       string `json:"file_path"`
}

func newTodoList(FilePath string) *TodoList {
	todoList := TodoList{
		FilePath: FilePath,
	}
	return &todoList
}

func (todoList *TodoList) doesActiveTasksContainsIndex(index int) bool {
	for _, task := range todoList.ActiveTasks {
		if task.Id == index {
			return true
		}
	}
	return false
}

func (todoList *TodoList) doesCompletedTasksContainsIndex(index int) bool {
	for _, task := range todoList.CompletedTasks {
		if task.Id == index {
			return true
		}
	}
	return false
}

func (todoList *TodoList) getUnusedIndex() int {
	i := 0
	for {
		if !todoList.doesActiveTasksContainsIndex(i) && !todoList.doesCompletedTasksContainsIndex(i) {
			return i
		}
		i++
	}
}

func (todoList *TodoList) AddTask(content string) error {
	task := NewTask(content, todoList.getUnusedIndex())
	todoList.ActiveTasks = append(todoList.ActiveTasks, *task)
	err := todoList.writeToFile()
	if err != nil {
		return err
	}
	return nil
}

func ReadTodoListFromFile(FilePath string) (*TodoList, error) {
	todoList := newTodoList(FilePath)
	marshalled, err := os.ReadFile(FilePath)
	if os.IsNotExist(err) {
		err = todoList.writeToFile()
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, errors.New(fmt.Sprint("Failed to read from json file due to error: ", err))
	} else {
		err = json.Unmarshal(marshalled, &todoList)
		if err != nil {
			return nil, errors.New(fmt.Sprint("Failed to deserialize todo list due to error: ", err.Error()))
		}
	}
	return todoList, nil
}

func (todoList *TodoList) writeToFile() error {
	marshalled, err := json.Marshal(todoList)
	if err != nil {
		return errors.New(fmt.Sprint("Failed to serialize todo list due to error: ", err))
	}
	err = os.WriteFile(todoList.FilePath, marshalled, 0644)
	if err != nil {
		return errors.New(fmt.Sprint("Failed to write serialized todo list to file due to error: ", err))
	}
	return nil
}

func (todoList *TodoList) PrintCompletedTasks() {
	fmt.Println("")
	printColumns()
	if len(todoList.CompletedTasks) == 0 {
		fmt.Println("No completed tasks")
	} else {
		for _, task := range todoList.CompletedTasks {
			task.PrintTask()
		}
	}
	fmt.Println("")
}

func (todoList *TodoList) PrintActiveTasks() {
	fmt.Println("")
	printColumns()
	if len(todoList.ActiveTasks) == 0 {
		fmt.Println("No active tasks")
	} else {
		for _, task := range todoList.ActiveTasks {
			task.PrintTask()
		}
	}
	fmt.Println("")
}

func (todoList *TodoList) getTaskIndexById(id int) (int, error) {
	for index, task := range todoList.ActiveTasks {
		if task.Id == id {
			return index, nil
		}
	}
	return -1, errors.New(fmt.Sprint("Failed to find task with id: ", id))
}

func (todoList *TodoList) MarkTaskCompleted(id int) error {
	comletedTaskIndex, err := todoList.getTaskIndexById(id)
	if err != nil {
		return err
	}
	completedTask := todoList.ActiveTasks[comletedTaskIndex]

	todoList.ActiveTasks = append(todoList.ActiveTasks[:comletedTaskIndex], todoList.ActiveTasks[comletedTaskIndex+1:]...)
	todoList.CompletedTasks = append(todoList.CompletedTasks, completedTask)
	err = todoList.writeToFile()
	if err != nil {
		return err
	}
	return nil
}
