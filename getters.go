package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func getBoxAction(report Report) {

	var win *widget.PopUp
	var cont *fyne.Container

	date := canvas.NewText("Дата", col)
	oper_number_card := canvas.NewText("Идентификатор пациента", col)

	oper_name := canvas.NewText("Название операции (описание)", col)
	oper_diangnos := canvas.NewText("Диагноз", col)
	oper_extreme := canvas.NewText("Срочность", col)
	oper_role := canvas.NewText("Роль в операции", col)

	date_ent := widget.NewEntry()
	date_ent.SetText(report.date)

	oper_number_card_ent := widget.NewEntry()
	oper_number_card_ent.SetText(report.number_card)

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
		check, date := checkDate(date_ent.Text)
		if !check {
			dialog.ShowInformation("Ошибка",
				"Неверный формат даты",
				w)
			w.Canvas().Focus(date_ent)
			return
		}
		date_ent.SetText(date)
		date_ent.Refresh()
		report.date = date_ent.Text
		report.oper_name = oper_name_ent.Text
		report.number_card = oper_number_card_ent.Text
		dbUpdateData(report)

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
		date,
		date_ent,
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
	win.Resize(fyne.NewSize(400, 400))
	win.Show()

}

func getBoxTMC(report Report) {

	var win *widget.PopUp
	var cont *fyne.Container
	date := canvas.NewText("Дата", col)
	district := canvas.NewText("ПСО", col)
	number_card := canvas.NewText("Идентификатор пациента", col)

	diagnos := canvas.NewText("Диагноз", col)
	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	date_ent := widget.NewEntry()
	date_ent.SetText(report.date)
	number_card_ent := widget.NewEntry()
	number_card_ent.SetText(report.number_card)
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
	diagnos_text_ent.SetText(report.diagnos_text)

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
		report.date = date_ent.Text
		report.number_card = number_card_ent.Text
		dbUpdateData(report)
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
		date,
		date_ent,
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
	win.Resize(fyne.NewSize(400, 400))

	win.Show()
}

func getBoxONMK(report Report) {
	var win *widget.PopUp
	var cont *fyne.Container
	date := canvas.NewText("Дата", col)
	number_card := canvas.NewText("Идентификатор пациента", col)
	diagnos := canvas.NewText("Диагноз", col)
	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	date_ent := widget.NewEntry()
	date_ent.SetText(report.date)
	number_card_ent := widget.NewEntry()
	number_card_ent.SetText(report.number_card)
	diagnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})
	diagnos_ent.PlaceHolder = report.diagnos

	diagnos_text_ent := widget.NewEntry()
	diagnos_text_ent.SetText(report.diagnos_text)
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
		report.date = date_ent.Text
		dbUpdateData(report)
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
		date,
		date_ent,
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

func getBoxPol(report Report) {
	var win *widget.PopUp
	var cont *fyne.Container
	date := canvas.NewText("Дата", col)
	number_card := canvas.NewText("Идентификатор пациента", col)
	diagnos := canvas.NewText("Диагноз", col)
	diagnos_text := canvas.NewText("Диагноз (описание)", col)
	result := canvas.NewText("Результат", col)

	date_ent := widget.NewEntry()
	date_ent.SetText(report.date)

	number_card_ent := widget.NewEntry()
	number_card_ent.SetText(report.number_card)

	diagnos_ent := widget.NewSelect(
		diagnosis,
		func(s string) {
			report.diagnos = s
		})
	diagnos_ent.PlaceHolder = report.diagnos

	diagnos_text_ent := widget.NewEntry()
	diagnos_text_ent.SetText(report.diagnos_text)
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
		report.date = date_ent.Text
		dbUpdateData(report)
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
		date,
		date_ent,
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

func getReportDocAll() {

	listAll := getEventAll()
	listDoc := make(map[string]ReportDoc)
	var res ReportDoc

	for _, doc := range docs {

		var count_tmc, count_stac, count_pol, count_oper_oper, count_oper_assist int

		for _, data := range listAll {
			if doc == data.doc {

				switch data.action {
				case actions[0]:
					count_tmc++
					break
				case actions[1]:
					count_stac++
					break
				case actions[2]:
					count_pol++
					break
				case actions[3]:
					{
						if data.oper_role == "оператор" {
							count_oper_oper++
							break
						} else {
							count_oper_assist++
							break
						}
					}

				}
			}

		}

		res.cons_tmc = count_tmc
		res.cons_stac = count_stac
		res.cons_pol = count_pol
		res.oper_oper = count_oper_oper
		res.oper_assist = count_oper_assist
		listDoc[doc] = res
	}

	// for i, data := range listDoc {
	// 	fmt.Println(i, data)
	// }
}

// func getReportDocPeriod(date_start, date_end string) {

// 	time_s, err := time.Parse("02.01.2006", date_start)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	time_e, err := time.Parse("02.01.2006", date_end)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	listAll := getEventAll()
// 	listDoc := make(map[string]ReportDoc)
// 	var res ReportDoc

// 	for _, doc := range docs {

// 		var count_tmc, count_stac, count_pol, count_oper_oper, count_oper_assist int

// 		for _, data := range listAll {
// 			if doc == data.doc {
// 				date_tmp, err := time.Parse("02.01.2006", data.date)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				switch data.action {
// 				case actions[0]:
// 					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
// 						count_tmc++
// 						break
// 					}
// 				case actions[1]:
// 					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

// 						count_stac++
// 						break
// 					}
// 				case actions[2]:
// 					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

// 						count_pol++
// 						break
// 					}
// 				case actions[3]:
// 					{
// 						if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
// 							if data.oper_role == "оператор" {
// 								count_oper_oper++
// 								break
// 							} else {
// 								count_oper_assist++
// 								break
// 							}

// 						}
// 					}

// 				}
// 			}

// 		}

// 		res.cons_tmc = count_tmc
// 		res.cons_stac = count_stac
// 		res.cons_pol = count_pol
// 		res.oper_oper = count_oper_oper
// 		res.oper_assist = count_oper_assist
// 		listDoc[doc] = res
// 	}

// 	for i, data := range listDoc {
// 		fmt.Println(i, data)
// 	}

// 	showReport()
// }

func getReportDocPeriod(date_start, date_end string) {

	time_s, err := time.Parse("02.01.2006", date_start)
	if err != nil {
		fmt.Println(err)
	}

	time_e, err := time.Parse("02.01.2006", date_end)
	if err != nil {
		fmt.Println(err)
	}

	listAll := getEventAll()

	var listDoc [8][7]string
	var count_docs [6]int
	var count_actions [6]int

	for i := 0; i < 1; i++ {
		for j := 0; j <= len(actions)+2; j++ {
			if j == 0 {
				listDoc[i][j] = "ФИО"
			} else if j == len(actions)+1 {
				listDoc[i][j] = "Ассистенции"

			} else if j == len(actions)+2 {
				listDoc[i][j] = "Всего"

			} else if j > 0 && j <= len(actions) {
				listDoc[i][j] = actions[j-1]
			}
		}

	}

	for i := 0; i <= len(docs)+1; i++ {
		if i > 0 && i <= len(docs) {
			listDoc[i][0] = docs[i-1]
		}
		if i == len(docs)+1 {
			listDoc[i][0] = "Всего"
		}
	}

	for i := 0; i < len(docs); i++ {

		var count_tmc, count_stac, count_pol, count_oper_oper, count_oper_assist int

		for j := 0; j < len(listAll); j++ {
			if docs[i] == listAll[j].doc {
				date_tmp, err := time.Parse("02.01.2006", listAll[j].date)
				if err != nil {
					fmt.Println(err)
				}
				switch listAll[j].action {
				case actions[0]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
						count_tmc++
						break
					}
				case actions[1]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

						count_stac++
						break
					}
				case actions[2]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

						count_pol++
						break
					}
				case actions[3]:
					{
						if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
							if listAll[j].oper_role == "оператор" {
								count_oper_oper++
								break
							} else {
								count_oper_assist++
								break
							}

						}
					}

				}
			}
			listDoc[i+1][1] = strconv.Itoa(count_tmc)
			listDoc[i+1][2] = strconv.Itoa(count_stac)
			listDoc[i+1][3] = strconv.Itoa(count_pol)
			listDoc[i+1][4] = strconv.Itoa(count_oper_oper)
			listDoc[i+1][5] = strconv.Itoa(count_oper_assist)
		}
		count_docs[i] = count_tmc + count_stac + count_pol + count_oper_oper + count_oper_assist
		count_actions[0] += count_tmc
		count_actions[1] += count_stac
		count_actions[2] += count_pol
		count_actions[3] += count_oper_oper
		count_actions[4] += count_oper_assist
		count_actions[5] += count_tmc + count_stac + count_pol + count_oper_oper + count_oper_assist

	}

	for i := len(docs) + 1; i < len(docs)+2; i++ {
		for j := 0; j < len(actions)+2; j++ {
			if j == len(actions)+2 {
				listDoc[i][j] = strconv.Itoa(count_actions[5])
			} else {
				listDoc[i][j+1] = strconv.Itoa(count_actions[j])
			}
		}
	}

	for i := len(actions) + 2; i <= len(actions)+2; i++ {
		for j := 0; j < len(docs); j++ {
			listDoc[j+1][i] = strconv.Itoa(count_docs[j])
		}
	}

	showReport(listDoc, date_start, date_end)
}

func getReportDoc(date_start, date_end string) {
	time_s, err := time.Parse("02.01.2006", date_start)
	if err != nil {
		fmt.Println(err)
	}

	time_e, err := time.Parse("02.01.2006", date_end)
	if err != nil {
		fmt.Println(err)
	}

	listAll := getEventAll()
	listDoc := make(map[string]ReportDoc)
	var res ReportDoc

	for _, doc := range docs {

		var count_tmc, count_stac, count_pol, count_oper_oper, count_oper_assist int

		for _, data := range listAll {
			if doc == data.doc {
				date_tmp, err := time.Parse("02.01.2006", data.date)
				if err != nil {
					fmt.Println(err)
				}
				switch data.action {
				case actions[0]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
						count_tmc++
						break
					}
				case actions[1]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

						count_stac++
						break
					}
				case actions[2]:
					if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {

						count_pol++
						break
					}
				case actions[3]:
					{
						if (time_s.Before(date_tmp) || time_s.Equal(date_tmp)) && (time_e.After(date_tmp) || time_e.Equal(date_tmp)) {
							if data.oper_role == "оператор" {
								count_oper_oper++
								break
							} else {
								count_oper_assist++
								break
							}

						}
					}

				}
			}

		}

		res.cons_tmc = count_tmc
		res.cons_stac = count_stac
		res.cons_pol = count_pol
		res.oper_oper = count_oper_oper
		res.oper_assist = count_oper_assist
		listDoc[doc] = res
	}

}

func showReport(arr [8][7]string, ds, de string) {
	win := a.NewWindow("report")

	d_start := canvas.NewText("Отчет деятельности нейрохирургов РСЦ с ", color.Opaque)
	d_end := canvas.NewText(" по ", color.Opaque)
	ds_label := canvas.NewText(ds, col)
	de_label := canvas.NewText(de, col)

	var cont *fyne.Container

	cont_date := container.NewHBox(d_start, ds_label, d_end, de_label)

	tabel := widget.NewTable(
		func() (int, int) { return len(arr), len(arr[0]) },
		func() fyne.CanvasObject { return widget.NewLabel("01234567801111213146541") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(arr[tci.Row][tci.Col])
		},
	)
	tabel.Resize(fyne.NewSize(800, 800))
	btn := widget.NewButton("Закрыть", func() {
		win.Close()
	})

	cont_date.Move(fyne.NewPos(10, 0))
	tabel.Move(fyne.NewPos(10, cont_date.Position().Y+35))
	tabel.Resize(fyne.NewSize(float32(len(arr[0])*220), float32(len(arr)*40)))

	btn.Move(fyne.NewPos(10, tabel.Position().Y+tabel.Size().Height+30))
	btn.Resize(fyne.NewSize(150, 40))
	cont = container.NewWithoutLayout(cont_date, tabel, btn)

	win.SetContent(cont)

	win.SetFullScreen(true)
	win.Show()

}

func getEventDocDate() {
	var mod *widget.PopUp
	var s_doc, s_date, s_diagnos string
	s_diagnos = ""
	var rep []Report
	var d []string
	var cont3 *fyne.Container
	for i := 0; i < len(docs); i++ {
		d = append(d, docs[i])
	}
	d = append(d, "Все")
	label_doc := widget.NewLabel("Выберите ответственного")
	sel_doc := widget.NewSelect(
		d,
		func(s string) {
			s_doc = s
		},
	)

	l_diagnos := widget.NewLabel("Диагноз")
	ent_diagnos := widget.NewSelect(
		diagnosis,
		func(s string) {
			s_diagnos = s
		})
	label_date := widget.NewLabel("Укажите дату")
	ent_date := widget.NewEntry()
	ent_date.SetText(timeConvert(time.Now()))
	label_notFound := widget.NewLabel("Не найдено")
	label_notFound.Hidden = true
	btn_search := widget.NewButton("Найти", func() {
		if ent_date.Text != "" {
			b, date := checkDate(ent_date.Text)

			if !b {
				dialog.NewInformation("Неверный формат даты", "Укажите дату в формате ДД.ММ.ГГГГ", w).Show()
				w.Canvas().Focus(ent_date)

				return
			} else {
				s_date = date
			}
		} else {
			s_date = ""
		}
		if s_doc == "" {
			dialog.NewInformation("Ошибка", "Выберите врача из списка", w).Show()
			w.Canvas().Focus(sel_doc)
			return
		}

		rep = *dbgetReports(s_doc, s_date, s_diagnos)
		if len(rep) > 0 {
			showList(rep, s_doc, s_date)
			label_notFound.Hidden = true
		} else {
			label_notFound.Hidden = false
			cont3.Refresh()
		}

	})

	btn_close := widget.NewButton("Закрыть", func() {
		mod.Hide()
	})

	cont1 := container.NewVBox(label_doc, sel_doc)
	cont2 := container.NewVBox(label_date, ent_date)
	cont4 := container.NewVBox(l_diagnos, ent_diagnos)
	cont3 = container.NewHBox(cont1, cont2, cont4, btn_search, btn_close)

	cont3.Resize(fyne.NewSize(600, 40))
	label_notFound.Resize(fyne.NewSize(100, 30))
	label_notFound.Move(fyne.NewPos(cont3.Position().X, cont3.Position().Y+cont3.Size().Height+30))
	cont := container.NewWithoutLayout(cont3, label_notFound)
	mod = widget.NewModalPopUp(cont, w.Canvas())
	mod.Resize(fyne.NewSize(600, 600))
	mod.Show()
}

func showList(rep []Report, doc, date string) {
	var mod *widget.PopUp

	var list *widget.List

	list = widget.NewList(
		func() int { return len(rep) + 1 },
		func() fyne.CanvasObject { return widget.NewButton("", func() {}) },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(rep) {
				obj.(*widget.Button).SetText(fmt.Sprintf("%d) %s Дата: %s Тип события: %s\nID: %s %s %s %s", id+1, rep[id].doc, rep[id].date, rep[id].action, rep[id].number_card, rep[id].diagnos, rep[id].oper_name, rep[id].oper_role))
				obj.(*widget.Button).Alignment = 1
				obj.(*widget.Button).OnTapped = func() {
					switch rep[id].action {
					case actions[0]:
						{
							getBoxTMC(rep[id])
						}
						break

					case actions[1]:
						{
							getBoxONMK(rep[id])
						}
						break

					case actions[2]:
						{
							getBoxPol(rep[id])
							break
						}

					case actions[3]:
						{
							getBoxAction(rep[id])

						}
						break
					}

				}
			}

		},
	)

	btn_close := widget.NewButton("Закрыть", func() {
		mod.Hide()
	})

	btn_refresh := widget.NewButton("Обновить", func() {
		mod.Hide()
		showList(*dbgetReports(doc, date, ""), doc, date)
	})

	str := "Всего записей: " + strconv.Itoa(len(rep))

	cont_btn := container.NewHBox(btn_close, btn_refresh, widget.NewLabel(str))
	contl := container.NewVScroll(list)
	cont_btn.Resize(fyne.NewSize(30, 50))
	contl.Resize(fyne.NewSize(570, 480))
	contl.Move(fyne.NewPos(10, cont_btn.Position().Y+cont_btn.Size().Height+10))
	cont := container.NewWithoutLayout(cont_btn, contl)
	mod = widget.NewModalPopUp(cont, w.Canvas())
	mod.Resize(fyne.NewSize(600, 600))
	mod.Show()

}

