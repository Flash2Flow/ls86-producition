$('#auth').on('click', function(){

    
    var login = document.getElementById('login_join').value;
    var password = document.getElementById('password_join').value;
    
    $.ajax({
        url: '/rest/auth/' + login + "/" + password,
        method: 'get',
        contentType: 'application/json',
        success: function (data) {
            var obj = JSON.parse(data);
            $('#accepted').show(200).delay(0).slideUp(1000);
            setTimeout(function(data){
                
                console.log("+++")
                window.location.replace("http://localhost:8080/auth/" + obj.Login + "/" + obj.AuthToken );
                
            }, 1500);
        },
        error: function () {
            console.log("---")
            $('#no-accepted-block').show(200).delay(3000).slideUp(1000);
        },
    });
})

$('#createdbtn').on('click', function(){

    
    var reglogin = document.getElementById('login').value;
    var regemail = document.getElementById('email').value;
    var regpassword = document.getElementById('password').value;
    
    $.ajax({
        url: '/rest/reg/' + reglogin + "/" + regemail + "/" + regpassword,
        method: 'get',
        headers: {"restType": "public", "title":"login"},
        contentType: 'application/json',
        success: function (data) {
            console.log(data);
            var obj = JSON.parse(data);
            $('#accepted').show(200).delay(0).slideUp(1000);
            setTimeout(function(){
                window.location.replace("http://localhost:8080/auth/" + obj.Login + "/" + obj.AuthToken);
                
            }, 1500);
        },
        error: function () {
            $('#no-accepted-block').show(200).delay(3000).slideUp(1000);
        },
    });
})