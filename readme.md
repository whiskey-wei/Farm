# 项目文档

## 运维
* 配置好golang环境，并安装好dep工具，执行go get github.com/whiskey-wei/Farm
* 创建farm数据库执行项目docs文件下的sql文件，创建表，并设置conf文件夹下的配置文件。
* 在项目目录下执行go run main.go项目开始运行

## 项目文件结构
* conf  
    配置文件
* image  
    一些截图
* middleware  
    登录中间件
* model  
    数据库模型
* pkg    
    * e  
        错误码和错误信息
    * seeting    
        读取配置文件
    * util  
        一些公用工具
* router  
    路由以及对应的api实现
* vendor  
    依赖包

## 相关接口  
		//产犊记录相关接口
		apiv1.GET("/calves", v1.GetCalves)
		apiv1.GET("/calves/:cow_id", v1.GetCalve)
		apiv1.POST("/calves", v1.AddCalve)
		apiv1.PUT("/calves/:id", v1.UpdateCalve)
		apiv1.DELETE("/calves/:id", v1.DeleteCalve)

		//交配记录相关接口
		apiv1.GET("/breeds", v1.GetBreeds)
		apiv1.GET("/breeds/:cow_id", v1.GetBreed)
		apiv1.POST("/breeds", v1.AddBreed)
		apiv1.PUT("/breeds/:id", v1.UpdateBreed)
		apiv1.DELETE("/breeds/:id", v1.DeleteBreed)

		//妊娠诊断记录相关接口
		apiv1.GET("/pregnancy", v1.GetPregnancys)
		apiv1.GET("/pregnancy/:cow_id", v1.GetPregnancy)
		apiv1.POST("/pregnancy", v1.AddPregnancy)
		apiv1.PUT("/pregnancy/:id", v1.UpdatePregnancy)
		apiv1.DELETE("/pregnancy/:id", v1.DeletePregnancy)

		//产奶记录相关接口
		apiv1.GET("/milk", v1.GetMilkProductions)
		apiv1.GET("/milk/:cow_id", v1.GetMilkProduction)
		apiv1.POST("/milk", v1.AddMilkProduction)
		apiv1.PUT("/milk/:id", v1.UpdateMilkProduction)
		apiv1.DELETE("/milk/:id", v1.DeleteMilkProduction)

		//生长发育记录相关接口
		apiv1.GET("/yak", v1.GetYaks)
		apiv1.GET("/yak/:yak_id", v1.GetYak)
		apiv1.POST("/yak", v1.AddYak)
		apiv1.PUT("/yak/:id", v1.UpdataYak)
		apiv1.DELETE("/yak/:id", v1.DeleteYak)

		//DHI测定相关接口
		apiv1.GET("/dhi", v1.GetDhis)
		apiv1.GET("/dhi/:cow_id", v1.GetDhi)
		apiv1.POST("/dhi", v1.AddDhi)
		apiv1.PUT("/dhi/:id", apiv1.UpdateDhi)
		apiv1.DELETE("/dhi/:id", v1.DeleteDhi)

		//用户相关接口
		apiv1.PUT("/user", v1.UpdateUser)
		apiv1.GET("/user", v1.GetUserSelf)

