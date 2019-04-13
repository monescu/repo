package main
// Tema 8 - Generati fractali Koch folosind o aplicatie Google GO
// Claudia Marinache - 10LD561
//
// comenzi pentru actualizare pachete folosite
//		go get -u github.com/llgcode/draw2d
//		go get -u github.com/jung-kurt/gofpdf
// lansare in executie:
//		go run tema3.go
// Referinte:
// 		https://godoc.org/github.com/llgcode/draw2d
// 		https://www.geeksforgeeks.org/koch-curve-koch-snowflake/

import (
	"github.com/llgcode/draw2d/draw2dpdf"
	"image/color"
	"math"
	"strconv"
	"fmt"
)

func main() {
	var latime_imagine, inaltime_imagine, dimensiune_fractal float64
	var	ordin_fractal int

	ordin_fractal = 4	//*** ordin_fractal = 4 complexitatea fractalului Koch ***
	
	latime_imagine = 600				// *** rezultatul generarii fractalului va fi salvat intr-un fisier pdf, deoarece suporta format vectorizat, pentru vizualizare de calitate
	inaltime_imagine = latime_imagine	// in format Letter A1, unde A1 este dimensiunea maxima suportata de draw2dpdf cu latimea de aprox 500pixeli

	dimensiune_fractal = 0.8 * inaltime_imagine //dimensiune_fractal = 480.0 dimensiunea in mm a fractalului generat
	
	// Initializare continut grafic
	continut := draw2dpdf.NewPdf("L", "mm", "A1")	// Letter mm format A1
	grafic := draw2dpdf.NewGraphicContext(continut)

	// Setari parametrii de culoare linie si continut, dimensiune linie
	grafic.SetFillColor(color.RGBA{0x00, 0xfc, 0xff, 0xff})    //  R, G, B, ?
	grafic.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})  //  R, G, B, ?
	grafic.SetLineWidth(0.001)

	fulg_de_nea_Koch(dimensiune_fractal, ordin_fractal, grafic, latime_imagine, inaltime_imagine)

	grafic.FillStroke()		// aplica culorile in zonele de arie inchisa

	// Salvare in fisier pdf
	var nume_fisier string 
	nume_fisier = "fractal_Koch_grad"
	nume_fisier += strconv.Itoa(ordin_fractal) + ".pdf"		// Itoa converteste int la string
	draw2dpdf.SaveToPdfFile( nume_fisier, continut)			// Salveaza in fisierul pdf cu numele compus mai sus
	
	fmt.Println()
	fmt.Println(nume_fisier)
}

func fulg_de_nea_Koch(dimensiune_fractal float64, ordin_fractal int, grafic *draw2dpdf.GraphicContext, latime_imagine float64, inaltime_imagine float64){	
	var x0, y0 float64
	x0 = latime_imagine / 2 - dimensiune_fractal / 2 // determina coordonatele de unde se incepe trasarea fractalului in functie de dimensiunea acestuia setata anterior pentru a fi pozitionat in centrul imaginii
	y0 = inaltime_imagine / 2 - math.Sqrt(3) * dimensiune_fractal / 6

	grafic.MoveTo(x0, y0)							// muta cursorul grafic in pozitia initiala de start

    TrasareLinieFractal(dimensiune_fractal, 0, ordin_fractal, grafic)
    TrasareLinieFractal(dimensiune_fractal, 120, ordin_fractal, grafic)
    TrasareLinieFractal(dimensiune_fractal, 240, ordin_fractal, grafic)	
}

func TrasareLinieFractal(lungime float64, unghi float64, ordin_fractal int, grafic *draw2dpdf.GraphicContext){
	if ordin_fractal == 0 {
        TrasareLinie_UnghiDistanta(lungime, unghi, grafic)
    } else {
        TrasareLinieFractal(lungime/3, unghi, ordin_fractal - 1, grafic)
        TrasareLinieFractal(lungime/3, unghi - 60, ordin_fractal - 1, grafic)
        TrasareLinieFractal(lungime/3, unghi + 60, ordin_fractal - 1, grafic)
        TrasareLinieFractal(lungime/3, unghi, ordin_fractal - 1, grafic)
    }
}

func TrasareLinie_UnghiDistanta(r float64, unghi_fi float64, grafic *draw2dpdf.GraphicContext){
	// se folosesc coordonatele polare r(Coordonata radiala) si unghi_fi
    var radiani, xx, yy, LastX, LastY float64

    radiani = unghi_fi / 180 * math.Pi			// converteste din grade in radiani
	LastX, LastY = grafic.LastPoint()			// actualizeaza noile coordonate de unde se continua trasarea unei noi linii
	
	xx = LastX + r * math.Cos(radiani)			// conversie din coordonate polare in coordonate carteziene
	yy = LastY + r * math.Sin(radiani)

	grafic.LineTo(xx, yy)						// traseaza o noua linie din pozitia LastX,LastY in pozitia xx,yy
	fmt.Print(".")								// afiseaza la consola .
}