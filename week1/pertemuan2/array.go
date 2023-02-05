// A.15. Array
// Contoh penerapan array:

var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"

fmt.Println(names[0], names[1], names[2], names[3])


// A.15.1. Pengisian Elemen Array yang Melebihi Alokasi Awal

var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"
names[4] = "ez" // baris kode ini menghasilkan error


// A.15.2. Inisialisasi Nilai Awal Array

var fruits = [4]string{"apple", "grape", "banana", "melon"}

fmt.Println("Jumlah element \t\t", len(fruits))
fmt.Println("Isi semua element \t", fruits)


// A.15.3. Inisialisasi Nilai Array Dengan Gaya Vertikal

var fruits [4]string

// cara horizontal
fruits  = [4]string{"apple", "grape", "banana", "melon"}

// cara vertikal
fruits  = [4]string{
    "apple",
    "grape",
    "banana",
    "melon",
}


// A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen

var numbers = [...]int{2, 3, 2, 4, 3}

fmt.Println("data array \t:", numbers)
fmt.Println("jumlah elemen \t:", len(numbers))


// A.15.5. Array Multidimensi

var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

fmt.Println("numbers1", numbers1)
fmt.Println("numbers2", numbers2)


// A.15.6. Perulangan Elemen Array Menggunakan Keyword for

var fruits = [4]string{"apple", "grape", "banana", "melon"}

for i := 0; i < len(fruits); i++ {
    fmt.Printf("elemen %d : %s\n", i, fruits[i])
}


// A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range

var fruits = [4]string{"apple", "grape", "banana", "melon"}

for i, fruit := range fruits {
    fmt.Printf("elemen %d : %s\n", i, fruit)
}


//
// A.15.8. Penggunaan Variabel Underscore _ Dalam for - range

var fruits = [4]string{"apple", "grape", "banana", "melon"}

for i, fruit := range fruits {
    fmt.Printf("nama buah : %s\n", fruit)
}

// Di sinilah salah satu kegunaan variabel pengangguran, atau underscore (_). Tampung saja nilai yang tidak ingin digunakan ke underscore.

var fruits = [4]string{"apple", "grape", "banana", "melon"}

for _, fruit := range fruits {
    fmt.Printf("nama buah : %s\n", fruit)
}


// A.15.9. Alokasi Elemen Array Menggunakan Keyword make
// Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make.

var fruits = make([]string, 2)
fruits[0] = "apple"
fruits[1] = "manggo"

fmt.Println(fruits)  // [apple manggo]