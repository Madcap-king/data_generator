# data_generator
该项目用于快速生成mongo模拟数据，大幅提高开发测试效率


背景
该项目诞生的背景是日常开发工作中，涉及数据库测试数据的模拟生成，人工添加太过于低效、繁琐，故想要开发一个自动生成测试数据的工具，目前只支持mongo数据自动生成


应用场景
后端开发、测试过程中自动批量生成数据库测试数据

快速启动
gin搭建的web服务，正常web运行，通过http调用即可



demo

结构体：
![image](https://user-images.githubusercontent.com/77766374/225532329-ae1229c0-c936-406a-803a-ec98b03c27ee.png)

支持mongo数据类型：
![image](https://user-images.githubusercontent.com/77766374/225533962-1753a05f-cb1c-47c4-bb76-9fe77284ee29.png)


post方法：
![image](https://user-images.githubusercontent.com/77766374/225533580-46236f52-349f-41ea-9e03-60ba8fe19445.png)



todo
docker镜像（待做）

多种文件格式解析（待做）

多种模拟数据生成（待做）
