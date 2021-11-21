/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 15:02
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
)

var trackTable = template.Must(template.New("Track").Parse(
	`<h1> Tracks </h1>
<table style="border: 1px solid salmon">
    <tr style='text-align: left'>
        <th onclick="submitform('Title')">Title
            <form action="" name="Title" method="post">
                <input type="hidden" name="orderby" value="Title"/>
            </form>
        </th>
        <th>Artist
            <form action="" name="Artist" method="post">
                <input type="hidden" name="orderby" value="Artist"/>
            </form>
        </th>
        <th>Album
            <form action="" name="Album" method="post">
                <input type="hidden" name="orderby" value="Album"/>
            </form>
        </th>
        <th onclick="submitform('Year')">Year
            <form action="" name="Year" method="post">
                <input type="hidden" name="orderby" value="Year"/>
            </form>
        </th>
        <th onclick="submitform('Length')">Length
            <form action="" name="Length" method="post">
                <input type="hidden" name="orderby" value="Length"/>
            </form>
        </th>
    </tr>
{{range .t}}
<tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
</tr>
{{end}}
</table>
	
<script>
	function submitform(formname){
		alert("click the col is:" + name)
		document.getElementsByName(name).submit()
	}
</script>
`))

type TracksTable struct {
	t       []*Track // exported
	primary string
	second  string
	third   string
}

func (x *TracksTable) Len() int {
	return len(x.t)
}

func (x *TracksTable) Less(i, j int) bool {
	key := x.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "Year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "Length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}
		if k == 0 {
			key = x.second
		} else if k == 1 {
			key = x.third
		}
	}
	return false
}

func (x *TracksTable) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func NewTrackTable(t []*Track, primary, second, third string) sort.Interface {
	return &TracksTable{
		t:       t,
		primary: primary,
		second:  second,
		third:   third,
	}
}

func printTrackTables(w io.Writer, x sort.Interface) {
	if x, ok := x.(*TracksTable); ok {
		trackTable.Execute(w, x)
	}
}

func SetPrimary(x sort.Interface, p string) {
	if x, ok := x.(*TracksTable); ok {
		x.primary, x.second, x.third = p, x.primary, x.second
	}
}

func main() {
	trackTables := NewTrackTable(tracks, "Title", "", "")
	sort.Sort(trackTables)

	http.HandleFunc("/sort", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			fmt.Printf("ParseForm: %v\n", err)
		}
		//fmt.Println(request)
		//form := request.Form
		//for k, v := range form {
		//	fmt.Println(k, v)
		//	if k == "orderby" {
		//		SetPrimary(trackTables, v[0])
		//	}
		//}
		t, _ := template.ParseFiles("./xx.html")
		//sort.Sort(trackTables)
		t.Execute(writer, trackTables)
		//printTrackTables(writer, trackTables)
	})

	http.HandleFunc("/sort_by_name", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			fmt.Printf("ParseForm: %v\n", err)
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("hello world"))
		//fmt.Println(request)
		//form := request.Form
		//SetPrimary(trackTables, form.Get("name"))
		//t, _ := template.ParseFiles("./xx.html")
		//sort.Sort(trackTables)
		//t.Execute(writer, trackTables)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
