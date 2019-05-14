// 添加项目
function addProject() {
    let name = $("#inputProject").val();
    let content = $("#inputProjectContent").val();
    let user_id = sessionStorage.getItem("user_id");
    if (name === "") {
        $("#project-err").html("请输入项目名称").css("color", "red");
        return
    }
    $.ajax({
        url: '/api/v1/project',
        type: 'POST',
        dataType: 'json',
        contentType: 'application/json;charset=UTF-8',
        data: JSON.stringify({
            "user_id": user_id,
            "project_name": name,
            "project_content": content,
        }),
        success: function (data, status) {
            if (data.code === 0) {
                $('#openProject').modal('hide');
                $("#msg").html(data.msg);
                $('.toast').toast('show');
                console.log(data.msg)
            } else {
                $('#openProject').modal('hide');
                $("#msg").html(data.msg);
                $('.toast').toast('show');
            }
        },
        error: function (data, status) {
            $("#msg").html("服务器错误，请稍后重试");
            $('.toast').toast('show');
        }
    });

}

// 点击编辑项目
function editProject() {
    console.log("点击了。");
    $('#editProject').on('show.bs.modal', function (e) {
        console.log("点击了。")
    });

}

// 删除项目
function deleteProject() {


}
