package main

import (
	"log"
	"os"
	"time"

	"github.com/raifpy/KaraAmbarKart"
)

func Xmain() {
	kart, _ := KaraAmbarKart.YeniKart()

	h, _ := os.Open("homer.png")
	defer h.Close()

	res, err := kart.AycicekYagi("Homer", "İvedik", "", h)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("_.png")
	KaraAmbarKart.Buf(res).WriteTo(f)
	f.Close()

}

func main() {
	t := time.Now()
	Xmain()
	log.Printf("%d ms", time.Since(t).Milliseconds())
}
