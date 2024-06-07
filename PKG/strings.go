package PKG

import (
	"bufio"
	"log"
	"os"
)

func Strings(fileName string, l rune, cache *[8]string) {

	file, err := os.Open(fileName)
	//fmt.Println("Openinging <====", fileName)
	if err != nil {
		log.Fatal("error reading banner file!!! ", err)
	}
	defer file.Close()

	//scanning file
	scanner := bufio.NewScanner(file)

	c := 0
	i := 0
	char := (int(l)-32)*9 + 2
	for scanner.Scan() {
		c++
		if c >= char && c < char+8 {
			cache[i] = cache[i] + scanner.Text()
			i++
		} else if c >= char+8 {
			return
		}

	}
}
