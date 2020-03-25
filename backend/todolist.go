package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/xid"
)

type todo struct {
	ID       string
	Task     string
	Complete bool
}

var todolist []todo

func sendGetResponse(w http.ResponseWriter, response []todo) {

	jsonresp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if string(jsonresp) == "null" {
		fmt.Fprintf(w, "[]")
	} else {
		fmt.Fprintf(w, string(jsonresp))
	}
}

func handleGetRequestEx(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uuid := ps.ByName("uuid")

	for i := range todolist {
		if todolist[i].ID == uuid {
			sendGetResponse(w, append([]todo{}, todolist[i]))
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sendGetResponse(w, todolist)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "", http.StatusNoContent)
		return
	}

	uuid := xid.New().String()
	task := todo{
		ID:       uuid,
		Task:     string(body),
		Complete: false,
	}

	todolist = append(todolist, task)

	fmt.Fprintf(w, uuid)
}

func handlePatchRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uuid := ps.ByName("uuid")

	for i := range todolist {
		if todolist[i].ID == uuid {
			todolist[i].Complete = !todolist[i].Complete
			fmt.Fprintf(w, http.StatusText(http.StatusOK))
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func handleDeleteRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todolist = nil
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}

func handleDeleteRequestEx(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uuid := ps.ByName("uuid")

	for i := range todolist {
		if todolist[i].ID == uuid {
			todolist[i] = todolist[len(todolist)-1]
			todolist = todolist[:len(todolist)-1]
			fmt.Fprintf(w, http.StatusText(http.StatusOK))
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
