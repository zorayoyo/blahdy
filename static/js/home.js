var globals = {};

$(document).ready(function () {
    globals.network = new lib.Network()
    globals.client = new BlahdyClient(globals.network)
    globals.blahTemplate = Hogan.compile('<li id="blah_{{Id}}" blah_id="{{Id}}" author_id="{{Author.Id}}" class="blah"><div class="avatar"></div><div class="blah_body"><div class="time">{{UpdateTimeHuman}}</div><div class="name">{{Author.Id}}</div><div class="text">{{Text}}</div></div></li>');

    globals.messageTemplate = Hogan.compile('<li id="{{Id}}" author_id="{{Author.Id}}" class="message"><div class="avatar"></div><div class="message_body"><div class="time">{{CreateTimeHuman}}</div><div class="name">{{Author.Id}}</div><div class="text">{{Text}}</div></div></li>');

    globals.profile_dialog = new widget.Dialog('#profile_dlg');
    globals.profile_dialog.resize(500, 450);
    globals.profile_dialog.create();
    initWelcome();
    
    $('#my_profile_avatar').click(function() {
        globals.profile_dialog.open();
        return false;
    });

    $('#respond .reply_hint').click(function() {
        $(this).hide();
        $('#respond .editor').fadeIn();
        return false;
    });

    $('#respond .save').click(function() {
        $('#respond .editor').hide();
        $('#respond .reply_hint').fadeIn();
        return false;
    });

    $('#message_save_button').click(function () {
        var msg = $('#message_tbox').val().trim()
        if (msg.length !== 0) {
            if (globals.selectedBlahId != null) {
                globals.client.createMessage(
                globals.selectedBlahId, msg,
                function (ret) {
                    loadBlahTimeline(globals.selectedBlahId);
                });
            }
        }
    });
});

function loadAllBlah() {
    globals.client.getAllBlahList(function (result) {
        renderBlahList(result);
    });
}

function renderBlahList(blahList) {
    $('.blah_list .blah').unbind();
    var arr = [];
    for (var i = 0; i < blahList.length; i += 1) {
        blahList[i].UpdateTimeHuman = toHumanTime(blahList[i].UpdateTime);
        arr.push(globals.blahTemplate.render(blahList[i]))
    }
    $('.blah_list').html(arr.join('\n'));
    $('.blah_list .blah').click(function () {
        $('.blah_list .blah.selected').removeClass('selected');
        var blahId = $(this).attr('blah_id');
        if (globals.selectedBlahId === blahId) {
            $('#content').hide();
            globals.selectedBlahId = null;
            $(this).removeClass('selected');
        } else {
            $(this).addClass('selected');
            $('#content').show();
            globals.selectedBlahId = blahId;
            loadBlahTimeline(blahId);
        }
    });
}

function loadBlahTimeline(blahId) {
    globals.client.getBlah(blahId, function (result) {
        renderBlahDetails(result);
    });
    globals.client.getBlahTimeline(blahId, function (result) {
        renderMessageList(result);
    });
}

function renderBlahDetails(blah) {
    var firstblood = $('#content .firstblood');
    blah.CreateTimeHuman = toHumanTime(blah.CreateTime);
    firstblood.find('.name').text(blah.Author.Id);
    firstblood.find('.time').text(blah.CreateTimeHuman);
    firstblood.find('.body').text(blah.Text);
}

function renderMessageList(messageList) {
    $('.discuss .message').unbind();
    var arr = [];
    for (var i = 0; i < messageList.length; i += 1) {
        messageList[i].CreateTimeHuman = toHumanTime(messageList[i].CreateTime);
        arr.push(globals.messageTemplate.render(messageList[i]))
    }
    $('.discuss').html(arr.join('\n'));
    $('.discuss .message').click(function () {
        $('.discuss .message.selected').removeClass('selected');
        $(this).addClass('selected');
    });
}

function toHumanTime(t) {
    var d = new Date(t*1000);
    return d.getFullYear()+'-'+d.getMonth()+'-'+d.getDay() + ' ' + d.toLocaleTimeString();
}


