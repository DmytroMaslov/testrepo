angular.module('golang')
	.config(['$stateProvider', '$urlRouterProvider',
	    function ($stateProvider, $urlRouterProvider) {
	    	var dir = 'client/app/templates/';

	        $urlRouterProvider.otherwise('/');

	        $stateProvider
	            .state('tasks', {
	                url: '/',
	                templateUrl: dir + 'tasks.html',
	                controller: 'TasksCtrl'
	             })            
	    }
]);