$(document).ready(function () {
  const $resultsBody = $("#resultsBody");

  const $players = $("#playersList input");
  $players.each(function () {
    const player = $(this).val();
    const html = `
    <tr>
      <th scope="row">
          <input type="checkbox"
                 class="form-check-input shouldIncludePlayer"
                 checked="checked"/>
      </th>
      <td>
        <p>${player}</p>
      </td>
      <td>
          <input type="number"
                 class="form-control"
                 value=""/>
      </td>
    </tr>`;
    $resultsBody.append(html);
  });

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
});
