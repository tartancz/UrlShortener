function updatePreviewUrl(e) {
  var ele = document.getElementById("PreviewURL");
  ele.innerText = `${window.location.host}/URL/${e}`;
}

function copyToCLipboard(text) {
  navigator.clipboard.writeText(text);
  var ele = document.getElementById("succesMessage");
  ele.style.display = "block";
  setTimeout(() => {
    ele.style.display = "none";
  }, 3000);
}
