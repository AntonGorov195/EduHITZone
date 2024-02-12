const isStudentInput = document.getElementById("is-student-input");
const studentForm = document.getElementById("student-form");
const sendCodeBtn = document.getElementById("send-code");

sendCodeBtn.addEventListener("click", (e) => e.preventDefault());
isStudentInput.addEventListener("change", isStundentToggle);

function isStundentToggle(event) {
	studentForm.classList.toggle("form-hidden");
}
