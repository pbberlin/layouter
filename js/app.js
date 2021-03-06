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
	.when('/tabbing-fontsizing', {
		controller: 'controller01',
		templateUrl: '/tpl-ng/tabbing-fontsizing.html'
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


var fctController04 = function($scope) {
	$scope.firstName = "John";
	$scope.lastName  = "Doe";
}
app01.controller('controller04', ['$scope', fctController04] );





var fctControllerForm = function($scope) {
	$scope.defaultData = {firstName: "John", lastName: "Doe", email: "Enter Valid Email"};
	$scope.fillDefaults = function() {
		$scope.user = angular.copy($scope.defaultData);
	};
	$scope.fillDefaults();
};
app01.controller('formController', ['$scope', fctControllerForm] );




var fctController05 = function($scope) {

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
app01.controller('controller05', ['$scope', fctController05] );



var fctController06 = function($scope,$window) {

	$scope.focusedElement = null;
	$scope.focusedElementId = null;

	$scope.keyCode = "";

	$scope.CSSStyle = "";
	$scope.fontSize = "";
	$scope.letterSpacing = 0;
	$scope.lineHeight = 125;



	$scope.setFontsize = function(focusEvent,focusOrBlur) {
		//console.log("element:",focusEvent)
		var el=  angular.element(focusEvent.target)
		// console.log("source:",el)

		if( focusOrBlur === 1){
			$scope.focusedElement = el;
			$scope.focusedElementId = el.attr("id");
			el.css({
				border: '0px solid red',
				backgroundColor: 'red'
			});		

			if ($window.getComputedStyle) {
				$scope.CSSStyle = $window.getComputedStyle(el[0], null)
			} else {
				$scope.CSSStyle = el[0].currentStyle
			}

			if( $scope.CSSStyle ) {
				$scope.fontSize = $scope.CSSStyle["fontSize"]
				console.log("edit fontsize - id:",$scope.focusedElementId, " font-size:",$scope.fontSize)
				// el.append("<h5>fontsize: "+$scope.fontSize+"</h5>")
			}


		} else {
			$scope.focusedElement = null;
			$scope.focusedElementId = "";
			$scope.fontSize = "";
			el.css({
				border: 'none',
				backgroundColor: 'transparent'
			});
			console.log("  edit fontsize cleared")
			// el.find("h5").remove()
		}


	}

	$scope.registerKeyEvent = function(keyEvent) {
		keyEvent.preventDefault();
		var el=  angular.element(keyEvent.target)
		var kc = $scope.keyCode = keyEvent.which;
		console.log("keycode:",kc,keyEvent.shiftKey , keyEvent.altKey )
		if ( kc === 108 || kc === 76 ) {
				if ( kc === 76  ) {
					$scope.lineHeight -= 5
				}
				if ( kc === 108 ) {
					$scope.lineHeight += 5
				}
				el.css({"lineHeight": $scope.lineHeight+"%"});
				el.children().css({"lineHeight": $scope.lineHeight+"%"});
				console.log("  lineHeight changed:",$scope.lineHeight);
		}
		if ( (kc === 42 || kc === 95) &&  keyEvent.shiftKey) {
				if ( kc === 95  ) {
					$scope.letterSpacing -= 1
				}
				if ( kc === 42 ) {
					$scope.letterSpacing += 1
				}
				el.css({"letterSpacing": $scope.letterSpacing+"px"});
				el.children().css({"letterSpacing": $scope.letterSpacing+"px"});
				console.log("  letter spacing changed:",$scope.letterSpacing);
		}
		if ( kc === 43 || kc === 45) {
				var valPlusUnit =  getNumPartOfCSSMeasure($scope.fontSize)
				if ( kc === 45 ) {
					valPlusUnit.val -= 1
				}
				if ( kc === 43 ) {
					valPlusUnit.val += 1
				}
				var newFontSize =  (valPlusUnit.val +  valPlusUnit.unit);
				el.css({"fontSize": newFontSize});
				el.children().css({"fontSize": newFontSize});
				// el.find("p").css();
				$scope.fontSize = newFontSize;
				console.log("  fontsize changed:",$scope.fontSize, " keycode:",kc );

		}
	}
};

app01.controller('controller06', ['$scope','$window', fctController06] );



console.log("/parsing app.js")



// We need a css property value parser
// http://glazman.org/JSCSSP/ is too heavy
// 15px => {val:15,unit:"px"}
function getNumPartOfCSSMeasure(measure){

	try {
		var digits = ""
		var unit = ""
		measure = measure.toLowerCase();

		for ( var i = 0; i < measure.length; i++) { 
			var c = measure.charAt(i); 
			var ascii = measure.charCodeAt(i); 
			if(ascii >= 97 && ascii <= 121){
				unit += c
			}
			if(ascii >= 46 && ascii <= 58){
				digits += c
			}
			// console.log(c, ascii)
		}
		// console.log(eval(digits), unit)
		return {val:eval(digits),unit:unit}
	}
	catch (e) {
		return {val:12,unit:"px"}
	}


}