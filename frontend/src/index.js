"use strict";

window.$ = window.jQuery = require('jquery');
require ('bootstrap');

let selectedDir = "";

function getUrlVars()
{
    var vars = [], hash;
    var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&');
    for(var i = 0; i < hashes.length; i++)
    {
        hash = hashes[i].split('=');
        vars.push(hash[0]);
        vars[hash[0]] = hash[1];
    }
    return vars;
}

$(() => {    
    const d = getUrlVars()["dir"];
    selectedDir = d ? d : "";
    $('#selected-dir').val(selectedDir);
});


window.showImage = (filename) => {
    filename = selectedDir + "/" + filename;
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

window.gotoDirectory = dir => {
    if(selectedDir != "") {
        dir = selectedDir + "/" + dir;
    }
    window.location.href = "/?dir=" + dir;
}

window.directoryUp = () => {
    if(selectedDir === "") {
       return;
    }
    selectedDir.split("/").length
    selectedDir = selectedDir.slice(0, selectedDir.lastIndexOf("/"));
    window.location.href = "/?dir=" + selectedDir
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