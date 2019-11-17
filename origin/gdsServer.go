package main

import (
	"context"
	"errors"
	"fmt"
	"go_chat/app/libs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(addr ...string) {

	// 项目启动地址
	var (
		address string
		err     error
	)

	// 判断地址是否符合规范，不符合则使用默认地址和端口
	address = resolveAddress(addr)

	// 阻塞程序用，防止 优雅退出 后续指令 还未 执行完毕 就退出了
	processed := make(chan struct{}, 1)

	if err = startServer(address, processed); err != nil {
		libs.DebugPrint(fmt.Sprintf("%v", err))
	}

	<-processed

}

func startServer(address string, processed chan struct{}) error {

	var err error

	server := &http.Server{
		Addr: address,
	}

	libs.DebugPrint("Listening and serving HTTP on %s\n", address)

	go func() {

		// 创建一个单信号 通讯量 channel
		c := make(chan os.Signal, 1)

		// 在此 设置 输入， 如果接收到 ctrl + c 终止程序信号， 则将信息传入c
		signal.Notify(c, os.Interrupt)

		// 在此阻塞， 如果上一步传入了信息，则这里将会 释放执行，如此便进入下面的 终止流程
		<-c

		// 自定义 超时 上下文
		ctx, cancel := context.WithTimeout(&ShutdownContext{Chan: nil}, 3*time.Second)

		defer cancel()

		// 将自定义上下文 绑定到 要运行的 http Server 实例，并执行了 shutdown方法，进行优雅退出 整个http 服务
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed, err: %v\n", libs.NewReportError(err))
		}

		libs.DebugPrint("server gracefully exit")

		// 优雅退出之后，关闭阻塞信道，可以让程序 最终结束
		close(processed)

	}()

	err = server.ListenAndServe()

	if http.ErrServerClosed != err {
		err = libs.NewReportError(errors.New(fmt.Sprintf("server not gracefully shutdown, err :%v", err)))
	}

	// 在程序优雅退出之后，可以关闭一系列的服务，比如 sql 服务等
	// TODO: Close something
	// content

	return err

}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			libs.DebugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		libs.DebugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too much parameters")
	}
}

// 实现 context.Context 接口
// Deadline() (deadline time.Time, ok bool)
// Done() <-chan struct{}
// Err() error
// Value(key interface{}) interface{}
// 让 http 优雅退出(graceful)
type ShutdownContext struct {
	Chan         chan struct{}
	DeadLineTime time.Time
}

func (s *ShutdownContext) Deadline() (deadline time.Time, ok bool) {

	deadline = s.DeadLineTime

	ok = true

	return

}

func (s *ShutdownContext) Done() <-chan struct{} {
	return s.Chan
}

func (s *ShutdownContext) Err() error {
	return nil
}

func (s *ShutdownContext) Value(key interface{}) interface{} {
	return nil
}
