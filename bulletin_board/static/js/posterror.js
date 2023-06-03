document.addEventListener("DOMContentLoaded", function () {
    // Define the validation rules
    var constraints = {
      title: {
        presence: {
          allowEmpty: false,
          message: "^Title is required",
        },
        length: {
          minimum: 5,
          maximum: 50,
          tooShort:"^Title is too short (minimum %{count} characters)",
          tooLong: "^Title is too long (maximum %{count} characters)",
        },
      },
      description: {
        presence: {
          allowEmpty: false,
          message: "^Description is required",
        },
        length: {
          minimum: 6,
          maximum: 255,
          tooShort: "^Description is too short (minimum %{count} characters)",
          tooLong: "^Description is too long (maximum %{count} characters)",
        },
      },
    };

    // Get the input elements
    var titleInput = document.getElementById("title");
    console.log(titleInput)
    var descriptionInput = document.getElementById("description");
    var postCreateBtn = document.getElementById("postCreateBtn");
    var postUpdateBtn=document.getElementById("postUpdateBtn");

    var titleError=document.getElementById("titleError")
    var desError=document.getElementById("descriptionError")

    // Attach input event listeners to the title and description inputs
    titleInput.addEventListener("input", validateForm);
    descriptionInput.addEventListener("input", validateForm);

    // Function to validate the form and enable/disable the create button
    function validateForm() {
      // Perform validation
      var validationResult = validate(
        {
          title: titleInput.value,
          description: descriptionInput.value,
        },
        constraints
      );

      // Clear previous error messages
      if (titleError) {
titleError.textContent = "";
}
if (desError) {
desError.textContent = "";
}


      // Display validation errors, if any
      if (validationResult) {
        var titleErrors = validationResult.title;
        var descriptionErrors = validationResult.description;

        if (titleErrors && titleErrors.length > 0) {
          titleError.textContent = titleErrors[0];
        } else {
          if (desError) {
    desError.textContent = descriptionErrors ? descriptionErrors[0] : "";
  }
        }
      }
      if(postCreateBtn){
        postCreateBtn.disabled = !!validationResult;
      }
      if(postUpdateBtn){
        postUpdateBtn.disabled = !!validationResult;
      }
    }
  });