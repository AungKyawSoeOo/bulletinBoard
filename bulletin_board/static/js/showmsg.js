function confirmDelete(id, title, description, status) {
  var confirmMessage =
    "Are you sure you want to delete the following post? <br>";
  confirmMessage += "ID: " + id + "<br>";
  confirmMessage += "Title: " + title + "<br>";
  confirmMessage += "Description: " + description + "<br>";
  confirmMessage += "Status: " + (status == 1 ? "Active" : "Inactive") + "<br>";

  return Swal.fire({
    title: "Confirmation",
    html: confirmMessage,
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#d33",
    cancelButtonColor: "#3085d6",
    confirmButtonText: "Delete",
  });
}

function showPostDetail(
  title,
  description,
  status,
  createdAt,
  createdBy,
  updatedAt,
  updatedBy
) {
  var detailMessage = "Title: " + title + "<br>";
  detailMessage += "Description: " + description + "<br>";
  detailMessage += "Status: " + (status == 1 ? "Active" : "Inactive") + "<br>";
  var datePart = createdAt.split(" ")[0];
  detailMessage += "Created Date: " + datePart + "<br>";
  detailMessage += "Created User: " + createdBy + "<br>";
  var updatePart = updatedAt.split(" ")[0];
  detailMessage += "Updated Date: " + updatePart + "<br>";
  detailMessage += "Updated User: " + updatedBy + "<br>";
  // detailMessage += "Updated User: "+ updatedBy + "<br>";

  Swal.fire({
    title: "Post Detail",
    html: detailMessage,
    icon: "info",
    confirmButtonText: "Close",
  });
}

document.addEventListener("DOMContentLoaded", function () {
  var postRows = document.getElementsByClassName("post-row");
  for (var i = 0; i < postRows.length; i++) {
    postRows[i].addEventListener("click", function () {
      var title = this.getAttribute("data-post-title");
      var description = this.getAttribute("data-post-description");
      var status = this.getAttribute("data-post-status");
      var createdAt = this.getAttribute("data-post-created-at");
      var createdBy = this.getAttribute("data-post-created-by");
      var updatedAt = this.getAttribute("data-post-updated-at");
      var updatedBy = this.getAttribute("data-post-updated-by");

      showPostDetail(
        title,
        description,
        status,
        createdAt,
        createdBy,
        updatedAt,
        updatedBy
      );
    });
  }
});
function highlightRow(row) {
  row.style.cursor = "pointer";
  row.style.opacity = "0.8";
  row.style.backgroundColor = "black"; // Set the background color of the row
  row.style.color = "white";
}

function unhighlightRow(row) {
  row.style.cursor = "default";
  row.style.opacity = "1";
  row.style.backgroundColor = ""; // Reset the background color to default
  row.style.color = "black";
}
