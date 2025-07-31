package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
)

type Report struct {
	id           int
	date         string
	doc          string
	action       string
	diagnos      string
	diagnos_text string
	district     string
	result       string
	number_card  string
	number_cons  string
	oper_name    string
	oper_type    string
	oper_extreme string
	oper_role    string
	date_full    string
}

type ReportDoc struct {
	id          int
	cons_tmc    int
	cons_stac   int
	cons_pol    int
	oper_oper   int
	oper_assist int
}

var report Report

var a fyne.App
var w fyne.Window

var actions = []string{}

var docs = []string{}

var districts = []string{}

var diagnosis = []string{}

var results = []string{}

var col = color.NRGBA{0, 255, 255, 255}

//
//
//
//
//
//
//
//
//
//

func list(path string) []string {

	var list []string

	f, err := os.Open(path)

	defer f.Close()

	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		list = append(list, sc.Text())
	}

	for i, data := range list {
		fmt.Sprintf("i=%d, data = %s", i, data)
	}

	return list
}

func load_data() {
	docs = list("list_doc.dat")
	districts = list("list_pso.dat")
	actions = list("list_actions.dat")
	diagnosis = list("list_diagnos.dat")
	results = list("list_results.dat")
}


type Data struct {
	fio        string
	pso        string
	date       string
	date_cons  string
	seria      string
	number     string
	date_pasp  string
	anamnes    string
	sop_pat    string
	concience  string
	pnoe       string
	pnoe_count string
	ad         string
	ps         string
	hh         string
	nerv       string
	ct         string
	questions  string
	contacts   string
	polis      string
	snils      string
	sop_check  string
	press      string
	report     string
	recommends string
	doc        string
}

var data Data