var globals = {};

$(document).ready(function () {
    globals.network = new lib.Network()
    globals.client = new BlahdyClient(globals.network)
    globals.client.getBlahList(function (result) {
        console.log(result);
        renderBlahList(result);
    });
    globals.blahTemplate = Hogan.compile('<li id="{{Id}}" author_id="{{AuthorId}}" class="blah"><div class="avatar"></div><div class="blah_body"><div class="name">Name</div><div class="text">{{Text}}</div></div></li>');
});

function renderBlahList(blahList) {
    var arr = [];
    for (var i = 0; i < blahList.length; i += 1) {
        arr.push(globals.blahTemplate.render(blahList[i]))
    }
    $('.blah_list').html(arr.join('\n'));
}
