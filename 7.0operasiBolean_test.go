/*
Operasi Bolean
Operator		Keterangan
===========================
&&				Dan
||				Atau
!				Kebalikan

Operasi &&
nilai_1		Operator	nilai_2		Hasil
-------------------------------------------
true		&&			true		true
true		&&			false		false
false		&&			true		false
false		&&			false		false
jika true && true hasilnya true , selain itu hasilnya false

Operasi ||
nilai_1		Operator	nilai_2		Hasil
-------------------------------------------
true		||			true		true
true		||			false		true
false		||			true		true
false		||			false		false

Operasi !
Operator	nilai_2		Hasil
-------------------------------
!			true		false
!			false		true
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestOperasiBolean(t *testing.T) {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 75

	var statusLulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(statusLulus)

	fmt.Println(nilaiAkhir > 80 && absensi > 75)
}
