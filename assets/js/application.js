require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {

  let trig = (name) => {
    console.log(name)
    let el = $(`[data-trigger="${name}"]`);
    return el;
  }

  let stop = () => {
      trig("transport.Play").removeClass("btn-success");
      trig("transport.Record").removeClass("btn-danger");
      trig("transport.Rewind").removeClass("btn-primary");
      trig("transport.Forward").removeClass("btn-primary");
  }

  let triggers = {
    "transport.Stop": stop,
    "transport.Play": () => {
      stop();
      trig("transport.Play").addClass("btn-success");
    },
    "transport.Forward": () => {
      trig("transport.Forward").addClass("btn-primary");
    },
    "transport.Rewind": () => {
      trig("transport.Rewind").addClass("btn-primary");
    },
    "transport.Record": () => {
      trig("transport.Play").addClass("btn-success");
      trig("transport.Record").addClass("btn-danger");
    }
  };

  var evtSource = new EventSource("/ws");

  evtSource.onmessage = function(e) {
    stop();
    let data = JSON.parse(e.data);
    triggers[data.data.trim()]();
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
