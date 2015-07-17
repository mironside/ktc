index.html
<div id="artist"><h2>JOHN OAKS</h2><br><h3> ETCHER &amp; SCREEN PRINTER</h3></div>
<div id="new">
<h5> New Work </h5>
<a href="VeiledThoughts"><img src="files/thoughts-cropped2.jpg"></a></div>
<div id="portfolio_container">
	<div id="portfolio_gallery">
		{{range .Files "jpg" "gif"}}
			<div class="item gallery"><img src="{{.}}"/></div>
		{{end}}
	</div>
</div>
