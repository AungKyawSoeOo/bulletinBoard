function search() {
  var input = document.getElementById("searchInput");
  var filter = input.value.toLowerCase();
  var table = document.getElementById("postTable");
  var rows = table.getElementsByTagName("tr");
  var noPostMsg = document.getElementById("noPostMsg");
  var foundPosts = 0;

  for (var i = 0; i < rows.length; i++) {
    var titleCell = rows[i].getElementsByTagName("td")[0];
    var descriptionCell = rows[i].getElementsByTagName("td")[1];
    if (titleCell || descriptionCell) {
      var title = titleCell.textContent.toLowerCase();
      var description = descriptionCell.textContent.toLowerCase();
      if (title.includes(filter) || description.includes(filter)) {
        rows[i].style.display = "";
        foundPosts++;
      } else {
        rows[i].style.display = "none";
      }
    }
  }
  if (foundPosts === 0) {
    noPostMsg.style.display = "block";
  } else {
    noPostMsg.style.display = "none";
  }
}
