package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"remoteblog/models"
	"remoteblog/server"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testCases = []struct {
	description string
	post        models.Post
}{
	{
		"post test",
		models.Post{Title: "1", Author: "2", Content: "3"},
	},
}

func TestServer(t *testing.T) {
	msg := `
	Description: %s
	Expected: %q
	Got: %q
	`
	lg := logrus.New()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		lg.WithError(err).Fatal("can't get new client")
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}

	if err = client.Ping(ctx, nil); err != nil {
		lg.WithError(err).Fatal("can't ping to db")
	}

	db := client.Database("testblog")

	testserver := server.New(ctx, nil, nil, db)
	for _, test := range testCases {
		// post test
		j, _ := json.Marshal(test.post)
		request := httptest.NewRequest(http.MethodPost, "/api/v1/posts", bytes.NewReader(j))
		response := httptest.NewRecorder()
		testserver.PostHandle(response, request)

		post := &models.Post{}
		json.NewDecoder(response.Body).Decode(post)
		request.Body.Close()
		response.Body.Reset()

		if post.Title != test.post.Title ||
			post.Author != test.post.Author ||
			post.Content != test.post.Content {
			t.Fatalf(msg, test.description, test.post, post)
		}
		tmp := template.HTML(test.post.Title)
		test.post.Title = test.post.Author
		test.post.Author = string(test.post.Content)
		test.post.Content = tmp
		test.post.ID = post.ID

		// put test
		j, _ = json.Marshal(test.post)
		request = httptest.NewRequest(http.MethodPut, "/api/v1/posts/"+test.post.ID.Hex(), bytes.NewReader(j))
		testserver.PutHandle(response, request)

		if h := response.Result().StatusCode; h != http.StatusOK {
			t.Fatalf(msg, http.StatusOK, h)
		}
		request.Body.Close()
		response.Body.Reset()

		// delete test
		request, _ = http.NewRequest(http.MethodDelete, "/api/v1/posts/"+test.post.ID.Hex(), bytes.NewReader(j))
		testserver.DeleteHandle(response, request)
		if h := response.Result().StatusCode; h != http.StatusOK {
			t.Fatalf(msg, http.StatusOK, h)
		}
		request.Body.Close()
		response.Body.Reset()
	}
	client.Disconnect(ctx)
}
