package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func setContent() *container.Split {

	date_label := widget.NewLabel("Дата")
	doc_label := widget.NewLabel("Ответственный")
	action_label := widget.NewLabel("Событие")

	date_ent := widget.NewEntry()
	date_ent.SetText(timeConvert(time.Now()))

	doc_ent := widget.NewSelect(
		docs,
		func(s string) {
			report.doc = s
		})

	doc_ent.PlaceHolder = "Select"

	action_ent := widget.NewSelect(
		actions,
		func(s string) {
			b, date := checkDate(date_ent.Text)

			if !b {
				dialog.ShowInformation("Ошибка",
					"Введите дату в формате ДД.ММ.ГГГГ",
					w)
				w.Canvas().Focus(date_ent)
				return
			}
			date_ent.Text = date
			date_ent.Refresh()

			if report.doc == "" {
				dialog.ShowInformation("Ошибка",
					"Укажите ответственного",
					w)
				w.Canvas().Focus(doc_ent)
				return
			}

			report.date = date_ent.Text
			switch s {

			case actions[0]:
				{
					// setBoxTMC()
					openFile()
					break
				}
			case actions[1]:
				{
					setBoxONMK()
					break
				}
			case actions[2]:
				{
					setBoxPol()
					break
				}

			case actions[3]:
				{
					setBoxAction()
					break
				}

			}

			report.action = s
		})

	action_ent.PlaceHolder = "Select"

	contUp := container.NewVBox(date_label, date_ent, doc_label, doc_ent, action_label, action_ent)

	list := getEventToday()
	listEvents := widget.NewList(
		func() int { return len(list) },
		func() fyne.CanvasObject { return widget.NewButton("", func() {}) },
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			switch list[id].action {

			case actions[0]:
				{
					obj.(*widget.Button).SetText(fmt.Sprintf("%d) - %s - %s - %s - %s - %s -%s ", id+1, list[id].doc, list[id].number_card, list[id].action, list[id].district, list[id].diagnos, list[id].result))
					obj.(*widget.Button).OnTapped = func() {
						getBoxTMC(list[id])
					}
					break
				}
			case actions[1]:
				{
					obj.(*widget.Button).SetText(fmt.Sprintf("%d) - %s - %s - %s - %s -%s ", id+1, list[id].doc, list[id].number_card, list[id].action, list[id].diagnos, list[id].result))
					obj.(*widget.Button).OnTapped = func() {

						getBoxONMK(list[id])
					}
					break
				}
			case actions[2]:
				{
					obj.(*widget.Button).SetText(fmt.Sprintf("%d) - %s - %s - %s - %s -%s ", id+1, list[id].doc, list[id].number_card, list[id].action, list[id].diagnos, list[id].result))
					obj.(*widget.Button).OnTapped = func() {

						getBoxPol(list[id])

					}
					break
				}

			case actions[3]:
				{
					obj.(*widget.Button).SetText(fmt.Sprintf("%d) - %s - %s - %s - %s -%s ", id+1, list[id].doc, list[id].number_card, list[id].action, list[id].diagnos, list[id].oper_role))
					obj.(*widget.Button).OnTapped = func() {

						getBoxAction(list[id])

					}
					break
				}

			}
			obj.(*widget.Button).Alignment = 1
		},
	)

	content := container.NewVSplit(contUp, listEvents)

	return content

}

func  clearReport() { //удаление данных  структуры Report

	report.action = ""
	report.diagnos = ""
	report.diagnos_text = ""
	report.district = ""
	report.result = ""
	report.number_card = ""
	report.number_cons = ""
	report.oper_name = ""
	report.oper_type = ""
	report.oper_extreme = ""
	report.oper_role = ""
	setContent()
}

func timeConvert(t time.Time) string { // преобразование time.Now()  в строку в формате дд.мм.гггг

	var str string

	str = t.Format("02.01.2006")

	return str
}
func checkDate(date string) (bool, string) {

	if len(date) > 10 {
		return false, date
	}

	now_year := time.Now().Format("06")
	year, _ := strconv.Atoi(now_year)

	date = strings.ReplaceAll(date, ",", ".")
	date = strings.ReplaceAll(date, "\\", ".")
	date = strings.ReplaceAll(date, "/", ".")
	date = strings.ReplaceAll(date, "-", ".")
	date = strings.ReplaceAll(date, "_", ".")
	date = strings.ReplaceAll(date, "*", ".")
	date = strings.ReplaceAll(date, "+", ".")
	date = strings.ReplaceAll(date, "@", ".")
	date = strings.ReplaceAll(date, "!", ".")
	date = strings.ReplaceAll(date, "#", ".")
	date = strings.ReplaceAll(date, "^", ".")

	str := strings.Split(date, ".")

	if len(str) < 2 {
		return false, date
	}

	for _, data := range str {

		strint, err := strconv.Atoi(data)
		if err != nil {
			return false, date
		}
		if strint == 0 {
			return false, date
		}
	}

	if len(str) == 2 {
		str = append(str, "."+now_year)
	}

	if len(str[2]) == 2 && year != 0 {
		temp, _ := strconv.Atoi(str[2])
		if temp > year {
			str[2] = "19" + str[2]
		} else {
			str[2] = "20" + str[2]
		}
	}

	if len(str[0]) > 2 || len(str[1]) > 2 || len(str[2]) != 4 {
		return false, date
	}

	if len(str[0]) < 2 && str[0] != "0" {
		str[0] = "0" + str[0]
	}
	if len(str[1]) < 2 && str[1] != "0" {
		str[1] = "0" + str[1]
	}

	strint0, _ := strconv.Atoi(str[0]) //day
	strint1, _ := strconv.Atoi(str[1]) // month
	strint2, _ := strconv.Atoi(str[2]) //year

	if strint1 > 12 {
		return false, date
	}

	if str[1] == "01" || str[1] == "03" || str[1] == "05" || str[1] == "07" || str[1] == "08" || str[1] == "10" || str[1] == "12" {
		if strint0 > 31 {
			return false, date
		}
	}

	if str[1] == "04" || str[1] == "06" || str[1] == "09" || str[1] == "11" {
		if strint0 > 30 {
			return false, date
		}
	}

	if str[1] == "02" && strint0 > 29 {
		return false, date
	}

	if str[1] == "02" && strint0 == 29 && !check29(strint2) {
		return false, date
	}

	result := str[0] + "." + str[1] + "." + str[2]
	return true, result
}

func check29(year int) bool { // проверка на високосный год

	for i := 1904; i <= year; i += 4 {
		if i == year {
			return true
		}
	}
	return false
}

func GeneratorId() string {

	list := DBGetAllId()

	var l = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	s := make([]rune, 8)

	for i := range s {
		s[i] = l[rand.Intn(len(l))]

		if s[0] == '0' {
			fmt.Println("первый знак равен 0")
			return GeneratorId()
		}
	}

	for _, value := range list {
		if value == string(s) {
			return GeneratorId()
			fmt.Println("Повтор id")
		}
	}
	// fmt.Println("создан id ", s)
	return string(s)

}
