"use strict";
layui.use(["okUtils", "table", "countUp", "okMock", 'okTab'], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;
    var okTab = layui.okTab();


    var id   = getUrlParam("id");
    var name = decodeURI(getUrlParam("name"));

    $("#project-title").html("项目详情 - " + name);
    $("#project_id").html(id);

    $("#chooseApi").attr('href', `projectApi.html?id=${id}&name=${name}`);
    $("#chooseException").attr('href', `projectException.html?id=${id}&name=${name}`);
    $("#chooseLogn").attr('href', `projectLog.html?id=${id}&name=${name}`);
});



