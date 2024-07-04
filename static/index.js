$(window).on("scroll", function () {
  if ($(this).scrollTop() > 100) {
    $("#back-to-top").fadeIn();
  } else {
    $("#back-to-top").fadeOut();
  }
});

$("#slider")
  .slider({
    id: "slider",
    min: 0,
    max: 9,
    range: true,
    value: [parseInt($("#levelMin").val()), parseInt($("#levelMax").val())],
    tooltip: "hide",
  })
  .on("change", function (data) {
    $("#levelMin").val(data.value.newValue[0]);
    $("#levelMax").val(data.value.newValue[1]);
    $("#slider")[0].dispatchEvent(new Event("change"));
  });
