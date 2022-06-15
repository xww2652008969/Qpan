package initservice

import (
	"Qpan/global"
	"fmt"
	"os"
)

func InitLog() {
	var err error
	global.QP_logfile, err = os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
}
