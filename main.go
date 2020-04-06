package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"go-project/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte(""))

	templateDirs = []string{"templates", "templates/partial"}
	templates    *template.Template
)

func getTemplates() (templates *template.Template, err error) {
	var allFiles []string
	for _, dir := range templateDirs {
		files2, _ := ioutil.ReadDir(dir)
		for _, file := range files2 {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				filePath := filepath.Join(dir, filename)
				allFiles = append(allFiles, filePath)
			}
		}
	}

	templates, err = template.New("").ParseFiles(allFiles...)
	return
}

func init() {
	templates, _ = getTemplates()
}

// Main

// GET "/"
func Index(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	if userSession := session.Values["user"]; userSession != nil {
		err := user.GetByUsername(userSession.(string))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data["user"] = user
	} else {
		data["user"] = nil
	}

	var articles []model.Article
	articles, err := model.GetAllArticles()

	data["articles"] = articles

	err = templates.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var err interface{}
	err = nil

	_, ok := r.URL.Query()["error"]
	if ok {
		err = "Username atau password salah."
	}

	errs := templates.ExecuteTemplate(w, "login", err)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
	}
}

// GET "/about"
func About(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	if userSession := session.Values["user"]; userSession != nil {
		err := user.GetByUsername(userSession.(string))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data["user"] = user
	} else {
		data["user"] = nil
	}

	errs := templates.ExecuteTemplate(w, "about", data)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
	}
}

// GET "/contactus"
func ContactUs(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	if userSession := session.Values["user"]; userSession != nil {
		err := user.GetByUsername(userSession.(string))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data["user"] = user
	} else {
		data["user"] = nil
	}

	errs := templates.ExecuteTemplate(w, "contactus", data)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
	}
}

// POST "/send-message"
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var message model.Message
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	if userSession := session.Values["user"]; userSession != nil {
		err := user.GetByUsername(userSession.(string))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data["user"] = user
	} else {
		data["user"] = nil
	}

	message.Name = r.FormValue("name")
	message.Email = r.FormValue("email")
	message.Message = r.FormValue("message")

	errMessage, err := message.Validate()
	if err {
		data["error"] = errMessage
		errs := templates.ExecuteTemplate(w, "contactus", data)
		if errs != nil {
			http.Error(w, errs.Error(), http.StatusInternalServerError)
		}
	}

	if err := message.Create(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/contactus?success=true", 302)
}

// POST "/login"
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user model.User
	session, err := store.Get(r, "blogss")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user.GetByUsername(r.FormValue("username"))

	if user.Auth(r.FormValue("password")) {
		session.Values["user"] = user.Username
		session.Save(r, w)

		if user.IsAdmin {
			http.Redirect(w, r, "/users", 302)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	} else {
		http.Redirect(w, r, "/login?error=auth", 302)
	}
}

// GET "/logout"
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "blogss")

	session.Values["user"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/login", 302)
}

// GET "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// POST "/register"
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	user.Name = r.FormValue("name")
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	errMessage, err := user.Validate()
	if err {
		t, _ := template.ParseFiles("templates/register.html")
		t.Execute(w, errMessage)
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/login?success=true", 302)
}

// Admin

// GET "/users"
func UserList(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	users, _ := model.GetAllUsers()
	data["users"] = users

	err := templates.ExecuteTemplate(w, "users", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET "/users/create"
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	err := templates.ExecuteTemplate(w, "userCreate", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// POST "/users/create"
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	user = model.User{}
	user.Name = r.FormValue("name")
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.IsAdmin, _ = strconv.ParseBool(r.FormValue("is_admin"))

	errMessage, err := user.Validate()
	if err {
		data["validateError"] = errMessage
		err := templates.ExecuteTemplate(w, "userCreate", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/users?success=true", 302)
}

// GET "/users/{id}/edit"
func UserEdit(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	idUser, _ := strconv.Atoi(mux.Vars(r)["id"])
	user.GetUser(idUser)
	data["userInfo"] = user

	err := templates.ExecuteTemplate(w, "userEdit", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// POST "/users/{id}/update"
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	nameVal := r.FormValue("name")
	isAdminVal := r.FormValue("is_admin")

	id := mux.Vars(r)["id"]
	idUser, _ := strconv.Atoi(id)
	user.GetUser(idUser)

	user.Name = nameVal
	user.IsAdmin, _ = strconv.ParseBool(isAdminVal)

	if err := user.Update(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/users?success=true", 302)
}

// POST "/users/{id}/delete"
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	id := mux.Vars(r)["id"]
	idUser, _ := strconv.Atoi(id)
	user.GetUser(idUser)

	if err := user.Delete(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/users?delete=true", 302)
}

func Articles(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	if userSession := session.Values["user"]; userSession != nil {
		err := user.GetByUsername(userSession.(string))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data["user"] = user
	} else {
		data["user"] = nil
	}

	var articles []model.Article
	articles, err := model.GetAllArticlesByUser(user.ID)

	data["articles"] = articles

	err = templates.ExecuteTemplate(w, "articles", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddArticle(w http.ResponseWriter, r *http.Request) {
	var user model.User
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	err := templates.ExecuteTemplate(w, "articleCreate", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	article.User.GetUser(user.ID)

	errMessage, errs := article.Validate()
	if errs {
		data["validateError"] = errMessage
		err := templates.ExecuteTemplate(w, "articleCreate", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	idArticle, err := article.Create()
	if err != nil {
		fmt.Println(err.Error())
	}

	returnURL := fmt.Sprintf("/articles/%+v", idArticle)
	http.Redirect(w, r, returnURL, 302)
}

func ArticleView(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	idArticle, _ := strconv.Atoi(mux.Vars(r)["id"])
	article.GetArticle(idArticle)
	data["article"] = article

	err := templates.ExecuteTemplate(w, "article", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ArticlePublish(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	id := mux.Vars(r)["id"]
	idArticle, _ := strconv.Atoi(id)
	article.GetArticle(idArticle)

	if err := article.Publish(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/articles?publish=ok", 302)
}

func ArticleUnpublish(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	id := mux.Vars(r)["id"]
	idArticle, _ := strconv.Atoi(id)
	article.GetArticle(idArticle)

	if err := article.Unpublish(); err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/articles?unpublish=ok", 302)
}

func ArticleEdit(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	idArticle, _ := strconv.Atoi(mux.Vars(r)["id"])
	article.GetArticle(idArticle)
	data["article"] = article

	err := templates.ExecuteTemplate(w, "articleEdit", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ArticleUpdate(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	idArticle, _ := strconv.Atoi(mux.Vars(r)["id"])
	article.GetArticle(idArticle)
	data["article"] = article

	article.Title = r.FormValue("title")
	article.Content = r.FormValue("content")
	article.User.GetUser(user.ID)

	errMessage, errs := article.Validate()
	if errs {
		data["validateError"] = errMessage
		err := templates.ExecuteTemplate(w, "articleEdit", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err := article.Update()
	if err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/articles", 302)
}

func ArticleDelete(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var article model.Article
	data := map[string]interface{}{}
	session, _ := store.Get(r, "blogss")

	user.GetByUsername(session.Values["user"].(string))
	if !user.IsAdmin {
		http.Redirect(w, r, "/", 302)
		return
	}
	data["user"] = user

	idArticle, _ := strconv.Atoi(mux.Vars(r)["id"])
	article.GetArticle(idArticle)
	data["article"] = article

	err := article.Delete()
	if err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/articles", 302)
}

func main() {
	r := mux.NewRouter()

	// r.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/login", Login).Methods("GET")
	r.HandleFunc("/login", Authenticate).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("GET")
	r.HandleFunc("/register", Register).Methods("GET")
	r.HandleFunc("/register", RegisterUser).Methods("POST")
	r.HandleFunc("/about", About).Methods("GET")
	r.HandleFunc("/contactus", ContactUs).Methods("GET")
	r.HandleFunc("/send-message", SendMessage).Methods("POST")

	r.HandleFunc("/articles", Articles).Methods("GET")
	r.HandleFunc("/articles/create", AddArticle).Methods("GET")
	r.HandleFunc("/articles/create", CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", ArticleView).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ArticleEdit).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}/update", ArticleUpdate).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}/publish", ArticlePublish).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}/unpublish", ArticleUnpublish).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ArticleDelete).Methods("POST")

	r.HandleFunc("/users", UserList).Methods("GET")
	r.HandleFunc("/users/create", AddUser).Methods("GET")
	r.HandleFunc("/users/create", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}/edit", UserEdit).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/update", UserUpdate).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}/delete", DeleteUser).Methods("POST")

	log.Println("Listening...")
	http.ListenAndServe(":3400", r)
}
