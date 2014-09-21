/*global window,angular*/
(function(window) {
	'use strict';
	var angular = window.angular, appName = 'new_blog_app';
	var new_blog_app = angular.module('new_blog_app', [ 'simditor' ]);
	
	new_blog_app.config(function($interpolateProvider) {
		$interpolateProvider.startSymbol('[[').endSymbol(']]');
	});

	new_blog_app.controller('mainController',
			function($scope) {
				$scope.editor = '';
			});
}(window));
