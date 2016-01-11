# API

[TOC]

## 获取验证码

``` 
{
  "CID": 91001,
  "PL": {
    "mobile": "18601761051"
  }
}

{
  "CID": 91002,
  "RC": 0,
  "PL": "799523"
}
```

## 注册

``` 
{
  "PL": {
    "Mobile": "18601761051",
    "PassWord": "123456",
    "RealName": "王猛",
    "IdCard": "412723198906060833",
    "AliyPayId": "8986462467@qq.com",
    "Captcha": "799523",
    "AliyPayName": "demon"
  },
  "CID": 91011
}

{
  "CID": 91012,
  "RC": 1004,
  "PL": null
}
```

## 登陆

``` 
{
  "CID": 91021,
  "PL": {
    "Mobile": "18601761052",
    "PassWord": "123456"
  }
}

{
  "CID": 91022,
  "RC": 0,
  "PL": {
    "SID": "8287168ee1440159d4ac6ac30439b0cd"
  }
}
```

## 登出

``` 
{
  "SID": "ec8a98ebe2d7ad2345ba803b5a7badb9",
  "CID": 91031
}

{
  "CID": 91032,
  "RC": 0
}
```



## 修改密码

``` 
{  "SID": "0d034b86508356706d6f9e020e7180fe",
  "CID": 91061,
  "PL": {
    "OldPswd": "123456",
	"NewPswd":"123456"
  }
}

{
  "CID": 91062,
  "RC": 0
}
```



## 查询用户信息

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91041
}

{
  "CID": 91042,
  "RC": 0,
  "PL": {
    "UID": "308101e05fee4a662166cfe9c8eab0ab",
    "Mobile": "18601761052",
    "RealName": "王猛",
    "IdCard": "41270219890510154X",
    "AliPayAccount": "89864624@qq.com",
    "AliPayName": "demon",
    "Active": 0,
    "Asset": 0,
    "Rate": 0,
    "Income": 0,
    "Total": 0,
    "CreateTime": 1451489343,
    "UpdateTime": 1451489343
  }
}
```

## 获取商品列表

``` 
{
  "SID": "8287168ee1440159d4ac6ac30439b0cd",
  "CID": 91101,
  "PL": {
    "Page": 1,
    "Size": 10
  }
}


{
  "CID": 91102,
  "RC": 0,
  "PL": {
    "Count": 2,
    "data": [
      {
        "GoodsId": 1,
        "ShopId": 1,
        "ShopName": "我的店",
        "State": 0,
        "Price": 100,
        "RequireLevel": 3,
        "ShopRequire": "买买卖阿",
        "ImageUrl": "http://www.baidu.com",
        "BrokerAge": 5,
        "CreateTime": 0,
        "UpdateTime": 0,
        "Quantity": 54,
        "LimitPurchaseQuantity": 1,
        "Memo": ""
      },
      {
        "GoodsId": 2,
        "ShopId": 1,
        "ShopName": "我的店",
        "State": 1,
        "Price": 100,
        "RequireLevel": 3,
        "ShopRequire": "买买卖阿",
        "ImageUrl": "http://www.baidu.com",
        "BrokerAge": 5,
        "CreateTime": 0,
        "UpdateTime": 0,
        "Quantity": 54,
        "LimitPurchaseQuantity": 1,
        "Memo": ""
      }
    ]
  }
}
```



## 获取商品分类列表

``` 
{
  "SID": "29ee4eed583aafc07e8d0143a90bef48",
  "CID": 91161,
  "PL": {
    "GoodsId": 1
  }
}

{
  "CID": 91162,
  "RC": 0,
  "PL": {
    "Count": 2,
    "Data": [
      {
        "CategroyId": 1,
        "GoodsId": 1,
        "ShopId": 1,
        "Price": 100,
        "Name": "大号",
        "EnableTime": 0,
        "TotalNum": 10,
        "OutNum": 5,
        "Memo": "我了个去你好好刷",
        "LimitPurchaseQuantity": 1
      },
      {
        "CategroyId": 2,
        "GoodsId": 1,
        "ShopId": 1,
        "Price": 120,
        "Name": "消耗",
        "EnableTime": 0,
        "TotalNum": 10,
        "OutNum": 4,
        "Memo": "好好干 挣钱过年\n",
        "LimitPurchaseQuantity": 1
      }
    ]
  }
}
```



## 购买商品

``` 
{
  "SID": "fb35b6bfb87dd381c2d65096c5ea7691",
  "CID": 91111,
  "PL": {
    "GoodsId": 6,
    "Count": 1
  }
}

{
  "CID": 91112,
  "RC": 0,
  "PL": {
    "OrderId": 8,
    "GoodsId": 2,
    "ShopId": 1,
    "UID": "308101e05fee4a662166cfe9c8eab0ab",
    "TaoBaoAccount": "",
    "State": 0,
    "ShopName": "我的店",
    "Price": 100,
    "RequireLevel": 3,
    "ShopRequire": "买买卖阿",
    "ImageUrl": "http://www.baidu.com",
    "BrokerAge": 5,
    "CreateTime": 1451621997,
    "UpdateTime": 1451621997,
    "Quantity": 1,
    "Memo": "接了一单"
  }
}
```

## 提交审核订单

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91121,
  "PL": {
    "OrderId": 4,
    "TaoBaoAccount": "838532366@qq.com"
  }
}

{
  "CID": 91122,
  "RC": 0
}
```

## 获取订单信息

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91121,
  "PL": {
    "OrderId": 4,
    "TaoBaoAccount": "838532366@qq.com"
  }
}

{
  "CID": 91142,
  "RC": 0,
  "PL": {
    "OrderId": 4,
    "GoodsId": 4,
    "ShopId": 1,
    "UID": "308101e05fee4a662166cfe9c8eab0ab",
    "TaoBaoAccount": "",
    "State": 1,
    "ShopName": "我的店",
    "Price": 100,
    "RequireLevel": 3,
    "ShopRequire": "买买卖阿",
    "ImageUrl": "http://www.baidu.com",
    "BrokerAge": 5,
    "CreateTime": 1451562272,
    "UpdateTime": 1451562272,
    "Quantity": 1,
    "Memo": "接了一单"
  }
}
```

## 获取订单列表

>  0: 正在完成的订单 1:正在审核的订单 2:已成功的订单 3:审核失败订单

``` 
{  
	"SID": "40757c560f2506beb30be23401917555",
  "CID": 91151,
  "PL": {
    "State": 1
  }
}

{
  "CID": 91142,
  "RC": 0,
  "PL": {
    "OrderId": 4,
    "GoodsId": 4,
    "ShopId": 1,
    "UID": "308101e05fee4a662166cfe9c8eab0ab",
    "TaoBaoAccount": "",
    "State": 1,
    "ShopName": "我的店",
    "Price": 100,
    "RequireLevel": 3,
    "ShopRequire": "买买卖阿",
    "ImageUrl": "http://www.baidu.com",
    "BrokerAge": 5,
    "CreateTime": 1451562272,
    "UpdateTime": 1451562272,
    "Quantity": 1,
    "Memo": "接了一单"
  }
}{
  "CID": 91152,
  "RC": 0,
  "PL": {
    "Count": 2,
    "Data": [
      {
        "OrderId": 2,
        "GoodsId": 4,
        "ShopId": 1,
        "UID": "308101e05fee4a662166cfe9c8eab0ab",
        "TaoBaoAccount": "",
        "State": 1,
        "ShopName": "我的店",
        "Price": 100,
        "RequireLevel": 3,
        "ShopRequire": "买买卖阿",
        "ImageUrl": "http://www.baidu.com",
        "BrokerAge": 5,
        "CreateTime": 1451562097,
        "UpdateTime": 1451562097,
        "Quantity": 1,
        "Memo": "接了一单"
      },
      {
        "OrderId": 4,
        "GoodsId": 4,
        "ShopId": 1,
        "UID": "308101e05fee4a662166cfe9c8eab0ab",
        "TaoBaoAccount": "",
        "State": 1,
        "ShopName": "我的店",
        "Price": 100,
        "RequireLevel": 3,
        "ShopRequire": "买买卖阿",
        "ImageUrl": "http://www.baidu.com",
        "BrokerAge": 5,
        "CreateTime": 1451562272,
        "UpdateTime": 1451562272,
        "Quantity": 1,
        "Memo": "接了一单"
      }
    ]
  }
}
```

## 删除订单

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91131,
  "PL": {
    "OrderId": 4
  }
}

{
  "CID": 91132,
  "RC": 0
}
```

## 获取淘宝账号列表

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91201
}

{
  "CID": 91202,
  "RC": 0,
  "PL": [
    {
      "Tid": 1,
      "UID": "308101e05fee4a662166cfe9c8eab0ab",
      "TaoBaoAccount": "838532366@qq.com",
      "CreateTime": 1451576420,
      "UpdateTime": 1451576420,
      "Memo": "wode taobao"
    },
    {
      "Tid": 2,
      "UID": "308101e05fee4a662166cfe9c8eab0ab",
      "TaoBaoAccount": "838532367@qq.com",
      "CreateTime": 1451576448,
      "UpdateTime": 1451576448,
      "Memo": "wode taobao"
    },
    {
      "Tid": 3,
      "UID": "308101e05fee4a662166cfe9c8eab0ab",
      "TaoBaoAccount": "838532368@qq.com",
      "CreateTime": 1451576452,
      "UpdateTime": 1451576452,
      "Memo": "wode taobao"
    }
  ]
}
```

## 更新淘宝账号

``` 
{  
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91211,
  "PL": {
    "Tid": 2,
	"TaoBaoAccount":"1213132@qq.com",
    "Memo":"我的淘宝账号"
  }
}
{
  "CID": 91212,
  "RC": 0
}
```

## 删除淘宝账号

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91221,
  "PL": {
    "Tid": 2
  }
}

{
  "CID": 91222,
  "RC": 0
}
```

## 添加淘宝账号

``` 
{
  "SID": "40757c560f2506beb30be23401917555",
  "CID": 91231,
  "PL": {
    "TaoBaoAccount": "838532368@qq.com",
    "Memo": "wode taobao"
  }
}

{
  "CID": 91232,
  "RC": 0
}
```

## 获取钱包历史列表

``` 
{
  "SID": "8287168ee1440159d4ac6ac30439b0cd",
  "CID": 91301,
  "PL": {
    "Mode": 1,
    "Page": 1,
    "Size": 10
  }
}

{
  "CID": 91302,
  "RC": 0,
  "PL": {
    "Count": 1,
    "Data": [
      {
        "Wid": 1,
        "UID": "308101e05fee4a662166cfe9c8eab0ab",
        "Amount": 10,
        "Categroy": 1,
        "CreateTime": 0,
        "Memo": "我曹了\n"
      }
    ]
  }
}
```