package comb

import (
	"errors"
	"strings"
)

const (
	L = 20 // максимальное количество символов
)

type str struct {
	s string
}

// допустимые символы
var pull = []string{

	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",

	"a", "b", "c", "d",
	"e", "f", "g", "h",
	"i", "j", "k", "l",
	"m", "n", "o", "p",
	"q", "r", "s", "t",
	"u", "v", "w", "x",
	"y", "z",

	//"-", "_", "^", "%",
}

//проверка строки на количество символов
//и на наличие неизвестных символов
func (st *str)check() error {

	//----------------------------------------
	//1. проверяем на количество
	if len(st.s) > L {
		return errors.New("Ошибка: Привышенно допустимое количество символов в строке ")
	}

	//-----------------------------------------
	//2. проверяем наличие неизвестных символов
	var r string      //символ заданной строки
	var fail []string // массив неизвестных символов

	// передираем символы в заданной строке
OSN:
	for _, y := range st.s {

		//rune to string
		r = string(y)

		// если нашли сивол из заданной строки в перечне допустимых
		for _, x := range pull {
			if r == x {
				continue OSN
			}
		}

		// если символ не найденн
		// добовляем в список неизвестных символов
		fail = append(fail, r)
	}

	// проверим были ли найденны неизвестные символы
	if len(fail) > 0 {
		return errors.New("Ошибка: присутствуют неизвестные символы: " + strings.Join(fail, ", "))
	}

	return nil

}

// возвращаем следующий символ
func lastS(s string) (string, error) {

	for i, r := range pull {
		if r == s {
			if i < len(pull)-1 {
				return pull[i+1], nil
			} else {
				return "-1", errors.New("последний символ")
			}
		}
	}

	return "", errors.New("Ошибка: последний символ заданной строки, не обноружен в допустимых ( " + s + " )")

}

// сброс символов на первые
// r - массив рун заданной строки
// i - порядковый номер элемента массива рун (r), с которого надо начать сброс на первый символ
func reset(r []rune, i int) []rune {

	y := []rune(pull[0])[0]

	for t := i; t < len(r); t++ {
		r[t] = y
	}

	return r

}

// добавляем одну позицию к текушей строуке, так как закончились возможные комбинации
// i - длинна прежней строки
func (st *str)add() string {

	var a string

	// формируем новую строку с добавлением еще одной позиции
	for y:=0; y<=len(st.s); y++{
		a += pull[0]
	}

	return a
}

// получаем следующую комбинацию
func (st *str)next() (string, error){

	na := []rune(st.s)

	for i := len(na) - 1; i >= 0; i-- {
		//fmt.Println(string(na[i]))

		l, err := lastS(string(na[i]))

		// если есть следующий символ
		if err == nil {
			na[i] = []rune(l)[0]
			return string(reset(na, i+1)), nil
		}

		//если это последний символ
		if err != nil {
			if l == "-1" {
				if i == 0 {
					// это значит что первый символ в заданной строке последний и его не возможно поменять на следующий
					// значит больше нет возможных комбинаций
					// значит надо добавить еще одну позицию если это позволяют условия
					if len(na) == L{
						return "", errors.New("Болше комбинаций не существует ")
					}
					return st.add(), nil
				}
			} else {
				return "", errors.New("Ошибка такого символа нет ")
			}
		}
	}

	return "", errors.New("Что то побло не так ")
}

// основная функция
func Next(s string) (string, error) {

	sroc := &str{s}

	// преверим переданную строку
	if err := sroc.check(); err != nil {
		return s, err
	}

	// получим новую комбинацию
	if n, err:=sroc.next();err!=nil{
		return s, err
	}else {
		return n, nil
	}


	return s, errors.New("Ошибка: основная функция Next завершилась ничем ")

}
