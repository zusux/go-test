package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func GetLog()  {
	t, err := tail.TailFile("nginx.log", tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll: true,
	})
	if err != nil{

	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
