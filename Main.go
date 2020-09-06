package main

import (
	"fmt"
	"math"
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

func sqrt(x float64) (float64, error) { // error as second arg
	if x < 0.0 {
		return 0.0, fmt.Errorf("can't root neg nos")
	}
	sq := 1.0
	for i := 0; i < 10; i++ {
		sq -= (sq*sq - x) / (2 * sq)
	}
	return sq, nil
}

func swap(x, y *string) (*string, *string) {
	*x += "*"
	*y += "+"
	return y, x
}

func sumint(values ...int) (rslt int) { //like *args  & default return
	fmt.Println(values)
	// rslt := 0
	for _, v := range values {
		rslt += v
	}
	return
}

func (anm Animal) roar() {
	fmt.Println(anm.Name, "Roaaaaaaar")
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
	fmt.Printf("%v ", ia)
	fmt.Printf("%v ", ib)
	fmt.Printf("%v \n", ic)
	fmt.Println("..... ॐ  .....")

	// variables --------------------
	// block scope only
	var i float64 = 65.6
	fmt.Println("Hello,", i)

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

	// Constants cannot be declared using the := syntax.
	const w, q int = 1, 2

	const c, python, java = true, false, "no!\n"
	fmt.Println(w, q, c, python, java)

	// Conditionals -----------
	fmt.Println("LOOPS")

	sum := 1
	for i := 0; i < 10; i++ {
		sum += sum
	}
	if true {
		fmt.Println(sum)
	}

	for i, j := 0, 0; i < 3; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	iloop := 0
	for iloop < 5 {
		fmt.Print(iloop)
		iloop++
	}
	fmt.Println("")

	s := []int{3, 2, 1, 0}
	for k, v := range s {
		fmt.Println(k, v)
	}

Loop: //label for break
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Print(i * j)
			if i*j >= 10 {
				break Loop //break tag used here
			}
		}
		fmt.Println("")
	}

	myNum := 0.41421
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.00001 {
		fmt.Println("\nThese are same no.")
	}

	fmt.Println("\nGo runs on switches ")
	// initializers switch
	switch os := runtime.GOOS; os { //break by default
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
	fmt.Println(t)
	switch { //tagless -> like if, if else, else
	case t.Hour() < 12 && t.Hour() > 4:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	var itrf interface{} = 1.
	switch itrf.(type) {
	case int:
		fmt.Println("itrf is a int")
	case float64:
		fmt.Println("itrf is a float64")
	case string:
		fmt.Println("itrf is a string")
	default:
		fmt.Println("itrf is a another type")
	}

	// byte conv
	bytestr := "ABCDEF"
	bconv := []byte(bytestr)
	fmt.Printf("\nStored as bytes %v -> %T\n", bconv, bconv)

	// defer/panic/recovery
	defer fmt.Println("\nDone")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
	defer fmt.Println("\nCounting")

	// num, den := 1, 0
	// ans := num / den
	// if den == 0 {
	// 	panic("Somthing bad") //like try/except
	// }
	// fmt.Println(ans)

	// primitives

	nbool := 1 == 1
	mbool := 1 == 2

	fmt.Printf("\n%v, %T\n", nbool, nbool)
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
	fmt.Printf("\n\nARRAY: %v\n", arr)

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
	fmt.Printf("\nSliceArr: %v\n", arrs)
	fmt.Printf("len of sliceArr: %v\n", len(arrs))
	fmt.Printf("Capacity of sliceArr: %v\n", cap(arrs))

	sa := arrs[3:7]
	fmt.Println(sa)
	fmt.Printf("\nlen of sa: %v\n", len(sa))
	fmt.Printf("Capacity of sa: %v\n", cap(sa))

	// built in make fn
	mkarr := make([]int, 3, 3) //initialized to 0, len 3, cap 3
	fmt.Println(mkarr)
	fmt.Printf("\nlen of mkarr: %v\n", len(mkarr))
	fmt.Printf("Capacity of mkarr: %v\n", cap(mkarr))

	mkarr = append(mkarr, 4, 5, 6, 7) //cap is doubled
	fmt.Println(mkarr)
	fmt.Printf("\nlen of mkarr: %v\n", len(mkarr))
	fmt.Printf("Capacity of mkarr: %v\n", cap(mkarr))

	mkarr = append(mkarr, sa...) // add slice using spread op
	fmt.Println(mkarr)
	fmt.Printf("\nlen of mkarr: %v\n", len(mkarr))
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
		fmt.Printf("\nPop of Goa: %v Is Key:%v\n\n", popGoa, ok)
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

	//Pointers
	fmt.Println("\n----Pointers----")
	// map and slices use pointers under the hood

	var ptra int = 42
	var ptrb *int = &ptra
	fmt.Println(ptra, &ptra, ptrb, *ptrb)
	ptra = 21
	fmt.Println(ptra, &ptra, ptrb, *ptrb)
	*ptrb = 10
	fmt.Println(ptra, &ptra, ptrb, *ptrb)

	ptrarr := [3]int{1, 2, 3}
	ptrc := &ptrarr[0]
	ptrd := &ptrarr[1] //cant perform pointer arithmetic
	fmt.Printf("\n%v %p %p\n", ptrarr, ptrc, ptrd)

	var anm *Animal
	fmt.Println(anm) // nil i.e not initialized
	anm = new(Animal)
	(*anm).Name = "Emu" // anm.Name will also work
	fmt.Println(*anm)

	//functions
	fmt.Println("")
	sqr, err := sqrt(25.125)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqr)
	}

	defsum := sumint(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is ", defsum)

	s1 := "hello"
	s2 := "world"
	var a, b *string = swap(&s1, &s2)
	fmt.Println(*a, *b)
	fmt.Println(s1, s2)

	// anonymous fn
	func(i int) {
		a := "\nHi" //local scope
		fmt.Println(a, "I'm Anonymous,", i)
	}(1)

	var f func() = func() {
		fmt.Println("I'm Anonymous, 2")
	}
	f()

	//methods

	// func(anm Animal) roar() {
	// 	fmt.Println(anm.Name, "Roaaaaaaar")
	// }
	lion := Animal{
		Name:   "Lion",
		Origin: "Africa",
	}
	lion.roar()

}
