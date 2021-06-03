package views

import "html/template"

//creation of viewtype
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layout/bootstrap.gohtml",
		"views/layout/navbar.gohtml",
		"views/layout/footer.gohtml",
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//createview type that we are going to use
type View struct {
	Template *template.Template
	Layout   string
}
