require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {

  let triggers = {
    "transport.Play": () => {
      $("[data-trigger='transport.Play']").addClass("btn-success");
    },
    "transport.Stop": () => {
      $("[data-trigger='transport.Play']").removeClass("btn-success");
    }
  };

  $("form").on("change", () => {
    $("form").submit();
  });

  $("button[data-trigger]").on("click", (e) => {
    e.preventDefault();
    let id = $(e.currentTarget).attr("data-trigger");
    let xhr = $.post("/trigger", {"events": id});
    xhr.done(() => {
      triggers[id]();
    });
  });
});
