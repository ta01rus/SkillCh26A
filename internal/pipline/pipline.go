package pipline

import (
	"context"
	"log"
)

// слушатель ошибок
func ErrListner(ctx context.Context) (<-chan error, context.Context) {
	out := make(chan error)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				log.Println("piipline out")
				return
			case err := <-out:
				if err != nil {
					cancel()
					return
				}
			}
		}
	}()
	return out, ctx
}

func Printer(ctx context.Context, chErr chan error) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				log.Println("piipline out")
				return
			case p := <-out:
				log.Println(p)
			}
		}
	}()
	return out
}

// получем и отдаем в сеть
func IntSource(ctx context.Context, chErr chan error) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				log.Println("piipline out")
				return
			case p := <-out:
				log.Println(p)
			}
		}
	}()

	return out
}
