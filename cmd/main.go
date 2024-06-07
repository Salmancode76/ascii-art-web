package main

import (
	"ascii/PKG"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler) // Register the handler function

	//styles := http.FileServer(http.Dir("stylesheets"))
	//http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", styles))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Ascii string
	}{
		Ascii: "",
	}
	plain_text := ""
	file_op := "standard"
	if r.Method == http.MethodPost {
		plain_text = r.FormValue("txtarea")
		file_op = r.FormValue("op")
		fmt.Println(file_op)

		//in out locations
		arr := [2]string{}
		arr[1] = plain_text
		arr[0] = "../banners/" + file_op + ".txt"
		//entry of txt
		fmt.Print(arr[1])
		for _, v := range arr[1] {
			if v < 32 || v > 126 {
				//log.Fatal("wrong text entered!!! ")
			}
		}

		if len(arr[1]) == 0 {
			return
		}

		sections := strings.Split(arr[1], "\\n")
		if len(sections) >= 2 && sections[0] == "" && sections[1] == "" {
			fmt.Println()
			sections = sections[2:]

		}
		for _, s := range sections {
			if s == "" {
				fmt.Println()
				continue
			}
			cache := [8]string{}
			for _, r := range s {
				PKG.Strings(arr[0], r, &cache)
			}
			data.Ascii = PKG.PrintA(&cache)
		}
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}
