package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbOpen() *sql.DB {

	db, err := sql.Open("mysql", "reportrsc_broughttop:ef71f2e1559c3a32ed14756453d2af5a733d3429@tcp(is6.h.filess.io:3306)/reportrsc_broughttop")
	if err != nil {
		panic(err)
	}
	return db
}

func dbOpenCarotid() *sql.DB {
	db, err := sql.Open("mysql", "carotids_womenuncle:e305c9455c88334f6f860750443f54c87b0b2492@tcp(7lu6h.h.filess.io:3307)/carotids_womenuncle")
	if err != nil {
		fmt.Println(err)
	}

	return db

}

func dbSetData(data Report) {
	db := dbOpen()
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `report` (`date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role`) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", report.date, report.doc, report.action, report.diagnos, report.diagnos_text, report.district, report.result, report.number_card, report.number_cons, report.oper_name, report.oper_type, report.oper_extreme, report.oper_role))
	defer insert.Close()
	if err != nil {
		panic(err)
	}

	if data.diagnos == diagnosis[6] && (data.result == results[1] || data.result == results[2] || data.result == results[4]) {
		db_carotids := dbOpenCarotid()
		defer db_carotids.Close()

		t := time.Now()
		str := t.Format("2006-01-02 15:04:05")
		id_pacient := GeneratorId()
		// fmt.Println(id_pacient)
		// fmt.Println(fmt.Sprintf("INSERT INTO `pacients` (`datecreate`, `surname`, `idPacient`, `fromhospital`, `diagnosis`,`lastchange` ) VALUES ('%s','%s', '%d','%s','%s', '%s')", str, data.id, id_pacient, data.district, data.diagnos_text, data.doc))

		ins, err := db_carotids.Query(fmt.Sprintf("INSERT INTO `pacients` (`datecreate`, `surname`, `idPacient`, `fromhospital`, `diagnosis`,`lastchange` ) VALUES ('%s','%s', '%s','%s','%s', '%s')", str, data.number_card, id_pacient, data.district, data.diagnos_text, data.doc))
		fmt.Println("carotids")
		defer ins.Close()
		if err != nil {
			fmt.Println(err)
		}
	}

	clearReport()
	// getEventToday()
	// w.SetContent(setContent())
}

func getEventToday() []Report {

	var reportsList []Report

	var id int = 0
	var temp Report

	db := dbOpen()

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE `date`='" + timeConvert(time.Now()) + "' "))

	defer res.Close()

	if err != nil {
		panic(err)
	}

	for res.Next() {

		result := res.Scan(&temp.id, &temp.date, &temp.doc, &temp.action, &temp.diagnos, &temp.diagnos_text, &temp.district, &temp.result, &temp.number_card, &temp.number_cons, &temp.oper_name, &temp.oper_type, &temp.oper_extreme, &temp.oper_role)

		if result != nil {
			panic(err)
		}

		reportsList = append(reportsList, temp)
		id++
	}

	return reportsList
}

func getEventAll() []Report {
	var listAll []Report
	db := dbOpen()
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report`"))
	defer res.Close()

	if err != nil {
		panic(err)
	}

	var temp Report

	for res.Next() {
		result := res.Scan(&temp.id, &temp.date, &temp.doc, &temp.action, &temp.diagnos, &temp.diagnos_text, &temp.district, &temp.result, &temp.number_card, &temp.number_cons, &temp.oper_name, &temp.oper_type, &temp.oper_extreme, &temp.oper_role)
		if result != nil {
			panic(err)
		}
		listAll = append(listAll, temp)
	}

	return listAll
}

func dbUpdateData(report Report) {
	db := dbOpen()
	defer db.Close()

	up, err := db.Query("UPDATE report SET `date`='"+report.date+"', `doc`='"+report.doc+"', `action`='"+report.action+"', `diagnos`='"+report.diagnos+"', `diagnos_text`='"+report.diagnos_text+"', `district`='"+report.district+"', `result`='"+report.result+"', `number_card`='"+report.number_card+"',  `number_cons`='"+report.number_cons+"', `oper_name`='"+report.oper_name+"', `oper_type`='"+report.oper_type+"', `oper_extreme`='"+report.oper_extreme+"', `oper_role`='"+report.oper_role+"' WHERE `id`=?", report.id)
	if err != nil {
		panic(err)
	}
	defer up.Close()
	getEventToday()
	w.SetContent(setContent())
}

func dbgetReports(doc, date, diagnos string) *[]Report {

	var rep []Report
	var temp Report
	var res *sql.Rows
	var err error

	db := dbOpen()
	defer db.Close()

	if doc == "Все" && date != "" && diagnos != "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE `date`='" + date + "' AND `diagnos` = '" + diagnos + "'"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc == "Все" && date == "" && diagnos == "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report`"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc == "Все" && date == "" && diagnos != "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE `diagnos` = '" + diagnos + "'"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc == "Все" && date != "" && diagnos == "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE `date`='" + date + "'"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc != "Все" && date != "" && diagnos != "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE `date`='" + date + "' AND `doc`='" + doc + "' AND `diagnos` = '" + diagnos + "'"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc != "Все" && date == "" && diagnos != "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE  `doc`='" + doc + "'  AND `diagnos` = '" + diagnos + "' ORDER BY `id`"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc != "Все" && date != "" && diagnos == "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE  `doc`='" + doc + "'  AND `date`='" + date + "' ORDER BY `id`"))
		if err != nil {
			fmt.Println(err)
		}

	} else if doc != "Все" && date == "" && diagnos == "" {
		res, err = db.Query(fmt.Sprintf("SELECT `id`, `date`, `doc`, `action`, `diagnos`, `diagnos_text`, `district`, `result`, `number_card`,  `number_cons`, `oper_name`, `oper_type`, `oper_extreme`, `oper_role` FROM `report` WHERE  `doc`='" + doc + "' ORDER BY `id`"))
		if err != nil {
			fmt.Println(err)
		}

	}

	for res.Next() {
		result := res.Scan(&temp.id, &temp.date, &temp.doc, &temp.action, &temp.diagnos, &temp.diagnos_text, &temp.district, &temp.result, &temp.number_card, &temp.number_cons, &temp.oper_name, &temp.oper_type, &temp.oper_extreme, &temp.oper_role)
		if err != nil {
			fmt.Println(result)
			panic(result)
		}
		rep = append(rep, temp)
	}
	res.Close()
	return &rep
}

func DBGetAllId() []string {
	var list []string

	db := dbOpenCarotid()
	defer db.Close()

	res, err := db.Query("SELECT `idPacient` FROM pacients")

	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		var v string

		res.Scan(&v)

		list = append(list, v)
	}
	return list

}
