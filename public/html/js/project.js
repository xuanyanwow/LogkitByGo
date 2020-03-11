"use strict";
layui.use(["okUtils", "table", "countUp", "okMock", 'okTab', 'layer'], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;
    var okTab = layui.okTab();
    var layer = layui.layer;

    function renderList() {
        okUtils.ajax("api/project/get_list", "get", null, true).done(function (response) {
            $("#list-bd").empty();
            $.each(response.data.list, function (index, item) {
                let html = `
                    <div class="layui-col-xs6 layui-col-md3">
                        <div class="layui-card">
                            <div class="ok-card-body project-one" data-id="${item.project_id}" data-name="${item.project_name}">
                                <div class="stat-heading">
                                    ${item.project_name}
                                    <span style="display:inline-block;float:right;" class="project-delete" data-id="${item.project_id}" >删除</span>
                                </div>
                            </div>
                        </div>
                    </div>
                `;
                $("#list-bd").append(html);
            });
        }).fail(function (error) {
            console.log(error);
        });
    }

    $("#list-bd").on("click", ".project-one", function () {
        let id = $(this).data("id");
        let name = $(this).data("name");
        let temName = encodeURI(name);
        let url = `pages/projectDetail.html?id=${id}&name=${name}`;
        let page = `<div lay-id="project_${id}" data-url="${url}"><cite>[项目] ${name} </cite></div>`;
        okTab.tabAdd(page);
    })

    $("#list-bd").on("click", ".project-delete", function () {
        okUtils.ajax("api/project/delete_one", "post", {
            project_id: $(this).data('id'),
        }, true).done(function (response) {
            layer.msg('删除成功');
            location.reload();
        }).fail(function (error) {
            console.log(error);
        });
        layer.close(index);
        return false; // 终止冒泡
    })

    $("#projectAdd").on('click', function () {
        layer.prompt({
            formType: 3,
            value: '',
            title: '请输入新项目名',
            area: ['200px', '150px'] //自定义文本域宽高
        }, function (value, index, elem) {
            okUtils.ajax("api/project/add", "post", {
                project_name: value,
            }, true).done(function (response) {
                location.reload();
            }).fail(function (error) {
                console.log(error);
            });
            layer.close(index);
        });
    });


    renderList();
});


