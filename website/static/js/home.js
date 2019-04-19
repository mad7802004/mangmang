function login(e) {

    var data = {

    };
    $.ajax({
        url: '/api/v1/user/loginAPW',
        type: 'POST',
        dataType: 'json',
        contentType: 'application/json;charset=UTF-8',
        data: JSON.stringify(data),
        success: function (data, status) {
            // console.log(data);
        }
    });

}
