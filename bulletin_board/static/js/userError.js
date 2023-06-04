var constraints = {
    username:{
      presence:{
        allowEmpty:false,
        message:"^Username is required"
      },
      length: {
            minimum: 4,
            maximum: 20,
            tooShort: "^Username is too short (minimum %{count} characters)",
            tooLong: "^Username is too long (maximum %{count} characters)"
        }
    },
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
              message: "^Password is empty"
          },
          length: {
      minimum: 6,
      maximum: 20,
      tooShort: "^Password is too short (minimum %{count} characters)",
      tooLong: "^Password is too long (maximum %{count} characters)"
  }
      },
      confirmPassword: {
        presence: {
              allowEmpty: false,
              message: "^Confirm password is empty"
          },
equality: {
  attribute: "password",
  message: "^Password and password confirmation do not match"
}
}
  };


  var usernameInput=document.getElementById('username');
  var emailInput = document.getElementById('email');
  var passwordInput = document.getElementById('password');
  var confirmInput=document.getElementById('confirmPassword')
  var createUserBtn = document.getElementById('createUserBtn');

  var nameError=document.getElementById('userError');
  var emailError=document.getElementById('emailError');
  var passwordError=document.getElementById('passwordError');
  var cpasswordError=document.getElementById('confirmPasswordError')
  // Attach input event listeners to the email and password inputs
  emailInput.addEventListener('input', validateForm);
  usernameInput.addEventListener('input',validateForm);
  passwordInput.addEventListener('input', validateForm);
  confirmInput.addEventListener('input',validateForm);


  function validateForm() {
var validationResult = validate({
username: usernameInput.value,
email: emailInput.value,
password: passwordInput.value,
confirmPassword: confirmInput.value
}, constraints);

// Clear previous error messages

if(emailError){
    emailError.textContent = '';
}
if(nameError){
    nameError.textContent = '';
}
if(passwordError){
passwordError.textContent = '';
}
if(cpasswordError){
    cpasswordError.textContent='';
}



if (validationResult) {
var userErrors = validationResult.username;
var emailErrors = validationResult.email;
var passwordErrors = validationResult.password;
var confirmPasswordErrors = validationResult.confirmPassword;

if (userErrors && userErrors.length > 0) {
  document.getElementById('userError').textContent = userErrors[0];
} else if (emailErrors && emailErrors.length > 0) {
  document.getElementById('emailError').textContent = emailErrors[0];
} else if (passwordErrors && passwordErrors.length > 0) {
  document.getElementById('passwordError').textContent = passwordErrors[0];
} else if (confirmPasswordErrors && confirmPasswordErrors.length > 0) {
  document.getElementById('confirmPasswordError').textContent = confirmPasswordErrors[0];
}
} else {
// Clear the error message when the password and confirm password match
document.getElementById('confirmPasswordError').textContent = '';
}

if(createUserBtn){
    createUserBtn.disabled = !!validationResult;
}

}