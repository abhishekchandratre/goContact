package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var fileName string = ""

//Render gohtml
func renderTemplate(w http.ResponseWriter, tmpl_str string, node *map[string][]string) {
	tmpl := make(map[string]*template.Template)
	tmpl_str = "../html/" + tmpl_str		//Path to html folder
	tmpl[tmpl_str] = template.Must(template.ParseFiles(tmpl_str, "../html/base.html"))
	tmpl[tmpl_str].ExecuteTemplate(w, "base", node)
}

func indexHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		test := map[string][]string{
			"contacts" : {"1;Abhishek Chandratre;+17049068013;abhishek.chandratre@gmail.com;516 barton creek drive, APT E;4", "2;Tejas Konduri;+17049068013;tejas.konduri@gmail.com;516 barton creek drive, APT E;3"},
		}
		renderTemplate(w, "index.gohtml", &test)
	}
}

func listContactHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		test := map[string][]string{
			"contacts" : {"1;Abhishek Chandratre;+17049068013;abhishek.chandratre@gmail.com;516 barton creek drive, APT E;0", "2;Tejas Konduri;+17049068013;tejas.konduri@gmail.com;516 barton creek drive, APT E;0"},
		}
		renderTemplate(w, "list.gohtml", &test)
	}
}

func addContactHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
			renderTemplate(w, "addContact.gohtml", nil)
		}
		if r.Method == "POST"{
			r.ParseForm()
			//Get input value of csv path
			name := r.Form["name"][0]
			phoneNo := r.Form["phoneNo"][0]
			email := r.Form["email"][0]
			address := r.Form["address"][0]

			var record string = name + ";" + phoneNo + ";" + email + ";" + address + ";0"
			fmt.Println("Adding record" + record)
			http.Redirect(w,r,"/listContact",http.StatusSeeOther)
		}
	}
}

func pathHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
			renderTemplate(w, "inputCsvFile.gohtml", nil)
		}
		if r.Method == "POST"{
			r.ParseForm()
			//Get input value of csv path
			fileName := r.Form["csvPath"][0]
			fmt.Println("Adding file "+ fileName + " For contact listing")
			http.Redirect(w,r,"/listContact",http.StatusSeeOther)
		}
	}
}

func editContactHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST"{
			r.ParseForm()
			//Get input value of csv path
			id := r.Form["numId"][0]
			fmt.Println("Editing ID[" + id +"].")
			http.Redirect(w,r,"/listContact",http.StatusSeeOther)
		}
	}
}

func deleteContactHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST"{
			r.ParseForm()
			//Get input value of csv path
			id := r.Form["numId"][0]
			fmt.Println("Deleting ID[" + id +"].")
			http.Redirect(w,r,"/listContact",http.StatusSeeOther)
		}
	}
}

func main() {
	fmt.Println("Starting Application...");
	http.HandleFunc("/addPath",pathHandler())
	http.HandleFunc("/addContact",addContactHandler())
	http.HandleFunc("/listContact", listContactHandler())
	http.HandleFunc("/deleteContact", deleteContactHandler())
	http.HandleFunc("/editContact", editContactHandler())
	http.HandleFunc("/",indexHandler())
	//Start listening
	http.ListenAndServe(":8080",nil)
}
