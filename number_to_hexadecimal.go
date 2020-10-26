package main

func toHex(num int) (result string) {
	if num == 0 {
		return "0"
	}

	hex := "0123456789abcdef"

	for number := num; number != 0; {
		remainder := number & 15
		result = string(hex[remainder]) + result
		number = int(uint32(number) >> 4)
	}

	return
}
