function updatePreviewUrl(e) {
  var ele = document.getElementById("PreviewURL");
  ele.innerText = `${window.location.host}/URL/${e}`;
}

