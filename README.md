## Basic Todo List app in Golang

## Compilation
```
go build
```
This should create `todo.exe` binary for windows and `todo` in macos

## Commands

### Create a new task

```
.\todo.exe add "First Task"
```

### View active tasks

```
.\todo.exe viewa
```

### Mark a task as completed 

```
.\todo.exe complete {task_id}
```
task_id should be fetched by running viewa command 
### View completed tasks

```
.\todo.exe viewc
```


