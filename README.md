# generation
生成代码工具, 帮助生成业务代码



template 里可以
tpl/ 下放模板
构造一个全的Struct 然后赋值，根据不同的模板用不同的值，

这里有的首字母大写和小写的问题，要想一下怎么处理比较合适

先分开写，后期再和在一起写

解析参数这个controller 写两个，一种直接用model 一种创建req 和 reply


生成代码工具, 帮助生成业务代码

1. 输入项目名称，生成项目结构基础结构
2. 生成启动项，以及启动项的基础结构，包括配置文件，基础组件.go和业务组件.go
3. 输入模块名，生成模块，生成项关联模块，在启动项中业务组件中生成使用代码
4. 根据一个规范生成API ,
5. 根据API 使用类似 swagger  生成API文档

1. 项目级别生成： 文件夹生成，是否生成main模块
2. 模块生成：输入模块名称，然后生成对应文件夹，以及选择生成模块的方式
3. 层次生成，生成MVC的某一层，或者是
4. 文件生成，选择某种模板，生成对应的文件
5. 互相之间进行关联，然后放入指定的路径中


etcd gorm 的超时时间直接在程序里进行写，不弄在外面了

* 项目：project
* 应用: application
* 微服务: service
* 单页：page 
* 模板：模板为go模板文件，默认模板写成常量不允许修改

