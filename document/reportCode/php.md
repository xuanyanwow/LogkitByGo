# php 上报数据 

```php

<?php
// 日志上报 单条
$send = [
    'action' => 'report',
    'data'   => [
        'project_id'   => 4,
        'log_category' => 'Api/PayBase/v2',
        'log_point'    => 'curl_begin',
        'log_sn'       => '666',
        'log_data'     => json_encode(["name" => "siam", "age" => 22]),
        'log_from'     => 'siam',
    ],
    'sign'   => md5('report'.'_siam'),
];
// api 统计批量
$send = [
    'action' => 'apis',
    'data'   => [
        [
            'project_id'    => 5,
            'api_category'  => 'wechat',
            'api_method'    => 'refund',
            'is_success'    => 1, // 1成功 0失败
            'consume_time'  => 15, // 耗时 15毫秒
            'user_from'     => 'payment',
            'user_identify' => '20200313001', // 记录标识  可以是订单号
            'api_response'  => '响应内容 可选',
            'api_param'     => json_encode(["test" => "tttt"]),
        ],
    ],
    'sign'   => md5('report'.'_siam'),
];

// json 然后处理十六进制
$sendStr = encode(json_encode($send, 256));

$sendStrArray = str_split(str_replace(' ', '', $sendStr), 2);  // 将16进制数据转换成两个一组的数组


try{
    $socket = socket_create(AF_INET, SOCK_STREAM, getprotobyname("tcp"));  // 创建Socket
    if (@socket_connect($socket, "127.0.0.1", 8999)) {  //连接
        $send = '';
        for ($j = 0; $j < count($sendStrArray); $j++) {
            $send .= chr(hexdec($sendStrArray[$j]));
        }
        socket_write($socket, $send);
    } else {
        echo "connection error";
    }

    socket_close($socket);  // 关闭Socket
}catch (Throwable $e){
    echo $e->getMessage();
}


function encode($string){
    $len = strlen($string);
    // 补位到4位
    $len = dechex($len);
    if (strlen($len) < 4) {
        $pd = 4 - strlen($len);
        $realyPd = '';
        while($pd>0){
            $realyPd .= "0";
            $pd--;
        }
        $realyPd.=$len;
    }else{
        $realyPd = $len;
    }
    return $realyPd . strToHex($string);
}

function strToHex($string)
{
    $hex = "";
    for ($i = 0; $i < strlen($string); $i++)
        $hex .= dechex(ord($string[$i]));
    $hex = strtoupper($hex);
    return $hex;
}

function hexToStr($hex)
{
    $string = "";
    for ($i = 0; $i < strlen($hex) - 1; $i += 2)
        $string .= chr(hexdec($hex[$i].$hex[$i + 1]));
    return $string;
}
```