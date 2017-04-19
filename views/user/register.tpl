<!DOCTYPE html>
<html lang="en">
<head>
    <title>注册</title>
    {{template "/inc/common.tpl"}}
    <link href="/static/css/theme_main.css" type="text/css" rel="stylesheet">
</head>
<body class="container">
<form class="form-signin" id="form-login" action="/user/login">
    <h2 class="form-signin-heading">注册</h2>
    <div class="form-group row">
        <div class="col-4 ">
            <label for="inputEmail" class="lable-register">用户名：</label>
        </div>
        <div class="col-8">
            <input type="text" id="inputEmail" class="form-control" placeholder="请输入用户名" name="username">
        </div>
    </div>
    <div class="form-group row ">
        <div class="col-4">
            <label for="inputPassword" class="lable-register">密码：</label>
        </div>
        <div class="col-8">
            <input type="password" id="inputPassword" class="form-control" placeholder="请输入密码" name="password">
        </div>
    </div>
    <div class="form-group row ">
        <div class="col-4">
            <label for="inputPassword" class="lable-register">确认密码：</label>
        </div>
        <div class="col-8">
            <input type="password" id="inputPassword" class="form-control" placeholder="请再次输入密码" name="password">
        </div>
    </div>
    <button class="btn btn-lg btn-primary btn-block" type="submit">注册</button>
</form>
</body>
</html>