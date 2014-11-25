"use strict";

var app01 = angular.module('layouterApp', ['ngRoute']);  


app01.config(['$routeProvider', function($routeProvider) {
	$routeProvider
	.when('/', {
		controller: 'controller01',
		templateUrl: '/corridor-set?vp=vp2 '
	})
	.when('/vp/:param1', {
		controller: 'controller01',
		templateUrl : function (params) {
			console.log(params); 
			return '/corridor-set?vp=' + params.param1;
		}
	})
	.when('/object-and-list', {
		controller: 'controller01',
		templateUrl: '/tpl-ng/object-and-list.html'
	})
	.when('/try', {
		controller: 'controller01',
		templateUrl: '/tpl-ng/try.html'
	})
	.when('/todo-list', {
		controller: 'controller01',
		templateUrl: '/tpl-ng/todo-list.html'
	})
	.when('/url-x', {
		controller: 'controller02',
		templateUrl: '/tpl-ng/try.html'
	})
	.otherwise({
		redirectTo: '/'
	});

}]);


// only the last  .controller('controller01' ... is fired
	app01.controller('controller01', 
		['$scope',  'Service01', 'Service02', '$routeParams',
		function($scope,  Service01, Service02 ,$routeParams) {
			Service01.viewSingleObject()
			.then(function(dynData) {
				$scope.title = dynData.title;
				$scope.description = dynData.description;
			});
			Service02.listObjects()
			.then(function(dynList) {
				$scope.list01 = dynList;
			});
	    //$scope.param1 = $routeParams.param1;
	    $scope.greeting = 'Hola!';
	    $scope.param1 = $routeParams.param1;
	    console.log("end of controller01 definition -  param1 = '" + $routeParams.param1 + "'");
	}
	]
	);



	app01.factory('Service01', ['$http', '$q', function($http, $q) {
		return {
			viewSingleObject: function() {
				return $http.get('/tpl-ng/01.json')
				.then(function(response) {
					if (response.data) {
						return response.data;
					} else {
						return $q.reject('No data in response.');
					}
				}, function(response) {
					return $q.reject('Server or connection error.');
				});
			}
		};
	}]);


	app01.factory('Service02', ['$http', '$q', function($http, $q) {
		return {
			listObjects: function() {

			// explicit list of objects
			var list = {}
			var post1 = {}
			var post2 = {}

			post1.title = 'Title Post 1'
			post2.title = 'Another Title'

			post1.content = 'content 1. content 1. content 1. content 1. content 1. '
			post2.content = 'content 2. content 2. content 2. content 2. content 2. content 2. content 2. '

			list['003'] = post1
			list['004'] = post2

			// equivalent object literal:
			list = [{title:"t1",content:"content 1 content 1"},{title:"t2",content:"content2 content2 content2 content2"}]

			// it has to be a "promise"
			return $http.get('/tpl-ng/01.json')
			.then(function(response) {
				return list
			}, function(response) {
				return list
			});
		}
	};
}]);




	var fctController02 = function nameNotNeeded($scope, $http) {

		$scope.user = "email or customer id"
		$scope.userLoggedIn = false

		$scope.requestNow = function() {
			$http.get('/tpl-ng/01.json', {
				username: $scope.user,
				password: $scope.password,
			})
			.success(function(data, status, headers, config) {
				$scope.userLoggedIn = data.isLoginValid;
				console.log(data);
			})
			.error(function(err, status, headers, config) {
				console.log("Well, this is embarassing.");
			});
		};


		$scope.spice = 'very';
		$scope.jalapenoSpicy = function() {
			$scope.spice = 'jalapeño';
		};   


		var spiceIter = function() {
			var flavors = ['tastless','spicy','revolting','sweety','juicy','sour']
			var cntr = 0
			return function() {
				cntr++
				idx = cntr % flavors.length
				flavor = flavors[idx]
				return flavor
			}
		//$scope.spice = 'jalapeño';
	};   
	var spiceIterInst = spiceIter()
	$scope.iterFlavors = function() {
		$scope.spice = spiceIterInst();
	};   

}



app01.controller('controller02', ['$scope',  '$http', fctController02] );


// -------------------------------------------------------------

var app02 = angular.module("app02", []);


var fctTmp1 = function($scope) {
	$scope.firstName = "John";
	$scope.lastName  = "Doe";
}
app01.controller('controller04', ['$scope', fctTmp1] );





var fctTmp2 = function($scope) {
	$scope.defaultData = {firstName: "John", lastName: "Doe", email: "Enter Valid Email"};
	$scope.fillDefaults = function() {
		$scope.user = angular.copy($scope.defaultData);
	};
	$scope.fillDefaults();
};
app01.controller('formController', ['$scope', fctTmp2] );




var fctTmp3 = function($scope) {

	$scope.newTodo = "blab blupp"

	$scope.alternator = false;

	$scope.todoList = [{task:"clean house",until:"22-12",done:false},
	{task:"fuck penguin",until:"24-12",done:true}];

	$scope.addTodo = function  () {
		var effectiveTask = "some new todo";
		if ($scope.newTodo.toString() != "") {
			effectiveTask = $scope.newTodo;
		}
		$scope.todoList.push({task:effectiveTask,until:"today",done:$scope.alternator})
		$scope.alternator = ! $scope.alternator;
	}

	$scope.removeSelected = function() {
		var updatedTodo = []
		angular.forEach($scope.todoList, function(value, key) {
			if (! value.done) {
				updatedTodo.push(value)
			}
		});		
		$scope.todoList = updatedTodo
	}
};
app01.controller('controller05', ['$scope', fctTmp3] );




console.log("/parsing app.js")
