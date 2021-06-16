package get

import (
	"fmt"
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
		return fmt.Errorf(" невозможно создать Request: %v", err)
	}

	// устанавливаем заголовок
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	// делаем запрос
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf(" не удалось сделать запрос: %v", err)
	}
	defer resp.Body.Close()

	// получаем заголовки и ответ
	s.Server = resp.Header["Server"]
	s.ServerPower = resp.Header["X-Powered-By"]
	s.ServerCode = resp.StatusCode

	//кодируем в нужную кодировку из заголовка
	coding, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return fmt.Errorf(" ошибка при установке кодировки для тела: %v", err)
	}

	// читаем тело
	h, err := ioutil.ReadAll(coding)
	if err != nil {
		s.Body = fmt.Sprintf(" не удалось прочитать тело: %v", err)
	} else {
		s.Body = string(h)
	}

	//fmt.Println(s.Url, " [", resp.StatusCode, "] ---", resp.Header["X-Powered-By"], "----------------------------", resp.Header["Server"])

	return nil

}
