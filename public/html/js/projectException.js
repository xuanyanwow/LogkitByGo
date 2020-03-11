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

    $("#project-title").html("异常统计 - " + name);

    let backUrl = `projectDetail.html?id=${id}&name=${name}`;

    $("#backBtn").click(function () {
        window.location.href = backUrl;
        return false;
    });


    function getChartData()
    {

        okUtils.ajax("/api/abnormal/get_static", "post", {
            project_id : id
        }, true).done(function (response) {
            console.log(response);

            let column = [];
            let data = [];

            $.each(response.data , function(index, item){
                column.push(item.ab_date);
                data.push(item.count);
            });

            var staticSource = {
                "title": {"text": ""},
                "tooltip": {"trigger": "axis", "axisPointer": {"type": "cross", "label": {"backgroundColor": "#6a7985"}}},
                "toolbox": {"feature": {"saveAsImage": {}}},
                "grid": {"left": "3%", "right": "4%", "bottom": "3%", "containLabel": true},
                "xAxis": [{"type": "category", "boundaryGap": false, "data": column}],
                "yAxis": [{"type": "value"}],
                "series": [
                    {"name": "Payment", "type": "line", "stack": "总量", "areaStyle": {}, "data": data},
                ]
            };

            showChart(staticSource);
        }).fail(function (error) {
            console.log(error);
        });


    }

    /**
     * 渲染图表
     */
    function showChart(staticSource) {

        var staticChart = echarts.init($("#exceptionChart")[0], "theme");
        staticChart.setOption(staticSource);
        okUtils.echartsResize([staticChart]);
    }

    /**
     * 列表
     */
    function initList() {
        let url = siamConfig.config('url') + "/api/abnormal/get_list";
        table.render({
            elem: '#lists'
            , height: 312
            , url: url //数据接口
            , where:{
                project_id : id
            }
            , page: true //开启分页
            , cols: [[ //表头
                { field: 'ab_id', title: 'ID', width: 80, fixed: 'left' }
                , { field: 'create_time', title: '时间', width:200}
                , { field: 'ab_class', title: '类', width:400}
                , { field: 'ab_message', title: '异常消息'}
                , {fixed: 'right', width:150, align:'center', toolbar: '#exceptionDo'} 
            ]]
            , response: {
                statusName: 'code' //规定数据状态的字段名称，默认：code
                ,statusCode: 200 //规定成功的状态码，默认：0
                ,countName: 'count' //规定数据总数的字段名称，默认：count
                ,dataName: 'data' //规定数据列表的字段名称，默认：data
            } 
            , parseData: function(res){ //res 即为原始返回的数据
                return {
                    "code": res.code, //解析接口状态
                    "msg": res.msg, //解析提示文本
                    "count": res.data.count, //解析数据长度
                    "data": res.data.list //解析数据列表
                };
            }
        });
    }

    // 行事件

    table.on("tool(lists)", function(obj){
        var data = obj.data; //获得当前行数据
        var layEvent = obj.event; //获得 lay-event 对应的值（也可以是表头的 event 参数对应的值）
        var tr = obj.tr; //获得当前行 tr 的 DOM 对象（如果有的话）

        okLayer.open("异常详情", `projectExceptionDetail.html?ab_id=${data.ab_id}`, "80%", "80%", function(layero){
        }, function(){

        });
    });

    getChartData();
    initList();
});
