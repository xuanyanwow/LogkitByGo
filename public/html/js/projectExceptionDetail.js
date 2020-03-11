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


    var ab_id = getUrlParam("ab_id");



    okUtils.ajax("/api/abnormal/get_detail", "post", {
        ab_id: ab_id
    }, true).done(function (response) {
        // 渲染数据
        console.log(response);
        renderData(response.data);
    }).fail(function (error) {
        console.log(error);
    });

    let ab_file_dom = $("#ab_file");
    let ab_class_dom = $("#ab_class");
    let ab_line_dom = $("#ab_line");
    let ab_message_dom = $("#ab_message");
    let ab_filesoureces_dom = $("#ab_filesoureces");
    let ab_stack_dom = $("#ab_stack");
    let ab_param_dom = $("#ab_param");

    function renderData(data){
        renderClass(data.ab_class);
        renderFileName(data.ab_file, data.ab_line);
        renderMessage(data.ab_message);
        renderCode(data.ab_fileresources);
        renderCallStask(data.ab_stack);
        renderParam(data.ab_data);
    }

    // 渲染异常类名
    function renderClass(title){
        ab_class_dom.html(title);
    }

    // 渲染文件名
    function renderFileName(file_name, file_line){
        ab_file_dom.html(paserFileName(file_name));
        ab_file_dom.attr('title', file_name);
        ab_line_dom.html(` ${file_line} `);
    }

    // 渲染消息
    function renderMessage(message){
        ab_message_dom.html(message);
    }

    // 渲染代码资源
    function renderCode(fileresources){
        if (fileresources.length>=1){
            $.each(fileresources, function(index, item){
                let html = "";
                if (item.is_hight){
                    html += `<span class="line-code">${item.line}</span><span class="light-code">${item.text}</span>\n`;
                }else{
                    html += `<span class="line-code">${item.line}</span>${item.text}\n`;
                }
                ab_filesoureces_dom.append(html);
            });
        }else{
            ab_filesoureces_dom.parent("pre").css("display", "none");
        }
    }

    // 渲染call stack
    function renderCallStask(stack){
        $.each(stack ,function (index,item ){
            ab_stack_dom.append(`<li>${index + 1}. ${item}</li>`);
        });
    }

    // 渲染参数
    function renderParam(param){
        $.each(param, function(index, item){
            let html = `<fieldset class="layui-elem-field">
                        <legend>${index}</legend>
                        <div class="layui-field-box">
                            <table class="layui-table" lay-skin="nob" lay-size="sm"> `;
                $.each(item, function(key, value){
                    html += `<tr>
                        <td class="field_name">${key}</td>
                        <td>${value}</td>
                    </tr>`;
                });
            html +=`</table></div></fieldset>`;
            ab_param_dom.append(html);
        });

    }

    // 解析文件名
    function paserFileName(file_name){
        let array = file_name.split("/");
        return array[array.length - 1];
    }
});
