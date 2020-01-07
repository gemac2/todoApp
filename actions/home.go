package actions

import (
	"net/http"

	"github.com/gemac2/todo/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// Todo ...
func Todo(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	todoes := models.Todoes{}
	tx.All(&todoes)

	c.Set("todoes", todoes)
	return c.Render(200, r.HTML("index.plush.html"))
}

// New ...
func New(c buffalo.Context) error {
	todo := models.Todo{}
	c.Set("todo", todo)
	return c.Render(200, r.HTML("new.plush.html"))
}

// Create ...
func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	todo := models.Todo{}

	c.Bind(&todo)

	tx.Create(&todo)

	c.Flash().Add("success", "Todo was created")
	return c.Redirect(http.StatusSeeOther, "/todo")
}

// Show ...
func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	todo := models.Todo{
		Title: string("primero"),
	}

	tx.First(todo)

	c.Set("todo", todo)
	return c.Render(200, r.HTML("show.plush.html"))
}
