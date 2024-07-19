package api

import (
	"context"
	"encoding/json"
	"fmt"
	"main/todo"
	"net/http"
	"sync"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

var todoList todo.List
var mutex sync.Mutex

func Serve(port string) {
	todoList = todo.NewList()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Route("/api/items", func(r chi.Router) {
		r.Use(itemListCtx)
		r.Get("/", listItems)
		r.Post("/", createItem)
	})

	router.Route("/api/item/{id}", func(r chi.Router) {
		r.Use(itemCtx)
		r.Put("/", updateItem)
		r.Delete("/", deleteItem)
	})

	router.Get("/", getIndex)
	router.Get("/items/create", getCreate)
	router.Route("/item/{id}", func(r chi.Router) {
		r.Use(itemCtx)
		r.Get("/", getEdit)
	})

	http.ListenAndServe(":8080", router)
}

func setList(list todo.List) {
	mutex.Lock()
	defer mutex.Unlock()

	todoList = list
}

func getList() todo.List {
	mutex.Lock()
	defer mutex.Unlock()

	return todoList
}

func itemListCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "todoList", getList())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func itemCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		list := getList()
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		idx := todo.GetIdx(list, id)
		if idx == -1 {
			fmt.Printf("unable to find item with id %s\n", id)
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "item", todo.GetItems(list)[idx])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item todo.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, item, err := todo.Add(todoList, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setList(list)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list := todo.GetItems(getList())

	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value("item").(todo.Item)

	var payloadItem todo.Item
	if err := json.NewDecoder(r.Body).Decode(&payloadItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payloadItem.ID = item.ID

	list, err := todo.Update(todoList, payloadItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Setting item")
		setList(list)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value("item").(todo.Item)

	list, err := todo.Delete(todoList, item.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		setList(list)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	items := todo.GetItems(getList())

	t, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		fmt.Printf("Internal server error when loading template: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getEdit(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value("item").(todo.Item)

	t, err := template.ParseFiles("templates/base.html", "templates/edit.html")
	if err != nil {
		fmt.Printf("Internal server error when loading template: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getCreate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/create.html")
	if err != nil {
		fmt.Printf("Internal server error when loading template: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
