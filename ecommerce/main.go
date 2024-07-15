package main

import (
	admin_models "ecommerce/admin/models"
	"ecommerce/config"
	"net/http"
)

func main() {
	admin_models.Product{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	/*
		admin_models.Product{
			Name:        "Kareli gömlek",
			Description: "Keten kareli gömlek",
		}.Add()
	*/
	/*
		product := admin_models.Product{}.Get("description = ?", "Keten kareli gömlek")
		fmt.Println(product.Name)
	*/
	/*
		fmt.Println(admin_models.Product{}.GetAll("description = ?", "Keten kareli gömlek"))
	*/

	/*
		product := admin_models.Product{}.Get(1)

		product.Updates(admin_models.Product{Name: "Oduncu Gömlek", Description: "Sade renkli oduncu gömlek"})
	*/
	/*
		product.Update("stock", 15)
	*/
	/*
		product := admin_models.Product{}.Get(2)
		product.Delete()
	*/
	http.ListenAndServe(":8080", config.Routes())
}
