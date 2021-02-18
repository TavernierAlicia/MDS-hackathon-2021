document.addEventListener('DOMContentLoaded', () => {
	document.body.classList.add('loaded');
})

document.addEventListener('scroll', () => {
	window.setTimeout(() => {
		if (window.scrollY > 20) {
			document.body.classList.add('scrolled')
		} else {
			document.body.classList.remove('scrolled')
		}
	}, 250)
})


function addClassNav() {
	if (window.scrollY > 20) {
	if (!document.body.classList[1]) {
			console.log("add scrolled")
			console.log(document.body.classList)
			document.body.classList.add('scrolled')
		}
	}
}


function removeClassNav() {
	if (window.scrollY > 20) {
	if (document.body.classList[1]) {
		if (document.body.classList[1].includes("scrolled")) {
			console.log("remove scrolled")
			console.log(document.body.classList)
			document.body.classList.remove('scrolled')
		}
		}
	}
}


function scrollButton() {
	if (window.scrollY > 20) {
		window.scrollTo({top:0, left:0, behavior:"smooth"})
	}
}




document.getElementById('menuButton').addEventListener('click', () => {
	document.body.classList.toggle('menuOpen');
})
