var globals = {};

$(document).ready(function () {
    globals.network = new lib.Network()
    globals.client = new BlahdyClient(globals.network)
    globals.blahTemplate = Hogan.compile('<li id="blah_{{Id}}" blah_id="{{Id}}" author_id="{{Author.Id}}" class="blah"><div class="avatar {{Author.Id}}"></div><div class="blah_body"><div class="time">{{UpdateTimeHuman}}</div><div class="name">{{Author.Id}}</div><div class="text">{{Text}}</div></div></li>');
    globals.messageTemplate = Hogan.compile('<li id="{{Id}}" author_id="{{Author.Id}}" class="message"><div class="avatar {{Author.Id}}"></div><div class="message_body"><div class="time">{{CreateTimeHuman}}</div><div class="name">{{Author.Id}}</div><div class="text">{{Text}}</div></div></li>');
    globals.participatorTemplate = Hogan.compile('<a href="" class="avatar {{Author.Id}}"></a>');

    globals.profile_dialog = new widget.Dialog('#profile_dlg');
    globals.profile_dialog.resize(500, 450);
    globals.profile_dialog.create();
    globals.compose_dialog = new widget.Dialog('#compose_dlg');
    globals.compose_dialog.resize(500, 300);
    globals.compose_dialog.create();
    initWelcome();
    
    $('#my_profile_avatar').click(function() {
        globals.profile_dialog.open();
        return false;
    });

    $('#toolbar_items a').click(function () {
        var target = $(this).attr('href').substring(1);
        $('#sidebar .sidebar_role:visible').hide();
        $('#sidebar .'+target).fadeIn();
        $('#toolbar_items .selected').removeClass('selected');
        $(this).addClass('selected');
        $('#content').fadeOut();
        if (target == 'search_b') {
            $('#sidebar .search_b li').hide();
        }
    });

    $('#sidebar .friend_b .search').keyup(function () {
        var text = $(this).val();
        $('#sidebar .friend_list li').each(function (i, o) {
            if ($(o).text().indexOf(text) == -1) {
                $(o).fadeOut();
            } else {
                $(o).fadeIn();
            }
        })
    });

    $('#sidebar .search_btn').click(function (ev) {
        var text = $('#search_tbox').val();
        if (text.length === 0) {
            $('#sidebar .search_b li').hide();
        } else {
            $('#sidebar .search_b li').each(function (i, o) {
                if ($(o).text().indexOf(text) == -1) {
                    $(o).fadeOut();
                } else {
                    $(o).fadeIn();
                }
            });
        }
    });

    $('#sidebar .blah_b .add_btn').click(function () {
        globals.compose_dialog.open();
    });

    $('#btn_compose_update').click(function () {
        var str = $('#compose_tbox').val();
        if (str.length === 0) {
            return
        }
        globals.client.createBlah(str, function (ret) {
            loadAllBlah(); 
            globals.compose_dialog.close('slide');
            $('#compose_tbox').val('');
        });
    });

    $('#respond .reply_hint').click(function() {
        $(this).hide();
        $('#respond .editor').fadeIn();
        globals.scrollbar.scroll_to($('.content_inner').get(0).scrollHeight);
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
                    $('#message_tbox').val('')
                    loadBlahTimeline(globals.selectedBlahId);
                    setTimeout(function () {
                        globals.scrollbar.scroll_to($('.content_inner').get(0).scrollHeight);
                    }, 500);
                });
            }
        }
    });
    
    $('#btn_profile_update').click(function () {
        globals.profile_dialog.close();
    });

    globals.scrollbar = new widget.Scrollbar($('.scrollbar_track'), $('.scrollbar_content'))
    globals.scrollbar.recalculate_layout();
    
    widget.Scrollbar.register();
});

function loadAllBlah() {
    globals.client.getAllBlahList(function (result) {
        renderBlahList(result);
    });
}

function renderBlahList(blahList) {
    if ($('#toolbar_items a.selected').attr('href') == '#search_b') {
        var container = $('.blah_list');
    } else {
        var container = $('.blah_list, .search_result');
    }
    container.find('.blah').unbind();
    var arr = [];
    for (var i = 0; i < blahList.length; i += 1) {
        blahList[i].UpdateTimeHuman = toHumanTime(blahList[i].UpdateTime);
        arr.push(globals.blahTemplate.render(blahList[i]))
    }
    container.html(arr.join('\n'));
    container.find('.blah').click(function () {
        container.find('.blah.selected').removeClass('selected');
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
    if (globals.selectedBlahId != null) {
        $('#blah_'+globals.selectedBlahId).addClass('selected');
    }
}

function loadBlahTimeline(blahId) {
    globals.client.getBlah(blahId, function (result) {
        renderBlahDetails(result);
    });
    globals.client.getBlahTimeline(blahId, function (result) {
        renderParticipators(result);
        renderMessageList(result);
    });
}

function renderBlahDetails(blah) {
    var firstblood = $('#content .firstblood');
    blah.CreateTimeHuman = toHumanTime(blah.CreateTime);
    firstblood.find('.name').text(blah.Author.Id);
    firstblood.find('.time').text(blah.CreateTimeHuman);
    firstblood.find('.body').text(blah.Text);
    firstblood.children('.avatar').removeClass().addClass('avatar').addClass(blah.Author.Id);
}

function renderMessageList(messageList) {
    $('.discuss .message').unbind();
    var arr = [];
    for (var i = 0; i < messageList.length; i += 1) {
        messageList[i].CreateTimeHuman = toHumanTime(messageList[i].CreateTime);
        arr.push(globals.messageTemplate.render(messageList[i]));
    }
    $('.discuss').html(arr.join('\n'));
    $('.discuss .message').click(function () {
        $('.discuss .message.selected').removeClass('selected');
        $(this).addClass('selected');
    });

    globals.scrollbar.recalculate_layout();
}

function renderParticipators (messageList) {
    $('#content .participators a').unbind();
    var arr = [];
    var set = {};
    for (var i = 0; i < messageList.length; i += 1) {
        messageList[i].CreateTimeHuman = toHumanTime(messageList[i].CreateTime);
        if (!set.hasOwnProperty(messageList[i].Author.Id)) {
            arr.push(globals.participatorTemplate.render(messageList[i]));
            set[messageList[i].Author.Id] = 1;
        }
    }
    $('#content .participators').html(arr.join('\n'));
}

function toHumanTime(t) {
    var d = new Date(t*1000);
    return d.getFullYear()+'-'+d.getMonth()+'-'+d.getDay() + ' ' + d.toLocaleTimeString();
}


