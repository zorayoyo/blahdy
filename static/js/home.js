var globals = {};

$(document).ready(function () {
    globals.network = new lib.Network()
    globals.client = new BlahdyClient(globals.network)
    globals.blahTemplate = Hogan.compile('<li id="{{Id}}" author_id="{{AuthorId}}" class="blah"><div class="avatar"></div><div class="blah_body"><div class="name">Name</div><div class="text">{{Text}}</div></div></li>');
    
    loadAllBlah(); 

    // bind events
    $('#sidebar .add_btn').click(function () {
        var str = prompt("Say something, pal.")();
        if (str === null || str.length === 0) {
            return
        }
        globals.client.createBlah(str, function (ret) {
            loadAllBlah(); 
        });
    });

    $('#welcome .signin').click(function () {
        var username = $('#welcome .signin_box .username').val();
        var password = $('#welcome .signin_box .password').val();
        globals.client.auth(
        username, password,
        function (ret) {
            if (ret !== 'error') {
                globals.client.token = ret;
                globals.client.username = username;
                passAuth(); 
            } else {
                alert('incorrect username or password');
            }
        });
    });

    $('#welcome .signup').click(function () {
        globals.client.createAccount(
        $('#welcome .signup_box .username').val(),
        $('#welcome .signup_box .password').val(),
        $('#welcome .signup_box .name').val(),
        function (ret) {
            if (ret !== 'error') {
                $('#welcome .welcome_mode_switch').click();
            } else {
                alert('incorrect username or password');
            }
        });
    });

    $('#welcome .welcome_mode_switch').click(function () {
        if ($('#welcome .signin_box:visible').length !== 0) {
            $('#welcome .signup_box').slideDown();
            $('#welcome .signin_box').slideUp();
            $(this).text('I already have an account.');
        } else {
            $('#welcome .signup_box').slideUp();
            $('#welcome .signin_box').slideDown();
            $(this).text('I don\'t have an account yet.');
        }
        return false;
    });
});

function passAuth() {
    $('#welcome').hide();
    $('#main').show();
}

function loadAllBlah() {
    globals.client.getAllBlahList(function (result) {
        renderBlahList(result);
    });
}

function renderBlahList(blahList) {
    var arr = [];
    for (var i = 0; i < blahList.length; i += 1) {
        arr.push(globals.blahTemplate.render(blahList[i]))
    }
    $('.blah_list').html(arr.join('\n'));
}
