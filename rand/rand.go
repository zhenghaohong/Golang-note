package main
import "fmt"
//import "math/rand"
import "time"
func main() {
 //    number := rand.Intn(100)
   // fmt.Println(number)
    fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
    fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
    fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
    str := "QB"
    nowTime := time.Now().Format("20060102")
    unix := time.Now().UnixNano() / 1e6
    DesignCode := fmt.Sprintf("%v%v%v", str, nowTime, unix)

    fmt.Println(DesignCode)
}


