function showUserDetail(userId, username, type, email, phone, dob, address) {
    // Format the user detail message
  
    var detailMessage = "Name: " + username + "<br>";
    detailMessage += "Type: " + (type === "0" ? "User" : "Admin") + "<br>";
    detailMessage += "Email: " + email + "<br>";
    detailMessage += "Phone: " + phone + "<br>";
    detailMessage += "Date of Birth: " + (dob ? dob : "") + "<br>";
    detailMessage += "Address: " + address + "<br>";

    // Display the user detail using a modal or any other method you prefer
    Swal.fire({
        title: 'User Detail',
        html: detailMessage,
        icon: 'info',
        confirmButtonText: 'Close'
    });
}

document.addEventListener("DOMContentLoaded", function() {
    var userRows = document.getElementsByClassName("user-row");
    for (var i = 0; i < userRows.length; i++) {
        userRows[i].addEventListener("click", function() {
            var userId = this.getAttribute("data-user-id");
            var username = this.getAttribute("data-username");
            var type = this.getAttribute("data-type");
            var phone = this.getAttribute("data-phone");
            var dob = this.getAttribute("data-dob");
            var birth=dob.split(" ")[0]
            var address = this.getAttribute("data-address");
            var email =this.getAttribute("data-email");

            showUserDetail(userId, username, type,email, phone, birth, address);
        });
    }
});