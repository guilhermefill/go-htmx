package pages

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/html/base.html", "public/html/components/navigation.html", "public/html/pages/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/html/base.html", "public/html/components/navigation.html", "public/html/pages/about.html")
	data := map[string]interface{}{
		"Title": "About",
		"Content": map[string]interface{}{
			"Name":        "About",
			"Description": "This is the about page",
		},
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/html/base.html", "public/html/components/navigation.html", "public/html/pages/contact.html")
	data := map[string]interface{}{
		"Title": "Contact",
		"Content": map[string]interface{}{
			"Name":        "Contact",
			"Description": "This is the contact page",
		},
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
