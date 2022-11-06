package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello, World!\n")
}

func postArticleHandler(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"Posting Article...\n")
}

func articleListHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Article List\n")
}

func articleDetailHandler(w http.ResponseWriter, req *http.Request)  {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func postNiceHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Posting Nice...\n")
}

func postCommentHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "Posting Comment...\n")
}