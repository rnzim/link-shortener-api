package controllers

import (
	"net/http"
	"shortLink/database"
	"shortLink/models"
	

	"github.com/gin-gonic/gin"
)

//Lista Todos os Links
func ListAll(ctx *gin.Context){
	var Links []models.Link
	database.DB.Find(&Links)
	ctx.JSON(200,Links)
}

//Lista Os Links Por Nome
//1-Pega O id
//2-Procura No Banco
//3-Retorna Not Foun Caso Não Exita É Ok Caso Exista
func ListByName(ctx *gin.Context){
	var Link models.Link
	Name:= ctx.Params.ByName("name")
	database.DB.Where(&models.Link{Name: Name}).First(&Link)
	if Link.ID == 0{
		ctx.JSON(404,gin.H{"Link Inexitente":Link})
		return
	}else{
		ctx.JSON(200,Link)
		return
		
	}

	
}
func CreateLink(ctx *gin.Context){
	var NewLink models.Link
	//Obtem Os Valores Do Metodo Post
	if err:= ctx.ShouldBindJSON(&NewLink); err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
	}
	
	database.DB.Where(&models.Link{Name: NewLink.Name}).First(&NewLink)
	//Verifica Se Existe Algum Link Com O Mesmo Nome
	if NewLink.ID == 0{

		database.DB.Create(&NewLink)

		ctx.JSON(200,NewLink)
		
		return
	}else{
		ctx.JSON(403,gin.H{"Nome Já Existente":NewLink})
		return
	}
}

func DeleteLink(ctx *gin.Context){
	var Link models.Link
	Id := ctx.Params.ByName("id")
	database.DB.First(&Link,Id)
	if Link.ID == 0{
		ctx.JSON(403,gin.H{"error":"Esse Link Não Existe"})
		return
	}else{
		database.DB.Delete(&Link,Id)
		ctx.JSON(200,gin.H{"status":"Ok"})
		return

	}	
}

func EditLink(ctx *gin.Context){
	var OldLink models.Link
	Id := ctx.Params.ByName("id")
	database.DB.First(&OldLink,Id)

	if OldLink.ID == 0{
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":"Not Found",
		})
	    return
	}

	if err := ctx.ShouldBindJSON(&OldLink); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}
	database.DB.Model(&OldLink).UpdateColumns(&OldLink)
	ctx.JSON(200,gin.H{
		"status":"Ok",
		"Link":OldLink,
	})
}

