package frame

//一些杂的字节处理函数
import (
	"encoding/binary"
	"fmt"
	_ "frame/logger"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func CheckCondition(flag bool) bool {
	if flag {
		return true
	}
	return false
}

func CheckConditionVoid(flag bool) {
	if !flag {
		return
	}
}

func ArrayToInt(array []byte) uint {
	if len(array) != 4 {
		return 0
	}
	var data uint32
	/*
	   data = uint( array[3] ) & 0xff
	   data = data | uint( array[3] ) & 0xff
	   data = data | uint( array[2] ) << 8 & 0xff00
	   data = data | uint( array[1] ) << 16 & 0xff0000
	   data = data | uint( array[0] ) << 24 & 0xff000000
	*/
	data = binary.BigEndian.Uint32(array[0:])
	return uint(data)
}

func IntToArray(data uint) []byte {
	array := make([]byte, 4, 4)
	/*
	   array[0] = byte( ( 0xff000000 & data ) >> 24 )
	   array[1] = byte( ( 0xff0000 & data ) >> 16 )
	   array[2] = byte( ( 0xff00 & data ) >> 8 )
	   array[3] = byte( ( 0xff & data ) )
	*/
	binary.BigEndian.PutUint32(array[0:], uint32(data))
	return array
}

/*测试代码
func main() {
	var a uint = 511
	array := make([]byte, 4, 4)
	array = IntToArray(a)
	b := ArrayToInt(array)
	logger.Info(array)
	logger.Info(b)
}
*/
