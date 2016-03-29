'use strict';

/**
 * @ngdoc function
 * @name openshiftConsole.controller:BuildsController
 * @description
 * # ProjectController
 * Controller of the openshiftConsole
 */
angular.module('openshiftConsole')
  .controller('BackingServiceInstanceController', function ($scope, $routeParams, $filter, AuthService, DataService,ProjectsService) {
    $scope.projectName = $routeParams.project;
    $scope.emptyMessage = '没有数据';
    $scope.alerts = {};
    AuthService.withUser().then(function() {
      loadbsi();
    });

    $scope.breadcrumbs = [
      {
        title: "Backing Service Instances",
        link: "project/" + $routeParams.project + "/browse/backingserviceinstances"
      }
    ];

    $scope.breadcrumbs.push({
      title: $routeParams.backingserviceinstance
    });

    var loadbsi = function(){
      ProjectsService
      .get($routeParams.project)
      .then(_.spread(function(project, context) {
        $scope.project = project;

            DataService.get("backingserviceinstances", $routeParams.backingserviceinstance, context).then(
              // success
              function(bsi) {
                $scope.loaded = true;
                $scope.backingserviceinstance = bsi;
                console.log('backingserviceinstance', bsi);

                // if (bsi.spec.binding.bind_deploymentconfig){
                //   DataService.get("deploymentconfigs", bsi.spec.binding.bind_deploymentconfig, context).then(
                //     function(dc){
                //       $scope.deploymentconfig = dc;
                //       console.log("deploymentconfigs", dc)
                //     }
                //   );
                // }
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


         })
        );
      };
    });




