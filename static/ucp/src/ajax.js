function getCookie(cookieName) {
    let cookie = {};
    document.cookie.split(';').forEach(function(el) {
      let [key,value] = el.split('=');
      cookie[key.trim()] = value;
    })
    return cookie[cookieName];
  }

$('#btn-ucp').on('click', function(){

    var id = getCookie("Id")
    var token = getCookie("hash")
    var nickname = document.getElementById('nickname').value;
    var floor = document.getElementById('floor').value;
    var age = document.getElementById('age').value;
    var nazi = document.getElementById('nazi').value;
    var skin = document.getElementById('skin').value;
    var country = document.getElementById('country').value;
    var quenta = document.getElementById('quenta').value;
    
    $.ajax({
        //{id}/{name}/{floor}/{age}/{nazi}/{skin}/{country}/{quenta}
        url: '/rest/ucp/create/' + id + "/" + nickname + "/" + floor + "/" + age + "/" + nazi + "/" + skin + "/" + country + "/" + quenta,
        method: 'get',
        headers: {"user-token": token},
        contentType: 'application/json',

        success: function (data) {
            var obj = JSON.parse(data);
            $('#accepted').show(200).delay(0).slideUp(1000);
            setTimeout(function(){
                window.location.replace("http://ls-86-rp.ru/home/" + nickname);
                
            }, 1500);
        },
        error: function () {
            $('#no-accepted-block').show(200).delay(3000).slideUp(1000);
        },
    });
})