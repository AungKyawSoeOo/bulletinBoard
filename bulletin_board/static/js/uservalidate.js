$(document).ready(function () {
  $("form").submit(function (event) {
    event.preventDefault(); // Prevent form submission

    // Clear previous error messages
    $(".error").text("");

    // Validate inputs
    var username = $("#username").val();
    var email = $("#email").val();
    var password = $("#password").val();
    var confirmPassword = $("#cpassword").val();

    // Perform validation and display error messages
    var isValid = true;

    if (username === "") {
      $("#usernameError").text("Username is required.");
      isValid = false;
    }

    if (email === "") {
      $("#emailError").text("Email is required.");
      isValid = false;
    }

    if (password === "") {
      $("#passwordError").text("Password is required.");
      isValid = false;
    }

    if (confirmPassword === "") {
      $("#confirmPasswordError").text("Password confirmation is required.");
      isValid = false;
    } else if (password !== confirmPassword) {
      $("#confirmPasswordError").text("Passwords do not match.");
      isValid = false;
    }

    // If all inputs are valid, proceed with form submission
    if (isValid) {
      $("createUserForm").unbind("submit").submit();
    }
  });
});
