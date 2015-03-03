##后台页面加载数据的时候不采用load，改用后台返回json，前端javascript解析json渲染。
##后台管理逐渐会采用API模式，基本只返回josn数据，这样PC端以及以后的手机端，包括开放部分资源的时候，可以统一。


# 建库要保证utf8格式的
# 关于不变参数传递的问题,参见[contentservice.GetContentList]中sort
# json自定义结构名称通过 `json: "xxx"`来定义，参见[Mediacontroller.UploadIn]