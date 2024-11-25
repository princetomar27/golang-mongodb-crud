package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/princetomar27/mogno-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main(){
	router := httprouter.New()

	userController := controllers.NewUserController(getSession())
	// router.GET("/users", userController.GetAllUsers)
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("localhost:8080",router)

}

func getSession() *mgo.Session{
	session, err := mgo.Dial("mgon")

	if err != nil{
		panic(err)
	}
	return session
}