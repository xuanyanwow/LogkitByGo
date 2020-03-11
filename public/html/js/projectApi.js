"use strict";
layui.use(["okUtils", "table", "countUp", "okMock", 'okTab', 'element', 'siamConfig', 'form'], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;
    var okTab = layui.okTab();
    var element = layui.element;
    var siamConfig = layui.siamConfig;
    var form = layui.form;


    var id   = getUrlParam("id");
    var name = decodeURI(getUrlParam("name"));

    $("#project-title").html("API统计 - " + name);

    let backUrl = `projectDetail.html?id=${id}&name=${name}`;

    $("#backBtn").click(function(){
        window.location.href = backUrl;
        return false;
    });

    var dateArray = [];

    function load_data()
    {
       okUtils.ajax("/api/api_log/overview", "post", {
           project_id:id
       }, true).done(function (res) {
           dateArray = res.data.date;
           // 重组格式
           let success_data = new Array(dateArray.length).fill(0);
           let fail_data = new Array(dateArray.length).fill(0);
           $.each(res.data.data, function(index,item){
               let i = dateArray.indexOf(item.time);
               if (i !== -1){
                   success_data[i] = item.success_times;
                   fail_data[i] = item.fail_times;
               }
           });
           requestCurve(success_data);
           errorCurve(fail_data);

           let qps_data = res.data.qps;
           render_qps(qps_data);
       }).fail(function (error) {
           console.log(error);
       })
    }

    function render_qps(data){
        $("#second_num").html(data.count);
        $("#avg_qps").html(data.qps);
    }


    // 总览图表

    function requestCurve(data) {
        let requestCurveOp = {
            tooltip: {
                trigger: 'axis'
            },
            xAxis: {
                type: 'category',
                data: dateArray,
            },
            yAxis: {
                type: 'value'
            },
            dataZoom: [
                {
                    type: 'slider',
                    show: true,
                    xAxisIndex: [0],
                    start: 0,
                    end: 7
                }
            ],
            series: [{
                data: data,
                type: 'line',
                smooth: true,
                symbol: "none",
            }]
        };
        let requestCurve = echarts.init($("#request-curve")[0], "theme");
        requestCurve.setOption(requestCurveOp);
        okUtils.echartsResize([requestCurve]);
    }


    function errorCurve(data) {
        let errorCurveOp = {
            tooltip: {
                trigger: 'axis'
            },
            xAxis: {
                type: 'category',
                data: dateArray
            },
            yAxis: {
                type: 'value'
            },
            dataZoom: [
                {
                    type: 'slider',
                    show: true,
                    xAxisIndex: [0],
                    start: 0,
                    end: 7
                }
            ],
            series: [{
                data: data,
                type: 'line',
                smooth: true,
                color:"red",
                symbol: "none",
            }],
        };
        let errorCurve = echarts.init($("#error-curve")[0], "theme");
        errorCurve.setOption(errorCurveOp);
        okUtils.echartsResize([errorCurve]);
    }

    function proportion()
    {
        let url = siamConfig.config('url') + "/api/api_log/proportion";
        table.render({
            elem: '#proportion'
            , height: 312
            , url: url //数据接口
            , where:{
                project_id : id
            }
            , page: false //开启分页
            , cols: [[ //表头
                { field: 'api_full', title: '接口名',width:300}
                , { field: 'num', title: '请求数'}
                , { field: 'fail_times', title: '失败次数'}
                , { field: 'avg_consume_time', title: '平均耗时(ms)'}
                , { field: 'proportion', title: '占比'}
                , { field: 'can_use', title: '可用性'}
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
    function user_from_list()
    {
        okUtils.ajax("/api/api_log/user_from_list", "post", {
            project_id :id
        }, true).done(function (res) {
            // 渲染选择列表
            let dom = $("select[name=user_from]");
            dom.empty();
            $.each(res.data.list, function(index, item){
                dom.append(`<option value='${item.user_from}'>${item.user_from}</option>`);
            });
            form.render();
        }).fail(function (error) {
            console.log(error);
        })
    }


    // 点击详情

    element.on('tab(api)', function(data){
        switch (data.index) {
            case 0:
                load_data();
                proportion();
                break;
            case 1:
                break;
            case 2:
                break
        }
    });
    form.on("submit(api_detail)", function(data){
        if (data.field.user_from === ''){
            layer.msg("请选择来源"); return false;
        }
        if (data.field.user_identify.length === 0){
            layer.msg("请填写标识");return false;
        }
        data.field.project_id = id;
        okUtils.ajax("/api/api_log/detail", "post", data.field, true).done(function (res) {
           // 渲染响应结果
            render_detail(res.data);
        }).fail(function (error) {
            console.log(error);
        });
        return false;
    });

    function render_detail(data){
        let not_exists_dom = $(".api-detail-not-exists");
        let result_dom = $(".api-detail-result");
        let api_full_dom = $("#api_full");
        let consume_time_dom = $("#consume_time");
        let is_success_dom = $("#is_success");
        let create_time_dom = $("#create_time");
        let api_param_table_dom = $("#api_param_table");
        let api_response_code_dom = $(".api_response_code");


        not_exists_dom.css("display", "none");
        result_dom.css("display", "none");
        if (data===null || data.length === 0 ){
            not_exists_dom.css("display", "block");
            return true;
        }else{
            result_dom.css("display", "block");
        }

        api_full_dom.html(data.api_full);
        consume_time_dom.html(data.consume_time);
        is_success_dom.html(data.is_success == 1 ? "成功" : "<span style='color:red'>失败</span>");
        create_time_dom.html(data.create_time);
        api_param_table_dom.empty();
        $.each(data.api_param, function(index, item){
            api_param_table_dom.append(`<tr><td class="field_name">${index}</td><td>${item}</td></tr>`)
        });

        api_response_code_dom.html(HTMLEncode(data.api_response));
    }

    function HTMLEncode(html) {
        var temp = document.createElement("div");
        (temp.textContent != null) ? (temp.textContent = html) : (temp.innerText = html);
        var output = temp.innerHTML;
        temp = null;
        return output;
    }

    load_data();
    proportion();
    // 搜索分组加载
    user_from_list();
});
