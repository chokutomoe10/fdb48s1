package main

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"
	"routing/connection"
	"routing/middleware"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	NodeJs      string
	Javascript  string
	ReactJs     string
	SocketIO    string
	Image       string
	Author      string
}

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
}

type UserIsLogin struct {
	IsLogin bool
	Name    string
}

var userIsLogin = UserIsLogin{}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("welldone"))))

	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/addmyproject", addmyproject)
	e.GET("/myproject", myproject)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-project", middleware.UploadFile(addProject))
	e.POST("/delete-project/:id", deleteProject)

	e.GET("/login", login)
	e.POST("/form-login", formLogin)

	e.GET("/register", register)
	e.POST("/form-register", formRegister)

	e.GET("/update-project-form/:id", updateProjectForm)
	e.POST("/update-project", middleware.UploadFile(updateProject))

	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	dataProjects, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_users.user_name, tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.image FROM tb_projects INNER JOIN tb_users ON tb_projects.author_id = tb_users.id")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project

	for dataProjects.Next() {
		var each = Project{}

		var tempAuthor sql.NullString

		err := dataProjects.Scan(&tempAuthor, &each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Author = tempAuthor.String

		each.NodeJs = ""
		each.Javascript = ""
		each.ReactJs = ""
		each.SocketIO = ""

		resultProjects = append(resultProjects, each)
	}

	if sess.Values["isLogin"] != true {
		userIsLogin.IsLogin = false
	} else {
		userIsLogin.IsLogin = true
		userIsLogin.Name = sess.Values["name"].(string)
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"],
		"FlashStatus":  sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	data := map[string]interface{}{
		"Projects":    resultProjects,
		"Flash":       flash,
		"UserIsLogin": userIsLogin,
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

	dataProjects, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_users.user_name, tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.image FROM tb_projects LEFT JOIN tb_users ON tb_projects.author_id = tb_users.id")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProjects []Project
	for dataProjects.Next() {
		var each = Project{}

		var tempAuthor sql.NullString

		err := dataProjects.Scan(&tempAuthor, &each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Author = tempAuthor.String

		each.NodeJs = ""
		each.Javascript = ""
		each.ReactJs = ""
		each.SocketIO = ""

		resultProjects = append(resultProjects, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userIsLogin.IsLogin = false
	} else {
		userIsLogin.IsLogin = true
		userIsLogin.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Projects":    resultProjects,
		"UserIsLogin": userIsLogin,
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

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	blogDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(),
		"SELECT tb_users.user_name, tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.image FROM tb_projects INNER JOIN tb_users ON tb_projects.author_id = tb_users.id WHERE tb_projects.id=$1", idToInt).Scan(&blogDetail.Author, &blogDetail.Id, &blogDetail.ProjectName, &blogDetail.StartDate, &blogDetail.EndDate, &blogDetail.Description, &blogDetail.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	if sess.Values["isLogin"] != true {
		userIsLogin.IsLogin = false
	} else {
		userIsLogin.IsLogin = true
		userIsLogin.Name = sess.Values["name"].(string)
	}

	// for index, data := range dataProjects {
	// 	if index == idToInt {
	// 		blogDetail = Project{
	// 			ProjectName: data.ProjectName,
	// 			StartDate:   data.StartDate,
	// 			EndDate:     data.EndDate,
	// 			Description: data.Description,
	// 			NodeJs:      data.NodeJs,
	// 			Javascript:  data.Javascript,
	// 			ReactJs:     data.ReactJs,
	// 			SocketIO:    data.SocketIO,
	// 		}
	// 	}
	// }

	data := map[string]interface{}{
		"Id":          id,
		"Project":     blogDetail,
		"UserIsLogin": userIsLogin,
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectname := c.FormValue("projectname")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	description := c.FormValue("description")
	// nodejs := c.FormValue("nodejs")
	// javascript := c.FormValue("javascript")
	// reactjs := c.FormValue("reactjs")
	// socketio := c.FormValue("socketio")
	// image := c.FormValue("uploadimage")
	image := c.Get("dataFile").(string)

	// newProject := Project{
	// 	ProjectName: projectname,
	// 	// StartDate:   startdate,
	// 	// EndDate:     enddate,
	// 	Description: description,
	// 	NodeJs:      nodejs,
	// 	Javascript:  javascript,
	// 	ReactJs:     reactjs,
	// 	SocketIO:    socketio,
	// }

	sess, _ := session.Get("session", c)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, image, author_id) VALUES ($1, $2, $3, $4, $5, $6)", projectname, startdate, enddate, description, image, sess.Values["id"].(int))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idtoInt, _ := strconv.Atoi(id)

	// dataProjects = append(dataProjects[:idtoInt], dataProjects[idtoInt+1:]...)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", idtoInt)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func updateProjectForm(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("html/updatemyproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idtoInt, _ := strconv.Atoi(id)

	updatedProject := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(),
		"SELECT id, name, start_date, end_date, description, image FROM tb_projects WHERE id=$1", idtoInt).Scan(&updatedProject.Id, &updatedProject.ProjectName, &updatedProject.StartDate, &updatedProject.EndDate, &updatedProject.Description, &updatedProject.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": updatedProject,
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProject(c echo.Context) error {
	id := c.FormValue("id")
	projectname := c.FormValue("projectname")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	description := c.FormValue("description")
	image := c.Get("dataFile").(string)
	// image := c.FormValue("uploadimage")

	// fmt.Println("id", id)
	// fmt.Println("proname", projectname)
	// fmt.Println("desc", description)
	// fmt.Println("img", image)

	// sess, _ := session.Get("session", c)

	idtoInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, image=$5 WHERE id=$6", projectname, startdate, enddate, description, image, idtoInt)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func login(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"],
		"FlashStatus":  sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func formLogin(c echo.Context) error {
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword")

	user := User{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, user_name, email, password FROM tb_users WHERE email=$1", inputEmail).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/login")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(inputPassword))

	if errPassword != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM -> berapa lama expired
	sess.Values["message"] = "Login berhasil!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/home")
}

func register(c echo.Context) error {
	tmpl, err := template.ParseFiles("html/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"],
		"FlashStatus":  sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func formRegister(c echo.Context) error {
	inputName := c.FormValue("inputName")
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, errQuery := connection.Conn.Exec(context.Background(), "INSERT INTO tb_users (user_name, email, password) VALUES ($1, $2, $3)", inputName, inputEmail, hashedPassword)

	if errQuery != nil {
		return redirectWithMessage(c, "Register gagal!", false, "/register")
	}

	return redirectWithMessage(c, "Register berhasil!", true, "/login")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Logout berhasil!", true, "/login")
}

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}
