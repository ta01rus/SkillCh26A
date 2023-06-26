package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ta01rus/SkillCh26A/pkg/container"
	"github.com/ta01rus/SkillCh26A/pkg/logger"
	rdr "github.com/ta01rus/SkillCh26A/pkg/reader"
)

const (
	BUFF_RING_SIZE   = 100
	BUFF_TIME_PERIOD = 5 * time.Second
)

func main() {
	var (
		lg          = logger.NewConsoleLoger()
		ticker      = time.NewTicker(BUFF_TIME_PERIOD)
		ctx, cancel = context.WithCancel(context.Background())
		errChan     = make(chan error)
	)
	defer cancel()
	defer close(errChan)

	// обработка ошибок
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-errChan:
				lg.Error("err: %s", err.Error())
			}
		}
	}(ctx)

	// источник из консоли
	source := func(ctx context.Context) <-chan int {
		out := make(chan int)
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(out)
					return
				default:
					fmt.Println("Ввидите числа через пробел:")
					ar, err := rdr.IntArrReader(os.Stdin)
					if err != nil {
						errChan <- err
						continue
					}
					for _, num := range ar {

						out <- num
					}
					ticker.Reset(BUFF_TIME_PERIOD)
				}
			}
		}()
		return out
	}

	in := source(ctx)

	//1 Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа).
	worker1 := func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(out)
					return
				case num := <-in:
					if num < 0 {
						continue
					}
					out <- num
				}
			}
		}()
		return out
	}

	in1 := worker1(ctx, in)

	//2 Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0.
	worker2 := func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(out)
					return
				case num := <-in:
					if num%3 > 0 || num == 0 {
						continue
					}
					out <- num
				}
			}
		}()
		return out
	}

	in2 := worker2(ctx, in1)

	/*
		  3 Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому,
			который был дан в качестве задания в 19 модуле.
			В этой стадии предусмотреть опустошение буфера (и соответственно, передачу этих
			данных, если они есть, дальше) с определённым интервалом во времени.
			Значения размера буфера и этого интервала времени сделать настраиваемыми
			(как мы делали: через константы или глобальные переменные).
	*/
	worker3 := func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		bufRing := container.NewIntRing(BUFF_RING_SIZE)

		go func() {
			for {
				select {
				case <-ctx.Done():
					close(out)
					return
				case <-ticker.C:
					i := bufRing.Get()
					for i != nil {
						out <- *i
						i = bufRing.Get()
					}
				case num := <-in:
					bufRing.Put(num)
				}
			}
		}()
		return out
	}

	in3 := worker3(ctx, in2)

	printer := func(ctx context.Context, in <-chan int) {
		var (
			out = make(chan int)
		)

		go func() {
			for {
				select {
				case <-ctx.Done():
					close(out)
					return
				case num := <-in:
					lg.Info("Число: %d\n", num)
				}
			}
		}()
	}

	printer(ctx, in3)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c

}
