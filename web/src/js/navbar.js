document.addEventListener("DOMContentLoaded", function () {
  fetch("partials/navbar.html")
    .then(response => response.text())
    .then(data => {
      document.getElementById("navbar-container").innerHTML = data;
    })
    .catch(error => console.error("Error fetching navbar:", error));
});

function toggleDropdown() {
  var dropdown = document.getElementById("user-menu");
  dropdown.classList.toggle("hidden");
}
