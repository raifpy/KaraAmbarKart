package KaraAmbarKart

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"

	"io"

	"github.com/fogleman/gg"
	"github.com/goki/freetype/truetype"
	"github.com/nfnt/resize"
)

//go:embed font.ttf
var font []byte

//go:embed kart.png
var kart []byte

type FontSecenek truetype.Options

type Secenekler struct {
	Renk color.Color
	Font []byte

	FontAyarlari *FontSecenek

	IsimX float64
	IsimY float64

	SoyIsimX float64
	SoyIsimY float64

	UyeNoX float64
	UyeNoY float64
}

type Kart struct {
	secenekler *Secenekler
	Font       *truetype.Font
	Kart       image.Image
}

func YeniKart(secenekler ...Secenekler) (kart_ *Kart, err error) {
	secenek := &Secenekler{}
	if len(secenekler) != 0 {
		secenek = &secenekler[0]
	}

	varsayilanSecenekler(secenek)

	kart_ = &Kart{secenekler: secenek}

	kart_.Kart, _, err = image.Decode(bytes.NewReader(kart))
	if err != nil {
		return
	}

	kart_.Font, err = truetype.Parse(secenek.Font)
	return
}
func (k *Kart) ZeytinYagi(isim, soyisim, uyeno string, photo image.Image) image.Image {
	ctx := gg.NewContextForImage(k.Kart)
	ctx.DrawImage(k.Kart, 0, 0)
	ctx.SetColor(k.secenekler.Renk)
	ctx.SetFontFace(truetype.NewFace(k.Font, (*truetype.Options)(k.secenekler.FontAyarlari)))

	ctx.DrawString(isim, k.secenekler.IsimX, k.secenekler.IsimY)
	ctx.DrawString(soyisim, k.secenekler.SoyIsimX, k.secenekler.SoyIsimY)
	ctx.DrawString(uyeno, k.secenekler.UyeNoX, k.secenekler.UyeNoY)
	ctx.DrawImage(ResizeImage(photo), 650, 203)
	return ctx.Image()
}

func (k *Kart) AycicekYagi(isim, soyisim, uyeno string, photo io.Reader) (image.Image, error) {
	im, _, err := image.Decode(photo)
	if err != nil {
		return nil, err
	}
	return k.ZeytinYagi(isim, soyisim, uyeno, im), nil
}

func (k *Kart) TereYagi(isim, soyisim, uyeno string, photo []byte) (image.Image, error) {
	im, _, err := image.Decode(bytes.NewReader(photo))
	if err != nil {
		return nil, err
	}
	return k.ZeytinYagi(isim, soyisim, uyeno, im), nil
}

func varsayilanSecenekler(s *Secenekler) {
	if s.Renk == nil {
		s.Renk = color.Black
	}

	if s.Font == nil {
		s.Font = font
	}

	if s.FontAyarlari == nil {
		s.FontAyarlari = &FontSecenek{Size: 40}
	}

	if s.IsimX <= 0 || s.IsimY <= 0 {
		s.IsimX = 160
		s.IsimY = 435
	}

	if s.SoyIsimX <= 0 || s.SoyIsimY <= 0 {
		s.SoyIsimX = 240
		s.SoyIsimY = 475
	}

	if s.UyeNoX <= 0 || s.UyeNoY <= 0 {
		s.UyeNoX = 245
		s.UyeNoY = 510
	}

}
func ResizeImage(im image.Image) image.Image {
	return resize.Resize(220, 314, im, resize.Bilinear)
}

func Buf(i image.Image) *bytes.Buffer {
	buf := &bytes.Buffer{}
	png.Encode(buf, i)
	return buf
}
