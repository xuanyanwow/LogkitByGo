# Logkig By Go


## php封包代码

```php
<?php


$sendStr = encode("132");

$sendStrArray = str_split(str_replace(' ', '', $sendStr), 2);  // 将16进制数据转换成两个一组的数组

$socket = socket_create(AF_INET, SOCK_STREAM, getprotobyname("tcp"));  // 创建Socket

if (socket_connect($socket, "127.0.0.1", 8999)) {  //连接
    $send = '';
    for ($j = 0; $j < count($sendStrArray); $j++) {
        $send .= chr(hexdec($sendStrArray[$j]));
    }
    socket_write($socket, $send);
} else {
    echo "connection error";
}
socket_close($socket);  // 关闭Socket


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