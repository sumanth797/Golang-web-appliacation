package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	Paths, Extension string = "views/layout/", ".gohtml"
)

//creation of viewtype
func NewView(layout string, files ...string) *View {
	// files = append(files,
	// 	"views/layout/bootstrap.gohtml",
	// 	"views/layout/navbar.gohtml",
	// 	"views/layout/footer.gohtml",
	// )
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
func layoutFiles() []string {
	files, err := filepath.Glob(Paths + "*" + Extension)
	if err != nil {
		panic(err)
	}
	return files
}

//createview type that we are going to use
type View struct {
	Template *template.Template
	Layout   string
}

//Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
