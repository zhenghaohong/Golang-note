package main
import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"log"
	"net/http"
)

func main()  {
	url:="https://www.cnblogs.com/ztshuai/p/13560155.html"
	resp, err := http.Get(url)
	// fmt.Println(resp)
	if resp.StatusCode != 200 {
		fmt.Println("请求错误")
	}
	
	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// html := `<html>
	// 			<body>
	// 				<h1 id="title">春晓</h1>
	// 				<p class="content1">
	// 				春眠不觉晓，
	// 				处处闻啼鸟。
	// 				夜来风雨声，
	// 				花落知多少。
	// 				</p>
	// 			</body>
	// 		</html>`

	dom,err := goquery.NewDocumentFromReader(strings.NewReader(url))
	fmt.Printf("%v",dom)
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("p").Each(func (i int, selection *goquery.Selection)  {
		fmt.Println(selection.Text())
	})


}
