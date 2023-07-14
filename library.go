package tes

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
//exported
type segitigaSamaSisi struct {
	alas, tinggi int
}

//exported
type persegiPanjang struct {
	panjang, lebar int
}

//exported
type tabung struct {
	jariJari, tinggi float64
}

//exported
type balok struct {
	panjang, lebar, tinggi int
}

//exported
type hitungBangunDatar interface {
	luas() float64
	keliling() float64
}

//exported
type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

//exported
func (s segitigaSamaSisi) luas() float64 {
	return float64(s.alas*s.tinggi) / 2
}

//exported
func (s segitigaSamaSisi) keliling() float64 {
	return float64(3 * s.alas)
}

//exported
func (p persegiPanjang) luas() float64 {
	return float64(p.panjang * p.lebar)
}

//exported
func (p persegiPanjang) keliling() float64 {
	return float64(2 * (p.panjang + p.lebar))
}

//exported
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

//exported
func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

//exported
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

//exported
func (b balok) luasPermukaan() float64 {
	return float64(2*(b.panjang*b.lebar) + 2*(b.panjang*b.tinggi) + 2*(b.lebar*b.tinggi))
}

// Soal 2
//exported
type phone struct {
	name  string
	brand string
	year  int
	color []string
}

//exported
type phoneData interface {
	displayData() string
}

//exported
func (p phone) displayData() string {
	data := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolor :%s", p.name, p.brand, p.year, strings.Join(p.color, ", "))
	return data
}

// Soal 3
//exported
func luasPersegi(sisi float64, showText bool) interface{} {
	if sisi == 0 {
		if showText {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return "nil"
		}
	}

	luas := sisi * sisi

	if showText {
		result := fmt.Sprintf("Luas persegi dengan sisi %.2f cm adalah %.2f cm", sisi, luas)
		return result
	} else {
		return luas
	}
}

// Soal 4

//exported
func stringifySlice(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, angka := range slice {
		strSlice[i] = fmt.Sprint(angka)
	}
	return strings.Join(strSlice, "+")
}
