const backToTop = document.getElementById("back-to-top");

window.addEventListener("scroll", () => {
  backToTop.style.display = window.scrollY > 100 ? "block" : "none";
});

var sliderC = new Slider("#slider", {
  id: "slider",
  min: 0,
  max: 9,
  range: true,
  value: [0, 9],
  tooltip: "hide",
}).on("change", function (data) {
  $("#levelMin").text(data.newValue[0]);
  $("#levelMax").text(data.newValue[1]);
});
