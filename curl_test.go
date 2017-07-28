package curl

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	c := New("http://opvbhjo1o.bkt.clouddn.com/2017/6/13/video_2017_06_13.mp4", true)
	c.SaveToFile("test.mp4")
	c.Progress(func(p ProgressStatus) {
		timeNeed := time.Duration(-1)
		if p.Speed > 0 {
			timeNeed = time.Duration(p.ContentLength/p.Speed) * time.Second
		}
		timeLeast := time.Duration(-1)
		if p.Speed > 0 {
			timeLeast = time.Duration(int64(float64(p.ContentLength)/float64(p.Speed)*float64(1-p.Percent))) * time.Second
		}

		// 打印下载状态
		fmt.Println(
			"Stat", p.Stat, // one of curl.Connecting / curl.Downloading / curl.Closed
			"speed", PrettySpeedString(p.Speed),
			"len", PrettySizeString(p.ContentLength),
			"got", PrettySizeString(p.Size),
			"time need", timeNeed,
			"time least", timeLeast,
			"percent", p.Percent,
			"paused", p.Paused,
		)
	}, time.Second)

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("强行关闭")
		err := c.ForceClose()
		fmt.Println(err)
	}()

	res, err := c.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("response:", res)
	}

	fmt.Println(os.Stat("test.mp4"))
}
