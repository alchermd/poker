function init() {
  addPlayerInput();
  addPlayerInput();
  addPlayerInput();
}

function addPlayerInput() {
  const html = `<input type="text"
                       class="form-control mb-2"
                       name="players"
                       id="players"/>`;
  $("#playersList").append(html);
}

$(document).ready(function () {
  init();
  $("#btnAddPlayer").on("click", addPlayerInput);
});

