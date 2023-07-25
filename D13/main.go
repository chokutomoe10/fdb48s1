package main

import (
	"context"
	"html/template"
	"net/http"
	"routing/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	NodeJs      string
	Javascript  string
	ReactJs     string
	SocketIO    string
}

var dataProjects = []Project{
	// {
	// 	ProjectName: "Project 1",
	// 	StartDate:   "20/07/2022",
	// 	EndDate:     "20/08/2022",
	// 	Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sed tristique purus orci libero egestas ut. Augue suspendisse blandit lorem massa ipsum, urna egestas mi, lacinia. Dui fusce etiam libero, lectus amet, risus molestie malesuada. Odio nam purus consectetur euismod congue leo quisque. Turpis amet sollicitudin nunc non in lectus dolor amet. Amet ullamcorper faucibus tincidunt accumsan ac adipiscing arcu. Donec tristique at proin maeceneas. Ante elit et iaculis ac sit",
	// 	NodeJs:      "Node Js",
	// 	Javascript:  "",
	// 	ReactJs:     "",
	// 	SocketIO:    "",
	// },
	// {
	// 	ProjectName: "Project 2",
	// 	StartDate:   "02/05/2022",
	// 	EndDate:     "27/05/2022",
	// 	Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sed tristique purus orci libero egestas ut. Augue suspendisse blandit lorem massa ipsum, urna egestas mi, lacinia. Dui fusce etiam libero, lectus amet, risus molestie malesuada. Odio nam purus consectetur euismod congue leo quisque. Turpis amet sollicitudin nunc non in lectus dolor amet. Amet ullamcorper faucibus tincidunt accumsan ac adipiscing arcu. Donec tristique at proin maeceneas. Ante elit et iaculis ac sit",
	// 	NodeJs:      "",
	// 	Javascript:  "Javascript",
	// 	ReactJs:     "React Js",
	// 	SocketIO:    "",
	// },
}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Static("/assets", "assets")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/addmyproject", addmyproject)
	e.GET("/myproject", myproject)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-project", addProject)
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errQuery := connection.Conn.Query(context.Background(), "SELECT name, start_date, end_date, description FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.ProjectName, &each.StartDate, &each.EndDate, &each.Description)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.NodeJs = ""
		each.Javascript = ""
		each.ReactJs = ""
		each.SocketIO = ""

		resultProjects = append(resultProjects, each)
	}

	data := map[string]interface{}{
		"Projects": resultProjects,
	}

	return tmpl.Execute(c.Response(), data)
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

	data := map[string]interface{}{
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(), data)
}

func addmyproject(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/addmyproject.html")

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

	idToInt, _ := strconv.Atoi(id)

	blogDetail := Project{}

	for index, data := range dataProjects {
		if index == idToInt {
			blogDetail = Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				NodeJs:      data.NodeJs,
				Javascript:  data.Javascript,
				ReactJs:     data.ReactJs,
				SocketIO:    data.SocketIO,
			}
		}
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": blogDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectname := c.FormValue("projectname")
	// startdate := c.FormValue("startdate")
	// enddate := c.FormValue("enddate")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	javascript := c.FormValue("javascript")
	reactjs := c.FormValue("reactjs")
	socketio := c.FormValue("socketio")

	newProject := Project{
		ProjectName: projectname,
		// StartDate:   startdate,
		// EndDate:     enddate,
		Description: description,
		NodeJs:      nodejs,
		Javascript:  javascript,
		ReactJs:     reactjs,
		SocketIO:    socketio,
	}

	dataProjects = append(dataProjects, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idtoInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idtoInt], dataProjects[idtoInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}
