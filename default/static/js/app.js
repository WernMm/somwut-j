'use strict';

/* App Module */

var somwutApp = angular.module('somwutApp', [
  'ngRoute',
  'ngSanitize',
  'portfolioCtrl'
]);

somwutApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/about', {
        templateUrl: 'partials/about.html'
      }).
      when('/portfolio', {
        templateUrl: 'partials/portfolio-list.html',
        controller: 'PortfolioListCtrl'
      }).
      when('/portfolio/:projectId', {
        templateUrl: 'partials/portfolio-detail.html',
        controller: 'PortfolioDetailCtrl'
      }).
      otherwise({
        redirectTo: '/about'
      });
  }]);