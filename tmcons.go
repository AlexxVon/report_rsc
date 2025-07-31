package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"baliance.com/gooxml/document"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jung-kurt/gofpdf"
	"github.com/nguyenthenguyen/docx"
)

func openFile() {

	window := a.NewWindow("Consultation")
	window.Resize(fyne.NewSize(800, 600))

	btn_open := widget.NewButton("Открыть файл запроса консультации", func() {
		dialog.ShowFileOpen(
			func(uc fyne.URIReadCloser, err error) {
				// fmt.Println("func oppening file start")
				// fmt.Println(uc.URI().Path())
				data = getFileData(uc.URI().Path())
				consWin(&data, window)

			}, window)

	},
	)

	btn_empty := widget.NewButton("Пустая консультативная форма", func() {
		var d Data
		d = checkdata(&d)
		consWin(&d, window)

	})

	btn_monkey := widget.NewButton("Заполнить обезьяну", func() {

		// dialog.ShowInformation("Внимание!", "Функция находится в разработке \nдо момент формирования локальной БД, \nс целью защиты личных данных пациентов", window)
		// вызов функции подразумевает внесение данных в базу учета действий врачей по статье
		//"телемедицинские консультации"
		if report.doc == "" {
			window.Close()
			dialog.ShowInformation("Внимание", "Укажите ответственного", w)
		} else {
			setBoxTMC()
			window.Close()
		}

	})

	btn_close := widget.NewButton("Закрыть", func() {
		window.Close()
	})
	cont := container.NewVBox(btn_open, btn_empty, btn_monkey, btn_close)
	window.SetContent(cont)

	window.Show()

}

func consWin(data *Data, w fyne.Window) {

	var win *widget.PopUp

	str := fmt.Sprintf("Анамнез заболевания: %s\nСопутствующая патология: %s \n%s\nСознание: %s\nДыхание: %s ЧДД: %s\nГемодинамика: %s\nВазопрессорная поддержка: %s\nТяжесть по Hunt-Hess: %s",
		data.anamnes,
		data.sop_check,
		data.sop_pat,
		data.concience,
		data.pnoe,
		data.pnoe_count,
		data.ad,
		data.press,
		data.hh,
	)

	var ot_x float32 = 15.0
	var ot_y float32 = 15.0
	size_mono := fyne.NewSize(600, 30)
	size_multi := fyne.NewSize(600, 120)
	size_label := fyne.NewSize(50, 30)

	titel := widget.NewLabel("Консультация нейрохирурга РСЦ")
	titel.Resize(fyne.NewSize(300, 15))
	titel.Move(fyne.NewPos(150, 15))

	ent_date := widget.NewEntry()
	ent_date.SetText(time.Now().Format("02.01.2006 15.04"))

	label_fio := widget.NewLabel("ФИО пациента")
	label_fio.Resize(size_label)
	label_fio.Move(fyne.NewPos(float32(ot_x), float32(titel.Position().Y+titel.MinSize().Height)+float32(ot_y)))

	ent_fio := widget.NewEntry()
	ent_fio.Resize(size_mono)
	ent_fio.Move(fyne.NewPos(label_fio.Position().X+label_fio.MinSize().Width+float32(ot_y), label_fio.Position().Y))
	ent_fio.SetText(data.fio)

	label_pso := widget.NewLabel("Находится на лечении в ")
	label_pso.Resize(size_label)
	label_pso.Move(fyne.NewPos(float32(ot_x), label_fio.MinSize().Height+label_fio.Position().Y+float32(ot_y)))

	ent_pso := widget.NewSelect(
		districts,
		func(s string) {
			data.pso = s
		})

	ent_pso.Resize(size_mono)
	ent_pso.Move(fyne.NewPos(label_pso.Position().X+label_pso.MinSize().Width+float32(ot_x), label_pso.Position().Y))
	ent_pso.PlaceHolder = data.pso

	label_pasp := widget.NewLabel("Паспортные данные")
	label_pasp.Resize(size_label)
	label_pasp.Move(fyne.NewPos(float32(ot_x), label_pso.Position().Y+label_pso.MinSize().Height+float32(ot_y)))

	ent_pasp := widget.NewMultiLineEntry()
	ent_pasp.Resize(size_multi)
	ent_pasp.Move(fyne.NewPos(label_pasp.Position().X+label_pasp.MinSize().Width+float32(ot_x), label_pasp.Position().Y))

	ent_pasp.SetText(fmt.Sprintf("паспорт: %s\nполис ОМС:  %s\nСНИЛС:  %s", data.seria, data.polis, data.snils))

	label_about_pacient := widget.NewLabel("Сведения о пациенте")
	label_about_pacient.Resize(size_label)
	label_about_pacient.Move(fyne.NewPos(ot_x, ot_y+float32(ent_pasp.Position().Y)+float32(ent_pasp.Size().Height)))

	ent_about_pacient := widget.NewMultiLineEntry()
	ent_about_pacient.Resize(size_multi)
	ent_about_pacient.Move(fyne.NewPos(label_about_pacient.Position().X+label_about_pacient.MinSize().Width+ot_x, label_about_pacient.Position().Y))
	ent_about_pacient.SetText(str)

	label_nerv := widget.NewLabel("Неврологический статус")
	label_nerv.Resize(size_label)
	label_nerv.Move(fyne.NewPos(ot_x, ent_about_pacient.Position().Y+ent_about_pacient.Size().Height+ot_y))

	ent_nerv := widget.NewMultiLineEntry()
	ent_nerv.Resize(size_multi)
	ent_nerv.Move(fyne.NewPos(ot_x+ot_x+label_nerv.Size().Width, label_nerv.Position().Y))
	ent_nerv.SetText(data.nerv)

	label_ct := widget.NewLabel("Нейровизуализация")
	label_ct.Resize(size_label)
	label_ct.Move(fyne.NewPos(ot_x, ent_nerv.Position().Y+ent_nerv.Size().Height+ot_y))

	ent_ct := widget.NewMultiLineEntry()
	ent_ct.Resize(size_multi)
	ent_ct.Move(fyne.NewPos(ot_x+ot_x+label_ct.Size().Width, label_ct.Position().Y))
	ent_ct.SetText(data.ct)

	label_contact := widget.NewLabel("Обратная связь")
	ent_contact := widget.NewMultiLineEntry()
	ent_contact.SetText(data.contacts)

	label_report := widget.NewLabel("Заключение")
	ent_report := widget.NewMultiLineEntry()
	ent_report.Wrapping = fyne.TextWrapOff

	label_rec := widget.NewLabel("Рекомендации")
	ent_rec := widget.NewMultiLineEntry()

	docs_cons := docs
	docs_cons = append(docs_cons, "консилиум")
	label_doc := widget.NewLabel("Консультант")
	ent_cons := widget.NewMultiLineEntry()
	ent_cons_item := widget.NewFormItem("", ent_cons)

	ent_doc := widget.NewSelect(
		docs_cons,
		func(s string) {
			if s != "консилиум" {
				data.doc = s
			} else {
				d := dialog.NewForm("Укажите состав консилиума",
					"Ok",
					"Закрыть",
					[]*widget.FormItem{ent_cons_item},
					func(b bool) {
						data.doc = ent_cons.Text
					},
					w)

				d.Show()
			}
		},
	)

	ent_doc.SetSelected(report.doc)

	btn_ok := widget.NewButton("Сформировать документ", func() {

		if ent_report.Text == "" {
			dialog.ShowInformation("Ошибка", "Поле не должно быть пустым", w)
			w.Canvas().Focus(ent_report)
			return
		}
		if ent_rec.Text == "" {
			dialog.ShowInformation("Ошибка", "Поле не должно быть пустым", w)
			w.Canvas().Focus(ent_rec)
			return
		}
		if data.doc == "" {
			dialog.ShowInformation("Ошибка", "Укажите ответственного", w)
			w.Canvas().Focus(ent_doc)
			return
		}

		data.recommends = ent_rec.Text
		data.report = ent_report.Text
		data.fio = ent_fio.Text
		data.pso = ent_pso.Selected
		data.seria = ent_pasp.Text
		data.anamnes = ent_about_pacient.Text
		data.nerv = ent_nerv.Text
		data.ct = ent_ct.Text
		data.date_cons = ent_date.Text

		path := "\\\\172.30.106.3\\share\\Нейрохирург.отд\\report_data_rsc\\TMcons"
		err := format_consultation(path, data)

		if err != nil {
			dialog.ShowError(err, w)
		} else {
			exec.Command("cmd", "/C", "explorer ", path).Run()
			win.Hide()
		}

	})

	btn_close := widget.NewButton("Закрыть", func() {
		win.Hide()
	})

	btn_box := container.NewHBox(btn_ok, btn_close)

	cont := container.NewVBox(titel,
		ent_date,
		label_fio,
		ent_fio,
		label_pso,
		ent_pso,
		label_pasp,
		ent_pasp,
		label_about_pacient,
		ent_about_pacient,
		label_nerv,
		ent_nerv,
		label_ct,
		ent_ct,
		label_contact,
		ent_contact,
		label_report,
		ent_report,
		label_rec,
		ent_rec,
		label_doc,
		ent_doc,
		btn_box)
	cont2 := container.NewVScroll(cont)
	win = widget.NewModalPopUp(cont2, w.Canvas())
	win.Resize(fyne.NewSize(800, 800))
	win.Show()
}

func getFileData(path string) Data {
	var data Data
	var text map[int]string
	text = make(map[int]string)

	pso := districts

	concience := [...]string{
		"Ясное",
		"Умеренное оглушение",
		"Глубокое оглушение",
		"Сопор",
		"Кома",
	}

	pnoe := []string{
		"Самостоятельное",
		"ИВЛ",
	}

	r, err := docx.ReadDocxFile(path)

	if err != nil {
		fmt.Println(err)
	}
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)

	// Or read from a filesystem object:
	// r, err := docx.ReadDocxFromFS(file string, fs fs.FS)

	defer r.Close()

	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	alldoc := docx1.GetContent()

	var str []string
	dropdown := strings.Split(alldoc, "FORMDROPDOWN")

	for i := 0; i < len(dropdown); i++ {
		if strings.Contains(dropdown[i], "списокпсо") {
			str = strings.Split(dropdown[i], "result w:val=\"")
			if len(str) > 1 && len(str[1]) > 0 {
				run1 := string(str[1][0])
				runint, _ := strconv.Atoi(run1)
				data.pso = pso[runint]
			}
		}
		if strings.Contains(dropdown[i], "списоксознание") {
			str = strings.Split(dropdown[i], "result w:val=\"")
			if len(str) > 1 && len(str[1]) > 0 {
				run1 := string(str[1][0])
				runint, _ := strconv.Atoi(run1)
				data.concience = concience[runint]
			}
		}
		if strings.Contains(dropdown[i], "списокдыхание") {
			str = strings.Split(dropdown[i], "result w:val=\"")
			if len(str) > 1 && len(str[1]) > 0 {
				run1 := string(str[1][0])
				runint, _ := strconv.Atoi(run1)
				data.pnoe = pnoe[runint]
			}
		}

		if strings.Contains(dropdown[i], "списокHuntHess") {
			str = strings.Split(dropdown[i], "result w:val=\"")
			if len(str) > 1 && len(str[1]) > 0 {
				run1 := string(str[1][0])
				data.hh = run1
			}
		}
	}

	checkbox := strings.Split(alldoc, "FORMCHECKBOX")

	for i := 0; i < len(checkbox); i++ {

		if strings.Contains(checkbox[i], "ФлажокАГ") && !strings.Contains(checkbox[i], "checked w:val") {
			data.sop_check += " Артериальная гипертония "
			continue
		}
		if strings.Contains(checkbox[i], "ФлажокИБС") && !strings.Contains(checkbox[i], "checked w:val") {
			data.sop_check += " ИБС "
			continue
		}
		if strings.Contains(checkbox[i], "ФлажокСД") && !strings.Contains(checkbox[i], "checked w:val") {
			data.sop_check += " СД "
			continue
		}
		if strings.Contains(checkbox[i], "ФлажокЖир") && !strings.Contains(checkbox[i], "checked w:val") {
			data.sop_check += " Ожирение "
			continue
		}

		if strings.Contains(checkbox[i], "ФлажокПрессоры") && !strings.Contains(checkbox[i], "checked w:val") {
			data.press = " да "
		}
		if strings.Contains(checkbox[i], "ФлажокПрессоры") && strings.Contains(checkbox[i], "checked w:val") {
			data.press = " нет "
		}

		// fmt.Println(data.press)
	}

	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	for _, tables := range doc.Tables() {

		for i, row := range tables.Rows() {

			for j, cell := range row.Cells() {
				if j == 1 {
					for _, para := range cell.Paragraphs() {

						for _, run := range para.Runs() {
							// fmt.Println(i, "row", j, "  cell ", run.Text())
							text[i] += run.Text()
						}
					}
				}
			}

		}

		data.ad = text[14]
		data.anamnes = text[8]
		// data.concience = text[24]

		data.contacts = text[22]

		data.ct = text[20]

		data.date = text[2]

		// data.date_pasp = text[12]

		data.fio = text[1]

		// data.hh = text[35]
		data.nerv = text[19]
		// data.number = text[11]
		// data.pnoe = text[25]
		data.pnoe_count = text[13]
		data.polis = text[5]
		// data.ps = text[30]
		data.questions = text[21]
		data.seria = text[4]
		data.snils = text[6]
		data.sop_pat = text[10]

	}

	data = checkdata(&data)
	return data
}

func checkdata(data *Data) Data {
	str := "нет данных"
	if data.ad == "" {
		data.ad = str
	}
	if data.anamnes == "" {
		data.anamnes = str
	}
	if data.concience == "" {
		data.concience = str
	}
	if data.contacts == "" {
		data.contacts = str
	}
	if data.ct == "" {
		data.ct = str
	}
	if data.date == "" {
		data.date = str
	}
	if data.date_pasp == "" {
		data.date_pasp = str
	}
	if data.fio == "" {
		data.fio = str
	}
	if data.hh == "" {
		data.hh = str
	}
	if data.nerv == "" {
		data.nerv = str
	}
	if data.number == "" {
		data.number = str
	}
	if data.pnoe == "" {
		data.pnoe = str
	}
	if data.pnoe_count == "" {
		data.pnoe_count = str
	}
	if data.polis == "" {
		data.polis = str
	}
	if data.ps == "" {
		data.ps = str
	}
	if data.pso == "" {
		data.pso = str
	}
	if data.questions == "" {
		data.questions = str
	}
	if data.seria == "" {
		data.seria = str
	}
	if data.snils == "" {
		data.snils = str
	}
	if data.sop_pat == "" {
		data.sop_pat = str
	}
	return *data
}

func format_consultation(path string, data *Data) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	text := 11.0
	// text_small := 9.0
	// text_big := 12.0

	pdf.AddFont("Helvetica", "", "font\\roman regular.json")
	pdf.AddFont("Helvetica", "B", "font\\roman bold.json")
	pdf.AddFont("Helvetica", "I", "font\\roman italic.json")
	pdf.AddFont("Helvetica", "BI", "font\\roman bolditalic.json")

	pdf.SetFillColor(0x99, 0x99, 0x99)
	pdf.SetFont("Helvetica", "", text)
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("font\\cp1251")
	ht := 4.5
	w := 270.0

	pdf.CellFormat(60, ht, tr(data.date_cons), "0", 0, "L", false, 0, "")
	pdf.SetFont("Helvetica", "BUI", text)
	pdf.CellFormat(60, ht+1, tr("Консультация нейрохирурга РСЦ"), "0", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)
	str := fmt.Sprintf("Ф.И.О. %s        Дата рождения %s", data.fio, data.date)
	pdf.SetFont("Helvetica", "", text)
	pdf.Cell(0, ht, tr(str))
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Паспортные данные"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.seria), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Анамнез"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.anamnes), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Неврологический статус"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.nerv), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Нейровизулизация"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.ct), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Заключение"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.report), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Рекомендации"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.recommends), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.SetFont("Helvetica", "BUI", text)
	pdf.Cell(0, ht, tr("Консультант"))
	pdf.Ln(-1)
	pdf.SetFont("Helvetica", "", text)
	pdf.MultiCell(w, ht, tr(data.doc), "0", "L", false)
	pdf.Ln(-1)
	pdf.Ln(-1)

	f := strings.Split(data.fio, " ")[0]

	str = path + "\\" + "TM_cons_" + f + "_" + data.date_cons + ".pdf"

	err := pdf.OutputFileAndClose(str)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
