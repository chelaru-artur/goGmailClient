package main

import (
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"net/http"
	_ "strconv"

	_ "github.com/gorilla/context"
	_ "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	fmt.Println("yo")
}
func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	id := session.Values["id"]
	fmt.Fprint(w, "Welcome!\n ", id)
}

func TestApi(w http.ResponseWriter, r *http.Request) {
	dbSess := GetMongoSession()
	session, _ := store.Get(r, "session-name")
	id := session.Values["id"].(string)
	oid := bson.ObjectIdHex(id)
	data := UserToken{}
	if err := dbSess.DB("").C("gmailTokens").FindId(oid).One(&data); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(data.TokenObj.Expiry)
	w.Write([]byte("Test complete!"))
}

func GetAuthUrl(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()
	authUrl := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	w.Write([]byte(authUrl))
}

func GetCtx(w http.ResponseWriter, r *http.Request) {
	//	config := context.Get(r, "config")
	w.Write([]byte("geted"))
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	config := GetConfig()
	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web %v", err)
	}
	fmt.Println(tok)
	mongo := GetMongoSession()
	newInsert := UserToken{
		Id:       bson.NewObjectId(),
		TokenObj: tok,
	}
	err = mongo.DB("").C("gmailTokens").Insert(newInsert)
	fmt.Println(err)
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Set some session values.
	session.Values["id"] = newInsert.Id.Hex()
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
	w.Write([]byte(r.FormValue("code")))
}
