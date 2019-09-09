package main

import (
	"fmt"
	"html/template"
	"io"
	"text/tabwriter"
)

var trackList = template.Must(template.New("tracklist").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="Access-Control-Allow-Origin" content="*">
    <title></title>
</head>
<body>
<h1> Track Info List </h1>
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
<script type="text/javascript">
  var ths = document.getElementsByTagName('th');    
  for(var i = 0; i < ths.length; i++){

      ths[i].onclick = function () {
        var xhr = new XMLHttpRequest();
        var url = "http://localhost:8000/?sort=" + this.innerHTML
        xhr.open('GET', url, true);
        xhr.send(null);
 
        xhr.onreadystatechange = processRequest;
        function processRequest(e) {
          if (xhr.readyState == 4) {
			window.location.reload();
          }
        }
      }   
  	}
</script>
</body>
</html>
`))

func htmlPrint(out io.Writer, tracks []*Track) {
	trackList.Execute(out, tracks)
}

func tabPrint(out io.Writer, tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(out, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
