package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/myproject", myproject)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-blog", addBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func myproject(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/myproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("html/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogDetail := map[string]interface{}{
		"Id":      id,
		"Title":   "Dumbways ID Batch 48 Stage 1",
		"Content": "Fullstack Developer",
	}

	return tmpl.Execute(c.Response(), blogDetail)
}

func addBlog(c echo.Context) error {
	projectname := c.FormValue("projectname")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	nextjs := c.FormValue("nextjs")
	reactjs := c.FormValue("reactjs")
	typescript := c.FormValue("typescript")

	fmt.Println("Title: ", projectname)
	fmt.Println("Start Date: ", startdate)
	fmt.Println("End Date: ", enddate)
	fmt.Println("Description: ", description)
	fmt.Println("Technologies: ", nodejs)
	fmt.Println("Technologies: ", nextjs)
	fmt.Println("Technologies: ", reactjs)
	fmt.Println("Technologies: ", typescript)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}
