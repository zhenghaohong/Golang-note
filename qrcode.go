package main
import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
)

const QrcodeSavePath = "./qrcode/"
func main() {
	fileName := "qr.png"
	content := "https://www.baidu.com"
	createQrCode(fileName, content)
	loadQrCode(fileName, content)
}

func createQrCode(fileName, content string) {
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	//u1 := uuid.NewV4()
	_dir := QrcodeSavePath+time.Now().Format("20060102")

	os.MkdirAll(_dir, os.ModePerm)
	//fileName := _dir+"/"+u1.String()+".png"
	fileName = _dir+"/"+"sdfsf.png"

	file, _ := os.Create(fileName)
	defer file.Close()

	png.Encode(file, qrCode)
}


func loadQrCode(fileName, content string) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	matrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("contrast result:", matrix.Content == content)
	fmt.Println(matrix.Content)

}
