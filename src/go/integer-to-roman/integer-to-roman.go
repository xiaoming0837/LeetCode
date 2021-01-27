package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(140))
}

func intToRoman(num int) string {
	m := num / 1000
	cm := (num - m*1000) / 900
	d := (num - m*1000 - cm*900) / 500
	cd := (num - m*1000 - cm*900 - d*500) / 400
	c := (num - m*1000 - cm*900 - d*500 - cd*400) / 100
	xc := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100) / 90
	l := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90) / 50
	xl := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50) / 40
	x := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50 - xl*40) / 10
	ix := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50 - xl*40 - x*10) / 9
	v := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50 - xl*40 - x*10 - ix*9) / 5
	iv := (num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50 - xl*40 - x*10 - ix*9 - v*5) / 4
	i := num - m*1000 - cm*900 - d*500 - cd*400 - c*100 - xc*90 - l*50 - xl*40 - x*10 - ix*9 - v*5 - iv*4

	p := func(num int, str string) (result string) {
		for i := 0; i < num; i++ {
			result += str
		}
		return result
	}

	return p(m, "M") + p(cm, "CM") + p(d, "D") + p(cd, "CD") + p(c, "C") + p(xc, "XC") + p(l, "L") + p(xl, "XL") + p(x, "X") + p(ix, "IX") + p(v, "V") + p(iv, "IV") + p(i, "I")
}
