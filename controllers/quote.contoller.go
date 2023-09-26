package controllers

import (
	"fmt"
	"net/http"
	"quoation-backend/models"
	"quoation-backend/services"

	"github.com/gin-gonic/gin"
)

type QuoteController struct {
	QuoteService services.QuoteService

}
// creating constructor
func New(quoteservice services.QuoteService) QuoteController{
  return QuoteController{
	QuoteService: quoteservice,
  }
}

func (qc *QuoteController) CreateQuote(ctx *gin.Context){
	var quote models.Quote
	if  err:=ctx.ShouldBindJSON(&quote); err!=nil{
       ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	   return
	}
	err := qc.QuoteService.CreateQuote(&quote)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"quote created successfully"})
  }
  
  func (qc *QuoteController) GetQuote(ctx *gin.Context) {
	quotename := ctx.Param("id")
	quote,err:= qc.QuoteService.GetQuote(&quotename)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,quote)
	}
  
	func (qc *QuoteController) GetAll(ctx *gin.Context) {
		quotes,err := qc.QuoteService.GetAll()
		if err!=nil{
			fmt.Println("Error:", err)
			ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
			return
		}
		fmt.Println("Quotes:", quotes)
		ctx.JSON(http.StatusOK,quotes)
	}
  
	func (qc *QuoteController) UpdateQuote(ctx *gin.Context){
		var quote models.Quote
		if err := ctx.ShouldBindJSON(&quote); err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
		}
		err := qc.QuoteService.UpdateQuote(&quote)
		if err!=nil{
			ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"message":"quote updated successfully"})
	}

	func (qc *QuoteController) DeleteQuote(ctx *gin.Context) {
		quoteid := ctx.Param("id")
		err := qc.QuoteService.DeleteQuote(&quoteid)
		if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"message":"quote deleted successfully"})
	}

	// receiver methods

	func (qc *QuoteController) RegisterQuoteRoutes(rg *gin.RouterGroup){
		quotesroutes := rg.Group("/quotes")
         quotesroutes.POST("/create",qc.CreateQuote)
         quotesroutes.GET("/get/:id",qc.GetQuote)
         quotesroutes.GET("/getall",qc.GetAll)
         quotesroutes.PUT("/update",qc.UpdateQuote)
         quotesroutes.DELETE("/delete/:id",qc.DeleteQuote)
	}
  
  
  
  
  
  
  
  