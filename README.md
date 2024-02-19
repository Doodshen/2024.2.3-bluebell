<!--
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-01-30 21:03:50
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-01-31 20:55:12
 * @FilePath: \web-app2\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
替代v1版本，使用结构体来绑定配置文件信息 

1. 使用viper读取配置文件信息，然后将配置信息反序列化到结构体中 
* setting模块中定义结构体，对应相应的配置信息 


2. 在main函数中将各个步骤中使用到的配置信息，使用settint.conf进行传递 

3. AppConfig表示整个项目中用到的配置，
* 不管使用的那种文件的配置信息，使用结构体type的方式将结构体中的字段和配置文件中的配置信息进行对应
* 所有的统一使用mapstrcuce结构体type
* viper函数调用onconfigchange进行回调，钩子函数，回调机制。在viper中使用OnConfigChange函数进行回调机制 



4. version3 viper读取配置文件的方法 、
* 方法1：指定配置文件路径(相对路径或者绝对路径)
* 方法2：指定配置文件名和配置文件的位置 


5. version4 
viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）：
上面两个函数查找的配置文件路径是执行该项目是的相对路径，
* 解决：解析命令行参数，指定配置文件路径 
* 1实现：setting.go中init（filename string）使用viper方式1指定配置文件路径来获取配置文件 
       filename是使用os.arg从命令行中读取配置文件路径 
* 2实现：filename使用flag包进行读取配置信息 
