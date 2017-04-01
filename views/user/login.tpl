<!DOCTYPE html>
<html lang="en">
<head>
    <title>登录</title>
    {{template "/inc/common.tpl"}}
    <link href="/static/css/theme_main.css" type="text/css" rel="stylesheet">
    <link href="/static/third_paty/bootstrapvalidator/css/bootstrapValidator.min.css" rel="stylesheet" type="text/css">
    <script src="/static/third_paty/bootstrapvalidator/js/bootstrapValidator.min.js" type="text/javascript"></script>
    <script src="/static/third_paty/bootstrapvalidator/js/language/zh_CN.js" type="text/javascript"></script>
</head>
<body class="container">
<form class="form-signin" id="form-login" action="/user/login">
    <h2 class="form-signin-heading">请登录</h2>
    <div class="form-group">
        <label for="inputEmail" class="sr-only">Email address</label>
        <input type="text" id="inputEmail" class="form-control" placeholder="账号" name="username">
    </div>
    <div class="form-group">
        <label for="inputPassword" class="sr-only">Password</label>
        <input type="password" id="inputPassword" class="form-control" placeholder="密码" name="password">
    </div>
    <div class="checkbox">
        <label>
            <input type="checkbox" value="remember-me"> 记住密码
        </label>
    </div>
    <div class="row">
        <div class="col-6">
            <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
        </div>
        <div class="col-6">
            <a href="/user/register" class="btn btn-lg btn-primary btn-block">注册</a>
        </div>
    </div>
</form>
<script type="text/javascript" src="/static/js/login.js"></script>
</body>
</html>