package spider

import (
	"fmt"
	"runtime"

	"github.com/kiberlom/spider/internal/spider/comb"
	"github.com/kiberlom/spider/internal/spider/get"
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

func Start() {

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
