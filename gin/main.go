package main

import (
	"log"
	"net/http"
	"fmt"
    "github.com/gin-gonic/gin"
	"time"
)

func main() {
   r := gin.New()
   r.Use(gin.Logger())
   r.Use(gin.Recovery())

   // gin.Context，封装了request和response
   r.GET("/", func(c *gin.Context) {
      c.String(http.StatusOK, "hello World!")
   })


   r.GET("/url",getUrlParams)

   // r.Static("/html","./html")

   //  r.StaticFS("/html", gin.Dir("/html", true))

   // group
   v1 := r.Group("/v1")
   {
   		v1.GET("/login",v1login)
   }

   v2 := r.Group("/v2")
   {
   		v2.GET("/login",v2login)
   }

   r.GET("/testShouldBind",shouldBindData)

   // XML、JSON、YAML 渲染
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

   r.GET("/moreJSON", func(c *gin.Context) {
   		// 数据结构
   		var msg struct{
   			Name string `json:"user"` // Name json 变成 user
			Message string
   			Age int
		}

		msg.Name  = "zhengteshuai"
		msg.Message= "hi"
		msg.Age = 18
	   // 注意 msg.Name 在 JSON 中会变成 "user"
	   // 将会输出： {"user": "Lena", "Message": "hey", "Number": 123}

		c.JSON(http.StatusOK,msg)
   })
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// static
	r.Static("/dir","./html")
	r.StaticFS("/html",http.Dir("html"))

   r.Run(":8000")
}

/*
	做api常用到的其实就是gin封装的各种render. 目前支持的有:

	func (c *Context) JSON(code int, obj interface{})
	func (c *Context) Protobuf(code int, obj interface{})
	func (c *Context) YAML(code int, obj interface{})
*/
func v1login(c *gin.Context) {
	c.JSON(200,gin.H{
		"msg":"v1 login",
	})
}

func v2login(c *gin.Context) {
	c.JSON(200,gin.H{
		"msg":"v2 login",
	})
}


func getUrlParams(c *gin.Context) {
	name := c.Query("name")
	contentType := c.Request.Header.Get("Content-Type")
	action := c.DefaultQuery("action", "more")
	fmt.Printf("page: %s; name: %s; contentType: %s", name, action, contentType)
	c.JSON(200, gin.H{
		"name":name,
		"action":action,
		"contentType":contentType,
	})
	//c.YAML(200, gin.H{
	//	"name":name,
	//	"action":action,
	//	"contentType":contentType,
	//})
	// c.String(http.StatusOK, name+" is "+ action)

}

// Must bind
/*
 *
*/

type Person struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int	`form:"age" json:"age"`
	Birthday time.Time `form:"birthday" json:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func shouldBindData(c *gin.Context)  {
	var data Person
	if c.ShouldBind(&data) == nil {
		log.Println(data.Name)
		log.Println(data.Age)
	}
	c.String(200,"1")

}



// 绑定html复选框 //
 type myForm struct {
 	Colors []string `form:"colors[]"`
 }

 func formHandler(c *gin.Context) {
 	var fakeForm myForm
 	c.ShouldBind(&fakeForm)
 	c.JSON(200, gin.H{"color": fakeForm.Colors})
 }
 // html
 var html = `
<form action="/" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red" />
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green" />
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue" />
    <input type="submit" />
</form>
`
// /--绑定html复选框 --/


