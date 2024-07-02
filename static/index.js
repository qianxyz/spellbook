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
});
