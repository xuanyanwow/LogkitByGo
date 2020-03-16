"use strict";
layui.use(["okUtils", "table", "countUp", "okMock", 'okTab', 'table', 'siamConfig', 'okLayer'], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;
    var okTab = layui.okTab();
    var table = layui.table;
    var siamConfig = layui.siamConfig;
    var okLayer = layui.okLayer;

    var id = getUrlParam("id");
    var name = decodeURI(getUrlParam("name"));

    $("#project-title").html("事务管理 - " + name);

    let backUrl = `projectDetail.html?id=${id}&name=${name}`;

    $("#backBtn").click(function () {
        window.location.href = backUrl;
        return false;
    });

    var todoNo = document.getElementById('todo-no');
    var sortableNo = Sortable.create(todoNo,{
        group: 'shared',
        animation: 150
    });
    var todoIng = document.getElementById('todo-ing');
    var sortableIng = Sortable.create(todoIng,{
        group: 'shared',
        animation: 150
    });
    var todoFinish = document.getElementById('todo-finish');
    var sortableFinish = Sortable.create(todoFinish,{
        group: 'shared',
        animation: 150
    });

    function isJSON(str) {
        if (typeof str == 'string') {
            try {
                var obj=JSON.parse(str);
                return !!(typeof obj == 'object' && obj);
            } catch(e) {
                return false;
            }
        }
    }

});
