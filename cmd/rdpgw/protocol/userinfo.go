package protocol

import (
	"fmt"
	"net/http"
	"time"
)

var scientists = []string{
        "Einstein",
        "Turing",
        "Lovelace",
        "Curie",
        "Franklin",
        "Hodgkin",
    }

type Identity interface {
        UserName() string
        SetUserName(string)
        DisplayName() string
        SetDisplayName(string)
        Domain() string
        SetDomain(string)
        Authenticated() bool
        SetAuthenticated(bool)
        AuthTime() time.Time
        SetAuthTime(time2 time.Time)
        SessionId() string
        SetAttribute(string, interface{})
        GetAttribute(string) interface{}
        Attributes() map[string]interface{}
        DelAttribute(string)
        Email() string
        SetEmail(string)
        Expiry() time.Time
        SetExpiry(time.Time)
        Marshal() ([]byte, error)
        Unmarshal([]byte) error
}


var sessionUser = map[string]string {
}

var targetserver = map[string]string{
        }

func addSessionUser(sessionid string,uid string) {

        sessionUser[sessionid]=uid

}


func delSessionUser(sessionid string) {

                //scientists = append(scientists,activeuser)
                delete(sessionUser,sessionid)

}



func addSessionTS(sessionid string,server string) {

	targetserver[sessionid]=server

}


func delSessionTS(sessionid string) {

	        //scientists = append(scientists,activeuser)
	        delete(targetserver,sessionid)

}


func UserInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	fmt.Fprint(w, scientists)
	fmt.Fprint(w, targetserver)
	fmt.Fprint(w, sessionUser)
}
