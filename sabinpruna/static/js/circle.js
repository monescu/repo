$(generateTable);

function generateTable() {
    table = $('<table class="table table-bordered"></table>').attr({ id: "basic_table" });
    let rows = $("#point_value").val();

    counter = 1
    for (let i = 0; i < rows; i++) {
        let row = $('<tr></tr>').appendTo(table);
        for (let j = 0; j < 2; j++) {
            if (j === 0) {
                $('<td></td>').text(`Point ${counter}`).appendTo(row);
            } else {
                $(`<td><input name="point-${counter}" type="text"></td>`).appendTo(row);
            }
        }
        counter++
    }
    table.appendTo("#circle_table");

    if (rows === "") {
        $('#find_circle').hide();
    } else {
        $('#find_circle').show();;
    }
}