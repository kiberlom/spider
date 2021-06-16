package spider

import (
	"context"
	"fmt"
	"sync"

	"github.com/kiberlom/spider/internal/spider/comb"
	"github.com/kiberlom/spider/internal/spider/get"
)

// конфигурирование паука
type Config struct {
	Thread   int
	Protocol string
	Domain   string
	Ctx      context.Context
}

// канал с очередью
var can chan recv

// структура в канале с очередью
type recv struct {
	i   int
	url string
}

// запускаем паука
func Start(config Config) {

	// создаем канал для очереди
	can = make(chan recv, config.Thread)

	//начальная строка
	a := "0"
	// просто счетчик
	i := 1

	wg := &sync.WaitGroup{}
	for i := 0; i < config.Thread; i++ {
		wg.Add(1)
		// запускаем поток запроса для одного домена
		go testSite(can, config, wg)
	}

loop:
	for {
		select {
		case <-config.Ctx.Done():
			//если пришел сигнал завершения работы
			break loop
		default:
			// занимаем место в канале
			t := recv{i, a}
			// передаем в канал данные
			can <- t
			// создаем следующую комбинацию
			ceb, err := comb.Next(a)
			if err != nil {
				fmt.Println(err)
				return
			}
			// сохраняем комбинацию
			a = ceb
			// увеличиваем счетчик
			i++

		}

	}

	// ожидаем завершения всех потоков
	wg.Wait()

}

// проверка домена
func testSite(c chan recv, config Config, wg *sync.WaitGroup) {
	defer wg.Done()
loop:
	for {
		select {
		case <-config.Ctx.Done():
			// сизнал завершения работы программы
			break loop
		case u := <-c:
			// опрашиваем url
			s := get.New(config.Protocol, u.url, config.Domain)
			err := s.Test()
			if err != nil {
				//fmt.Println(err)
				continue
			}

			//fmt.Println(u.i, ".  ", s.Url, "  -", s.ServerCode, "   [", err, "]")
			fmt.Printf("%d  %s [%d] %v\n", u.i, s.Url, s.ServerCode, err)
		}

	}

}
