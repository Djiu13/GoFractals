package main

import (
	"net/http"
	"html/template"
	"fracture"
	"fracture/algorithm"
	"strconv"
)

type formValue struct {
	A, B float64
	Ox, Oy, P float64
	W, H int
	Action, FractaleName string
}

const (
	PAS_ZOOMIN  float64 = 0.9
	PAS_ZOOMOUT float64 = 0.9
)

func getAction(r *http.Request) string {
	actions := [...]string{"MoveLeft", "MoveDown", "MoveUp", "MoveRight", "ZoomIn", "ZoomOut"}

	for _, a := range actions {
		if r.FormValue(a) != "" {
			return a
		}
	}

	return ""
}

func getFormValue(r *http.Request) *formValue {
	values := new(formValue)
	v, err := strconv.ParseFloat(r.FormValue("ox"), 64)
	if err == nil {
		values.Ox = v
	}
	v, err = strconv.ParseFloat(r.FormValue("oy"), 64)
	if err == nil {
		values.Oy = v
	}
	v, err = strconv.ParseFloat(r.FormValue("p"), 64)
	if err == nil {
		values.P = v
	}
	v, err = strconv.ParseFloat(r.FormValue("a"), 64)
	if err == nil {
		values.A = v
	}
	v, err = strconv.ParseFloat(r.FormValue("b"), 64)
	if err == nil {
		values.B = v
	}
	var vInt int
	vInt, err = strconv.Atoi(r.FormValue("w"))
	if err == nil {
		values.W = vInt
	}
	vInt, err = strconv.Atoi(r.FormValue("h"))
	if err == nil {
		values.H = vInt
	}

	values.FractaleName = r.FormValue("fractaleName")
	values.Action = getAction(r)

	return values
}

func getFractaleRoutine(fractaleName string, c complex128) algorithm.Routine {
	if fractaleName == "MandelBrot" {
		return nil //algorithm.MandelbrotRoutine(-0.8 + 0.156i)
	}
	if fractaleName == "Julia" {
		return algorithm.GetJuliaRoutine(c)
	}

	return nil
}

func handleAction(value *formValue) {
	if value.Action == "MoveRight" {
		value.Ox += value.P*100
	} else if value.Action == "MoveLeft" {
		value.Ox -= value.P*100
	} else if value.Action == "MoveDown" {
		value.Oy -= value.P*100
	} else if value.Action == "MoveUp" {
		value.Oy += value.P*100
	} else if value.Action == "ZoomIn" {
		print("Before ZoomIn: ", value.P, "\n")
			value.P *= PAS_ZOOMIN/2
		print("After ZoomIn: ", value.P, "\n")
	} else if value.Action == "ZoomOut" {
		value.P *= PAS_ZOOMOUT*2
	}
}

func generateImage(value formValue) {
	c := complex(value.A, value.B)
		routine := getFractaleRoutine(value.FractaleName, c)
		if routine == nil {
		print("routine is null !!\n, value.FractaleName: ", value.FractaleName, "\n")
	}
	fracture.GetImage(routine, value.W, value.H, value.Ox, value.Oy, value.P, "image.png")
}

func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("index.html")
	indexTemplate.Execute(w, r)
}

func viewPNG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, "image.png")
}

/*
func mandelbrot(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "fractale.html")
}
*/

func print_values(v *formValue) {
	print("--------\n")
		print("ox: ", v.Ox, "\n")
		print("v.Oy: ", v.Oy, "\n")
		print("v.P: ", v.P, "\n")
		print("v.W: ", v.W, "\n")
		print("v.H: ", v.H, "\n")
		print("v.FractaleName: ", v.FractaleName, "\n")
		print("v.Action: ", v.Action, "\n")
		print("--------\n")
}

func mandelbrot(w http.ResponseWriter, r *http.Request) {
	values := getFormValue(r)
	print("Before HandleAction")
	print_values(values)
		handleAction(values)
	print("After HandleAction")
	print_values(values)
		generateImage(*values)
    t, err := template.ParseFiles("fractale.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, values)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/image.png", viewPNG)
	http.HandleFunc("/mandelbrot", mandelbrot)
    http.ListenAndServe(":8080", nil)
}
