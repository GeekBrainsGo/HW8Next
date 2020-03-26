{{define "page"}}
    {{template "header" .}}
    <h3>Post</h3>
    <section>
        <h5>{{.Data.Title}}</h5>
        <small>Created on {{.Data.ID.Timestamp.Local.Format "2-01-2006 15:04:05"}} by {{.Data.Author}}</small>
        <p>{{.Data.Content}}</p>
        <a type="button" class="btn btn-outline-info" href="/edit/{{.Data.ID.Hex}}">Edit</a>
        <a type="button" class="btn btn-outline-primary" href="/">Back</a>
    </section>
    {{template "footer"}}
{{end}}
