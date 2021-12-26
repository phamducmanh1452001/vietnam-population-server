const chooseFile = document.getElementById("file-portrait");
const imgPreview = document.getElementById("img-preview");

chooseFile.addEventListener("change", function () {
    getImgData();
});

var formData;

function getImgData() {
    const files = chooseFile.files[0];
    formData = new FormData();
    formData.append('file-portrait', files);
    if (files) {
        const fileReader = new FileReader();
        fileReader.readAsDataURL(files);
        fileReader.addEventListener("load", function () {
            imgPreview.style.display = "block";
            imgPreview.innerHTML = '<img src="' + this.result + '"' + ' width=200 ' + ' />';
        });
    }
}