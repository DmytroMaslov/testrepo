angular.module('golang')
	.factory('tasks', function ($http, DEFAULTS) {
		this.tasks = DEFAULTS;
	    return {
	    	tasks: this.tasks,
	        runTask : function (id) {
	            return $http.post('/task/' + id, this.tasks[id]);
	        },

	        runAll: function (id) {
	        	return $http.post('/tasks', this.tasks);
	        }
	    }
	});