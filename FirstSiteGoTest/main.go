package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<!doctype html>
<html>
<head>
	<title>First page in GoLang</title>
<style type="text/css">
*{
	margin: 0;
	padding: 0;
}
#site{
	width:1200px;
	margin:0 auto;
	border: 5px solid red;
}
header{
	border: 2px solid blue;
	height: 300px;
}
nav{
	border: 2px solid green;
	height: 50px;
}
footer{
	border: 2px solid yellow;
	height: 300px;
}
</style>
</head>
<div id="site">
<header>
<h1>Header - First page</h1>
</header>
<nav>
	<h1>NavBar</h1>
</nav>
<div>
<h1>Content</h1>
</div>
<footer>
<h1>Content</h1>
</footer>
</div>
</html>
`)
}
