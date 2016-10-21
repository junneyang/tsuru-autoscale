// Copyright 2016 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"html/template"
	"net/http"

	"github.com/ajg/form"
	"github.com/gorilla/mux"
	"github.com/tsuru/tsuru-autoscale/alarm"
)

func alarmHandler(w http.ResponseWriter, r *http.Request) error {
	t, err := template.ParseFiles("web/templates/alarm/list.html")
	if err != nil {
		return err
	}
	a, err := alarm.FindAlarmBy(nil)
	if err != nil {
		return err
	}
	return t.Execute(w, a)
}

func alarmDetailHandler(w http.ResponseWriter, r *http.Request) error {
	t, err := template.ParseFiles("web/templates/alarm/detail.html")
	if err != nil {
		return err
	}
	vars := mux.Vars(r)
	a, err := alarm.FindAlarmByName(vars["name"])
	if err != nil {
		return err
	}
	return t.Execute(w, a)
}

func alarmAdd(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return err
		}
		var a alarm.Alarm
		d := form.NewDecoder(nil)
		d.IgnoreCase(true)
		d.IgnoreUnknownKeys(true)
		err = d.DecodeValues(&a, r.Form)
		if err != nil {
			return err
		}
		err = alarm.NewAlarm(&a)
		if err != nil {
			return err
		}
		http.Redirect(w, r, "/web/alarm", 302)
		return nil
	}
	t, err := template.ParseFiles("web/templates/alarm/add.html")
	if err != nil {
		return err
	}
	return t.Execute(w, nil)
}
