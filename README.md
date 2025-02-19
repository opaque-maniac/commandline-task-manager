# commandline-task-manager
Simple task manager for the commandline
It is a simple task manager for the commandline. It allows you to add, remove, list todo tasks.
It persists the tasks in a file so that you can access them later.

## Installation
1. Run the following command to clone the repository.
``` shell
https://github.com/opaque-maniac/commandline-task-manager.git
```

2. Run the following command to install the dependencies.
``` shell
go build -o <executable_name>
```

3. Run the following command to run the application.
``` shell
./<executable_name> help
```

## Usage
1. To get help, run the following command.
``` shell
./<executable_name> help
```

2. To add a task, run the following command.
``` shell
./<executable_name> add <task>
```

3. To list all the tasks, run the following command.
``` shell
./<executable_name> list
```

4. To remove a task, run the following command.
``` shell
./<executable_name> remove <task_name>
```

5. To remove all the tasks, run the following command.
``` shell
./<executable_name> remove-all
```
