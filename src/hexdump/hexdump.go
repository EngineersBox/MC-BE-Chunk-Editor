package hexhump

import "fmt"

func charFormatBytes(data []byte) string {
	buf := ""
	for b, _ := range data {
		char := "."
		if b >= 0x20 && b <= 0x7f {
			buf += fmt.Sprintf("%x", b)
		}
	}
	return buf
}

func createByteSeparatedLiteral(data []byte) string {
	buf := ""
	for i := 0; i < len(data); i += 2 {
		buf += fmt.Sprintf(" %x", data[i])
		if (i + 1) != len(data) {
			buf += fmt.Sprintf("%x", data[i+1])
		}
	}
	return buf
}

func createHexdumpText(data []byte) string {
	dwordCount := len(data) / 16
	residual := len(data) % 16
	buf := ""
	for i := 0; i < dwordCount; i++ {
		address := fmt.Sprintf("000000%x", i*0x10)
		buf += fmt.Sprintf(
			"%s: %s\n",
			address[len(address)-1:],
			createByteSeparatedLiteral(data[i*16:(i+1)*16]),
		)
	}
	residualByteString := fmt.Sprintf("000000%x", (dwordCount)*0x10)
	residualByteString = residualByteString[len(residualByteString)-1:]
	for i := 0; i < residual/2; i += 2 {
		if (i+1)*16 == len(data) {
			residualByteString += fmt.Sprintf(" %x", data[(i*16)-1])
			break
		}
		residualByteString += fmt.Sprintf(" %x%x", data[(i*16)+dwordCount-1], data[((i+1)+16)-1])
	}
	return buf + residualByteString
}
