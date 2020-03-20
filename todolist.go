package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func removeslash(path string) string {
	if path[0] == '/' {
		path = path[1:]
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}

func handler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(removeslash(r.URL.Path), "/")
	if len(paths) > 2 {
		fmt.Fprintf(w, "[ERROR] Too many parameters")
		return
	}

	switch r.Method {
	case http.MethodGet:

		if len(paths) == 1 {
			// return status of all items
			fmt.Fprintf(w, returnTodoList())
			return
		}

		// return status of specific item
		if status, ok := todolist[paths[1]]; ok {
			if status {
				fmt.Fprintf(w, "[SUCCESS] Task is pending")
			} else {
				fmt.Fprintf(w, "[SUCCESS] Task is complete")
			}
		} else {
			fmt.Fprintf(w, "[ERROR] Task not present in the list")
		}

	case http.MethodPost:

		if len(paths) == 1 {
			fmt.Fprintf(w, "[ERROR] Task cannot be empty")
			return
		}

		if _, ok := todolist[paths[1]]; ok {
			fmt.Fprintf(w, "[ERROR] Task already in the list")
		} else {
			// add new item and mark as pending
			todolist[string(paths[1])] = true
			fmt.Fprintf(w, "[SUCCESS] Task added to the list")
		}

	case http.MethodPatch:

		if len(paths) == 1 {
			fmt.Fprintf(w, "[ERROR] Task cannot be empty")
			return
		}

		if status, ok := todolist[paths[1]]; ok {
			// flip the status of the item
			todolist[string(paths[1])] = !status
			fmt.Fprintf(w, "[SUCCESS] Task status has been changes")
		} else {
			fmt.Fprintf(w, "[ERROR] Task not present in the list")
		}

	case http.MethodDelete:

		if len(paths) == 1 {
			// delete all items
			todolist = make(map[string]bool)
			fmt.Fprintf(w, "[SUCCESS] Task list has been cleared")
			return
		}

		// delete specific item
		if _, ok := todolist[paths[1]]; ok {
			delete(todolist, paths[1])
			fmt.Fprintf(w, "[SUCCESS] Task deleted from the list")
		} else {
			fmt.Fprintf(w, "[ERROR] Task not present in the list")
		}

	default:
		fmt.Fprintf(w, "[ERROR] HTTP method not supported")
	}
}
