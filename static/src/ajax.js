$('#regs, #auth').click(function () {

    if (this.id == 'regs') {

        var reglogin = document.getElementById('reglogin').value;
        var regemail = document.getElementById('regemail').value;
        var regpassword = document.getElementById('regpassword').value;

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
    }
    else if (this.id == 'auth') {

        var authlogin = document.getElementById('authlogin').value;
        var authpassword = document.getElementById('authpassword').value;

        $.ajax({
            url: '/rest/auth/' + authlogin + "/" + authpassword,
            method: 'get',
            contentType: 'application/json',
            headers: {"rest-token": "537003"},
            success: function (data) {
                var obj = JSON.parse(data);
                $('#accepted').show(200).delay(0).slideUp(1000);
                //redirect on auth page with id + data.AuthToken
                //setTimeout(location.replace("http://localhost:3030/auth"), 30000);
                setTimeout(function(){
                    window.location.replace("http://localhost:3030/auth/" + obj.Login + "/" + obj.AuthToken);
                }, 1500);

            },
            error: function () {
                $('#no-accepted-block').show(200).delay(3000).slideUp(1000);
            },
        });
    }
 });