$('#btn-send-ucp').click(function () {

    var pers_name = document.getElementById('pers_name').value;
    var pers_age = document.getElementById('pers_age').value;
    var pers_floor = document.getElementById('pers_floor').value;
    var pers_nazi = document.getElementById('pers_nazi').value;
    var pers_skin = document.getElementById('pers_skin').value;
    var pers_country = document.getElementById('pers_country').value;
    var pers_quenta = document.getElementById('pers_quenta').value;

    $.ajax({
        url: '/rest/reg/' + reglogin + "/" + regemail + "/" + regpassword,
        method: 'get',
        contentType: 'application/json',
        headers: {"rest-token": "537003"},
        success: function (data) {
            console.log(data);
            $('#accepted').show(200).delay(3000).slideUp(1000);
        },
        error: function () {
            $('#no-accepted-block').show(200).delay(3000).slideUp(1000);
        },
    });


    setTimeout(function(){
        window.location.replace("http://localhost:3030/auth/" + obj.Login + "/" + obj.AuthToken);
        
    }, 1500);
 });