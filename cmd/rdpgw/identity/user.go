package identity

import (
	"bytes"
	"encoding/gob"
	"github.com/google/uuid"
	"time"
	"log"
)

type User struct {
	authenticated   bool
	domain          string
	userName        string
	displayName     string
	email           string
	authTime        time.Time
	sessionId       string
	expiry          time.Time
	attributes      map[string]interface{}
	groupMembership map[string]bool
	userGroups      []interface{}
	userAcceptedHosts      []string
}

type user struct {
	Authenticated   bool
	UserName        string
	Domain          string
	DisplayName     string
	Email           string
	AuthTime        time.Time
	SessionId       string
	Expiry          time.Time
	Attributes      map[string]interface{}
	GroupMembership map[string]bool
	UserGroups      []interface{}
	UserAcceptedHosts      []string
}

func NewUser() *User {
	uuid := uuid.New().String()
	return &User{
		attributes:      make(map[string]interface{}),
		groupMembership: make(map[string]bool),
		sessionId:       uuid,
	}
}

func (u *User) UserName() string {
	return u.userName
}



func (u *User) SetUserName(s string) {
	u.userName = s
}



func (u *User) DisplayName() string {
	if u.displayName == "" {
		return u.userName
	}
	return u.displayName
}

func (u *User) SetDisplayName(s string) {
	u.displayName = s
}

func (u *User) Domain() string {
	return u.domain
}

func (u *User) SetDomain(s string) {
	u.domain = s
}

func (u *User) Authenticated() bool {
	return u.authenticated
}

func (u *User) SetAuthenticated(b bool) {
	u.authenticated = b
}

func (u *User) AuthTime() time.Time {
	return u.authTime
}

func (u *User) SetAuthTime(t time.Time) {
	u.authTime = t
}

func (u *User) SessionId() string {
	return u.sessionId
}

func (u *User) SetAttribute(s string, i interface{}) {
	u.attributes[s] = i
}



func (u *User) SetAcceptedHost(s []string) {
        u.userAcceptedHosts = s
        log.Printf("SET ACCEPTED HOSTS  %s by policy handler", u.userAcceptedHosts)
}

func (u *User) AcceptedHosts() []string  {
	log.Printf("SAVED HOSTS  %s by policy handler", u.userAcceptedHosts)
        return u.userAcceptedHosts
}


func (u *User) SetGroups(s []interface{}) {
        u.userGroups = s
	log.Printf("SET GROUPS  %s by policy handler", s)
}

func (u *User) Groups() []interface{}  {
        return u.userGroups
}



func (u *User) GetAttribute(s string) interface{} {
	if found, ok := u.attributes[s]; ok {
		return found
	}
	return nil
}



func (u *User) Attributes() map[string]interface{} {
	return u.attributes
}



func (u *User) DelAttribute(s string) {
	delete(u.attributes, s)
}




func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(s string) {
	u.email = s
}

func (u *User) Expiry() time.Time {
	return u.expiry
}

func (u *User) SetExpiry(t time.Time) {
	u.expiry = t
}

func (u *User) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	uu := user{
		Authenticated:   u.authenticated,
		UserName:        u.userName,
		Domain:          u.domain,
		DisplayName:     u.displayName,
		Email:           u.email,
		AuthTime:        u.authTime,
		SessionId:       u.sessionId,
		Expiry:          u.expiry,
		Attributes:      u.attributes,
		GroupMembership: u.groupMembership,
		UserGroups:      u.userGroups,
		UserAcceptedHosts:      u.userAcceptedHosts,
	}
	err := enc.Encode(uu)

	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func (u *User) Unmarshal(b []byte) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var uu user
	err := dec.Decode(&uu)
	if err != nil {
		return err
	}
	u.sessionId = uu.SessionId
	u.userName = uu.UserName
	u.domain = uu.Domain
	u.displayName = uu.DisplayName
	u.email = uu.Email
	u.authenticated = uu.Authenticated
	u.authTime = uu.AuthTime
	u.expiry = uu.Expiry
	u.attributes = uu.Attributes
	u.userGroups= uu.UserGroups
	u.userAcceptedHosts= uu.UserAcceptedHosts

	return nil
}
