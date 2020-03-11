<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <title>安装向导</title>

    <link rel="stylesheet" type="text/css" href="https://www.layuicdn.com/layui-v2.5.5/css/layui.css" />
    <script src="https://www.layuicdn.com/layui-v2.5.5/layui.js" v="layui" e="layui"></script>

    <style>
    body{
        background:#f6f5f1
    }
    #install-card{
        margin-top:30px;
    }
    #install-tips{
        width: 100%;
        color: #79C48C;
        margin:0 auto;
        text-align:center;
        font-size:2em;
    }
    </style>

</head>
<body>
<div class="layui-container">
  <div class="layui-row">
    <div class="layui-col-xs12 layui-col-sm6 layui-col-md6 layui-col-md-offset3">
        <div class="layui-card" id="install-card">
          <div class="layui-card-header">
            <div id="install-tips">Siam LogKit By Go 安装向导</div>
          </div>
          <div class="layui-card-body">
            数据库配置 <br/>{{.databaseInfo}}<br/><br/>
            Mysql版本号<br/>  {{.version}} <br/> <br/>

            能正确看到数据库配置和版本号即可点击进行创建表结构完成安装
            <br/>
            <br/>
            <a href="/install_run_sql" class="layui-btn layui-btn-normal">开始安装</a>
          </div>
        </div>
    </div>
  </div>
</div>
</body>
</html>