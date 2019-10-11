{{define "page"}}
    {{template "header" .}}
    <h3>Post</h3>
    <section post-id="new">
        <article class="form-group">
            <label>Title</label>
            <input required type="text" name="title" class="form-control">
        </article>
        <article class="form-group">
            <label>Author</label>
            <input required type="text" name="author" class="form-control">
        </article>
        <article class="form-group">
            <label>Body</label>
            <textarea required name="body" class="form-control"></textarea>
        </article>
        <button type="button" class="btn btn-outline-info" onclick="createPost()">Create</button>
        <a type="button" class="btn btn-outline-primary" href="/">Back</a>
    </section>
    {{template "newscript"}}
    {{template "footer"}}
{{end}}
