/**
 * Created by ghy on 2017/3/28.
 */

var wsHost = 'ws://localhost:8080/chat?email='+urlParam('email'),
    conn;

if (window["WebSocket"]) {
    conn = new WebSocket(wsHost);
    conn.onclose = function(evt) {
        appendLog($("<div><b>Connection closed.</b></div>"))
    }
    conn.onmessage = function(evt) {
        var data = JSON.parse(evt.data);
        switch(data.MType) {
            case "text_mtype":
                appendMsg(data.TextMessage)
                break;
            case "status_mtype":
                updateUsers(data.UserStatus)
                break;
            default:
        }

    }
} else {
    appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
}

function updateUsers(data){
    console.log(data)
}

function appendMsg(data){
    var msg = $('<li><time>'+data.Time+'</time><div>'+data.Content+'</div></li>'),
        user = data.UserInfo
    if(user){
        msg.prepend('<div class="user-avt"><a href="mailto:'+user.Email+'" title="'+user.Name+'"><img src="'+user.Gravatar+'"/> </a></div>')
    }
    $("#msgContent").append(msg)
    $(window).scrollTop(0)
}
function appendLog(dom){
    return $(document.body).append(dom);
}

$(function(){
    $("#subForm").on('submit',function(){
        var _msgVal = $("#msgTxt").val().trim()
        if(_msgVal){
            conn.send(_msgVal);
            $("#msgTxt").val('')
        }
        return false;
    })
})



function urlParam(name){
    var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
    return results?results[1] : '';
}

