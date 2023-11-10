package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "log"
)
func showIndexPage(c *gin.Context) {
  articles := getAllArticles()
  render(c,gin.H{
	  "title":"Home Page",
	  "payload":articles},"index.html")
}

func getArticle(c *gin.Context){
	log.Println("getArticle function was invoked")
	if articleID,err := strconv.Atoi(c.Param("article_id"));err==nil{
		if article,err := getArticleById(articleID);err==nil{
			render(c,gin.H{
				"title":"Home Page",
				"payload":article},"index.html")
			
		}else{
		c.AbortWithError(http.StatusNotFound,err)
	        }
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}
func render(c *gin.Context, data gin.H, templateName string) {

	format:= c.Request.Header.Get("Accept")
	log.Println(format)
	if(format=="application/json"){
		c.JSON(http.StatusOK,data["payload"])
	}else if(format == "application/xml"){
		c.XML(http.StatusOK,data["payload"])
	}else{
		c.HTML(http.StatusOK,templateName,data)
	}

//	switch c.Request.Header.Get("Accept") {
//	case "application/json":
//	  // Respond with JSON
//	  log.Println("Json method was invoked")
//	  c.JSON(http.StatusOK, data["payload"])
//	case "application/xml":
//	  // Respond with XML
//	  log.Println("Switch case for application/xml was invoked")
//	  c.XML(http.StatusOK, data["payload"])
//	default:
//	  // Respond with HTML
//	  log.Println("HTML method invoked")
//	  c.HTML(http.StatusOK, templateName, data)
//	}
  
  }
