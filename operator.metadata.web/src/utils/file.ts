function dataURLToBlob(fileDataURL) {
  let arr = fileDataURL.split(','),
    mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]),
    n = bstr.length,
    u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new Blob([u8arr], { type: mime });
}

function fileToDataURL(file) {
  let reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = function (e) {
    return reader.result;
  };
}

export { dataURLToBlob, fileToDataURL };
