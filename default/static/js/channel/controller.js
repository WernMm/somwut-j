'use strict';

var dictCtrl = angular.module('dictCtrl', []);

dictCtrl.controller('DictFormCtrl', ['$scope' ,'$http',
  function($scope, $http) {
    $scope.dictRecord = {};
    $scope.words = {};

    $scope.save = function(dictWord) {
      $scope.dictRecord = angular.copy(dictWord);
      var res = $http.post('/rest/dict/saveword', dictWord);
      res.success(function(data, status, headers, config) {
        alert(data);
      });

      res.error(function(data, status, headers, config) {
        alert('failure message: '+ JSON.stringify({data: data}));
      });
    }

    $scope.search = function(dictWord) {
      $scope.words = {};
      var res = $http.post('/rest/dict/searchword', dictWord);
      res.success(function(data, status, headers, config) {
        $scope.words = data;
        alert(JSON.stringify({data: data}));
      });

      res.error(function(data, status, headers, config) {
        alert('failure message: '+ JSON.stringify({data: data}));
      });
    }

    var getListWord = function() {
      var res = $http.post('/rest/dict/searchword', null);
      res.success(function(data, status, headers, config) {
        $scope.words = data;
        alert(JSON.stringify({data: data}));
      });

      res.error(function(data, status, headers, config) {
        alert('failure message: '+ JSON.stringify({data: data}));
      });
    }

    getListWord();

  }]);

dictCtrl.controller('PortfolioDetailCtrl', ['$scope', '$routeParams', '$http',
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