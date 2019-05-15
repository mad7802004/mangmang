$(document).ready(function () {
    let url = document.location.pathname;
    let splitUrl = url.split('/');
    if (splitUrl[1] === "" || splitUrl[1] === "home") {
        $("#home").addClass("active");
    } else if (splitUrl[1] === "project") {
        $("#project").addClass("active");
    }
});