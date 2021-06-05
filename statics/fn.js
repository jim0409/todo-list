const scheme = "http://"
const url = "127.0.0.1"
const port = ":8000"
const pageLimit = "10"
let basePageNum = 0
const baseAddr = scheme + url + port
let listContent = []
let totalPage = 1

function renderList(str, refresh) {
	if (refresh) {
		listContent = []
	}
	let htmlStr = ''

	let obj = {}
	JSON.parse(str, function (k, v) {
		if (k != "") {
			if (k == "Title") {
				obj.title = v
			} else if (k == "Content") {
				obj.content = v
			} else {
				obj.id = k
				listContent.push(obj)
				// refresh obj
				obj = {}
			}
		}
	})

	var dc = document.getElementById("list")
	listContent.forEach(function (item) {
		htmlStr = htmlStr + `
				<form>
					<p id=${item.id}>id: ${item.id}</p>
					<p>title: ${item.title}</p>
					<p>content: ${item.content}</p>
					<input id="upti_${item.id}"">update title</input>
					<button type="submit" onclick="updateTitle(${item.id})">update</button>
					<input id="upct_${item.id}"">update content</input>
					<button type="submit" onclick="updateContent(${item.id})">update</button>
					<button type="submit" onclick="delNoteByID(${item.id})">delete</button>
				</form>
				`
	})
	dc.innerHTML = htmlStr
}

function renderPageInfo() {
	var dc = document.getElementById("pageInfo")
	let showPage = basePageNum + 1
	dc.innerHTML = "Total Page: " + totalPage + ", Current Page: " + showPage
	if (showPage == 1) {
		return
	}
	if (showPage > totalPage) {
		// loop back to origin page
		getNotes(0)
	}
}

function getPageCount(pageLimit) {
	let addr = baseAddr + "/allpage?limit=" + pageLimit
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			totalPage = this.responseText
			if (totalPage == 0) {
				totalPage = 1
			}
			renderPageInfo()
		}
	}
	xhttp.open("GET", addr)
	xhttp.send()
}

function updateTitle(id) {
	const updateId = "#upti_" + id
	const title = document.querySelector(updateId).value

	let addr = baseAddr + "/note/" + id
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			renderList(this.responseText, true)
		}
	}
	xhttp.open("PUT", addr)
	xhttp.setRequestHeader("Content-Type", "application/json")
	xhttp.send(JSON.stringify({ "title": title }))
}

function updateContent(id) {
	const updateId = "#upct_" + id
	const content = document.querySelector(updateId).value

	let addr = baseAddr + "/note/" + id
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			renderList(this.responseText, true)
		}
	}
	xhttp.open("PUT", addr)
	xhttp.setRequestHeader("Content-Type", "application/json")
	xhttp.send(JSON.stringify({ "content": content }))
}

function delNoteByID(id) {
	let addr = baseAddr + "/note/" + id
	deleting(addr)
}

function getNotes(pageNum) {
	if (pageNum < 0) {
		return
	}
	if (pageNum === 0) {
		basePageNum = 0
	}
	basePageNum = pageNum
	let addr = baseAddr + "/note?limit=" + pageLimit + '&page=' + pageNum
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			renderList(this.responseText, true)
			renderPageInfo()
		}
	}
	xhttp.open("GET", addr, true)
	xhttp.send()
}

function addNote() {
	const title = document.querySelector('#title').value
	const content = document.querySelector('#content').value;
	let addr = baseAddr + "/note"
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			renderList(this.responseText, true)
		}
	}
	xhttp.open("POST", addr)
	xhttp.setRequestHeader("Content-Type", "application/json")
	xhttp.send(JSON.stringify({
		"title": title,
		"content": content,
	}))
}

function deleting(addr) {
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			renderList(this.responseText, true)
		}
	}
	xhttp.open("DELETE", addr, true)
	xhttp.send()
}

function delNote() {
	const id = document.querySelector('#delid').value
	let addr = baseAddr + "/note/" + id
	deleting(addr)
}