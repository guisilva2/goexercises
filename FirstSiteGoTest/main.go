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
body {
	font-family: Tahoma;
}
.topbar{
	height:60px;
	background: #054B7C;
	border-bottom: 5px solid #1492CD;
	width: 100%;
	border-radius: 0 0 18px 18px;
}
nav{
	width: 1200px;
	margin: 0 auto;
	height: 60px;
}
.logo {
	float: left;
	width: 150px;
	background: #ccc;
	display: block;
	color: white;
	height: 60px;
}
.menu{
	float:right;
	margin: 0;
	padding: 0;
	list-style-type: none;
}
.menu li {
	background: #054B7C;
	border-right: 1px solid #043759;
	display: inline-block;
	float: left;
}
li:first-child{
	border-left: 1px solid #043759;
}
.menu li a {
	padding: 20px;
	display: block;
	color: white;
	text-decoration: none;
}
.menu li a:hover {
	color: #ccc;
	text-decoration: underline;
	background: #043759;
}
header{
	height: 500px;
	background: url('https://cdn.lynda.com/static/landing/images/hero/CloudDeveloper_1200x630-1503427707377.jpg');
	background-position: 50% 50%;
	background-size: cover;
	width:100%;
	padding-top:35px;
}
.header-box{
	padding:50px;
	text-align: center;
	font-size:30px;
	width:600px;
	height: 200px;
	background: #054B7C;
	border-radius: 25px;
	opacity: 0.8;
	margin: 75px auto;
}
.content{
	margin: 0 auto;
	height: 800px;
	padding: 15px;
	width:1200px;
	box-shadow: 0 18px 20px #ccc;
}
.footer{
	height: 300px;
	background: #054B7C;
	width:100%;
	border-top: 18px solid #1492CD;
}
</style>
</head>
<body>
<div class="topbar">
	<nav>
		<div class="logo">
		<h3></h3>
		</div>

	 	<ul class="menu">
	 	<li><a href="#">Link 1</a></li>
	 	<li><a href="#">Link 2</a></li>
	 	<li><a href="#">Link 3</a></li>
	 	<li><a href="#">Link 4</a></li>
	 	<li><a href="#">Link 5</a></li>
	 </ul>
 	</nav>
</div>

<header>
	<div class="header-box">
	Welcome!<br/>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Phasellus euismod massa nec nulla varius aliquam,  varius aliquam.</p></div>
</header>

<div class="content">
<h1>Content</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Phasellus euismod massa nec nulla varius aliquam,  varius aliquam.</p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Phasellus euismod massa nec nulla varius aliquam,  varius aliquam.</p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Phasellus euismod massa nec nulla varius aliquam,  varius aliquam.</p>
</div>
<div class="footer"><!-- <footer>  -->
&nbsp;
</div>
	
</body>
</html>
`)
}
