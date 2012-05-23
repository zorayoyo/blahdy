var globals = {};

$(document).ready(function () {
    globals.network = new lib.Network()
    globals.client = new BlahdyClient(globals.network)
    globals.blahTemplate = Hogan.compile('<li id="{{Id}}" author_id="{{Author.Id}}" class="blah"><div class="avatar"></div><div class="blah_body"><div class="time">{{UpdateTimeHuman}}</div><div class="name">{{Author.Id}}</div><div class="text">{{Text}}</div></div></li>');
    
    initWelcome();
    
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
        blahList[i].UpdateTimeHuman = new Date(blahList[i].UpdateTime).toLocaleTimeString();
        console.log(blahList[i].Text)
        arr.push(globals.blahTemplate.render(blahList[i]))
    }
    $('.blah_list').html(arr.join('\n'));
    $('.blah_list .blah').click(function () {
        $('.blah_list .blah.selected').removeClass('selected');
        $(this).addClass('selected');
    });
}
