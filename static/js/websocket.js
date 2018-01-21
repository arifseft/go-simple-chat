var socket;

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('wss://' + window.location.host + '/chat/join?uname=' + $('#uname').text() + '&group=' + $('#group').text());
        
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        if(data.Group == $('#group').text()) {
            switch (data.Type) {
            case 0: // JOIN
                if (data.User == $('#uname').text()) {
                    $("#chatbox li").last().after("<li>You joined the chat room.</li>");
                } else {
                    $("#chatbox li").last().after("<li><b>" + data.User + "</b> joined the chat room.</li>");
                }
                break;
            case 1: // LEAVE
                $("#chatbox li").last().after("<li>" + data.User + " left the chat room.</li>");
                break;
            case 2: // MESSAGE
                $("#chatbox li").last().after("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
                break;
            }
        }
    };

    // Send messages.
    var postContent = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val("");
    }

    $('#sendbtn').click(function () {
        postContent();
    });
});