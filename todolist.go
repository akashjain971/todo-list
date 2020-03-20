package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var todolist = make(map[string]bool)

type response struct {
	Complete []string
	Pending  []string
}

func returnTodoList() string {
	complete := make([]string, 0)
	pending := make([]string, 0)
	for item, todo := range todolist {
		if todo {
			pending = append(pending, item)
		} else {
			complete = append(complete, item)
		}
	}

	resp := response{
		Complete: complete,
		Pending:  pending,
	}
	jsonresp, err := json.Marshal(resp)
	if err != nil {
		return "[ERROR] Failed to parse the list"
	}
	return string(jsonresp)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	task := ps.ByName("task")
	if task == "" {
		// return status of all items
		fmt.Fprintf(w, returnTodoList())
		return
	}

	// return status of specific item
	if status, ok := todolist[task]; ok {
		if status {
			fmt.Fprintf(w, "[SUCCESS] Task is pending")
		} else {
			fmt.Fprintf(w, "[SUCCESS] Task is complete")
		}
	} else {
		fmt.Fprintf(w, "[ERROR] Task not present in the list")
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	task := ps.ByName("task")

	if _, ok := todolist[task]; ok {
		fmt.Fprintf(w, "[ERROR] Task already in the list")
	} else {
		// add new item and mark as pending
		todolist[task] = true
		fmt.Fprintf(w, "[SUCCESS] Task added to the list")
	}
}

func handlePatchRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	task := ps.ByName("task")

	if status, ok := todolist[task]; ok {
		// flip the status of the item
		todolist[task] = !status
		fmt.Fprintf(w, "[SUCCESS] Task status has been changed")
	} else {
		fmt.Fprintf(w, "[ERROR] Task not present in the list")
	}
}

func handleDeleteRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	task := ps.ByName("task")

	if task == "" {
		// delete all items
		todolist = make(map[string]bool)
		fmt.Fprintf(w, "[SUCCESS] Task list has been cleared")
		return
	}

	// delete specific item
	if _, ok := todolist[task]; ok {
		delete(todolist, task)
		fmt.Fprintf(w, "[SUCCESS] Task deleted from the list")
	} else {
		fmt.Fprintf(w, "[ERROR] Task not present in the list")
	}
}
