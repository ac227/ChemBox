package main

import (
	"embed"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed data/*
var f embed.FS

// CalcMr calculates the molecular mass based on the chemical formula
func calcMr(formula string) float64 {
	var elementPattern = regexp.MustCompile("([A-Z][a-z]?)(\\d*)")
	matcher := elementPattern.FindAllStringSubmatch(formula, -1)
	mass := 0.0
	for _, match := range matcher {
		elementSymbol := match[1]
		elementCount, _ := strconv.Atoi(match[2])
		if elementCount == 0 {
			elementCount = 1
		}
		mass += getAr(elementSymbol) * float64(elementCount)
	}
	return mass
}

// GetAr returns the relative atomic mass of an element
func getAr(formula string) float64 {
	ar := 0.0
	arrA := []string{"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne", "Na", "Mg", "Al", "Si", "P", "S", "Cl", "Ar", "K", "Ca", "Sc", "Ti", "V", "Cr", "Mn", "Fe", "Co", "Ni", "Cu", "Zn", "Ga", "Ge", "As", "Se", "Br", "Kr", "Rb", "Sr", "Y", "Zr", "Nb", "Mo", "Tc", "Ru", "Rh", "Pd", "Ag", "Cd", "In", "Sn", "Sb", "Te", "I", "Xe", "Cs", "Ba", "La", "Ce", "Pr", "Nd", "Pm", "Sm", "Eu", "Gd", "Tb", "Dy", "Ho", "Er", "Tm", "Yb", "Lu", "Hf", "Ta", "W", "Re", "Os", "Ir", "Pt", "Au", "Hg", "Tl", "Pb", "Bi", "Po", "At", "Rn", "Fr", "Ra", "Ac", "Th", "Pa", "U", "Np", "Pu", "Am", "Cm", "Bk", "Cf", "Es", "Fm", "Md", "No", "Lr", "Rf", "Db", "Sg", "Bh", "Hs", "Mt", "Ds", "Rg", "Cn", "Fl", "Lv"}
	arrB := []float64{1, 4, 7, 9, 11, 12, 14, 16, 19, 20, 23, 24, 27, 28, 31, 32, 35.5, 40, 39, 40, 45, 48, 51, 52, 55, 56, 59, 59, 64, 65, 70, 73, 75, 79, 80, 84, 85, 87, 89, 91, 93, 96, 98, 101, 103, 106, 108, 112, 115, 119, 122, 128, 127, 131, 133, 137, 139, 140, 141, 144, 145, 150, 152, 157, 159, 162, 165, 167, 169, 173, 175, 178, 181, 184, 186, 190, 192, 195, 197, 201, 204, 207, 209, 209, 210, 222, 223, 226, 227, 232, 231, 238, 239, 243, 245, 247, 249, 253, 254, 259, 260, 261, 264, 269, 270, 273, 274, 272, 278, 283, 282, 287, 291, 295}

	for i, element := range arrA {
		if element == formula {
			ar = arrB[i]
			break
		}
	}
	return ar
}

// GetData retrieves data for a given atomic number and keyword
func getData(atomicNumber int, keyword string) string {
	var result string
	i := 0

	file, err := f.ReadFile("/ChemistryElementInfomation.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "error"
	}

	var data struct {
		Rows []struct {
			Cells []string `xml:"Cell"`
		} `xml:"Row"`
	}
	err = xml.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return "error"
	}

	if atomicNumber-1 < len(data.Rows) {
		switch keyword {
		case "Symbol":
			i = 1
		case "Name":
			i = 2
		case "AtomicMass":
			i = 3
		case "CPKHexColor":
			i = 4
		case "ElectronConfiguration":
			i = 5
		case "Electronegativity":
			i = 6
		case "AtomicRadius":
			i = 7
		case "IonizationEnergy":
			i = 8
		case "ElectronAffinity":
			i = 9
		case "OxidationStates":
			i = 10
		case "StandardState":
			i = 11
		case "MeltingPoint":
			i = 12
		case "BoilingPoint":
			i = 13
		case "Density":
			i = 14
		case "GroupBlock":
			i = 15
		case "YearDiscovered":
			i = 16
		default:
			fmt.Println("Unknown keyword")
			return "error"
		}
		result = data.Rows[atomicNumber-1].Cells[i]
	}
	return result
}

// GetDataInt retrieves integer data for a given atomic number and keyword
func getDataInt(atomicNumber int, keyword string) int {
	result, _ := strconv.Atoi(getData(atomicNumber, keyword))
	return result
}

// GetDataDouble retrieves float64 data for a given atomic number and keyword
func getDataDouble(atomicNumber int, keyword string) float64 {
	result, _ := strconv.ParseFloat(getData(atomicNumber, keyword), 64)
	return result
}

func main() {
	fmt.Println(calcMr("H2O"))
}
