## Shop Receipt

## What this project tries to resolve?

```
需求描述

商店里进行购物结算时会使用收银机系统，这台收银机会在结算时根据客户的购物车中的商品和商店正在进行的优惠活动进行结算和打印购物小票。

已知商品信息包含：名称，数量单位，单价，类别和条形码（伪）。 
已知我们可以对收银机进行设置，使之支持各种优惠。

我们需要实现一个名为打印小票的小模块，收银机会将输入的数据转换成一个JSON数据然后一次性传给我们这个小模块，我们将从控制台中输出结算清单的文本。

输入格式（样例）：

[
    'ITEM000001',
    'ITEM000001',
    'ITEM000001',
    'ITEM000001',
    'ITEM000001',
    'ITEM000003-2',
    'ITEM000005',
    'ITEM000005',
    'ITEM000005'
]

其中对'ITEM000003-2'来说,"-"之前的是标准的条形码,"-"之后的是数量。 
当我们购买需要称量的物品的时候,由称量的机器生成此类条形码,收银机负责识别生成小票。

该商店正在对部分商品进行“买二赠一”的优惠活动和对部分商品进行95折的优惠活动。其中：

“买二赠一”是指，每当买进两个商品，就可以免费再买一个相同商品。
“95折”是指，在计算小计的时候按单价的95%计算每个商品。
每一种优惠都详细标记了哪些条形码对应的商品可以享受此优惠。
店员设置，当“95折”和“买二赠一”发生冲突的时候，也就是一款商品既符合享受“买二赠一”优惠的条件，又符合享受“95折”优惠的条件时，只享受“买二赠一”优惠。

要求写代码支持上述的功能，并根据输入和设置的不同，输出下列小票。

小票内容及格式（样例）：

当购买的商品中，有符合“买二赠一”优惠条件的商品时：
***<没钱赚商店>购物清单***
名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：6.00(元)
名称：羽毛球，数量：5个，单价：1.00(元)，小计：4.00(元)
名称：苹果，数量：2斤，单价：5.50(元)，小计：11.00(元)
----------------------
买二赠一商品：
名称：可口可乐，数量：1瓶
名称：羽毛球，数量：1个
----------------------
总计：21.00(元)
节省：4.00(元)
**********************

 

当购买的商品中，没有符合“买二赠一”优惠条件的商品时：
***<没钱赚商店>购物清单***
名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：9.00(元)
名称：羽毛球，数量：5个，单价：1.00(元)，小计：5.00(元)
名称：苹果，数量：2斤，单价：5.50(元)，小计：11.00(元)
----------------------
总计：25.00(元)
**********************

 

当购买的商品中，有符合“95折”优惠条件的商品时
***<没钱赚商店>购物清单***
名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：9.00(元)
名称：羽毛球，数量：5个，单价：1.00(元)，小计：5.00(元)
名称：苹果，数量：2斤，单价：5.50(元)，小计：10.45(元)，节省0.55(元)
----------------------
总计：24.45(元)
节省：0.55(元)
**********************

 

当购买的商品中，有符合“95折”优惠条件的商品，又有符合“买二赠一”优惠条件的商品时
***<没钱赚商店>购物清单***
名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：6.00(元)
名称：羽毛球，数量：6个，单价：1.00(元)，小计：4.00(元)
名称：苹果，数量：2斤，单价：5.50(元)，小计：10.45(元)，节省0.55(元)
----------------------
买二赠一商品：
名称：可口可乐，数量：1瓶
名称：羽毛球，数量：2个
----------------------
总计：20.45(元)
节省：5.55(元)
**********************
```

### Install
Assuming that you have installed go1.5 and `$GOPATH` has been set.

```shell
go get github.com/Focinfi/shop_receipt/app/models
```

### Test
```shell
go test github.com/Focinfi/shop_receipt/app/models -v
```

### Project Structure
This project assumes that it will be used as a backend server.

1. `app/models`:
  1. `product.go` defines Product struct
  1. `lineitem.go` defines LineItem struct as line-item for receipt with subtotal, total calculation methods
  1. `promotion.go` defines Promotion and PromotionType for products 
  1. `favorable_lineitem.go` defines FavorableLineItems struct for line-items of the favorable products part in receipt
  1. `receipt.go` defines Receipt struct for shopping list
  1. `receipt_test.go` testing for receipt.go

1.  `app/views` contains all views:
  1. receipt.tmpl view template for Receipt

1. `config`:
  1. `app.go` contains global settings: current language and currency
  1. `tarnslation.go` contains translation map

1. `libs`:
  1. `utils.go` contains helper methods
  1. `view_funcMap.go` contains a map of methods used in template
  1. `translation.go` contains translate method
