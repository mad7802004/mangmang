// 登录
function login() {
    let phone = $('#inputPhone').val();
    let passWord = $("#inputPassWord").val();
    if (checkPhone(phone) && checkPassWord(passWord)) {
        $.ajax({
            url: '/api/v1/user/loginAPW',
            type: 'POST',
            dataType: 'json',
            contentType: 'application/json;charset=UTF-8',
            data: JSON.stringify({"phone": phone, "pass_word": passWord}),
            success: function (data, status) {
                if (data.code === 0) {
                    sessionStorage.setItem("token", data.data.user_id);
                    document.cookie = "user_id=" + data.data.user_id;
                    window.location.href = "/home";
                } else {
                    $("#errMsg").html(data.msg).css("visibility", "");
                }
            },
            error: function (data, status) {
                $("#errMsg").html("服务器发生错误").css("visibility", "");
            }
        });
    }


}

// 检验账号是否正确
function checkPhone(phone) {
    let reg = /^0?(13[0-9]|15[012356789]|17[0678]|18[0123456789]|14[57])[0-9]{8}$/;
    if (phone === "") {
        $("#errMsg").html("手机号码不能为空").css("visibility", "");
        return false
    } else if (phone.length < 11) {
        $("#errMsg").html("手机号码长度有误").css("visibility", "");
        return false
    } else if (!reg.test(phone)) {
        $("#errMsg").html("手机号不存在").css("visibility", "");
        return false
    } else {
        $("#errMsg").css("visibility", "hidden");
        return true
    }
}

// 检验密码是否正确
function checkPassWord(passWord) {
    if (passWord === "") {
        $("#errMsg").html("密码不能为空").css("visibility", "");
        return false
    } else {
        $("#errMsg").css("visibility", "hidden");
        return true
    }

}
