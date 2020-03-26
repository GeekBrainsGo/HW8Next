{{define "page"}}
{{template "header" .}}
{{range .Data}}
<article class="well">
    <h3>{{.Title}}</h3>
    <small>Created on {{.ID.Timestamp.Local.Format "2-01-2006 15:04:05"}} by {{.Author}}</small>
    <div style="overflow:hidden;max-height:160px">
        <p>{{.Content}}</p>
    </div>
    <a class="btn btn-outline-primary btn-sm" href="/view/{{.ID.Hex}}">read more</a>
</article>
{{end}}
{{template "footer"}}
{{end}}