'use strict';

/* App Module */

var dictApp = angular.module('dictApp', [
  'ngRoute',
  'dictCtrl'
]);

dictApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/', {
        templateUrl: 'partials/dict/dict-form.html',
        controller: 'DictFormCtrl'
      }).
      when('/portfolio/:projectId', {
        templateUrl: 'partials/portfolio-detail.html',
        controller: 'PortfolioDetailCtrl'
      }).
      otherwise({
        redirectTo: '/'
      });
  }]);