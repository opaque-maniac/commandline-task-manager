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
	http.ServeFile(w, r, "static/index.html")
}

// submit handler
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

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

	resp := SuccessListResp{Data: todos}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}
}

// update handler
func updateHandler(w http.ResponseWriter, r *http.Request) {
	taskName := getTaskName(r.URL.Path, "/update/")

	if r.Method != http.MethodPut {
		sendError(w, fmt.Errorf("Invalid request method"), http.StatusBadRequest)
		return
	}

	var body UpdateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	if err := logic.UpdateTask(taskName, body.NewName); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
}

// remove handler
func removeHandler(w http.ResponseWriter, r *http.Request) {
	taskName := getTaskName(r.URL.Path, "/remove/")
	if err := logic.RemoveTask(taskName); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	sendGenericResponse(w, "Task removed")
}

// complete handler
func completeHandler(w http.ResponseWriter, r *http.Request) {
	taskName := getTaskName(r.URL.Path, "/complete/")

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

	sendGenericResponse(w, "Task completed")
}

// remove all
func removeAllHandler(w http.ResponseWriter, r *http.Request) {
	if err := logic.RemoveAllTasks(); err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(NoReturnSuccessResp{Message: "All tasks removed"}); err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/submit", submitHandler)
	mux.HandleFunc("/update/", updateHandler)
	mux.HandleFunc("/remove/", removeHandler)
	mux.HandleFunc("/complete/", completeHandler)
	mux.HandleFunc("/remove-all", removeAllHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

    fmt.Println("Server started at http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
