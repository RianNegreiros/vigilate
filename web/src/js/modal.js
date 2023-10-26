document.addEventListener("DOMContentLoaded", function () {
  fetch("partials/modal.html")
    .then(response => response.text())
    .then(data => {
      document.getElementById("modal-container").innerHTML = data;
    })
    .catch(error => console.error("Error fetching modal:", error));
});

function toggleModal() {
  var dropdown = document.getElementById("default-modal");
  dropdown.classList.toggle("hidden");
}