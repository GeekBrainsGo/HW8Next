{{define "header"}}
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>{{.Title}}</title>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.3.1/cerulean/bootstrap.min.css" media="screen">
        <script src="https://bootswatch.com/_vendor/jquery/dist/jquery.min.js"></script>
        <script src="https://bootswatch.com/_vendor/bootstrap/dist/js/bootstrap.min.js"></script>
    </head>
    <body>
        {{template "navbar" .}}
        <div class="container">
{{end}}

{{define "footer"}}
        </div>
    </body>
</html>
{{end}}

{{define "navbar"}}
        <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
            <a class="navbar-brand" href="#">{{.Title}}</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarColor01">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item">
                    <a class="nav-link" href="/">Home </a>
                    </li>
                    <li class="nav-item">
                    <a class="nav-link" href="/new">Add Post </a>
                    </li>
                </ul>
            </div>
      </nav>
{{end}}


{{define "newscript"}}
<script>
    function createPost() {
        let p = document.querySelector('section[post-id="new"]')
        let title = p.querySelector('input[name="title"]').value
        let author = p.querySelector('input[name="author"]').value
        let content = p.querySelector('textarea[name="body"]').value
        fetch('/api/v1/posts', {
            method: 'POST',
            body: JSON.stringify({
                title,
                author,
                content
            })
            }).then(resp => {
                resp.json().then(data => {
                    window.location = "/view/"+data.id
            })
        })
    }
</script>
{{end}}


{{define "editscript"}}
<script>
    function updatePost(id) {
        let p = document.querySelector(`section[post-id='${id}']`)
        let title = p.querySelector('input[name="title"]').value
        let author = p.querySelector('input[name="author"]').value
        let content = p.querySelector('textarea[name="body"]').value
        let edit = false
        fetch(`/api/v1/posts/${id}`, {
            method: 'PUT',
            body: JSON.stringify({
                id,
                title,
                author,
                edit,
                content
            })
        }).then(resp => {
            resp.json().then(data => {
                window.location = "/view/"+data.id
            })
        })
    }

    function deletePost(id) {
        fetch(`/api/v1/posts/${id}`, {
            method: 'DELETE', 
            body: JSON.stringify({
                id
            })
        }).then(resp => {
                window.location = "/"
        })
    }
</script>
{{end}}