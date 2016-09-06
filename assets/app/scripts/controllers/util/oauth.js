'use strict';

/**
 * @ngdoc function
 * @name openshiftConsole.controller:OAuthController
 * @description
 * # OAuthController
 * Controller of the openshiftConsole
 */
angular.module('openshiftConsole')
  .controller('OAuthController', function ($scope, $location, $q, RedirectLoginService, DataService, AuthService, Logger) {
    var authLogger = Logger.get("auth");

    // Initialize to a no-op function.
    // Needed to let the view confirm a login when the state is unverified. 
    $scope.completeLogin = function(){};
    $scope.cancelLogin = function() {
      $location.replace();
      $location.url("./");
    };

    RedirectLoginService.finish()
    .then(function(data) {
      var token = data.token;
      var then = data.then;
      var verified = data.verified;
      var ttl = data.ttl;

      // Try to fetch the user
      var opts = {errorNotification: false, http: {auth: {token: token, triggerLogin: false}}};
      authLogger.log("OAuthController, got token, fetching user", opts);

      DataService.get("users", "~", {}, opts)
      .then(function(user) {
        // Set the new user and token in the auth service
        authLogger.log("OAuthController, got user", user);


        //daoVoice
        daovoice('init', {
          app_id: "b31d2fb1",
          user_id: "user.metadata.uid", // 必填: 该用户在您系统上的唯一ID
          //email: "daovoice@example.com", // 选填:  该用户在您系统上的主邮箱
          name: user.metadata.name, // 选填: 用户名
          signed_up: parseInt((new Date(user.metadata.creationTimestamp)).getTime() / 1000) // 选填: 用户的注册时间，用Unix时间戳表示
        });
        daovoice('update');


        $scope.completeLogin = function() {
          // Persist the user
          AuthService.setUser(user, token, ttl);
          
          // Redirect to original destination (or default to './')
          var destination = then || './';
          if (URI(destination).is('absolute')) {
            authLogger.log("OAuthController, invalid absolute redirect", destination);
            destination = './';
          }
          authLogger.log("OAuthController, redirecting", destination);
          $location.replace();
          $location.url(destination);
        };
        
        if (verified) {
          // Automatically complete
          $scope.completeLogin();
        } else {
          // Require the UI to prompt
          $scope.confirmUser = user;
          
          // Additionally, give the UI info about the user being overridden
          var currentUser = AuthService.UserStore().getUser();
          if (currentUser && currentUser.metadata.name !== user.metadata.name) {
            $scope.overriddenUser = currentUser;
          }
        }
      })
      .catch(function(rejection) {
        // Handle an API error response fetching the user
        var redirect = URI('error').query({error: 'user_fetch_failed'}).toString();
        authLogger.error("OAuthController, error fetching user", rejection, "redirecting", redirect);
        $location.replace();
        $location.url(redirect);

        //daoVoice
        daovoice('init', {
          app_id: "b31d2fb1"
        });
        daovoice('update');
      });

    })
    .catch(function(rejection) {
      var redirect = URI('error').query({
        error: rejection.error || "",
        error_description: rejection.error_description || "",
        error_uri: rejection.error_uri || ""
      }).toString();
      authLogger.error("OAuthController, error", rejection, "redirecting", redirect);
      $location.replace();
      $location.url(redirect);

      //daoVoice
      daovoice('init', {
        app_id: "cd9644b0"
      });
      daovoice('update');
    });

  });
