'use strict';

var dictCtrl = angular.module('dictCtrl', []);

dictCtrl.controller('DictFormCtrl', ['$scope' ,'$http', 'BASE',
  function($scope, $http, base) {
    $scope.dictRecord = {};
    $scope.words = {};

    $scope.save = function(dictWord) {
      $scope.dictRecord = angular.copy(dictWord);
      var res = $http.post(base+'/rest/dict/saveword', dictWord);
      res.success(function(data, status, headers, config) {
        // alert(data);
      });

      res.error(function(data, status, headers, config) {
        alert('failure message: '+ JSON.stringify({data: data}));
      });
    }

    $scope.reset = function() {
      $scope.dictWord = {
        meaning: '',
        word: ''
      }
    }

    $scope.edit = function(dictWord) {
      $scope.dictWord = {
        meaning: dictWord.Meaning,
        word: dictWord.Word
      }
    }

    $scope.search = function(dictWord) {
      $scope.words = {};
      var res = $http.post(base+'/rest/dict/searchword', dictWord);
      res.success(function(data, status, headers, config) {
        $scope.words = data;
        // alert(JSON.stringify({data: data}));
      });

      res.error(function(data, status, headers, config) {
        alert('failure message: '+ JSON.stringify({data: data}));
      });
    }

    var getListWord = function() {
      var res = $http.post(base+'/rest/dict/searchword', null);
      res.success(function(data, status, headers, config) {
        $scope.words = data;
        // alert(JSON.stringify({data: data}));
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
    	return $http.get(base+'/rest/portfolio/'+$scope.projectId).
    	success(function(data) {
        $scope.project = data.Project;
    	});
    };
    getProject();
  }]);