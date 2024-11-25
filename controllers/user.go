package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/princetomar27/mogno-golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{session:s} 
}


func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){ 
		w.WriteHeader(http.StatusNotFound)
	}
	
	oid := bson.ObjectIdHex(id)
	user := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&user);
	err != nil{
		w.WriteHeader(404)
		return
	}

	userObj, err := json.Marshal(user)
	if err!= nil{
        w.WriteHeader(500)
		fmt.Println(err)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userObj)
}

// CreateUser

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request,_  httprouter.Params){
	// 1. Create user object
	user := models.User{}
	// 2. decode the user json to User object
	json.NewDecoder(r.Body).Decode(&user)

	//3. Assign the id to the user object using bson
	user.Id = bson.NewObjectId()

	// 4. Save the user to the database
	uc.session.DB("mongo-golang").C("users").Insert(user)

	// 5. Return the user object with status code 201
	userObj, error := json.Marshal(user)

	if error != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userObj)
}

// DeleteUser