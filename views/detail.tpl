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
    <div class="col-md-9 bj-left" role="main">

      <div class="page-header">
        <h1 id="type">{{str2html .content.ContentTitle}}</h1>
        <small> <b>发布者：</b>{{.content.CreateId | getUsernameById }}&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <b>发布：</b>{{date .content.CreateTime "Y-m-d H:i:s" }}</small> </div>
      <div>
      <a title=" {{.content.ContentTitle}} " target="_blank" href="#"><img src="{{if .content.MediaPath}}{{.content.MediaPath}}{{else}}img/biaoti-pic.jpg{{end}}" class="img-responsive" alt="Responsive image"></a>
    </div>
      <div class="bj-article">
        {{str2html .content.ContentMain}}
      </div>
      <p><div class="bshare-custom"><a title="分享到QQ空间" class="bshare-qzone"></a><a title="分享到新浪微博" class="bshare-sinaminiblog"></a><a title="分享到人人网" class="bshare-renren"></a><a title="分享到腾讯微博" class="bshare-qqmb"></a><a title="分享到网易微博" class="bshare-neteasemb"></a><a title="更多平台" class="bshare-more bshare-more-icon more-style-sharethis"></a><span class="BSHARE_COUNT bshare-share-count">0</span></div><script type="text/javascript" charset="utf-8" src="http://static.bshare.cn/b/buttonLite.js#style=-1&amp;uuid=&amp;pophcol=2&amp;lang=zh"></script><script type="text/javascript" charset="utf-8" src="http://static.bshare.cn/b/bshareC0.js"></script></p>
      <div class="bj-bl20"></div>
      {{template "base/plugin/changyan-comment.tpl"}}
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
<script src="/js/application.js"></script>
</body>
</html>