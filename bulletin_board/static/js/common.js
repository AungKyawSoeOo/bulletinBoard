function showConfirmation() {
  document.getElementById("overlay").style.display = "block";
  document.getElementById("confirmTitle").textContent =
    document.getElementById("title").value;
  document.getElementById("confirmDescription").textContent =
    document.getElementById("description").value;

  var statusCheckbox = document.getElementById("status");
  var statusValue = statusCheckbox.checked ? "Active" : "Inactive";

  document.getElementById("confirmStatus").textContent = statusValue;
}

function showUserConfirmation() {
  document.getElementById("overlay").style.display = "block";
  document.getElementById("confirmName").textContent =
    document.getElementById("username").value;
  document.getElementById("confirmEmail").textContent =
    document.getElementById("email").value;
  var type = document.getElementById("type").value;
  type == "0"
    ? (document.getElementById("confirmType").textContent = "User")
    : (document.getElementById("confirmType").textContent = "Admin");

  document.getElementById("confirmPhone").textContent =
    document.getElementById("phone").value;
  document.getElementById("confirmDob").textContent =
    document.getElementById("dob").value;
  document.getElementById("confirmAddress").textContent =
    document.getElementById("address").value;
}
// User

function confirmUserCreate() {
  document.getElementById("createUserForm").submit();
}

function confirmUserUpdate() {
  document.getElementById("UserupdateForm").submit();
}

function confirmCreate() {
  document.getElementById("createForm").submit();
}

function confirmUpdate() {
  document.getElementById("updateForm").submit();
}

function cancelCreate() {
  document.getElementById("overlay").style.display = "none";
}
