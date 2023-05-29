function showUserDetail(
  userId,
  username,
  type,
  email,
  phone,
  dob,
  address,
  created_date,
  createdBy,
  updated_date,
  updatedBy
) {
  // Format the user detail message

  var detailMessage = "Name: " + username + "<br>";
  detailMessage += "Type: " + (type === "0" ? "User" : "Admin") + "<br>";
  detailMessage += "Email: " + email + "<br>";
  detailMessage += "Phone: " + phone + "<br>";
  detailMessage += "Date of Birth: " + (dob ? dob : "") + "<br>";
  detailMessage += "Address: " + address + "<br>";
  var create_part = created_date.split(" ")[0];
  detailMessage += "Created Date: " + create_part + "<br>";
  detailMessage += "Created User: " + createdBy + "<br>";
  var update_part = updated_date.split(" ")[0];
  detailMessage += "Updated Date: " + update_part + "<br>";
  detailMessage += "Updated User: " + updatedBy + "<br>";

  Swal.fire({
    title: "User Detail",
    html: detailMessage,
    icon: "info",
    confirmButtonText: "Close",
  });
}

document.addEventListener("DOMContentLoaded", function () {
  var userRows = document.getElementsByClassName("user-row");
  for (var i = 0; i < userRows.length; i++) {
    userRows[i].addEventListener("click", function () {
      var userId = this.getAttribute("data-user-id");
      var username = this.getAttribute("data-username");
      var type = this.getAttribute("data-type");
      var phone = this.getAttribute("data-phone");
      var dob = this.getAttribute("data-dob");
      var birth = dob.split(" ")[0];
      var address = this.getAttribute("data-address");
      var email = this.getAttribute("data-email");
      var created_date = this.getAttribute("data-created-date");
      var created_by = this.getAttribute("data-user-created-by");
      var updated_date = this.getAttribute("data-updated-date");
      var updated_by = this.getAttribute("data-user-updated-by");

      showUserDetail(
        userId,
        username,
        type,
        email,
        phone,
        birth,
        address,
        created_date,
        created_by,
        updated_date,
        updated_by
      );
    });
  }
});
