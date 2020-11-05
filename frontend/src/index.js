"use strict";

window.$ = window.jQuery = require('jquery');
require ('bootstrap');

window.showImage = (filename) => {
    var img = document.createElement('img');
    img.id = "display-img"
    img.src = "/getfile?filename=" + filename;
    $('#img-wrapper').fadeIn(200);

    const imgHolder = $('#img-wrapper #img-holder');
    imgHolder.html(img);

}

window.closeImage = () => {
    $('#img-wrapper').hide();
}

$(document).on('keyup',(e) => {
    if (e.key === "Escape") {
        closeImage();
   }
});

$(document).on('click', (event) => { 
    var target = $(event.target)[0];
    if(target.id === 'display-img') {
        return;
    }
    
    switch(target.id) {
        case 'display-img': return;
        case 'img-wrapper': closeImage();
        case 'img-holder': closeImage();
        default: return;
    }
});