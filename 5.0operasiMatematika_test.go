package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestOperasiMatematika(t *testing.T) {
	//Operasi matematika + , - , * , / , %
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	/* Augmented Assignments
	operasi matematika		augmented Assignments
	d = d + 5				d += 5
	d = d - 5				d -= 5
	d = d * 5				d *= 5
	d = d / 5				d /= 5
	d = d % 5				d %= 5
	*/
	var d = 10
	d += 5 //d=d+5
	fmt.Println(d)

	/* Unary Operator
	++		====>	e = e + 1
	--		====>	e = e - 1
	-		====>	negative
	+		====>	positive
	!		====>	bolean kebalikan
	*/
	var e = 15
	var negative = -100

	e++

	fmt.Println(e)
	fmt.Println(negative)
}
