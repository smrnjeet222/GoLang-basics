package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

// k:=6 not possible in global scope
var (
	// lowercase 1st letter var are scoped to the package
	k     int8   = 6
	fname string = "Simrajneet"
	lname string = "Singh"
	// uppercase 1st letter expose to outer levels
	Number int = 99
)

func swap(x, y string) (string, string) {
	return y, x
}

func sqrt(x float64) float64 {
	sq := 1.0
	for i := 0; i < 10; i++ {
		sq -= (sq*sq - x) / (2 * sq)
	}
	return sq
}

// Struct

// Vertex is for shapes
type Vertex struct {
	X      int
	Y      int
	Shapes []string
}

// Animal is ...
type Animal struct {
	Name   string `required:"true" max:"100"` // tags
	Origin string
}

// Bird is  child of animal
type Bird struct {
	Animal
	CanFly bool
}

const (
	// _  = iota   // skip 0
	ia = iota
	ib // by default ib=iota i.e increments ++
	ic // by default ic=iota (if scoped to const block)
)

func main() {
	// enumerated constansts (iota)
	fmt.Printf("%v\n", ia)
	fmt.Printf("%v\n", ib)
	fmt.Printf("%v\n", ic)

	// variables --------------------
	// block scope only
	var i float64 = 65.6
	fmt.Println("Hello, ॐ ,", i)

	j := 84
	fmt.Printf("\nj = %v -> %T", j, j)
	j = int(i)
	fmt.Printf("\nj = %v -> %T", j, j)

	fmt.Printf("\n\nk = %v -> %T", k, k)
	k := 6.4 // can redeclare for type convertion in diff scope i.e shadow them
	fmt.Printf("\nk = %v -> %T", k, k)

	var str string
	// str = string(j)
	// fmt.Printf("\n\nstr = %v -> %T (to ASCII)", str, str)

	str = strconv.Itoa(j)
	fmt.Printf("\nstr = %v -> %T (not ASCII)", str, str)

	fmt.Printf("\n\n%v -> %T\n", fname, fname)

	var a, b string = swap("hello", "world")
	fmt.Println(a, b)

	// Constants cannot be declared using the := syntax.
	const w, q int = 1, 2

	const c, python, java = true, false, "no!"
	fmt.Println(w, q, c, python, java)

	// Conditionals -----------
	sum := 1
	for i := 0; i < 10; i++ {
		sum += sum
	}
	if true {
		fmt.Println(sum)
	}

	fmt.Println(sqrt(2))

	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	bytestr := "ABCDEF"
	bconv := []byte(bytestr)
	fmt.Printf("Stored as bytes %v -> %T\n", bconv, bconv)

	fmt.Println("counting")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Print(i) // at last / reverse
	// }

	fmt.Println("done")

	// primitives

	nbool := 1 == 1
	mbool := 1 == 2

	fmt.Printf("%v, %T\n", nbool, nbool)
	fmt.Printf("%v, %T\n", mbool, mbool)

	bita := 10                // 1010
	bitb := 3                 // 0011
	fmt.Println(bita & bitb)  // 0010
	fmt.Println(bita | bitb)  // 1011
	fmt.Println(bita ^ bitb)  // 1001
	fmt.Println(bita &^ bitb) // 0100
	fmt.Println(bita << 3)    // 10*2^3 = 80
	fmt.Println(bita >> 3)    // 10/2^3 = 1

	r := 'a' // rune (int32) encoded in utf32
	fmt.Printf("%v, %T\n", r, r)

	coma := 1 + 2i   // real(a) + imag(a)i
	comb := 2 + 5.2i // or complex(2, 5.2)
	fmt.Println(coma + comb)
	fmt.Println(coma - comb)
	fmt.Println(coma * comb)
	fmt.Println(coma / comb)

	// Arrays & Slices

	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("\nARRAY: %v\n", arr)

	var stud [3]string
	stud[0] = "Simranjeet"
	fmt.Printf("Students: %v\n", stud)
	fmt.Printf("no.of students: %v\n", len(stud))

	arr2 := arr // arr considered as value itself
	arr2[2] = 0 // new arr is created
	fmt.Println(arr)
	fmt.Println(arr2)

	arrref := &arr //use reference to prevent copy
	arrref[3] = 0
	fmt.Println(arrref)
	fmt.Println(arr)

	// slice
	arrs := []int{8, 7, 6, 5, 4, 3, 2, 1, 0}
	barrs := arrs //slice naturally use refernces
	barrs[0] = 0
	fmt.Printf("SliceArr: %v\n", arrs)
	fmt.Printf("len of sliceArr: %v\n", len(arrs))
	fmt.Printf("Capacity of sliceArr: %v\n", cap(arrs))

	sa := arrs[3:7]
	fmt.Println(sa)
	fmt.Printf("len of sa: %v\n", len(sa))
	fmt.Printf("Capacity of sa: %v\n", cap(sa))

	// built in make fn
	mkarr := make([]int, 3, 3) //initialized to 0, len 3, cap 3
	fmt.Println(mkarr)
	fmt.Printf("len of mkarr: %v\n", len(mkarr))
	fmt.Printf("Capacity of mkarr: %v\n", cap(mkarr))

	mkarr = append(mkarr, 4, 5, 6, 7) //cap is doubled
	fmt.Println(mkarr)
	fmt.Printf("len of mkarr: %v\n", len(mkarr))
	fmt.Printf("Capacity of mkarr: %v\n", cap(mkarr))

	mkarr = append(mkarr, sa...) // add slice using spread op
	fmt.Println(mkarr)
	fmt.Printf("len of mkarr: %v\n", len(mkarr))
	fmt.Printf("Capacity of mkarr: %v\n", cap(mkarr))

	mkarr = mkarr[:len(mkarr)-1] // pop
	fmt.Println(mkarr)

	cutarr := append(mkarr[:2], mkarr[7:]...) //cutting middle
	fmt.Println(cutarr)                       // but this ruins the main arr (be carefull)

	// Maps
	statePop := make(map[string]int, 5)
	statePop = map[string]int{ //array is valid key type but not slice
		"Delhi": 4663161436,
		"Goa":   346243,
		"Ohio":  462463,
	}
	statePop["GGN"] = 452352
	delete(statePop, "Ohio")
	fmt.Println(statePop) //not in order though

	// sp := statePop  // passed by reference statePop is also changed
	// fmt.Printf("Pop of Goa: %v\n", statePop["Goa"]) //for no key return 0

	if popGoa, ok := statePop["Goa"]; ok {
		fmt.Printf("Pop of Goa: %v Is Key:%v\n", popGoa, ok)
	} // ok is false if key not present

	//  Struct  (strict are vaue types & copy is passed like arrays)
	v := Vertex{Y: 1, X: 2, Shapes: []string{"circle", "square", "rec"}}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	//anonymous struct
	prs := struct {
		name string
		age  int
	}{
		name: "Simranjeet",
		age:  20,
	}
	fmt.Println(prs)

	//struct composition/embeding  and Tags
	brd := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		CanFly: false,
	}

	fmt.Println(brd)

	tag := reflect.TypeOf(Animal{})
	field, _ := tag.FieldByName("Name")
	fmt.Println(field.Tag)

}
