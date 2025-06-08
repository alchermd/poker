// TODO: DRY this up.
function addPlayerInput() {
  const html = `<input type="text"
                       class="form-control mb-2"
                       name="players"
                       id="players"/>`;
  $("#playersList").append(html);
}

$(document).ready(function () {
  $("#btnAddPlayer").on("click", addPlayerInput);
});

