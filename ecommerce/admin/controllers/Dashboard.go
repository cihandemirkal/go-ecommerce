package controllers

import (
	"ecommerce/admin/helpers"
	"ecommerce/admin/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Products"] = models.Product{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)

	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()

	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	name := r.FormValue("product-name")
	description := r.FormValue("product-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("product-category"))
	price, _ := strconv.Atoi(r.FormValue("product-price"))
	stock, _ := strconv.Atoi(r.FormValue("product-stock"))

	// UPLOAD

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("product-picture")

	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	// Upload End

	if err != nil {
		fmt.Println(err)
		return
	}

	models.Product{
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
		Price:       price,
		Stock:       stock,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	product := models.Product{}.Get(params.ByName("id"))
	product.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Product"] = models.Product{}.Get(params.ByName("id"))
	data["Categories"] = models.Product{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	product := models.Product{}.Get(params.ByName("id"))

	name := r.FormValue("product-name")
	description := r.FormValue("product-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("product-category"))
	price, _ := strconv.Atoi(r.FormValue("product-price"))
	stock, _ := strconv.Atoi(r.FormValue("product-stock"))
	is_selected := r.FormValue("is_selected")

	var picture_url string

	if is_selected == "1" {
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("product-picture")

		if err != nil {
			fmt.Println(err)
			return
		}

		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
		picture_url = "uploads/" + header.Filename
		os.Remove(product.Picture_url)
	} else {
		picture_url = product.Picture_url
	}

	product.Updates(models.Product{
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
		Price:       price,
		Stock:       stock,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
