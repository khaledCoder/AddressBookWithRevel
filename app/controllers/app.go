package controllers

import (
	"github.com/revel/revel"
	"github.com/go-sql-driver/mysql"
	"AddressBookWithRevel/app/models"
)
var DB *mysql.MySQLDriver
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
 var Contacts =models.FindAllUsers()
	return c.Render(Contacts)
}
func (c App) FindUser(id int)revel.Result{
 	var user = models.FindUser(id)
	return c.RenderJson(user)
}
func (c App) SaveUserInfo()revel.Result{
	var user  models.Users
	c.Params.Bind(&user.Pk,"idHiddenVal")
	c.Params.Bind(&user.Name,"name")
	c.Params.Bind(&user.Mobile,"mobile")
	c.Validation.Required(user.Name).Message("name is required!")
	c.Validation.MinSize(user.Mobile, 11).Message("mobile number is not valid!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	if(user.Pk!=0){
		models.UpdateUserInfo(user)
	}else{
		models.CreateUser(user.Name,user.Mobile)
}
	return c.Redirect(App.Index)
}
func (c App) DeleteUser(id int) revel.Result{
	models.DeleteUser(id)
	return c.RenderJson(id)
}