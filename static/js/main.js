function updatePreviewUrl() {
  var ele = document.getElementById("PreviewURL");
  ele.innerText = `${window.location.host}/URL/${
    document.getElementById("ShortenUrl").value
  }`;
}

document
  .getElementById("ShortenUrl")
  ?.addEventListener("input", updatePreviewUrl);

function copyToClipboard() {
  var text = document.getElementById("FullUrl").innerText;
  navigator.clipboard.writeText(text);
  var ele = document.getElementById("succesMessage");
  ele.style.display = "block";
  setTimeout(() => {
    ele.style.display = "none";
  }, 3000);
}

document.getElementById("btnCopy")?.addEventListener("click", copyToClipboard);
