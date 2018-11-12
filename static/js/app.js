document.addEventListener("DOMContentLoaded", function(ev) {
    console.log("DOM content loaded");

    // interactive elements
    var siBtn = document.querySelector("#sign-in");
    var rememberCB = document.querySelector("#remember");
    var email = document.querySelector("#email");
    var passwd = document.querySelector("#passwd");
    var uploadF = document.querySelector("#upload-file");
    var uploadBtn = document.querySelector("#upload");
    var notes = document.querySelector("#notes");
    var item0 = document.querySelector("#item0");
    var brownfox = document.querySelector("#brownfox");
    var enItem = document.querySelector("#en-item");
    
    siBtn.onclick = function(ev) {
        var fd = new FormData();
        fd.append('email',email.value);
        fd.append('remember',rememberCB.checked);
        fd.append('passwd',passwd.value);
        fetch('/signin', {
            method: 'POST',
            body: fd
        })
        .then(response => response.text())
        .catch(error => console.error('Error:', error))
        .then(response => console.log('Success:',response));
    };

    uploadBtn.onclick = function(ev) {
        var fd = new FormData();
        fd.append('uploadfile',uploadF.files[0]);
        console.log('upload file:',uploadF.files[0]);
        fetch('/uploadfile', {
            method: 'POST',
            body: fd
        })
        .then(response => response.text())
        .catch(error => console.error('File error:',error))
        .then(response => notes.textContent = response);
    };

    item0.onclick = function(ev) {
        console.log("item0: clicked");
    };

    brownfox.onclick = function(ev) {
        console.log("brown fox clicked");
    };

    enItem.onclick = function(ev) {
        console.log("enabled item clicked");
    };
});

// vim: set ts=4 sts=4 sw=4 et:
