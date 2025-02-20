package web

import "commandline-taskmanager/logic"

type ErrorResp struct {
	Error string `json:"error"`
}

type SuccessListResp struct {
	Data []logic.Todo `json:"data"`
}

type NoReturnSuccessResp struct {
	Message string `json:"message"`
}

type SuccessResp struct {
	Data logic.Todo `json:"data"`
}

type UpdateBody struct {
	NewName string `json:"newName"`
}

type NewTaskBody struct {
	Task string `json:"task"`
}
