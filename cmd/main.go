package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var (
		output          string
		first           float64
		second          float64
		op              string
		result          float64
		equalPressed    bool
		opButtonPressed bool
	)
	ic, _ := fyne.LoadResourceFromPath("icon.png")
	calcApp := app.New()
	calcWindow := calcApp.NewWindow("poor calc")
	calcApp.SetIcon(ic)
	calcWindow.SetFixedSize(true)
	calcWindow.Resize(fyne.Size{Width: 240, Height: 300})

	input := widget.NewLabel(output)

	buttons := container.NewGridWithColumns(4,
		widget.NewButton("C", func() {
			if len(output) > 0 && !equalPressed {
				output = output[:len(output)-1]
				input.SetText(output)
			}
		}),
		widget.NewButton("±", func() {
			if output != "0" {
				num, _ := strconv.ParseFloat(output, 64)
				num *= -1
				output = strconv.FormatFloat(num, 'f', -1, 64)
				input.SetText(output)
			}
		}),
		widget.NewButton("%", func() {
			if output != "0" {
				num, _ := strconv.ParseFloat(output, 64)
				num /= 100
				output = strconv.FormatFloat(num, 'f', -1, 64)
				input.SetText(output)
			}
		}),
		widget.NewButton("÷", func() {
			first, _ = strconv.ParseFloat(output, 64)
			op = "/"
			output = "0"
			input.SetText(output)
			opButtonPressed = true
		}),

		widget.NewButton("7", appendNum(&output, "7", input)),
		widget.NewButton("8", appendNum(&output, "8", input)),
		widget.NewButton("9", appendNum(&output, "9", input)),
		widget.NewButton("×", func() {
			first, _ = strconv.ParseFloat(output, 64)
			op = "*"
			output = "0"
			input.SetText(output)
			opButtonPressed = true
		}),

		widget.NewButton("4", appendNum(&output, "4", input)),
		widget.NewButton("5", appendNum(&output, "5", input)),
		widget.NewButton("6", appendNum(&output, "6", input)),
		widget.NewButton("-", func() {
			first, _ = strconv.ParseFloat(output, 64)
			op = "-"
			output = "0"
			input.SetText(output)
			opButtonPressed = true
		}),

		widget.NewButton("1", appendNum(&output, "1", input)),
		widget.NewButton("2", appendNum(&output, "2", input)),
		widget.NewButton("3", appendNum(&output, "3", input)),
		widget.NewButton("+", func() {
			first, _ = strconv.ParseFloat(output, 64)
			op = "+"
			output = "0"
			input.SetText(output)
			opButtonPressed = true
		}),

		widget.NewButton("AC", func() {
			output = "0"
			input.SetText(output)
			equalPressed = false
		}),

		widget.NewButton("0", appendNum(&output, "0", input)),
		widget.NewButton(",", func() {
			if !contains(output, ".") {
				output += "."
				input.SetText(output)
			}
		}),
		widget.NewButton("=", func() {
			if op != "" {
				second, _ = strconv.ParseFloat(output, 64)
				switch op {
				case "+":
					result = first + second
				case "-":
					result = first - second
				case "*":
					result = first * second
				case "/":
					if second != 0 {
						result = first / second
					} else {
						output = "Error"
						input.SetText(output)
						return
					}
				}
				output = strconv.FormatFloat(result, 'f', -1, 64)
				input.SetText(output)

				equalPressed = true

				opButtonPressed = false
			}
		}),
	)

	spacer := layout.NewSpacer()
	spacer.Resize(fyne.Size{Width: 240, Height: 10})

	calcWindow.SetContent(container.NewVBox(
		input,
		spacer,
		buttons,
	))

	calcWindow.Show()

	calcWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyReturn:
			if op != "" {
				second, _ = strconv.ParseFloat(output, 64)
				switch op {
				case "+":
					result = first + second
				case "-":
					result = first - second
				case "*":
					result = first * second
				case "/":
					if second != 0 {
						result = first / second
					} else {
						output = "Error"
						input.SetText(output)
						return
					}
				}
				output = strconv.FormatFloat(result, 'f', -1, 64)
				input.SetText(output)

				equalPressed = true

				opButtonPressed = false
			}
		case fyne.Key0, fyne.Key1, fyne.Key2, fyne.Key3, fyne.Key4, fyne.Key5, fyne.Key6, fyne.Key7, fyne.Key8, fyne.Key9:
			if !opButtonPressed {
				appendNum(&output, string(ev.Name), input)()
				opButtonPressed = false
			}
		}
	})

	calcApp.Run()
}

func appendNum(output *string, num string, input *widget.Label) func() {
	return func() {
		if *output == "0" {
			*output = num
		} else {
			*output += num
		}
		input.SetText(*output)
	}
}

func contains(s, substr string) bool {
	return s != "" && substr != "" && len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}
