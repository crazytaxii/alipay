# 支付宝手机支付SDK API文档

## 公共参数
### 请求地址
|环境|HTTPS请求地址|
|---|---|
|正式环境|https://openapi.alipay.com/gateway.do|

### 公共请求参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|app_id|string|是|32|支付宝分配给开发者的应用ID|
|method|string|是|128|接口名称|
|format|string|否|40|仅支持JSON|
|charset|string|是|10|请求使用的编码格式，如utf-8,gbk,gb2312等|
|sign_type|string|是|10|商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2|
|sign|string|是|344|商户请求参数的签名串，详见签名|
|timestamp|string|是|19|发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"|
|version|string|是|3|调用的接口版本，固定为：1.0|
|app_auth_token|string|否|40|详见应用授权概述|
|biz_content|string|是||请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档|

### 公共响应参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|code|string|是|-|网关返回码,详见文档|
|msg|string|是|-|网关返回码描述,详见文档|
|sub_code|string|否|-|业务返回码，参见具体的API接口文档|
|sub_msg|string|否|-|业务返回码描述，参见具体的API接口文档|
|sign|string|是|-|签名,详见文档|

## app支付
### 请求参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|timeout_express|string|可选|6|该笔订单允许的最晚付款时间，逾期将关闭交易。|
|total_amount|string|可选|9|订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]|
|seller_id|string|可选|16|收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID|
|product_code|string|可选|64|销售产品码，商家和支付宝签约的产品码|
|body|string|可选|128|对一笔交易的具体描述信息。|
|subject|string|可选|256|商品的标题/交易标题/订单标题/订单关键字等。|
|out_trade_no|string|可选|64|商户网站唯一订单号|
|time_expire|string|可选||绝对超时时间，格式为yyyy-MM-dd HH:mm。|
|goods_type|string|可选|2|商品主类型:0-虚拟类商品,1-实物类商品|
|promo_params|string|可选|512|优惠参数|
|passback_params|string|可选|512|公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。|
|enable_pay_channels|string|可选|128|可用渠道，用户只能在指定渠道范围内支付|
|store_id|string|可选|32|商户门店编号|
|specified_channel|string|可选|128|指定渠道，目前仅支持传入pcredit。若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。|
|disable_pay_channels|string|可选|128|禁用渠道，用户不可用指定渠道支付当有多个渠道时用“,”分隔注，与enable_pay_channels互斥|
|business_params|string|可选|512|商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式|

### 响应参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|out_trade_no|string|必填|64|商户网站唯一订单号|
|trade_no|string|必填|64|该交易在支付宝系统中的交易流水号|
|total_amount|string|必填|9|该笔订单的资金总额|
|seller_id|string|必填|16|收款支付宝账号对应的支付宝唯一用户号。|

## 退款
### 请求参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|out_trade_no|string|特殊可选|64|订单支付时传入的商户订单号,不能和 trade_no同时为空。|
|trade_no|string|特殊可选|64|支付宝交易号，和商户订单号不能同时为空|
|refund_amount|string|必选|9|需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数|
|refund_currency|string|可选|8|订单退款币种信息|
|refund_reason|string|可选|256|退款的原因说明|
|out_request_no|string|可选|64|标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。|
|operator_id|string|可选|30|商户的操作员编号|
|store_id|string|可选|32|商户的门店编号|
|terminal_id|string|可选|32|商户的终端编号|
|goods_detail|GoodsDetail[]|可选||退款包含的商品列表信息，Json格式。其它说明详见：“商品明细说明”|
|refund_royalty_parameters|OpenApiRoyaltyDetailInfoPojo[]|可选||退分账明细信息|

### 响应参数
|参数|类型|是否必填|最大长度|描述|
|---|---|---|---|---|
|trade_no|string|必填|64|2013112011001004330000121536|
|out_trade_no|string|必填|64|商户订单号|
|buyer_logon_id|string|必填|100|用户的登录id|
|fund_change|string|必填|1|本次退款是否发生了资金变化|
|refund_fee|string|必填|11|退款总金额|
|refund_currency|string|选填|8|退款币种信息|
|gmt_refund_pay|timestamp|必填|32|退款支付时间|
|refund_detail_item_list|TradeFundBill|选填|
|store_name|string|选填|512|交易在支付时候的门店名称|
|buyer_user_id|string|必填|28|买家在支付宝的用户id|
refund_preset_paytool_list|PresetPayToolInfo|选填||退回的前置资产列表|
|present_refund_buyer_amount|string|选填|11|本次退款金额中买家退款金额|
|present_refund_discount_amount|string|选填|11|本次退款金额中平台优惠退款金额|
|present_refund_mdiscount_amount|string|选填|11|本次退款金额中商家优惠退款金额|

##
