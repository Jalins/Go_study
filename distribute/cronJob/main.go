package cronJob

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	cronExpress *cronexpr.Expression
	nextTime time.Time
}

func main() {

	var (
		expr *cronexpr.Expression
		err error
		cronJob *CronJob
		cronJobMap map[string]*CronJob
	)

	cronJobMap = make(map[string]*CronJob)

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	cronJob = &CronJob{
		expr,
		expr.Next(time.Now()),
	}

	cronJobMap["job1"] = cronJob


	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	cronJob = &CronJob{
		expr,
		expr.Next(time.Now()),
	}

	cronJobMap["job2"] = cronJob


	go func() {
		for{
			for key, value := range cronJobMap {
				if value.nextTime.Before(time.Now()) || value.nextTime.Equal(time.Now()) {
					go func() {
						fmt.Println("任务被调度：", key)
					}()

					value.nextTime = value.cronExpress.Next(time.Now())
					fmt.Println(key, "下次执行时间为：", value.nextTime)
				}

			}

			// 阻塞for循环， 控制for循环遍历速度，防止循环次数过多从而占用系统资源
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}
		}

	}()


	time.Sleep(100 * time.Second)

}
