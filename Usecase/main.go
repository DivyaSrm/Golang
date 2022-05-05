package main

import (
	"Usecase/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//category route
	s.HandleFunc("/createCategory", app.CreateCategory).Methods("POST")
	s.HandleFunc("/getAllCategory", app.GetAllCategory).Methods("GET")
	s.HandleFunc("/getCategory", app.GetCategory).Methods("POST")
	s.HandleFunc("/updateCategory", app.UpdateCategory).Methods("PUT")
	s.HandleFunc("/deleteCategory/{id}", app.DeleteCategory).Methods("DELETE")

	//subcategory routes
	s.HandleFunc("/createSubCategory", app.CreateSubCategory).Methods("POST")
	s.HandleFunc("/getAllSubCategory", app.GetAllSubCategory).Methods("GET")
	s.HandleFunc("/getSub", app.GetSubCategory).Methods("POST")
	s.HandleFunc("/updatesubCategory", app.UpdateSubCategory).Methods("PUT")
	s.HandleFunc("/deleteSubCategory/{id}", app.DeleteSubCategory).Methods("DELETE")

	//brand routes
	s.HandleFunc("/createBrand", app.CreateBrand).Methods("POST")
	s.HandleFunc("/getAllBrand", app.GetAllBrand).Methods("GET")
	s.HandleFunc("/getBrand", app.GetBrand).Methods("POST")
	s.HandleFunc("/updateBrand", app.UpdateBrand).Methods("PUT")
	s.HandleFunc("/deleteBrand/{id}", app.DeleteBrand).Methods("DELETE")

	//variant routes
	s.HandleFunc("/createVariant", app.CreateVariant).Methods("POST")
	s.HandleFunc("/getAllVariant", app.GetAllVariant).Methods("GET")
	s.HandleFunc("/getVariant", app.GetVariant).Methods("POST")
	s.HandleFunc("/updateVariant", app.UpdateVariant).Methods("PUT")
	s.HandleFunc("/deleteVariant/{id}", app.DeleteVariant).Methods("DELETE")

	//product routes
	s.HandleFunc("/createProduct", app.CreateProduct).Methods("POST")
	s.HandleFunc("/getAllProducts", app.GetAllProducts).Methods("GET")
	s.HandleFunc("/getProduct", app.GetProduct).Methods("POST")
	s.HandleFunc("/updateProduct", app.UpdateProduct).Methods("PUT")
	s.HandleFunc("/deleteProduct{id}", app.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
