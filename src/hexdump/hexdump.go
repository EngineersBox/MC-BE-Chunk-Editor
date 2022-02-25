package hexhump

import "fmt"

func charFormatBytes(data []byte) string {
	buf := ""
	for _, b := range data {
		char := "."
		if b >= 0x20 && b <= 0x7f {
			char = fmt.Sprintf("%c", b)
		}
		buf += char
	}
	return buf
}

func createByteSeparatedLiteral(data []byte) string {
	buf := ""
	for i := 0; i < len(data); i += 2 {
		buf += fmt.Sprintf(" %02x", data[i])
		if (i + 1) != len(data) {
			buf += fmt.Sprintf("%02x", data[i+1])
		}
	}
	return buf
}

func CreateHexdumpText(data []byte) string {
	dwordCount := len(data) / 16
	residual := len(data) % 16
	buf := ""
	for i := 0; i < dwordCount; i++ {
		buf += fmt.Sprintf(
			"%s: %s  %s\n",
			fmt.Sprintf("%08x", i*0x10),
			createByteSeparatedLiteral(data[i*16:(i+1)*16]),
			charFormatBytes(data[i*16:(i+1)*16]),
		)
	}
	residualByteString := fmt.Sprintf("%08x: ", (dwordCount)*0x10)
	for i := 0; i < residual/2; i += 2 {
		if (i+1)*16 == len(data) {
			residualByteString += fmt.Sprintf(
				"%-39s%s",
				fmt.Sprintf("%02x", data[(i*16)-1]),
				charFormatBytes(data[(i*16)-1:]),
			)
			break
		}
		residualByteString += fmt.Sprintf(
			" %-39s  %s",
			fmt.Sprintf("%02x%02x", data[(i*16)+dwordCount-1], data[((i+1)+16)-1]),
			charFormatBytes(data[dwordCount*16:]),
		)
	}
	return buf + residualByteString
}
