package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	errG, ctx := errgroup.WithContext(context.Background())
	errG.Go(func() (err error) {
		errChan := make(chan error)
		serv := http.Server{Addr: "127.0.0.1:8080"}
		go func() {
			fmt.Println("server start")
			errChan <- serv.ListenAndServe()
		}()

		select {
		case e := <-errChan:
			err = errors.New("server err -> " + e.Error())
		case <-ctx.Done():
			serv.Close()
			fmt.Println("server stopped")
		}

		return
	})

	errG.Go(func() (err error) {
		sigChan := make(chan os.Signal)
		//忽略不同信号处理
		signal.Notify(sigChan)

		select {
		case s := <-sigChan:
			err = errors.New("signal triggered -> " + s.String())
		case <-ctx.Done():
			signal.Stop(sigChan)
			fmt.Println("signal notify stopped")
		}

		return
	})

	if err := errG.Wait(); err != nil {
		fmt.Println(err)
	}
}
