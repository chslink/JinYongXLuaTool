# 金庸群侠传X 脚本加密解密工具

## 使用办法

金庸群侠传X脚本加解密工具
无双/后宫 
目前支持111 119,其他版本没测试,主要魔改版本也太多了


下载 jxt.exe 放到lua/xml 脚本所在文件夹。 直接执行就好，记得注意备份原有文件。

或者用--help查看高级参数 制定路径/版本/密钥等
 ```shell
  # 查看帮助
 jxt.exe --help 
 # 加密脚本
 jxt.exe -mode e
 # 解密脚本 
 jxt.exe -mode d
 ```
## 20240716更新
重构了解密部分逻辑
新增支持版本 `20230909` 加密暂时不可用，有点懒得写
