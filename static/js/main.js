$(document).ready(function(){
  $.ajaxSetup({
  		cache:false,//prevent browser cache result to redirect  failed.
  		headers:{"edgex-club-token":window.localStorage.getItem("edgex-club-token")},
  		statusCode: {
  			302: function() {
  				window.location.href='/login.html?ran='+Math.random(); //prevent browser cache result to redirect  failed.
  			}
  		}
  	});
  edgexClubMainModule.loginIsVaild = edgexClubMainModule.checkLogin();
  //edgexClubMainModule.loadArticleList();
  $("div.header_login").on("click",function(){
    mainModuleBtnGroup.login();
  });
  $("div.header_user").on("click",function(){
    mainModuleBtnGroup.user();
  });
  $("div.header_post").on("click",function(){
    mainModuleBtnGroup.post();
  });
  // $("div.header_user").mouseover(function(){
  //
  //   $("div.user_hide_info").show()
  // });
  // $("div.header_user").mouseleave(function(){
  //   $("div.user_hide_info").hide()
  // });

  if (edgexClubMainModule.loginIsVaild) {
    $.ajax({
      url:"/api/v1/message/count",
      type:"GET",
      success:function(data){
        if(data == 0){
          return;
        }else if(data > 100){
          $("div.header_user span.badge").text(data+"+");
        }else{
          $("div.header_user span.badge").text(data);
        }
        $("div.header_user span.badge").show();
      }
    });
  }
});

var edgexClubMainModule = {
  loginIsVaild:false,
  checkLogin:function(){
    var token = window.localStorage.getItem("edgex-club-token");
    var user = JSON.parse(window.localStorage.getItem("edgex-club-userInfo"));
    var isVaild = false;
    if(token){
      $.ajax({
        url:"/api/v1/isvalid/"+token,
        type:"GET",
        async:false,
        success:function(data){
          if (data == 1) {
            isVaild = true;
            $("div.header_login").hide();
            $("div.header_user img").prop("src",user["avatarUrl"])
            $("div.header_user").show();
          }else{
            isVaild = false;
          }
        }
      });
      return isVaild;
    }else{
      $("div.header_user").hide();
      return isVaild;
    }
  }
}

var mainModuleBtnGroup = {
  login:function(){
    window.location.href = "/login.html"
  },
  user:function(){
    if(edgexClubMainModule.loginIsVaild){
      var user = JSON.parse(window.localStorage.getItem("edgex-club-userInfo"));
      window.location.href = "/users/" + user["name"];
    }else{
      window.location.href = "/login.html"
    }
  },
  post:function(){
    if(edgexClubMainModule.loginIsVaild){
      var user = JSON.parse(window.localStorage.getItem("edgex-club-userInfo"));
      window.location.href = "/article/edit/"+user["name"]+"/new";
    }else{
      window.location.href = "/login.html"
    }
  }
}

testData=[{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/article.html",
  created:1520996873655,
},{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/article.html",
  created:1520996873655,
},{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/article.html",
  created:1520996873655,
},{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/test.html",
  created:1520996873655,
},{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/test.html",
  created:1520996873655,
},{
  id:"5aa89209e4b01d97205d7f4c",
  title:"这个是文章的title",
  path:"/articles/test.html",
  created:1520996873655,
}
]
