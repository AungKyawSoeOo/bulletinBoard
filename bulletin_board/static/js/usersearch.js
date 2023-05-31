function searchUser() {
  var input = document.getElementById("searchInputUser");
  var startDateInput = document.getElementById("start-date");
  var endDateInput = document.getElementById("end-date");
  var searchButton = document.getElementById("date-search");

  if (input.value !== "") {
    startDateInput.disabled = true;
    endDateInput.disabled = true;
    searchButton.disabled = true;
  } else {
    startDateInput.disabled = false;
    endDateInput.disabled = false;
    searchButton.disabled = false;
  }

  var filter = input.value.toLowerCase();
  var table = document.getElementById("userTable");
  var rows = table.getElementsByTagName("tr");
  var noUserMsg = document.getElementById("noUserMsg");
  var foundUsers = 0;

  for (var i = 0; i < rows.length; i++) {
    var nameCell = rows[i].getElementsByTagName("td")[0];
    var addressCell = rows[i].getElementsByTagName("td")[4];
    if (nameCell || addressCell) {
      var name = nameCell.textContent.toLowerCase();
      var address = addressCell.textContent.toLowerCase();
      if (name.includes(filter) || address.includes(filter)) {
        rows[i].style.display = "";
        foundUsers++;
      } else {
        rows[i].style.display = "none";
      }
    }
  }

  if (foundUsers === 0) {
    noUserMsg.style.display = "block";
  } else {
    noUserMsg.style.display = "none";
  }
}

var searchButton = document.getElementById("date-search");
var startDateInput = document.getElementById("start-date");
var endDateInput = document.getElementById("end-date");
if (startDateInput.value === "" && endDateInput.value === "") {
  searchButton.disabled = true;
} else {
  searchButton.disabled = false;
}

function searchByDate() {
  var startDateInput = document.getElementById("start-date");
  var endDateInput = document.getElementById("end-date");
  var input = document.getElementById("searchInputUser");
  var clearDateBtn = document.getElementById("clear-date-btn");
  var searchButton = document.getElementById("date-search");

  if (startDateInput.value !== "" || endDateInput.value !== "") {
    input.disabled = true;
    clearDateBtn.style.display = "block";
  } else {
    input.disabled = false;
    clearDateBtn.style.display = "block";
    searchButton.style.display = "block";
  }

  //   if (startDateInput.value === "" && endDateInput.value === "") {
  //     searchButton.disabled = true;
  //   } else {
  //     searchButton.disabled = false;
  //   }
  var startDate = new Date(startDateInput.value);
  var endDate = new Date(endDateInput.value);
  var table = document.getElementById("userTable");
  var rows = table.getElementsByTagName("tr");
  var noUserMsg = document.getElementById("noUserMsg");
  var foundUsers = 0;

  for (var i = 0; i < rows.length; i++) {
    var dateCell = rows[i].getElementsByTagName("td")[6];
    if (dateCell) {
      var dateStr = dateCell.textContent;
      var date = new Date(dateStr);

      if (date >= startDate && date <= endDate) {
        rows[i].style.display = "";
        foundUsers++;
      } else {
        rows[i].style.display = "none";
      }
    }
  }

  if (foundUsers === 0) {
    noUserMsg.style.display = "block";
  } else {
    noUserMsg.style.display = "none";
  }
}

function clearDate() {
  document.getElementById("start-date").value = "";
  document.getElementById("end-date").value = "";
  document.getElementById("searchInputUser").disabled = false;
  document.getElementById("clear-date-btn").style.display = "none";
  searchUser();
}
