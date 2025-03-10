// delete task
export async function deleteTodo(task) {
    try {
        const url = `/api/remove/${task}`;
        const options = {
            method: "Delete",
            headers: {
                "Content-Type": "application/json"
            },
        };

        const resposne = await fetch(url, options);

        if (!resposne.ok) {
            throw new Error("Error deleting task");
        }
        
        const data = await resposne.json();
        return data
    } catch(error) {
        throw error
    }
}

// send new task
export async function sendNewTodo(todo) {
    if (todo) {
        try {
            const url = "/api/submit"
            const body =  JSON.stringify({task: todo})

            const options = {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body,
            };

            const response = await fetch(url, options)

            if (!response.ok) {
                throw new Error("Error sending new todo")
            }

            const data = await response.json();
            return data;
        } catch(error) {
            console.log(error instanceof Error);
            throw error;
        }
    }
}


// complete
export async function completeTodo(task) {
    try {
        const url = `/api/complete/${task}`;
        const options = {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
        };

        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error("Failed to complete task");
        }

        const data = await response.json();
        console.log(data);
        return data;
    } catch(error) {
        throw error;
    }
}

// Update
export async function updateTask(task, newTask) {
    try {
        const url = `/api/update/${task}`;
        const options = {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({newName: newTask})
        }

        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error("Failed to update task");
        }

        const data = await response.json();
        return data;
    } catch(error) {
        throw error;
    }
}

// remove all
export async function removeAll() {
    try {
        const url = "/api/remove-all";
        const options = {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
        }

        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error("Failed to remove all tasks");
        }

        const data = await response.json();
        return data;
    } catch(error) {
        throw error;
    }
}
