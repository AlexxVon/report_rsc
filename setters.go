package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func setBoxAction() {

	var win *widget.PopUp
	var cont *fyne.Container

	oper_number_card := canvas.NewText("Идентификатор пациента", col)

	oper_name := canvas.NewText("Название операции (описание)", col)
	oper_diangnos := canvas.NewText("Диагноз", col)
	oper_extreme := canvas.NewText("Срочность", col)
	oper_role := canvas.NewText("Роль в операции", col)

	oper_number_card_ent := widget.NewEntry()
	oper_number_card_ent.Text = report.number_card
	oper_name_ent := widget.NewEntry()
	oper_name_ent.SetText(report.oper_name)
	oper_diangnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})

	oper_diangnos_ent.PlaceHolder = report.diagnos

	oper_extreme_ent := widget.NewSelect(
		[]string{
			"экстренная",
			"плановая",
		},
		func(s string) {
			report.oper_extreme = s
		})

	oper_extreme_ent.PlaceHolder = report.oper_extreme

	oper_role_ent := widget.NewSelect(
		[]string{
			"оператор",
			"ассистент",
		},
		func(s string) {
			report.oper_role = s
		})

	oper_role_ent.PlaceHolder = report.oper_role

	btn_ok := widget.NewButton("Ок", func() {

		if oper_number_card_ent.Text == "" {
			dialog.ShowInformation("Ошибка", "Укажите Идентификатор пациента", w)
			w.Canvas().Focus(oper_number_card_ent)
			return
		}
		if report.diagnos == "" {
			dialog.ShowInformation("Ошибка", "Укажите диагноз", w)
			w.Canvas().Focus(oper_diangnos_ent)
			return
		}
		if report.oper_extreme == "" {
			dialog.ShowInformation("Ошибка", "Укажите срочность операции", w)
			w.Canvas().Focus(oper_extreme_ent)
			return
		}
		if report.oper_role == "" {
			dialog.ShowInformation("Ошибка", "Укажите роль в операции", w)
			w.Canvas().Focus(oper_role_ent)
			return
		}

		report.oper_name = oper_name_ent.Text
		report.number_card = oper_number_card_ent.Text
		dbSetData(report)

		clearReport()
		dialog.ShowInformation("", "Сохранено успешно", w)
		w.Canvas().Refresh(w.Content())
		win.Hide()
	})

	btn_cancel := widget.NewButton("Cancel", func() {
		win.Hide()
		clearReport()
	})

	btn_box := container.NewHBox(btn_ok, btn_cancel)
	cont = container.NewVBox(
		oper_number_card,
		oper_number_card_ent,
		oper_name,
		oper_name_ent,
		oper_diangnos,
		oper_diangnos_ent,
		oper_extreme,
		oper_extreme_ent,
		oper_role,
		oper_role_ent,
		btn_box,
	)

	win = widget.NewModalPopUp(cont, w.Canvas())
	win.Resize(fyne.NewSize(400, 600))
	win.Show()
}

func setBoxTMC() {
	fmt.Println(data.doc)
	var win *widget.PopUp
	var cont *fyne.Container
	district := canvas.NewText("ПСО", col)
	diagnos := canvas.NewText("Диагноз", col)
	number_card := canvas.NewText("Идентификатор пациента", col)

	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	district_ent := widget.NewSelect(
		districts,
		func(s string) {
			report.district = s
		})
	district_ent.PlaceHolder = report.district

	diagnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})
	diagnos_ent.PlaceHolder = report.diagnos

	diagnos_text_ent := widget.NewEntry()

	result_ent := widget.NewSelect(
		results,
		func(s string) {
			report.result = s
		})

	result_ent.PlaceHolder = report.result
	number_card_ent := widget.NewEntry()
	number_card_ent.Text = report.number_card
	btn_ok := widget.NewButton("Ок", func() {

		if report.diagnos == "" {
			dialog.ShowInformation("Ошибка", "Укажите диагноз", w)
			w.Canvas().Focus(diagnos_ent)
			return
		}
		if report.district == "" {
			dialog.ShowInformation("Ошибка", "Укажите название ПСО", w)
			w.Canvas().Focus(district_ent)
			return
		}
		if report.result == "" {
			dialog.ShowInformation("Ошибка", "Укажите результат", w)
			w.Canvas().Focus(result_ent)
			return
		}
		report.diagnos_text = diagnos_text_ent.Text
		report.number_card = number_card_ent.Text

		if report.doc == "" {
			report.doc = data.doc
		}
		dbSetData(report)

		clearReport()
		win.Hide()
		dialog.ShowInformation("", "Сохранено успешно", w)
	})

	btn_cancel := widget.NewButton("Cancel", func() {
		win.Hide()
		clearReport()
	})

	btn_box := container.NewHBox(btn_ok, btn_cancel)
	cont = container.NewVBox(
		district,
		district_ent,
		number_card,
		number_card_ent,
		diagnos,
		diagnos_ent,
		diagnos_text,
		diagnos_text_ent,
		result,
		result_ent,
		btn_box,
	)

	win = widget.NewModalPopUp(cont, w.Canvas())
	win.Resize(fyne.NewSize(400, 600))

	win.Show()
}

func setBoxONMK() {
	var win *widget.PopUp
	var cont *fyne.Container
	number_card := canvas.NewText("Идентификатор пациента", col)
	diagnos := canvas.NewText("Диагноз", col)
	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	number_card_ent := widget.NewEntry()
	number_card_ent.Text = report.number_card
	diagnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})

	diagnos_text_ent := widget.NewEntry()
	result_ent := widget.NewSelect(
		results,
		func(s string) {
			report.result = s
		})

	result_ent.PlaceHolder = report.result

	btn_ok := widget.NewButton("Ок", func() {

		if report.diagnos == "" {
			dialog.ShowInformation("Ошибка", "Укажите диагноз", w)
			w.Canvas().Focus(diagnos_ent)
			return
		}
		if number_card_ent.Text == "" {
			dialog.ShowInformation("Ошибка", "Укажите Идентификатор пациента", w)
			w.Canvas().Focus(number_card_ent)
			return
		}
		if report.result == "" {
			dialog.ShowInformation("Ошибка", "Укажите результат", w)
			w.Canvas().Focus(result_ent)
			return
		}
		report.diagnos_text = diagnos_text_ent.Text
		report.number_card = number_card_ent.Text
		dbSetData(report)
		clearReport()
		win.Hide()
		dialog.ShowInformation("", "Сохранено успешно", w)
	})

	btn_cancel := widget.NewButton("Cancel", func() {
		win.Hide()
		clearReport()
	})

	btn_box := container.NewHBox(btn_ok, btn_cancel)
	cont = container.NewVBox(
		number_card,
		number_card_ent,
		diagnos,
		diagnos_ent,
		diagnos_text,
		diagnos_text_ent,
		result,
		result_ent,
		btn_box,
	)

	win = widget.NewModalPopUp(cont, w.Canvas())
	win.Resize(fyne.NewSize(400, 600))

	win.Show()
}

func setBoxPol() {
	var win *widget.PopUp
	var cont *fyne.Container
	number_card := canvas.NewText("Идентификатор пациента", col)
	diagnos := canvas.NewText("Диагноз", col)
	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	number_card_ent := widget.NewEntry()
	number_card_ent.Text = report.number_card
	diagnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})

	diagnos_text_ent := widget.NewEntry()
	result_ent := widget.NewSelect(
		results,
		func(s string) {
			report.result = s
		})

	result_ent.PlaceHolder = report.result

	btn_ok := widget.NewButton("Ок", func() {

		if report.diagnos == "" {
			dialog.ShowInformation("Ошибка", "Укажите диагноз", w)
			w.Canvas().Focus(diagnos_ent)
			return
		}
		if number_card_ent.Text == "" {
			dialog.ShowInformation("Ошибка", "Укажите Идентификатор пациента", w)
			w.Canvas().Focus(number_card_ent)
			return
		}
		if report.result == "" {
			dialog.ShowInformation("Ошибка", "Укажите результат", w)
			w.Canvas().Focus(result_ent)
			return
		}
		report.diagnos_text = diagnos_text_ent.Text
		report.number_card = number_card_ent.Text
		dbSetData(report)
		clearReport()
		win.Hide()
		dialog.ShowInformation("", "Сохранено успешно", w)

	})

	btn_cancel := widget.NewButton("Cancel", func() {
		win.Hide()
		clearReport()
	})

	btn_box := container.NewHBox(btn_ok, btn_cancel)
	cont = container.NewVBox(
		number_card,
		number_card_ent,
		diagnos,
		diagnos_ent,
		diagnos_text,
		diagnos_text_ent,
		result,
		result_ent,
		btn_box,
	)

	win = widget.NewModalPopUp(cont, w.Canvas())
	win.Resize(fyne.NewSize(400, 600))

	win.Show()
}

func setMenu() *fyne.MainMenu {

	m1 := fyne.NewMenuItem("", func() {})
	m1_1 := fyne.NewMenuItem("Новая консультация ПСО", func() {
		openFile()
	})

	m2 := fyne.NewMenuItem("Отчет за период", func() {
		showModalWinPeriod()
	})
	m3 := fyne.NewMenuItem("Отчет за сегодня", func() {
		getReportDocPeriod(timeConvert(time.Now()), timeConvert(time.Now()))
	})

	m4 := fyne.NewMenuItem("Отчет за год", func() {
		getReportDocPeriod(fmt.Sprintf("01.01.%v", time.Now().Year()), fmt.Sprintf("31.12.%v", time.Now().Year()))
	})

	m3_1 := fyne.NewMenuItem("Поиск событий",
		func() {
			getEventDocDate()
		})

	menu_1 := fyne.NewMenu("Файл", m1, m1_1)
	menu_2 := fyne.NewMenu("Отчеты", m3, m4, m2)
	menu_3 := fyne.NewMenu("Поиск", m3_1)
	main_menu := fyne.NewMainMenu(menu_1, menu_2, menu_3)

	return main_menu

}

func showModalWinPeriod() {
	var win *widget.PopUp

	ds_label := widget.NewLabel("Дата начала")
	de_label := widget.NewLabel("Дата окончания")
	ds_ent := widget.NewEntry()
	de_ent := widget.NewEntry()

	btn_ok := widget.NewButton("Ok", func() {

		check_ds, ds := checkDate(ds_ent.Text)
		check_de, de := checkDate(de_ent.Text)
		if check_de && check_ds {
			win.Hide()
			fmt.Println("data has downloaded")
			getReportDocPeriod(ds, de)
		} else if !check_ds {
			dialog.ShowInformation("Ошибка", "Неверный формат даты", w)
			w.Canvas().Focus(ds_ent)
		} else {
			dialog.ShowInformation("Ошибка", "Неверный формат даты", w)
			w.Canvas().Focus(de_ent)
		}
	})

	btn_cancel := widget.NewButton("Cancel", func() {
		win.Hide()
	})

	cont3 := container.NewHBox(btn_ok, btn_cancel)
	cont := container.NewVBox(ds_label, ds_ent, de_label, de_ent, cont3)

	win = widget.NewModalPopUp(cont, w.Canvas())
	win.Resize(fyne.NewSize(300, 100))
	win.Show()
}
