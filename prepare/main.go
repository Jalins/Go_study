package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd

		err error
		output []byte
	)

	// 生成的ctx与cancelFunc之间存在一个channel用于传递信息
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		cmd = exec.CommandContext(ctx, "bash", "-c", "sleep 2; echo hello")

		if output, err = cmd.CombinedOutput(); err != nil {
			fmt.Println(err)
			return
		}


	}()


	time.Sleep(1*time.Second)

	cancelFunc()




}
