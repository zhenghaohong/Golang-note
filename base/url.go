package main
import(
	"fmt"
	"net/url"
	// "strings"

)

func main(){
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User.Username)
	fmt.Println(u)
}



