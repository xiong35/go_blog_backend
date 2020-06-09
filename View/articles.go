package view

/* import (
    "net/http"

    "gin/Models"
	"github.com/gin-gonic/gin"
)

func TestInsert(c *gin.Context) {
	var testModel models.Test

	err := c.ShouldBindJSON(&testModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := testModel.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})

}
*/
