var constraints = {
  username: {
    presence: {
      allowEmpty: false,
      message: "^Username is required",
    },
    length: {
      minimum: 4,
      maximum: 20,
      tooShort: "^Username is too short (minimum %{count} characters)",
      tooLong: "^Username is too long (maximum %{count} characters)",
    },
  },
  email: {
    presence: {
      allowEmpty: false,
      message: "^Email is required",
    },
    email: true,
    length: {
      maximum: 50,
      tooLong: "^Email is too long (maximum %{count} characters)",
    },
  },
  password: {
    presence: {
      allowEmpty: false,
      message: "^Password is empty",
    },
    length: {
      minimum: 6,
      maximum: 20,
      tooShort: "^Password is too short (minimum %{count} characters)",
      tooLong: "^Password is too long (maximum %{count} characters)",
    },
  },
  confirmPassword: {
    presence: {
      allowEmpty: false,
      message: "^Confirm password is empty",
    },
    equality: {
      attribute: "password",
      message: "^Password and password confirmation do not match",
    },
  },
};

var usernameInput = document.getElementById("username");
var emailInput = document.getElementById("email");
var passwordInput = document.getElementById("password");
var confirmInput = document.getElementById("confirmPassword");
var photoInput = document.getElementById("photo");
var createUserBtn = document.getElementById("createUserBtn");
var updateUserBtn = document.getElementById("updateUserBtn");

var nameError = document.getElementById("userError");
var emailError = document.getElementById("emailError");
var passwordError = document.getElementById("passwordError");
var cpasswordError = document.getElementById("confirmPasswordError");
// Attach input event listeners to the email and password inputs
emailInput.addEventListener("input", validateForm);
usernameInput.addEventListener("input", validateForm);
if (passwordInput) {
  passwordInput.addEventListener("input", validateForm);
}
if (confirmInput) {
  confirmInput.addEventListener("input", validateForm);
}

photoInput.addEventListener("input", validateForm);

function validateForm() {
  if (passwordInput && confirmInput) {
    var validationResult = validate(
      {
        username: usernameInput.value,
        email: emailInput.value,
        password: passwordInput.value,
        confirmPassword: confirmInput.value,
        photo: photoInput.files[0],
      },
      constraints
    );
  } else {
    var validationResult = validate(
      {
        username: usernameInput.value,
        email: emailInput.value,
        photo: photoInput.files[0],
      },
      constraints
    );
  }

  // Clear previous error messages
  if (emailError) {
    emailError.textContent = "";
  }
  if (nameError) {
    nameError.textContent = "";
  }
  if (passwordError) {
    passwordError.textContent = "";
  }
  if (cpasswordError) {
    cpasswordError.textContent = "";
  }
  if (profileError) {
    profileError.textContent = "";
  }

  if (validationResult) {
    var userErrors = validationResult.username;
    var emailErrors = validationResult.email;
    var passwordErrors = validationResult.password;
    var confirmPasswordErrors = validationResult.confirmPassword;
    var photoErrors = validationResult.photo;

    if (userErrors && userErrors.length > 0) {
      document.getElementById("userError").textContent = userErrors[0];
    } else if (emailErrors && emailErrors.length > 0) {
      document.getElementById("emailError").textContent = emailErrors[0];
    } else if (passwordInput && passwordErrors && passwordErrors.length > 0) {
      document.getElementById("passwordError").textContent = passwordErrors[0];
    } else if (
      confirmInput &&
      confirmPasswordErrors &&
      confirmPasswordErrors.length > 0
    ) {
      document.getElementById("confirmPasswordError").textContent =
        confirmPasswordErrors[0];
    } else if (photoErrors && photoErrors.length > 0) {
      document.getElementById("profileError").textContent = photoErrors[0];
    }
  } else {
    // Clear the error messages
    document.getElementById("userError").textContent = "";
    document.getElementById("emailError").textContent = "";
    document.getElementById("passwordError").textContent = "";
    document.getElementById("confirmPasswordError").textContent = "";
    document.getElementById("profileError").textContent = "";
  }

  // Check if the selected file is valid
  var selectedFile = photoInput.files[0];
  var isValidFile = validateFile(selectedFile);

  if (!isValidFile) {
    // Display the file error message
    document.getElementById("profileError").textContent = "Invalid file format";
  }

  console.log(isValidFile);
  // Disable the create button if there are any validation errors
  if (createUserBtn) {
    createUserBtn.disabled = !!validationResult || !isValidFile;
  }
  if (updateUserBtn) {
    updateUserBtn.disabled = !!validationResult || !isValidFile;
  }
}

function validateFile(file) {
  if (!file) {
    // No file selected
    return true;
  }

  var allowedExtensions = ["jpeg", "jpg", "png", "gif"];
  var fileExtension = file.name.split(".").pop().toLowerCase();

  // Check file extension
  if (!allowedExtensions.includes(fileExtension)) {
    return false;
  }

  // Perform other file validations (e.g., size, MIME type) if needed

  return true;
}
