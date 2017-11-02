'use strict';

var portfolioCtrl = angular.module('portfolioCtrl', []);

portfolioCtrl.controller('PortfolioListCtrl', ['$scope' ,'$http',
  function($scope, $http) {
    $scope.projects = [];
    var logError = function(data, status) {
      console.log('code '+status+': '+data);
    };

    var getPortfolioList = function() {
      return $http.get('/rest/portfolio/all').
          success(function(data) {
          // var chunkProjects = [];
          // var chunkSize = 3;
          // for(var i=0; i<data.Projects.length; i+=chunkSize) {
          //   chunkProjects.push(data.Projects.slice(i, i+chunkSize));
          // }
          // $scope.chunkProjects = chunkProjects;
          // alert(data.Projects);
          $scope.projects = data.Projects;
        }).
        error(logError);
    };

    getPortfolioList();
  }]);

portfolioCtrl.controller('PortfolioDetailCtrl', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    $scope.projectId = $routeParams.projectId;

    var getProject = function() {
    	return $http.get('/rest/portfolio/'+$scope.projectId).
    	success(function(data) {
        $scope.project = data.Project;
    	});
    };
    getProject();
  }]);