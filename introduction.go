/*--- Basic Introduction for GoLang ---*/
package main

import "fmt"

func main() {
	/*--- Multiple var ---*/
	var (
		firstName = "Erwin"
		lastName  = "Ardiantha"
	)

	/*--- Multiple const ---*/
	const (
		age     = 29
		address = "Jakarta"
	)

	fmt.Println("var & const")
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(address)
	fmt.Println("----------")

	/*--- int conversion ---*/
	fmt.Println("\n>> number conversion")
	var nilai32 = 1000000
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)
	fmt.Println("64", nilai64)
	fmt.Println("8", nilai8)

	/*--- String conversion ---*/
	fmt.Println("\n>> string conversion")
	var name = "rrwin"
	var e = name[0]         // --> byte
	var eString = string(e) // --> array location
	fmt.Println(name)
	fmt.Println(eString)

	/*--- Type Declaration ---*/
	fmt.Println("\n>> type declaration")
	type NoKtp string
	type Married bool

	var noKtpErwin NoKtp = "123123"
	var isMarried Married = true

	fmt.Println(noKtpErwin, isMarried)

	/*--- Math Operation ---*/
	fmt.Println("\n>> math operation")
	var i = 10
	i += 10 // 20
	i++     // 21
	fmt.Println(i)

	/*--- Comparation Operator ---*/
	fmt.Println("\n>> comparation operator")
	var resultName bool = firstName == lastName
	fmt.Println(resultName)

	var value1 = 100
	var value2 = 200 // bisa langsung dikompare juga
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 != value2)

	/*--- Array ---*/
	fmt.Println("\n>> Array")
	// cara manual
	var siswa [2]string
	siswa[0] = "Ayu"
	siswa[1] = "Indira"

	fmt.Println(siswa)
	fmt.Println(siswa[0], siswa[1])

	// cara langsung
	var guru = [4]string{
		"Pak Taufik", "Bu Keo", "Samsuding", "",
	}
	fmt.Println(guru)
	fmt.Println(len(siswa))
	fmt.Println(len(guru))

	/*--- Slice ---*/
	fmt.Println("\n>> Slice")
	// iniate array
	var months = [...]string{
		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
	}

	// make slice from array
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// kalau ganti isi array monthnya, slice nya ikut keganti
	months[4] = "bukanbulan"
	fmt.Println(slice1)

	// ganti isi slice monthnya
	slice1[0] = "ayambebek"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "sapigoreng")
	fmt.Println(slice3)
	slice3[1] = "bukan november, "
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// menggunakan fungsi make untuk buat slice
	rak := make([]string, 2, 5)
	rak[0] = "Sam"
	rak[1] = "Sul"
	fmt.Println(rak)

	// copy slice | kapasitas harus sama
	copySlice := make([]string, len(rak), cap(rak))
	copy(copySlice, rak)
	fmt.Println(copySlice)

	/*--- Map ---*/
	fmt.Println("\n>> Map")
	person := map[string]string{
		"name":    "Erwin",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Lecturer"

	fmt.Println(person["title"])

	var book = make(map[string]string)
	book["title"] = "Einstein"
	book["price"] = "300000"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
