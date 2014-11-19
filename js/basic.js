var Person = function function_name (name, age) {
	this.name = name;
	this.age = age;
}

Person.prototype.say = function(prefix) {
	alert( prefix + " - " + this.name + " - " + this.age)
};

var ich = new Person("Peter",27);
//alert(ich.say("Präambel: "));

