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

const $sort = $("input[name=sort]");
const $sortLevelIcon = $("#thLevel i");
const $sortNameIcon = $("#thName i");

const sortIcons = {
  level: ["fa-solid fa-sort-up", "fa-solid fa-sort"],
  "-level": ["fa-solid fa-sort-down", "fa-solid fa-sort"],
  name: ["fa-solid fa-sort", "fa-solid fa-sort-up"],
  "-name": ["fa-solid fa-sort", "fa-solid fa-sort-down"],
};

function refreshSortIcons() {
  const [levelIcon, nameIcon] = sortIcons[$sort.val()];
  $sortLevelIcon.attr("class", levelIcon);
  $sortNameIcon.attr("class", nameIcon);
}

refreshSortIcons();

$("#thLevel").on("touchstart mousedown", function (event) {
  event.preventDefault();
  $sort.val($sort.val() === "level" ? "-level" : "level");
  refreshSortIcons();
  $sort[0].dispatchEvent(new Event("change"));
});

$("#thName").on("touchstart mousedown", function (event) {
  event.preventDefault();
  $sort.val($sort.val() === "name" ? "-name" : "name");
  refreshSortIcons();
  $sort[0].dispatchEvent(new Event("change"));
});
