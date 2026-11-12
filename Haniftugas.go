package main

import (
	"fmt"
)

type Laporan interface {
	CetakLaporan() string
}

type Pertandingan struct {
	HomeTeam string
	AwayTeam string
	HomeGoal int
	AwayGoal int
	Stadion  string
}

func (p Pertandingan) TentukanPemenang() string {
	if p.HomeGoal > p.AwayGoal {
		return p.HomeTeam + " Menang"
	} else if p.HomeGoal < p.AwayGoal {
		return p.AwayTeam + " Menang"
	}
	return "Pertandingan Berakhir Seri"
}

func (p Pertandingan) CetakLaporan() string {
	hasil := p.TentukanPemenang()
	return fmt.Sprintf("Stadion: %s | %s [%d - %d] %s --> Hasil: %s",
		p.Stadion, p.HomeTeam, p.HomeGoal, p.AwayGoal, p.AwayTeam, hasil)
}

func main() {

	daftarPertandingan := []Pertandingan{
		{HomeTeam: "Barca", AwayTeam: "Persija", HomeGoal: 2, AwayGoal: 1, Stadion: "Camp Nou"},
		{HomeTeam: "Arema FC", AwayTeam: "Manchester United", HomeGoal: 2, AwayGoal: 2, Stadion: "Kanjuruhan"},
		{HomeTeam: "Semen Padang", AwayTeam: "Liverpool", HomeGoal: 1, AwayGoal: 0, Stadion: " H Agus Salim"},
	}

	fmt.Println("=== HASIL PERTANDINGAN LiGA Tarkam ===")
	for i, match := range daftarPertandingan {

		var laporan Laporan = match
		fmt.Printf("%d. %s\n", i+1, laporan.CetakLaporan())
	}
}
