package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
fmt.Println(albums) // [{1 Blue Train John Coltrane 56.99} {2 Jeru Gerry Mulligan 17.99} {3 Sarah Vaughan and Clifford Brown Sarah Vaughan 39.99}]
fmt.Println(albums[1]) //{2 Jeru Gerry Mulligan 17.99} 

router := gin.Default()
router.GET("/albums", getAlbums)
router.GET("/albums/:id", getAlbumByID)
router.POST("/albums", postAlbums)
router.DELETE("/albums/:id", deleteAlbumByID)
router.PUT("/albums/:id", putAlbumsByID)
router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums) // IndentedJSON- to serialize the struct into JSON
}

 
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

 
    foundAlbum := isIDFoundInAlbums(id, c)
    if foundAlbum {
    	return
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


/** Loop over the list of albums, looking for
an album whose ID value matches the parameter. **/
func isIDFoundInAlbums(id string, c *gin.Context) bool {
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return true
		}
	}
	return false
}


// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("id to be deleted", id)
	i, err := strconv.Atoi(id) // converting string to int
	fmt.Println("Object to be deleted", albums[i])
	if err != nil {
        // ... handle error
		fmt.Println("Error occured while converting the :id provided as path param")
        panic(err)
    }
	for _, a := range albums {
        if a.ID == id {
			deletedAlbums := remove(albums, i)
            c.IndentedJSON(http.StatusOK, deletedAlbums)
            return
        }
    }

}

func remove(slice []album, s int) []album {
    return append(slice[:s], slice[s+1:]...)
}

// PUT is used to replace the entire resource
func putAlbumsByID(c *gin.Context)  {
    id := c.Param("id")
    var putAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&putAlbum); err != nil {
        return
    }
    fmt.Println("Id values to be replaced", id)
    fmt.Println("Put Request Body", putAlbum)
 
    i, err := strconv.Atoi(id) // converting string to int  
    if err != nil {
        // ... handle error
		fmt.Println("Error occured while converting the :id provided as path param")
        panic(err)
    }

    albums[i] = putAlbum // replacing the complete element inside the albumbs array
    c.IndentedJSON(http.StatusCreated, albums)
  
}

// PATCH is used to replace only the part of the resource
func patchAlbumbsByID(c *gin.Context){

}