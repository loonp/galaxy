<header role="banner" class="navbar navbar-inverse navbar-fixed-top bs-docs-nav">
  <div class="container">
    <div class="navbar-header">
      <button data-target=".bs-navbar-collapse" data-toggle="collapse" type="button" class="navbar-toggle"> <span class="sr-only">Toggle navigation</span> <span class="icon-bar"></span> <span class="icon-bar"></span> <span class="icon-bar"></span> </button>
      <a class="navbar-brand" href="../">loonp.com</a> </div>
    <nav role="navigation" class="navbar-collapse bs-navbar-collapse in" style="height: auto;">
      <ul class="nav navbar-nav">
        <li><a href="/#">首页</a> </li>
        {{range $key,$val := .channels}} 
          <li><a href="/?channelId={{$val.Id}}">{{$val.ChannelName}}</a> </li>
        {{ end}}
      </ul>
      <ul class="nav navbar-nav navbar-right">
        <li> <a href="../about">关于我们</a> </li>
      </ul>
    </nav>
  </div>
</header>