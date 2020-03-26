{{define "page"}}
    {{template "header" .}}
    <h3>Post</h3>
    <section post-id="{{.Data.ID.Hex}}">
        <article class="form-group">
            <label>Title</label>
            <input required type="text" name="title" class="form-control" value="{{.Data.Title}}">
        </article>
        <article class="form-group">
            <label>Author</label>
            <input required type="text" name="author" class="form-control" value="{{.Data.Author}}">
        </article>
        <article class="form-group">
            <label>Body</label>
            <textarea required name="body" class="form-control">{{.Data.Content}}</textarea>
        </article>
        <button type="button" class="btn btn-outline-info" onclick="updatePost('{{.Data.ID.Hex}}')">Update</button>
        <button type="button" class="btn btn-outline-danger" onclick="deletePost('{{.Data.ID.Hex}}')">Delete</button>
        <a type="button" class="btn btn-outline-primary" href="/">Back</a>
    </section>
    {{template "editscript"}}
    {{template "footer"}}
{{end}}
