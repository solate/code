

# 患者端

医生端基本 PATH： /patient

## 接口索引

* [医生详情](#医生详情)


## 接口定义

### 医生详情

* Note

获得某个医生的基础信息，

* URL

/doctor/info

* Method

 GET

* Request

| 字段                   | 类型     | 是否必填 | 说明          |
| --------------------  | ------ | ---- | ----------- |
| doctorId              | int | 是    | 医生id          |
| doctorName            | string | 是    | 医生姓名          |


* Response

```
{
    "data": {
        "id" : 1,
        "avatar" : "头像图片" //xxx
        "name" : "张三", //姓名

    },
   "errmsg": "success",
   "errcode": 200
}


```


| 字段                   | 类型     | 是否必填 | 说明          |
| --------------------  | ------ | ---- | ----------- |
| avatar              | int | 是    | 头像图片          |
| name          | string | 是    | 姓名          |




### 患者详情

* Note

获得某个患者的基础信息，

* URL

/patient/info

* Method

 GET

* Request

| 字段                   | 类型     | 是否必填 | 说明          |
| --------------------  | ------ | ---- | ----------- |
| patientId              | int | 是    | 医生id          |
| patientName            | string | 是    | 医生姓名          |


* Response

```
{
    "data": {
        "id" : 1,
        "avatar" : "头像图片" //xxx
        "name" : "张三", //姓名

    },
   "errmsg": "success",
   "errcode": 200
}


```


| 字段                   | 类型     | 是否必填 | 说明          |
| --------------------  | ------ | ---- | ----------- |
| patient              | int | 是    | 头像图片          |
| patient          | string | 是    | 姓名          |