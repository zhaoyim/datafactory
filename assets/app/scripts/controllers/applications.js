'use strict';

/**
 * @ngdoc function
 * @name openshiftConsole.controller:ApplicationsController
 * @description
 * # ApplicationsController
 * Controller of the openshiftConsole
 */
angular.module('openshiftConsole')
  .controller('ApplicationsController', function ($routeParams, $scope, AlertMessageService, DataService, ProjectsService, $filter, LabelFilter, Logger, $location, $anchorScroll) {
    $scope.projectName = $routeParams.project;
    $scope.apps = {};
    $scope.unfilteredApps = {};
    $scope.routesByApp = {};
    $scope.routes = {};
    $scope.labelSuggestions = {};
    $scope.alerts = $scope.alerts || {};
    $scope.emptyMessage = "Loading...";
    $scope.emptyMessageRoutes = "Loading...";

    // get and clear any alerts
    AlertMessageService.getAlerts().forEach(function(alert) {
      $scope.alerts[alert.name] = alert.data;
    });
    AlertMessageService.clearAlerts();

    var watches = [];

    ProjectsService
      .get($routeParams.project)
      .then(_.spread(function(project, context) {
        $scope.project = project;
        watches.push(DataService.watch("applications", context, function(apps, action) {
          $scope.unfilteredApps = apps.by("metadata.name");
          LabelFilter.addLabelSuggestionsFromResources($scope.unfilteredApps, $scope.labelSuggestions);
          LabelFilter.setLabelSuggestions($scope.labelSuggestions);
          $scope.apps = LabelFilter.getLabelSelector().select($scope.unfilteredApps);
          $scope.emptyMessage = "No applications to show";
          updateFilterWarning();

          // Scroll to anchor on first load if location has a hash.
          if (!action && $location.hash()) {
            // Wait until the digest loop completes.
            setTimeout($anchorScroll, 10);
          }

          Logger.log("applications (subscribe)", $scope.unfilteredApps);
        }));

        //todo what's this
        watches.push(DataService.watch("routes", context, function(routes){
          $scope.routes = routes.by("metadata.name");
          $scope.emptyMessageRoutes = "No routes to show";
          $scope.routesByApp = routesByApp($scope.apps);
          Logger.log("routes (subscribe)", $scope.routesByApp);
        }));

        function routesByApp(routes) {
          var routeMap = {};
          angular.forEach(routes, function(route, routeName){
            var to = route.spec.to;
            if (to.kind === "Application") {
              routeMap[to.name] = routeMap[to.name] || {};
              routeMap[to.name][routeName] = route;
            }
          });
          return routeMap;
        }

        function updateFilterWarning() {
          if (!LabelFilter.getLabelSelector().isEmpty() && $.isEmptyObject($scope.apps)  && !$.isEmptyObject($scope.unfilteredApps)) {
            $scope.alerts["applications"] = {
              type: "warning",
              details: "The active filters are hiding all applications."
            };
          }
          else {
            delete $scope.alerts["applications"];
          }
        }

        LabelFilter.onActiveFiltersChanged(function(labelSelector) {
          // trigger a digest loop
          $scope.$apply(function() {
            $scope.apps = labelSelector.select($scope.unfilteredApps);
            updateFilterWarning();
          });
        });

        $scope.$on('$destroy', function(){
          DataService.unwatchAll(watches);
        });

      }));
  });
