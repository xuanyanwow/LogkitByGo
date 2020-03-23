# java上报数据

```java
package test;

import java.io.InputStream;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.Socket;

public class LogKit {

	protected static String ip = "127.0.0.1";
	protected static int portS = 8999;

	public static void main(String args[]) {

		// 数据json  看数据格式说明

		String json = "{\"action\":\"report\",\"data\":{\"project_id\":4,\"log_category\":\"Api/PayBase/v2\",\"log_point\":\"curl_begin\",\"log_sn\":\"666\",\"log_data\":\"{\\\"name\\\":\\\"siam\\\",\\\"age\\\":22}\",\"log_from\":\"siam\"}}";
		
//		String cmdInfor = "01 1f 7B 22 61 63 74 69 6F 6E 22 3A 22 61 70 69 73 22 2C 22 64 61 74 61 22 3A 5B 7B 22 70 72 6F 6A 65 63 74 5F 69 64 22 3A 31 2C 22 61 70 69 5F 63 61 74 65 67 6F 72 79 22 3A 22 77 65 63 68 61 74 22 2C 22 61 70 69 5F 6D 65 74 68 6F 64 22 3A 22 72 65 66 75 6E 64 22 2C 22 69 73 5F 73 75 63 63 65 73 73 22 3A 31 2C 22 63 6F 6E 73 75 6D 65 5F 74 69 6D 65 22 3A 31 35 2C 22 75 73 65 72 5F 66 72 6F 6D 22 3A 22 70 61 79 6D 65 6E 74 22 2C 22 75 73 65 72 5F 69 64 65 6E 74 69 66 79 22 3A 22 32 30 32 30 30 33 31 33 30 30 31 22 2C 22 61 70 69 5F 72 65 73 70 6F 6E 73 65 22 3A 22 E5 93 8D E5 BA 94 E5 86 85 E5 AE B9 20 E5 8F AF E9 80 89 22 2C 22 61 70 69 5F 70 61 72 61 6D 22 3A 22 7B 5C 22 74 65 73 74 5C 22 3A 5C 22 74 74 74 74 5C 22 7D 22 7D 5D 2C 22 73 69 67 6E 22 3A 22 63 36 37 33 34 38 61 62 62 33 66 34 36 66 34 32 65 32 61 30 33 35 33 33 38 31 63 34 38 63 34 34 22 7D ";

		int cmdLength = json.length();

      	String head = LengHead(cmdLength);
		String cmdInfor = str2HexStr(json);
		// 传输格式  长度头+内容

		send(head+cmdInfor);
	}
	
	public static String LengHead(int cmdLength) {
		int i = 0;
        char[] S = new char[4];
        
		while(cmdLength!=0)  
        {
            int t=cmdLength%16;
            if(t >=0 && t<10)  
            {
                S[i] = (char)(t+'0');
                i++;
            }
            else
            {
                S[i] = (char)(t+'A'-10);
                i++;
            }
            cmdLength=cmdLength/16;
        } 
		String tem = String.valueOf(S);
		
		return reverse2(tem);
	}

	public static String reverse2(String s) {
		int length = s.length();
		String reverse = "";
		for (int i = 0; i < length; i++){
			// siam 标注 判断char是否为空，如果为空就补0
			if (s.charAt(i) == '\0' ){
				reverse = "0" + reverse;
			}else{
				reverse = s.charAt(i) + reverse;
			}
			
		}
		return reverse;
	}

	public static String send(String cmdInfor){ 
		System.out.println(cmdInfor);
		String strReturn = null;
		try {  
			//要连接的服务端IP地址  
			String host = ip; 
			//要连接的服务端对应的监听端口  
			int port = portS;
			//将十六进制的字符串转换成字节数组
			byte[] cmdInfor2 = hexStrToBinaryStr(cmdInfor);
			
			//1.建立客户端socket连接，指定服务器位置及端口  
			Socket clientSocket =new Socket(host,port);  
			
			//2.得到socket读写流  
			OutputStream os=clientSocket.getOutputStream();  
			PrintWriter pw=new PrintWriter(os);  
			//输入流  
			InputStream is=clientSocket.getInputStream();  
			
			//3.利用流按照一定的操作，对socket进行读写操作  
			os.write(cmdInfor2);  
			os.flush();  
			clientSocket.shutdownOutput();  
			
			//接收服务器的响应 
			// int line = 0; 
			// byte[] buf = new byte[is.available()];
			
			//  //接收收到的数据  siam标注 这里只是上报 不需要接收
			// while((line=is.read(buf))!=-1){
			// 	//将字节数组转换成十六进制的字符串
			// 	strReturn= BinaryToHexString(buf);
			// }

			//4.关闭资源  
			is.close();  
			pw.close();  
			os.close();  
			clientSocket.close();
		}catch (Exception e){
			e.printStackTrace();
		}
		return strReturn;
	}  
	/**
    * 字符串转换成为16进制(无需Unicode编码)
    * @param str
    * @return
    */
    public static String str2HexStr(String str) {
        char[] chars = "0123456789ABCDEF".toCharArray();
        StringBuilder sb = new StringBuilder("");
        byte[] bs = str.getBytes();
        int bit;
        for (int i = 0; i < bs.length; i++) {
            bit = (bs[i] & 0x0f0) >> 4;
            sb.append(chars[bit]);
            bit = bs[i] & 0x0f;
            sb.append(chars[bit]);
            // sb.append(' ');
        }
        return sb.toString().trim();
    }

	/**
	 * 将十六进制的字符串转换成字节数组
	 *
	 * @param hexString
	 * @return
	 */
	public static byte[] hexStrToBinaryStr(String hexString) {
 

		hexString = hexString.replaceAll(" ", "");
 
		int len = hexString.length();
		int index = 0;
 
		byte[] bytes = new byte[len / 2];
 
		while (index < len) {
 
			String sub = hexString.substring(index, index + 2);
 
			bytes[index/2] = (byte)Integer.parseInt(sub,16);
 
			index += 2;
		}
 
 
		return bytes;
	}
	
	/**
	 * 将字节数组转换成十六进制的字符串
	 *
	 * @return
	 */
	public static String BinaryToHexString(byte[] bytes) {
	    String hexStr = "0123456789ABCDEF";
	    String result = "";
	    String hex = "";
	    for (byte b : bytes) {
	        hex = String.valueOf(hexStr.charAt((b & 0xF0) >> 4));
	        hex += String.valueOf(hexStr.charAt(b & 0x0F));
	        result += hex + " ";
	    }
	    return result;
	}
}
```