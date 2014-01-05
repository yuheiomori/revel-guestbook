package controllers

import (
	"github.com/robfig/revel"
	"revel-guestbook/app/models"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var greetings []models.Greeting
	_, err := Dbm.Select(&greetings, "select * from greetings order by CreateAt desc")
	if err != nil {
		panic(err)
	}
	return c.Render(greetings)
}

func (c App) Post(name, comment string) revel.Result {

	greeting := models.Greeting{
		Name:     name,
		Comment:  comment,
		CreateAt: time.Now().UnixNano(),
	}
	Dbm.Insert(&greeting)
	return c.Redirect("/")
}
