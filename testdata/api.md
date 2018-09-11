

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
| doctorId   | string | 门诊号 | 是 |       |


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


| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| doctorName   | string | 门诊号 | 是 |       |


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
| patientId   | string | 门诊号 | 是 |       |

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


| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| patientName   | string | 门诊号 | 是 |       |

