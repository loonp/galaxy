<!DOCTYPE html>
<html dir="ltr" lang="zh-CN">

{{template "base/head.tpl"}}

<body>
{{template "base/header.tpl" .}}

{{template "base/content-header.tpl"}}
<!--主体内容部分-->
<div class="container bj-container">
  <div class="row"> 
    <!--left内容部分 开始-->
    <div class="col-md-9 bj-left" role="main" id="left">
     {{range $key,$val := .contents}} 
      <div class="row bj-bom20 ">
        <div class="col-md-4"><a title=" {{$val.ContentTitle}} " target="_blank" href="#"><img src="{{if $val.MediaPath}}{{$val.MediaPath}}{{else}}img/biaoti-pic.jpg{{end}}" class="img-responsive" alt="Responsive image"></a></div>
        <div class="col-md-8">
          <h3 class="text-success"><a title="{{$val.ContentTitle}}" target="_blank" href="/detail?contentId={{$val.Id}}">{{$val.ContentTitle}}</a></h3>
          <div class="bj-bl10">loonp.com</div>
          <small><b>发布者：</b>{{$val.CreateId | getUsernameById }}&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <b>发布：</b>{{date $val.CreateTime "Y-m-d H:i:s" }}</small>
          <div class="bj-bl10">loonp.com</div>
          <p class="text-muted"><a target="_blank" href="/detail?contentId={{$val.Id}}">{{str2html $val.ContentSummary }}</a></p>
          <div class="bj-bl10">loonp.com</div>
          <P><b>标签：</b>{{$val.Id | getTagNamesByCid}}</P> 
        </div>
      </div>
      {{ end}}
      
      <div class="bs-example">
<!--         <ul class="pagination bj-fenye">
          <li class="disabled"><a href="#">«</a></li>
          <li class="active"><a href="#">1 <span class="sr-only">(current)</span></a></li>
          <li><a href="#">2</a></li>
          <li><a href="#">3</a></li>
          <li><a href="#">4</a></li>
          <li><a href="#">5</a></li>
          <li><a href="#">6</a></li>
          <li><a href="#">7</a></li>
          <li><a href="#">8</a></li>
          <li><a href="#">9</a></li>
          <li><a href="#">……</a></li>
          <li><a href="#">24</a></li>
          <li><a href="#">»</a></li>
        </ul> -->
        {{str2html .pagebar}}
      </div>
    </div>
    <!--right广告 列表部分 开始-->
<!--     <div class="col-md-3 bj-right">
      <div class="row" role="complementary">
        <div class="bs-sidebar hidden-print affix-top" role="complementary">
          <form class="form-inline" role="form">
            <div class="form-group">
              <input type="站内检索" class="form-control" id="exampleInputEmail2" placeholder="站内检索">
            </div>
            <button type="submit" class="btn btn-default">Search</button>
          </form>
          <div class="bj-bl20"></div>
          <p><img src="img/广告1.jpg" class="img-responsive" alt="Responsive image"></p>
          <p><img src="img/广告2.jpg" class="img-responsive" alt="Responsive image"></p>
        </div>
      </div>
    </div> -->
  </div>
</div>
{{template "base/footer.tpl"}}
<script src="http://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script> 
<script src="http://cdn.bootcss.com/twitter-bootstrap/3.0.3/js/bootstrap.min.js"></script> 
<script src="js/application.js"></script>
<script type="text/javascript">
function jump(page){
  $("#left").html("loading...");
  $("#left").load(
  "/",  //请求路径
  {"page":page});
}
</script>
</body>
</html>