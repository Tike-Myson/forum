package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func CreateUsersTable(db *sql.DB) {
	userTable, err := db.Prepare(createUsersTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = userTable.Exec()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Users table created in DB")
}

func insertUserIntoDB(db *sql.DB, user User) {
	insertUser, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertUser.Exec(
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("New user inserted into DB")
}

func InsertUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Println(err)
	}

	user.CreatedAt = time.Now()
	user.Password = string(hashedPassword)
	// Check for duplicate parameters in DB

	insertUserIntoDB(db, user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New user sent to clientside")
}

//func authenticateUser(email string, password string) int {
//	var id int
//	var hashedPassword string
//
//	row := db.QueryRow(getUserSQL, email)
//	err := row.Scan(&id, &hashedPassword)
//	if err != nil {
//		fmt.Println("Error occured while scanning through db", err)
//		return 0
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
//	if err != nil {
//		fmt.Println("Invalid credentials")
//		return 0
//	}
//	return id
//}
//
//func LoginUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "POST")
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//
//	var user models.User
//	err = json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		fmt.Println(err)
//		return
//	}
//
//	// cookie, err := r.Cookie(cookieName)
//	// if err != nil {
//	// 	if err == http.ErrNoCookie {
//	// 		w.WriteHeader(http.StatusUnauthorized)
//	// 		return
//	// 	}
//	// 	w.WriteHeader(http.StatusBadRequest)
//	// 	return
//	// }
//	// sessionToken = cookie.Value
//
//
//}
//
//func setCookie(w http.ResponseWriter) {
//	sessionToken := uuid.NewV4().String()
//
//	cookie := &http.Cookie{
//		Name: cookieName,
//		Value: sessionToken,
//		Expires: time.Now().Add(120 * time.Second),
//	}
//
//	http.SetCookie(w, cookie)
//}
//
//func GetUser(id int) models.User {
//	var user models.User
//	return user
//}