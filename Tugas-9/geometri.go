package geometri

import (
	"math"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type hitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s segitigaSamaSisi) Luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) Keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) Luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) Keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) Volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) Volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) LuasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

func DisplayInfo(d displayData) string {
	return d.Display()
}