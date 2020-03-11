"use strict";
layui.use(["okUtils", "table", "countUp", "okMock"], function () {
    var countUp = layui.countUp;
    var table = layui.table;
    var okUtils = layui.okUtils;
    var okMock = layui.okMock;
    var $ = layui.jquery;

    /**
     * 数据统计
     */
    function statText() {
        var elem_nums = $(".stat-text");

        let data = [];
        okUtils.ajax("/api/console/get_data", "post", {}, true).done(function(res){
            data = res.data.min;
            elem_nums.each(function (i, j) {
                !new countUp({
                    target: j,
                    endVal: data[i]
                }).start();
            });
        }).fail(function (error){
            console.log(error);
        });
    }

    var userSourceOption = {
        "title": {"text": ""},
        "tooltip": {"trigger": "axis", "axisPointer": {"type": "cross", "label": {"backgroundColor": "#6a7985"}}},
        "legend": {"data": ["Payment", "娃娃机", "车票"]},
        "toolbox": {"feature": {"saveAsImage": {}}},
        "grid": {"left": "3%", "right": "4%", "bottom": "3%", "containLabel": true},
        "xAxis": [{"type": "category", "boundaryGap": false, "data": ["周一", "周二", "周三", "周四", "周五", "周六", "周日"]}],
        "yAxis": [{"type": "value"}],
        "series": [
            {"name": "Payment", "type": "line", "stack": "总量", "areaStyle": {}, "data": [0,0,1,8,0,0,1]},
            {"name": "娃娃机", "type": "line", "stack": "总量", "areaStyle": {}, "data": [3,5,7,1,0,0,0]},
            {"name": "车票", "type": "line", "stack": "总量", "areaStyle": {}, "data": [6,8,15,10,5,0,0]}
        ]
    };

    /**
     * 用户访问
     */
    function userSource() {
        var userSourceMap = echarts.init($("#userSourceMap")[0], "theme");
        userSourceMap.setOption(userSourceOption);
        okUtils.echartsResize([userSourceMap]);
    }

    // statText();
    // userSource();
});


