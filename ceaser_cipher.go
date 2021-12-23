package main

import "fmt"

func CeaserCipher(s string, key int) string {
	new_str := []byte{}
	for _, ch := range s {
		fmt.Println(string(ch))
		if ch >= 'a' && ch <= 'z' {
			val := int(ch) + key
			if val > 'z' {
				val = int('a') - 1 + (val % int('z'))
			}
			new_str = append(new_str, byte(val))
			continue
		}
		new_str = append(new_str, byte(ch))
	}
	return string(new_str)
}
