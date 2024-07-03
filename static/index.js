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
    value: [0, 9],
    tooltip: "hide",
  })
  .on("change", function (data) {
    $("#levelMin").text(data.value.newValue[0]);
    $("#levelMax").text(data.value.newValue[1]);
    document.body.dispatchEvent(new Event("changeLevel"));
  });
