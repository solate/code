

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

| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| doctorId   | string | 医生id | 是 |       |
| doctorName   | string | 医生姓名 | 否 |       |

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

__single__ //返回值类型是单层结构


| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| doctorName   | string | 医生姓名 | 是 |       |
| doctorAge   | string | 医生年龄 | 否 |       |

### 患者详情

* Note

获得某个患者的基础信息，

* URL

/patient/info

* Method

 GET

* Request

| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| patientId   | string | 患者id | 是 |       |
| patientName   | string | 患者姓名 | 否 |       |

* Response

```
{
    "data": {
       list : [{
            "id" : 1,
            "avatar" : "头像图片" //xxx
       },{
            "id" : 1,
            "avatar" : "头像图片" //xxx
       }]

    },
   "errmsg": "success",
   "errcode": 200
}


```

__list__ //返回值为list结构

| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| patientName   | string | 门诊号 | 是 |       |
| patientAge   | string | 患者年龄 | 否 |       |
