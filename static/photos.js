var less = ["ArrowUp", "ArrowLeft", "k", "h"];
var more = ["ArrowDown", "ArrowRight", "j", "l"];

var container = document.getElementById("container")
var max = document.getElementsByClassName("thumb").length - 1
var cursor = -1
 
function toggle(i) {
	if (typeof i !== "number" || isNaN(i) || i < 0 || i > max) {
		return
	}

	cursor = i	
	console.log("going to", i)
	var url = "static/photos/" + i + ".png";

	// Remove currently shown image, if any
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
		if (!img.classList.contains("show") && history.replaceState) {
			history.replaceState(null, null, " ")
			return
		}
	}

	if (history.replaceState) {
		history.replaceState(null, "Picture #" + i, "#" + i)
	}
}


function keylisten(e) {
	var next = cursor
	if (less.includes(e.key)) {
		next -= 1
	} else if (more.includes(e.key)) {
		next += 1 
	} 

	toggle(next)
}

var hash = window.location.hash
if (hash.length > 0) {
	var id = parseInt(hash.substring(1))
	toggle(id)
}

window.addEventListener("keydown", keylisten, true)
