"use strict";

window.$ = window.jQuery = require('jquery');
require ('bootstrap');

window.selectedDir = '';
window.loading = false;

window.registerEventlisteners = () => {
    $('#pending-dir-name').on('keyup', (e) => {
        if(e.keyCode != 13){
            return;
        }

        saveNewFolder();
    });
    $('#pending-dir-name').on('keypress', (e) => {
        return e.which != 13;
    });
    
    $('#pending-path-name').on('keyup', (e) => {
        if(e.keyCode != 13){
            return;
        }

        saveRename();
    });
    $('#pending-path-name').on('keypress', (e) => {
        return e.which != 13;
    });
};

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

    if (selectedDir == '') {
        $('#up-wrapper').hide();
    }
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

window.saveNewFolder = () => {
    if(loading) {
        return;
    }
    const name = $('#pending-dir-name').text().trim();
    if (name === "") {
        return;
    }
    loading = true;
    $.ajax({
        url: '/mkdir',
        method: 'post',
        data: {
            name: selectedDir + "/" + name,
        },
        success: res => {
            loading = false;
            $('#pending-directory').html('directory created ...');
            setTimeout(() => {
                window.location.reload();
            }, 1000)                    
        },
        error: err => {
            loading = false;
            alert(err.responseText)
        },
    });
};

window.newFolder = () => {
    if ($('#pending-directory').length == 0) {
        const dom = `<div class="col-12 mt-2" id="pending-directory">
                <div class="d-inline-block pointer">
                    <div class="ffolder small cyan float-left"></div>
                    <div class="float-left ml-2">
                        <div class="mt-2" id="pending-dir-name" contentEditable >New Folder</div>                           
                    </div>
                </div>
            </div>`;
        $('#directories').prepend(dom);
    }
   
    $('#pending-dir-name').focus();
    
    const sel = window.getSelection();
    const ele = $('#pending-dir-name')[0];

    if(sel.toString() == ''){ //no text selection
        let range = document.createRange(); //range object
        range.selectNodeContents(ele); //sets Range
        sel.removeAllRanges(); //remove all ranges from selection
        sel.addRange(range);//add Range to a Selection.
    }

    registerEventlisteners();
};

window.deleteItem = (path) => {
    path = selectedDir + "/" + path;
    const cnf = window.confirm("Are you sure to delete the selected file/directory?\nThis action can't be undone.");
    if(!cnf) {
        return;
    }

    loading = true;
    $.ajax({
        url: '/del?path=' + path,
        method: 'delete',
        success: res => {
            loading = false;
            alert(res);
            setTimeout(() => {
                window.location.reload();
            }, 1000) 
        },
        error: err => {
            loading = false;
            alert(err.responseText)
        },
    });

};

window.rename = (ele) => {
    $('#pending-path-name').attr("id", "");
    ele = $(ele);
    window.oldPathName = ele.text().trim();
    
    ele.attr("contentEditable", true);
    ele.attr("id", "pending-path-name");
    registerEventlisteners();
}

window.saveRename = () => {
    if(loading) {
        return;
    }
    const name = $('#pending-path-name').text().trim();
    if (name === "") {
        return;
    }
    loading = true;
    $.ajax({
        url: '/mv',
        method: 'post',
        data: {
            old_path: selectedDir + "/" + oldPathName,
            new_path: selectedDir + "/" + name,
        },
        success: res => {
            loading = false;
            setTimeout(() => {
                window.location.reload();
            }, 1000)                    
        },
        error: err => {
            loading = false;
            if(err.status == 422 ) {
                $('#pending-path-name').text(window.oldPathName);
            }
            
            alert(err.responseText)
        },
    });
}