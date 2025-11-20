package main

import (
	"fmt"

	go_say_hello "github.com/Akuutauf-Portofolio/go-say-hello/v2"
)

// get dependencies
// cara untuk mengambil dependencies bisa menggunakan perintah go get <nama-modules>
// contoh : go get github.com/Akuutauf-Portofolio/go-say-hello
// namun pastikan repository modules yang ingin ditambahkan itu bersifat public

// upgrade dependencies 
// bisa dilakukan dengan mengubah version dependencies ke yang terbaru (v1.0.0)
// kemudian lakukan update di terminal dengan perintah 'go get'\
// sekecil apapun perubahan, maka nanti versinya juga harus diperbarui, karena kalau tidak ketika tidak diperbarui
// maka tidak akan bisa di update dependencies nya (nyangkut ke versi yang lama)

func main() {
	fmt.Println(go_say_hello.SayHello("Taufik"))
}