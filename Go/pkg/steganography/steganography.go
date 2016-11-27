package steganography

import (
	"../bits"
	"fmt"
	"os"
	//"io/ioutil"
	"encoding/binary"
	//"strconv"
	//"image"
	//"golang.org/x/image/bmp"
)

/*
func main() {
	fmt.Println("Done")
	ChangeBit()
	fmt.Println("Writing message")
	WriteMessage("Hola como estas viejo??", "../../uploads/bmp.bmp")
	fmt.Println("Reading message")
	ReadMessage("../../uploads/bmp.bmp")
}

/*
func ReadAndWrite(){
	number := 4
	bits.SetBit(&number, 0)
	fmt.Println(number)
	fmt.Println(bits.HasBit(number, 0))
	bits.ClearBit(&number, 0)
	fmt.Println(number)
	fmt.Println(bits.HasBit(number, 0))
	text, _ := ioutil.ReadFile("../../uploads/sample.txt")
	fmt.Println(string(text))

	file, _ := os.OpenFile("../../uploads/sample.txt", os.O_WRONLY, 0666)
	defer file.Close()
	some := []byte("Adios")
	file.Write(some)
	file.Sync()
}
*/

func getPixelSizeInBytes(file *os.File) uint64 {
	offset := make([]byte, 2)
	beforeSeek, _ := file.Seek(0, 1)
	file.Seek(28, 0)
	file.Read(offset)
	file.Seek(beforeSeek, 0)
	return byteToInt64(offset) / 8
}

func getDataOffset(file *os.File) uint64 {
	offset := make([]byte, 4)
	beforeSeek, _ := file.Seek(0, 1)
	file.Seek(10, 0)
	file.Read(offset)
	file.Seek(beforeSeek, 0)
	return byteToInt64(offset)
}

func byteToInt64(bytes []byte) uint64 {
	entero := make([]byte, 8)
	copy(entero, bytes)
	return binary.LittleEndian.Uint64(entero)
}

func ChangeBit() {
	file, _ := os.Open("../../uploads/bmp.bmp")
	defer file.Close()
	fmt.Printf("Data offset: %d Pixel Size: %d \n", getDataOffset(file), getPixelSizeInBytes(file))
	message := "Holas"
	dst := bytesToBits([]byte(message), len(message))
	message = string(bitsToBytes(dst, len(message)*8))
	fmt.Println("Message from chage bit: " + message)
}

func bytesToBits(bytes []byte, length int) []bool {
	dst := make([]bool, length*8)
	for i := 0; i < length; i++ {
		character := bytes[i]
		for j := 0; j < 8; j++ {
			dst[i*8+j] = bits.HasBit(character, uint(j))
			//fmt.Println(dst[i*8 + j])
		}
	}
	return dst
}

func bitsToBytes(src []bool, length int) []byte {
	strLen := length / 8
	str := make([]byte, strLen)
	fmt.Println(length / 8)
	for i := 0; i < strLen; i++ {
		for j := 0; j < 8; j++ {
			if src[i*8+j] {
				str[i] = bits.SetBit(str[i], uint(j))
			} else {
				str[i] = bits.ClearBit(str[i], uint(j))
			}
		}
	}
	return str
}

///*
func WriteMessage(message, filename string) {
	file, _ := os.OpenFile(filename, os.O_RDWR, 0666)
	defer file.Close()
	messageLenBits := make([]byte, 4)
	binary.LittleEndian.PutUint32(messageLenBits, uint32(len(message)))

	messageBits := bytesToBits([]byte(message), len(message))

	writeBits(bytesToBits(messageLenBits, 4), 4*8, 0, file)
	writeBits(messageBits, len(message)*8, 4*8, file)
}

func writeBits(bitsToStore []bool, length, offset int, file *os.File) {
	dataOffset := getDataOffset(file)
	pixelSize := getPixelSizeInBytes(file)
	beforeSeek, _ := file.Seek(0, 1)
	whereToRW := int64(dataOffset + uint64(offset))
	file.Seek(whereToRW, 0)
	for i := 0; i < length; i++ {
		beforeRead, _ := file.Seek(0, 1)
		lectura := make([]byte, pixelSize)
		file.Read(lectura)
		file.Seek(beforeRead, 0)

		if bitsToStore[i] {
			lectura[0] = bits.SetBit(lectura[0], 0)
		} else {
			lectura[0] = bits.ClearBit(lectura[0], 0)
		}

		file.Write(lectura)

		file.Seek(beforeRead+int64(pixelSize), 0)
	}

	file.Seek(beforeSeek, 0)
}

//*/

//*
func ReadMessage(filename string) string {
	fmt.Println("Starting read")
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()

	messageLenBytes := bitsToBytes(readBits(4*8, 0, file), 4*8)
	messageLen := int(byteToInt64(messageLenBytes))

	fmt.Printf("Message to read Len: %d\n", messageLen)

	messageBits := readBits(messageLen*8, 4*8, file)
	message := bitsToBytes(messageBits, messageLen*8)

	fmt.Println("Finishing read")

	fmt.Println("Message read: " + string(message))

	return string(message)
}

func readBits(length, offset int, file *os.File) []bool {
	dataOffset := getDataOffset(file)
	pixelSize := getPixelSizeInBytes(file)
	beforeSeek, _ := file.Seek(0, 1)
	whereToRW := int64(dataOffset + uint64(offset))
	file.Seek(whereToRW, 0)

	lectura := make([]bool, length)

	for i := 0; i < length; i++ {
		pixel := make([]byte, pixelSize)
		file.Read(pixel)
		lectura[i] = bits.HasBit(pixel[0], 0)
	}

	file.Seek(beforeSeek, 0)
	return lectura
}

//*/
