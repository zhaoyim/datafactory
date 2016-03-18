'use strict';

/**
 * @ngdoc function
 * @name openshiftConsole.controller:ApplicationController
 * @description
 * Controller of the openshiftConsole
 */
angular.module('openshiftConsole')
  .controller('ApplicationController', function ($scope, $routeParams, DataService, ProjectsService, $filter) {
    $scope.projectName = $routeParams.project;
    $scope.app = null;
    $scope.alerts = {};
    $scope.renderOptions = $scope.renderOptions || {};
    $scope.renderOptions.hideFilterWidget = true;
    $scope.breadcrumbs = [
      {
        title: "Applications",
        link: "project/" + $routeParams.project + "/browse/applications"
      },
      {
        title: $routeParams.application
      }
    ];

    var watches = [];

    ProjectsService
      .get($routeParams.project)
      .then(_.spread(function(project, context) {
        $scope.project = project;
        DataService.get("applications", $routeParams.application, context).then(
          // success
          function(app) {
            $scope.loaded = true;
            $scope.app = app;

            // If we found the item successfully, watch for changes on it
            watches.push(DataService.watchObject("applications", $routeParams.application, context, function(app, action) {
              if (action === "DELETED") {
                $scope.alerts["deleted"] = {
                  type: "warning",
                  message: "This application has been deleted."
                };
              }
              $scope.app = app;
            }));
          },
          // failure
          function(e) {
            $scope.loaded = true;
            $scope.alerts["load"] = {
              type: "error",
              message: "The application details could not be loaded.",
              details: "Reason: " + $filter('getErrorDetails')(e)
            };
          }
        );

        //todo so what's this
        watches.push(DataService.watch("routes", context, function(routes) {
          $scope.routesForApp = [];
          angular.forEach(routes.by("metadata.name"), function(route) {
            if (route.spec.to.kind === "Application" &&
              route.spec.to.name === $routeParams.application) {
              $scope.routesForApp.push(route);
            }
          });

          Logger.log("routes (subscribe)", $scope.routesByApp);
        }));

        $scope.$on('$destroy', function(){
          DataService.unwatchAll(watches);
        });

      }));
  });
