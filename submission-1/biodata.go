package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func createPerson(person []Person, people ...[]string) []Person {
	for _, p := range people {
		person = append(person, Person{
			nama:      p[0],
			alamat:    p[1],
			pekerjaan: p[2],
			alasan:    p[3],
		})
	}
	return person
}

func getArgs() string {
	return os.Args[1]
}

func printData(args int, person []Person) {
	if args <= len(person) {
		fmt.Printf("Absen-%v \n", args)
		fmt.Println(strings.Repeat("=", 30))
		args -= 1
		fmt.Println("Nama\t\t: ", person[args].nama)
		fmt.Println("Alamat\t\t: ", person[args].alamat)
		fmt.Println("Pekerjaan\t: ", person[args].pekerjaan)
		fmt.Println("Alasan\t\t: ", person[args].alasan)
	} else {
		fmt.Println("Nomor absen tidak valid!")
	}
}

func main() {
	data := [][]string{
		{"Dodi", "Jakarta", "Software Engineer", "I love coding"},
		{"Siti", "Bandung", "Data Scientist", "I love data"},
		{"Ridwan", "Surabaya", "DevOps Engineer", "I love ops"},
		{"Rizky", "Jogja", "System Engineer", "I love system"},
		{"Budi", "Tangerang", "Network Engineer", "I love network"},
		{"Toni", "Bekasi", "Security Engineer", "I love security"},
	}

	person := createPerson([]Person{}, data...)
	absen, err := strconv.Atoi(getArgs())

	if err == nil {
		printData(absen, person)
	} else {
		fmt.Println("Mohon masukkan nomor absen yang benar!")
	}
}
