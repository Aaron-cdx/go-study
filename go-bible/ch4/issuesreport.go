/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 11:18
 * @Motto: Keep thinking, keep coding!
when use the go run, can add the dependencies in the suffix, and add the params in the dependencies
and use the > xxx.html can auto-generate the html file.
*/

package main

import (
	"html/template"
	"log"
	"os"
	"study-go/go-bible/github"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.
	Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
