package main

import (
	"fmt"
	"github.com/kiberlom/spider/comb"
	"github.com/kiberlom/spider/get"
)

const (
	GOR  = 500     // количество одновреммено работающий потоков
	PROT = "https" // протокол
	DOM  = "ua"    // доменная зона
)

// канал с очередью
var can chan int = make(chan int, GOR)

func testSite(name string, c chan int, i int) {

	// освобождаем место в канале
	defer func() {
		<-c
	}()

	// опрашиваем url
	s := get.New(PROT, name, DOM)
	err := s.Test()
	if err == nil {

	}
	fmt.Println(i, ". ", s.Url, "  -", s.ServerCode, "   [", err, "]")

}

func main() {

	//начальная строка
	a := "0"
	// просто счетчик
	i := 1

	for {

		// если есть возможность создать поток
		if len(can) < GOR {
			// занимаем место в канале
			can <- 1

			// создаем гарутину опроса по url
			go testSite(a, can, i)

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

}
