package main

import (
	"com/onlineCourse/gorilla/mux"
	"com/onlineCourse/metadataService/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/courses").HandlerFunc(handler.GetCourses)
	sub.Methods("POST").Path("/courses").HandlerFunc(handler.SaveCourse)
	sub.Methods("GET").Path("/courses/{name}").HandlerFunc(handler.GetCourse)
	sub.Methods("PUT").Path("/courses/{name}").HandlerFunc(handler.UpdateCourse)
	sub.Methods("DELETE").Path("/courses/{name}").HandlerFunc(handler.DeleteCourse)

	log.Fatal(http.ListenAndServe(":3000", router))
}
