package main

import (
	"bufio"
	"encoding/json"
	"time"
	"net"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Wakeup struct {
	ID int64 `gorm:"column:id;primary_key" form:"id" json:"id" example:"1" `
	// 设备id
	DeviceId int64 `gorm:"column:device_id" form:"device_id" json:"device_id"`

	// 胸腔位置
	ChestLocation string `gorm:"column:chest_location" form:"chest_location" json:"chest_location"`
	// 复数
	Complex string `gorm:"column:complex" form:"complex" json:"complex"`
	// 振幅
	Amplitude string `gorm:"column:amplitude" form:"amplitude" json:"amplitude"`
	// 相位差
	PhaseDifference string `gorm:"column:phase_difference" form:"phase_difference" json:"phase_difference"`
	// 均方差
	MeanSquare string `gorm:"column:mean_square" form:"mean_square" json:"mean_square"`
	// 胸腔运动
	ChestMove string `gorm:"column:chest_move" form:"chest_move" json:"chest_move"`

	// 呼吸
	Breathing bool `gorm:"column:breathing" form:"breathing"  json:"breathing"`
	// 心率
	HeartRate bool `gorm:"column:heart_rate" form:"heart_rate"  json:"heart_rate"`
	// 呼气
	Exhale bool `gorm:"column:exhale" form:"exhale"  json:"exhale"`
	// 吸气
	Inhale bool `gorm:"column:inhale" form:"inhale"  json:"inhale"`
	// 在床
	OnBed bool `gorm:"column:on_bed" form:"on_bed"  json:"on_bed"`

	// 体动
	BodyMove int `gorm:"column:body_move" form:"body_move"  json:"body_move"`

	// 配置
	Config int `gorm:"column:config" form:"config"  json:"config"`

	// 刷新率
	RefreshRate int `gorm:"column:refresh_rate" form:"refresh_rate"  json:"refresh_rate"`

	// 增益
	Gain string `gorm:"column:gain" form:"gain" json:"gain"`
	// 起点
	StartPoint string `gorm:"column:start_point" form:"start_point" json:"start_point"`
	// 范围
	Range string `gorm:"column:range" form:"range" json:"range"`
	// 步长
	Step string `gorm:"column:step" form:"step" json:"step"`

	//创建时间
	CreateTime int64 `gorm:"column:create_time" json:"createTime"`

	//是否删除
	IsDel int `gorm:"column:is_del" json:"isDel"`

}


var Db *sqlx.DB
func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ecus")
	//database, err := sqlx.Open("mysql", "beta:root_beta@tcp(127.0.0.1:3306)/beta")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}


// TCP server端

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("读取客户端数据失败, err:", err)
			break
		}
		recvStr := string(buf[:n])
		if strings.Index(recvStr,"{")==0 {
			// var dat map[string]interface{}
			var dat Wakeup
			if err := json.Unmarshal([]byte(recvStr), &dat); err == nil {
				//fmt.Printf("%#v", dat)
				amplitude := dat.Amplitude
				body_move := dat.BodyMove
				breathing := dat.Breathing
				chest_location := dat.ChestLocation
				chest_move := dat.ChestMove
				complex := dat.Complex
				config := dat.Config
				device_id := dat.DeviceId
				exhale := dat.Exhale
				gain := dat.Gain
				heart_rate := dat.HeartRate
				inhale := dat.Inhale
				mean_square := dat.MeanSquare
				on_bed := dat.OnBed
				phase_difference := dat.PhaseDifference
				ranges := dat.Range
				refresh_rate := dat.RefreshRate
				start_point := dat.StartPoint
				step := dat.Step
				create_time := time.Now().Unix()

				r, err := Db.Exec("insert into e_wakeup(amplitude, body_move, breathing,chest_location, chest_move, complex, config, device_id, exhale, gain, heart_rate, " +
					"inhale, mean_square, on_bed, phase_difference, ranges, refresh_rate, start_point, step, create_time)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
					amplitude,body_move,breathing,chest_location, chest_move, complex, config, device_id, exhale, gain, heart_rate,
					inhale, mean_square, on_bed, phase_difference, ranges, refresh_rate,start_point, step, create_time)
				if err != nil {
					fmt.Println("exec failed, ", err)
					return
				}
				id, err := r.LastInsertId()
				if err != nil {
					fmt.Println("exec failed, ", err)
					return
				}
				fmt.Println("insert success:", id)
				//conn.Write([]byte(recvStr))
			}

		}

		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}

}


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	//listen, err := net.Listen("tcp", "0.0.0.0:8003")
	if err != nil {
		fmt.Println("监听失败, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("接收数据失败, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接

	}
}

