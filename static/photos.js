var container = document.getElementById("container")
var max = document.getElementsByClassName("thumb").length - 1
var cursor = -1
 
function toggle(i) {
	cursor = i	
	console.log("going to", i)
	var url = "static/photos/" + i + ".png";
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

var less = ["ArrowUp", "ArrowLeft", "k", "h"];
var more = ["ArrowDown", "ArrowRight", "j", "l"];

function keylisten(e) {
	var next = cursor
	if (less.includes(e.key)) {
		next = Math.max(cursor - 1, 0)
	} else if (more.includes(e.key)) {
		next = Math.min(cursor + 1, max)
	} 

	if (next !== cursor) {
		toggle(next)
	}
}

window.addEventListener("keydown", keylisten, true)
