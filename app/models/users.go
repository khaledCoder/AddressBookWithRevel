package models

import (
	"database/sql"
	"fmt"
)

type Users struct{
	Pk int
	Name string
	Mobile string
}
type page struct{

	Contacts[]Users
}

func CreateUser(name string , mobile string)Users{
	db,err :=sql.Open("mysql","root:123456@/addressBook")
	checkErr(err)
	data,err := db.Prepare("INSERT INTO contacts (name, mobile) VALUES (?,?)")
	checkErr(err)
	res,err := data.Exec(name,mobile)
	checkErr(err)
	id,err := res.LastInsertId()
	checkErr(err)
var userID int
	userID=int(id)
	var user = FindUser(userID)
	return user
}
func FindAllUsers() page{
	p:=page{Contacts:[]Users{}}
	DB,err := sql.Open("mysql","root:123456@/addressBook")
	defer DB.Close()

	rows, err := DB.Query("SELECT * FROM contacts")
	checkErr(err)
	for rows.Next() {
		var b Users
		err = rows.Scan(&b.Pk , &b.Name, &b.Mobile)
		p.Contacts=append(p.Contacts,b)
		checkErr(err)
	}
	return p
}

func FindUser(id int) Users{

	db,err :=sql.Open("mysql","root:123456@/addressBook")
	checkErr(err)
	var userData Users
	err = db.QueryRow("SELECT * FROM contacts where id=?", id).Scan(&userData.Pk, &userData.Name, &userData.Mobile)
	checkErr(err)
	return userData
}

func UpdateUserInfo(user Users){
	db,err :=sql.Open("mysql","root:123456@/addressBook")
	checkErr(err)
	data,err:=db.Prepare("UPDATE contacts SET name=?,mobile=? WHERE id=?")
	res,err:= data.Exec(user.Name,user.Mobile,user.Pk)
	fmt.Println(res)
	checkErr(err)

}
func DeleteUser(id int){
	fmt.Println(id)
	db,err :=sql.Open("mysql","root:123456@/addressBook")
	checkErr(err)
	data,err:=db.Prepare("DELETE FROM contacts WHERE id=?")
	res,err:=data.Exec(id)
	checkErr(err)
	fmt.Println(res)
}

func checkErr(err error)  {
	if err!=nil{
		fmt.Println(err.Error())
	}
}