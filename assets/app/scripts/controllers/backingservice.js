'use strict';

/**
 * @ngdoc function
 * @name openshiftConsole.controller:BackingserviceController
 * @description
 * Controller of the openshiftConsole
 */
angular.module('openshiftConsole')
  .controller('BackingserviceController', function ($scope, $routeParams, AuthService, DataService, $filter) {


    $scope.breadcrumbs = [
      {
        title: "Backing Services",
        link: "backingservices"
      }
    ];

    $scope.breadcrumbs.push({
      title: $routeParams.backingservice
    });

    AuthService.withUser().then(function() {
      loadBackingService();
    });

    var loadBackingService = function() {
      $scope.namespace = 'openshift';
      DataService.get("backingservices", $routeParams.backingservice, $scope).then(
        // success
        function(backingservice) {
          $scope.loaded = true;
          $scope.backingservice = backingservice;
          console.log('backingservice', backingservice);

          var buildNumber = $filter("annotation")(backingservice, "buildNumber");
          if (buildNumber) {
            $scope.breadcrumbs[2].title = "#" + buildNumber;
          }
        },
        // failure
        function(e) {
          $scope.loaded = true;
          $scope.alerts["load"] = {
            type: "error",
            message: "The build details could not be loaded.",
            details: "Reason: " + $filter('getErrorDetails')(e)
          };
        }
      );
    };
  });
