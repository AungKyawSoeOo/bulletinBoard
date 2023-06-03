   // Define the validation rules
   var constraints = {
    email: {
        presence: {
            allowEmpty: false,
            message: "^Email is required"
        },
        email: true,
        length: {
    maximum: 50,
    tooLong: "^Email is too long (maximum %{count} characters)"
}
    },
    password: {
        presence: {
            allowEmpty: false,
            message: "^Password is required"
        },
        length: {
    minimum: 6,
    maximum: 20,
    tooShort: "^Password is too short (minimum %{count} characters)",
    tooLong: "^Password is too long (maximum %{count} characters)"
}
    }
};

// Get the email and password input elements
var emailInput = document.getElementById('email');
var passwordInput = document.getElementById('password');
var loginButton = document.getElementById('loginButton');

// Attach input event listeners to the email and password inputs
emailInput.addEventListener('input', validateForm);
passwordInput.addEventListener('input', validateForm);

// Function to validate the form and enable/disable the login button
function validateForm() {
    // Perform validation
    var validationResult = validate({
        email: emailInput.value,
        password: passwordInput.value
    }, constraints);

    // Clear previous error messages
    document.getElementById('emailError').textContent = '';
    document.getElementById('passwordError').textContent='';
    // Display validation errors, if any
    if (validationResult) {
var emailErrors = validationResult.email;
var passwordErrors = validationResult.password;

if (emailErrors && emailErrors.length > 0) {
    document.getElementById('emailError').textContent = emailErrors[0];
    // document.getElementById('passwordError').style.display = 'none'; // Hide the password error message
} else {
    document.getElementById('passwordError').textContent = passwordErrors ? passwordErrors[0] : '';
    // document.getElementById('passwordError').style.display = 'block'; // Show the password error message
}
}

// Enable/disable the login button
loginButton.disabled = !!(validationResult);
}