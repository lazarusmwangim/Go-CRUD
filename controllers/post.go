package controllers

import (
	"fmt"
	"net/http"
	"swan/models"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPost := models.Post{Title: input.Title, Content: input.Content}

	models.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

type MyPost struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func FindPostsRaw(c *gin.Context) {
	var posts []MyPost
	// models.DB.Raw("select c_name, id_number from customer_master").Rows()
	rows, err := models.DB.Raw("select title, content, date(created_at) from posts").Rows()

	if err != nil {
		fmt.Println(err)
		panic("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		var title, content string
		var created_at string

		err := rows.Scan(&title, &content, &created_at)

		if err != nil {
			panic(err.Error())
		}

		f1 := MyPost{Title: title, Content: content, CreatedAt: created_at}
		posts = append(posts, f1)
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
	// c.JSON(http.StatusOK, gin.H{"data": "posts"})
}
