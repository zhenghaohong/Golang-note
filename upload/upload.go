package upload
import(
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"os"
)

// gin 上传文件
func InsertImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	guid := uuid.New().String()
	filPath := "static/img/" + guid + ".jpg"
	for _, file := range files {
		log.Println(file.Filename)
		// 上传文件至指定目录
		_ = c.SaveUploadedFile(file, filPath)
	}
	//ImgUrl := "/" + filPath
	//app.OK(c, ImgUrl, "upload success")
	// app.OK(c, filPath, "upload success")
	// return "upload success"
	c.JSON(200, gin.H{"message": "upload success"})
}

// 删除文件
func DeleteImg(c *gin.Context) {
	fileName := c.Request.FormValue("fileName")
	err := os.Remove(fileName)
	if err != nil {
		c.JSON(500, gin.H{"message": "delete fail"})
	}
	c.JSON(200, gin.H{"message": "delete success"})
	// tools.HasError(err, msg.DeletedFail, 500)
	// app.OK(c, nil, msg.DeletedSuccess)

}

