package main

import (
	"fmt"
	"github.com/kiberlom/spider/comb"
	"github.com/kiberlom/spider/get"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

const (
	GOR  = 1000    // количество одновреммено работающий потоков
	PROT = "https" // протокол
	DOM  = "cz"    // доменная зона
)

// канал с очередью
var can = make(chan recv, GOR)

type recv struct {
	i   int
	url string
}

// проверка домена
func testSite(c chan recv) {

	for {
		select {
		case u := <-c:
			// опрашиваем url
			s := get.New(PROT, u.url, DOM)
			err := s.Test()
			if err != nil {
				fmt.Println(err)
				continue
			}

			g := runtime.NumGoroutine()
			fmt.Println(u.i, ".  [", g, "]  ", s.Url, "  -", s.ServerCode, "   [", err, "]")
		}

	}

}

func main() {

	// pprof
	go http.ListenAndServe(":8080", nil)

	//начальная строка
	a := "0"
	// просто счетчик
	i := 1

	for i := 0; i < GOR; i++ {
		go testSite(can)
	}

	for {

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
