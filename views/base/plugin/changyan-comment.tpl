<!--兼容版，可保证页面完全兼容-->
<div id="SOHUCS"></div>
<script>
  (function(){
    var appid = 'cyr6kMmSR',
    conf = 'prod_622dd664eb50cc4d5dde39d815742821';
    var doc = document,
    s = doc.createElement('script'),
    h = doc.getElementsByTagName('head')[0] || doc.head || doc.documentElement;
    s.type = 'text/javascript';
    s.charset = 'utf-8';
    s.src =  'http://assets.changyan.sohu.com/upload/changyan.js?conf='+ conf +'&appid=' + appid;
    h.insertBefore(s,h.firstChild);
  })()
</script>