"use strict";
layui.use(["okUtils", "table", "countUp", "okMock", 'okTab', 'table', 'siamConfig', 'okLayer', 'form'], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;
    var okTab = layui.okTab();
    var table = layui.table;
    var siamConfig = layui.siamConfig;
    var okLayer = layui.okLayer;
    var form = layui.form;


    var id = getUrlParam("id");
    var name = decodeURI(getUrlParam("name"));

    $("#project-title").html("日志查询 - " + name);

    let backUrl = `projectDetail.html?id=${id}&name=${name}`;

    $("#backBtn").click(function () {
        window.location.href = backUrl;
        return false;
    });

    form.on("submit(logs)", function(data){
        if (data.field.log_sn.length === 0){
            layer.msg("请填写标识");return false;
        }
        okUtils.ajax("api/logs/query", "post", {
            project_id: id,
            log_sn: data.field.log_sn
        }, true).done(function (res) {
            // 渲染响应结果
            let bodyDom = $("#tbody");
            bodyDom.empty();

            if (res.data.length<=0){
                layer.alert("查询不到该单号结果");
            }else{
                $.each(res.data, function (idenx, item) {
                    let html = `
                        <tr>
                            <td>${item.log_category}</td>
                            <td>${item.log_point}</td>
                            <td>${item.log_sn}</td>
                            <td>${item.log_from}</td>
                            <td>${item.create_at}</td>
                            <td><div class="layui-btn layui-btn-xs layui-btn-normal detail-btn">展开</div></td>
                        </tr>
                        <tr class="detail-bd">
                            <td colspan="6" >
<pre data-detail-id="${item.id}">
${item.log_data}
</pre>
                            </td>
                        </tr>
                    `;
                    bodyDom.append(html);

                    // 判断数据是否为json
                    if (isJSON(item.log_data)){
                        $(`pre[data-detail-id=${item.id}]`).JSONView(item.log_data);
                    }
                })
            }
        }).fail(function (error) {
            console.log(error);
        });
        return false;
    });

    $("#tbody").on("click", ".detail-btn",function () {
        // 是展开还是收缩
        var open = $(this).attr("data-open");

        if (open === "2" || open === undefined){ // 当前没有打开 所以要打开
            $(this).html("收缩");
            $(this).parent("td").parent("tr").next("tr").removeClass("detail-bd");
            $(this).attr("data-open", "1");
        }else{// 当前已经打开  所以要关闭
            $(this).html("展开");
            $(this).attr("data-open", "2");
            $(this).parent("td").parent("tr").next("tr").addClass("detail-bd");
        }
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
