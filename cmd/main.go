package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kiberlom/spider/internal/config"
	"github.com/kiberlom/spider/internal/spider"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("конфиг не загружен: %v\n", err)
	}

	wg := &sync.WaitGroup{}
	// завершение работы
	ctx, cancel := context.WithCancel(context.Background())

	// Ожидание на получение сигналов от системы для завершения работы
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
		<-done
		log.Println("Завершение работы")
		cancel()
	}()

	// запускаем паука
	wg.Add(1)
	go func() {
		defer wg.Done()
		spider.Start(spider.Config{
			Thread:   c.GetInt("spider.thread"),
			Protocol: c.GetString("spider.protocol"),
			Domain:   c.GetString("spider.domain"),
			Ctx:      ctx,
		})
	}()

	wg.Wait()

	fmt.Println("-= Работа завершена =-")

}
