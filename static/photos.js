var container = document.getElementById("container")
function toggle(e){
	var url = "static/photos/" + e.id + ".png";
	var show = document.querySelector(".show")
	if (show !== null && show.id !== url) {
		show.classList.remove("show")
	}
	var img = document.getElementById(url);
	if (img === null) {
		img = document.createElement("img");
		img.classList.add("large", "show")
		img.id = url
		img.src = url
		container.appendChild(img)
	} else {
		img.classList.toggle("show")
	}
}

