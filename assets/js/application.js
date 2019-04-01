require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {
    $("form").on("change", function() {
        $("form").submit();
    })
});