angular.module('findpaswd', [])
  .controller('findpswdController', function($scope, $http) {

      $scope.getcaptcha = function(){

        var mobile = $scope.userphone;
        if(mobile == undefined ||mobile.length  != 11){
          $.alert("输入的手机号码有误,请重新输入");
          return;
        }

        $http({
          url:'/ranbb/api',
          method:'POST',
          data: {
            CID: 91001,
            PL: {
              Mobile: $scope.userphone
            }
          }
        }).success(function(data,header,config,status){
          $.alert(data.PL.Captcha);
        }).error(function(data,header,config,status){
          $.alert(data);
        });
      };

      $scope.commit = function(){
        if($scope.userphone == undefined ||$scope.userphone.length  != 11){
          $.alert("手机号码有误,请重新输入");
          return;
        }
        if($scope.idcard == undefined ||$scope.idcard.length  != 18){
          $.alert("身份证号码错误,请重新输入");
          return;
        }
        if($scope.captcha == undefined ||$scope.captcha.length  != 6){
          $.alert("验证码错误");
          return;
        }

        $http({
          url:'/ranbb/api',
          method:'POST',
          data: {
            CID: 91051,
            PL: {
              Mobile: $scope.userphone,
              IdCard:$scope.idcard,
              Captcha:$scope.captcha
            }
          }
        }).success(function(data,header,config,status){
          showData(data.RC);
        }).error(function(data,header,config,status){
          $.alert("网络访问失败");
        });


      };

      function showData(data){
        if (data == 0){
          $.alert('新密码已经发送到您的手机,请返回登陆');
        }
        if (data == 1001){
          $.alert("数据有误,请检查自己的数据")
        }
        if (data == 1003){
          $.alert("发送失败")
        }
        if (data == 1004){
          $.alert("验证码有误")
        }
        if (data == 1005){
          $.alert("错误的身份证")
        }
        if (data == 1010){
          $.alert("用户不存在")
        }
        if (data == 1031){
          $.alert("身份证不符")
        }
        if (data == 1032){
          $.alert("找回密码失败")
        }
      }
  });
