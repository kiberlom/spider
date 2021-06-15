package get

import (
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

type setSite struct {
	Proto       string
	Name        string
	Dom         string
	Url         string
	ServerCode  int
	Server      []string
	ServerPower []string
	Body        string
}

func New(protokol, name, dom string) setSite {

	s := setSite{}
	s.Proto = protokol
	s.Name = name
	s.Dom = dom

	// формеруем путь
	s.Url = s.Proto + "://" + s.Name + "." + s.Dom

	return s

}

// запрос на сервер
func (s *setSite) Test() error {

	// создаем клиента
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// создаем запрос
	r, err := http.NewRequest("GET", s.Url, nil)
	if err != nil {
		return err
	}

	// устанавливаем заголовок
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	// делаем запрос
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	// получаем заголовки и ответ
	s.Server = resp.Header["Server"]
	s.ServerPower = resp.Header["X-Powered-By"]
	s.ServerCode = resp.StatusCode

	//кодируем в нужную кодировку из заголовка
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	// читаем тело
	if h, err := ioutil.ReadAll(utf8); err != nil {
		s.Body = "error read"
	} else {
		s.Body = string(h)
	}

	//fmt.Println(s.Url, " [", resp.StatusCode, "] ---", resp.Header["X-Powered-By"], "----------------------------", resp.Header["Server"])

	return nil

}
