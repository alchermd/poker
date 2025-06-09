$(document).ready(function () {
  // Populate the player list.
  const $players = $("#playersList input");
  $players.each(function () {
    const player = $(this).val();
    const html = `
    <tr>
      <th scope="row">
          <input type="checkbox"
                 class="form-check-input shouldIncludePlayer"
                 name="shouldIncludePlayer"
                 checked="checked"/>
      </th>
      <td>
        <p>${player}</p>
        <input type="hidden" name="player" value="${player}">
      </td>
      <td>
          <input type="number"
                 name="result"
                 class="form-control"/>
      </td>
    </tr>`;
    const $resultsBody = $("#resultsBody");
    $resultsBody.append(html);
  });

  // Handle checkbox change.
  const $shouldIncludePlayer = $(".shouldIncludePlayer");
  $shouldIncludePlayer.on("change", function () {
    const $checkbox = $(this);
    const shouldInclude = $checkbox.prop("checked");
    const $playerRow = $checkbox.closest("tr");
    const $inputs = $playerRow.find(`input:not([type="checkbox"])`);
    if (shouldInclude) {
      $inputs.removeAttr("disabled");
    } else {
      $inputs.attr("disabled", "disabled");
    }
  });

  // Handle form submission
  const $formLogSession = $("#formLogSession");
  $formLogSession.on("submit", function (e) {
    e.preventDefault();
    const payload = $(this).serializeArray();
    const formData = new FormData();
    payload.forEach(({name, value}) => formData.append(name, value));

    const gameID = $("#gameID").val();
    fetch(`/games/${gameID}/sessions`, {
      method: "POST",
      body: formData,
    })
      .then(res => {
        window.location = `/games/${gameID}`;
      })
      .catch(e => {
        console.log(e);
        alert("Something went wrong.")
      });
  });
});
