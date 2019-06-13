## 接口文档

注：/api/v1，/admin的路由会有登录检验，在头部带上token，debug的时候可以在router中把apiv1.Use(jwt.JWT())注释掉就可以不用中间件。

* 获取token(后期修改)  GET
/user?username=xxx&password=xxx
    * Response:
        code:
        msg:
        data: token:xxx

* 用户注册  POST 
/user
    * Request:
        username:xxx
        password:xxx
        e_mail:xxx
        telephone_number:xxx
    * Response:
        code:
        msg:
        data:

* 修改个人信息 PUT
/api/v1/user
    * Request:
        username:xxx
        password:xxx
        e_mail:xxx
        telephone_number:xxx
    * Response:
        code:
        msg:

* 用户获取自身信息 GET
/api/v1/user
    * Response
        {
    "code": 200,
    "data": {
        "id": 2,
        "username": "whiskey",
        "password": "123456",
        "role": 1,
        "telephone_number": "",
        "e_mail": ""
    },
    "msg": "ok"
}

* 获取第n页的产犊列表  GET  
当没有page时会返回全部数据
/api/v1/calves?page=n
    * Response:
        {
    "code": 200,
    "data": {
        "list": [
            {
                "id": 3,
                "cow_id": 7,
                "fetus_organ": "腿",
                "is_complete": 1,
                "is_abortion": 2,
                "yak_id": 11,
                "yak_index": 1,
                "milk_production": 2,
                "cream": 3,
                "protein": 4,
                "birth_time": "2019-04-21 14:46:43",
                "flowing_time": "2019-04-23 14:46:43",
                "fetus_time": "2019-04-23 14:46:43",
                "fetus_birth_time": "2019-04-23 14:46:43",
                "placenta_time": "2019-04-23 14:46:43"
            },
            {
                "id": 4,
                "cow_id": 101,
                "fetus_organ": "head",
                "is_complete": 1,
                "is_abortion": 2,
                "yak_id": 11,
                "yak_index": 1,
                "milk_production": 2,
                "cream": 3,
                "protein": 4,
                "birth_time": "2019-04-21 14:46:43",
                "flowing_time": "2019-04-23 14:46:43",
                "fetus_time": "2019-04-23 14:46:43",
                "fetus_birth_time": "2019-04-23 14:46:43",
                "placenta_time": "2019-04-23 14:46:43"
            }
        ],
        "total": 2
    },
    "msg": "ok"
}

* 通过cow_i获取产犊记录     GET
/api/v1/calves/:cow_id
    * Response:
        {
    "code": 200,
    "data": [
        {
            "id": 4,
            "cow_id": 101,
            "fetus_organ": "head",
            "is_complete": 1,
            "is_abortion": 2,
            "yak_id": 11,
            "yak_index": 1,
            "milk_production": 2,
            "cream": 3,
            "protein": 4,
            "birth_time": "2019-04-21 14:46:43",
            "flowing_time": "2019-04-23 14:46:43",
            "fetus_time": "2019-04-23 14:46:43",
            "fetus_birth_time": "2019-04-23 14:46:43",
            "placenta_time": "2019-04-23 14:46:43"
        }
    ],
    "msg": "ok"
}

* 增加产犊记录 POST
/api/v1/calves
    * Request
        产犊记录里除了id的所有内容，字段名去看下model的后面的json标示中的字段。
    * Response
        code:
        msg:


