document.body.addEventListener("htmx:responseError", function (e) {
	const errFooter = document.getElementById("err-footer");
	errFooter.textContent = "An error has occured: " + e.detail.xhr.responseText;
	errFooter.style.display = "block";
});
document.body.addEventListener("htmx:beforeSend", function () {
	const errFooter = document.getElementById("err-footer");
	errFooter.style.display = "none";
	errFooter.innerText = "Failed to hide the error message.";
});
