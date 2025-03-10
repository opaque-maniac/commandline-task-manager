package web

import (
	"commandline-taskmanager/logic"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getTaskName(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

// send Error
func sendError(w http.ResponseWriter, err error, statusCode int) {
	fmt.Printf("Error: %s\n", err.Error())
	resp := ErrorResp{Error: err.Error()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func sendGenericResponse(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(NoReturnSuccessResp{Message: data}); err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}
}

// handle home
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// submit handler
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	fmt.Printf("Request method: %s\n", r.Method)

	var body NewTaskBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	if err := logic.AddTask(body.Task); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	sendGenericResponse(w, "Task added")
}

// handle data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := logic.ReadData()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	fmt.Printf("Request method: %s\n", r.Method)
	resp := SuccessListResp{Data: todos}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}
}

// update handler
func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	taskName := getTaskName(r.URL.Path, "/api/update/")

	var body UpdateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Printf("Request method: %s\n", r.Method)

	if err := logic.UpdateTask(taskName, body.NewName); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	sendGenericResponse(w, "Task updated")
}

// remove handler
func removeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	taskName := getTaskName(r.URL.Path, "/api/remove/")
	if err := logic.RemoveTask(taskName); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Printf("Request method: %s\n", r.Method)

	sendGenericResponse(w, "Task removed")
}

// complete handler
func completeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	taskName := getTaskName(r.URL.Path, "/api/complete/")

	list, err := logic.ReadData()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	var task logic.Todo

	for _, v := range list {
		if v.Task == taskName {
			task = v
		}
	}

	if task.Task == "" {
		sendGenericResponse(w, "Task not found")
		return
	}

	if task.Completed {
		err = logic.UnCompleteTask(taskName)
	} else {
		err = logic.CompleteTask(taskName)
	}

	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Printf("Request method: %s\n", r.Method)

	sendGenericResponse(w, "Task completed")
}

// remove all
func removeAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	if err := logic.RemoveAllTasks(); err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Request method: %s\n", r.Method)

	sendGenericResponse(w, "All tasks removed")
}

func Start(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/api/data", dataHandler)
	mux.HandleFunc("/api/submit", submitHandler)
	mux.HandleFunc("/api/update/", updateHandler)
	mux.HandleFunc("/api/remove/", removeHandler)
	mux.HandleFunc("/api/complete/", completeHandler)
	mux.HandleFunc("/api/remove-all", removeAllHandler)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	fmt.Println("Server started at http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
