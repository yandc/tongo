package boc

import (
	"fmt"
	"testing"
)

func TestOn(t *testing.T) {

	//fmt.Println(math.Ceil(float64(1023) / float64(8)))
	var str = NewBitString(8 * 10)
	//
	str.Print()
	//str.WriteBigUint(big.NewInt(255), 8)
	str.WriteCoins(77)
	//str.WriteInt(-2, 8)
	//str.WriteBigInt(big.NewInt(-128), 8)
	//str.WriteUint(25, 8)
	//str.WriteByte(2)
	str.Print()

	var reader = NewBitStringReader(&str)

	var num = reader.ReadCoins()
	fmt.Println(num)

	//var n = big.NewInt(1)
	//fmt.Println("77777", n.BitLen())
	//
	//for i := n.BitLen() - 1; i != 0; i-- {
	//	fmt.Println(n.Bit(i))
	//}

	//str.WriteBigInt(big.NewInt(127), 8)
	//str.On(0)
	//
	//str.On(8)
	//str.Print()
	//
	//str.Off(8)
	//str.Print()
	//
	//str.Toggle(8)
	//str.Print()
	//str.Toggle(8)
	//str.Print()

}