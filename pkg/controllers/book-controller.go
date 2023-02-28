package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maksonviini/Go-Book-Management-System/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
}


func Create(c *gin.Context) {

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	id, err := models.Insert(book)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Inserido com sucesso! ID: %d", id),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	  })
}

func Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))


	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	var book models.Book

	err = c.ShouldBindJSON(&book)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	rows, err := models.Update(int64(id), book)

	var resp map[string]any

	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram atualizados %d registros", rows)
	}

	resp = map[string]any{
		"Error":   false,
		"Message": "Dados atualizados com sucesso!",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	  })
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	rows, err := models.Delete(int64(id))

	var resp map[string]any

	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros", rows)
	}

	resp = map[string]any{
		"Error":   false,
		"Message": "Dados removidos com sucesso!",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp,
	  })
}

func GetAll(c *gin.Context) {
	book, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": book,
	  })
}

func Get(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	book, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": book,
	  })
}
