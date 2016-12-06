var TableRow = require('./TableRow');
var util = require('util');

exports.test = function(){
	var x = new TableRow(0,0,0,0);
	console.log(x);

	var data = `{"edges" :
				[
					{"from" : "a", "to" : "b", "weight" : 6},
					{"from" : "a", "to" : "g", "weight" : 8},
					{"from" : "a", "to" : "d", "weight" : 10},
					{"from" : "d", "to" : "e", "weight" : 6},
					{"from" : "e", "to" : "b", "weight" : 15},
					{"from" : "b", "to" : "h", "weight" : 13},
					{"from" : "b", "to" : "c", "weight" : 11},
					{"from" : "c", "to" : "h", "weight" : 3},
					{"from" : "g", "to" : "h", "weight" : 5},
					{"from" : "g", "to" : "i", "weight" : 5},
					{"from" : "e", "to" : "f", "weight" : 2},
					{"from" : "f", "to" : "g", "weight" : 4},
					{"from" : "f", "to" : "i", "weight" : 6},
					{"from" : "h", "to" : "i", "weight" : 7}
				]}`

	//console.log(data);

	//var graph = new Graph('hola');
	//var graph = new Graph(JSON.parse(data));
	var graph = JSON.parse(data);
	console.log(graph.edges.length);
	console.log(graph);
	var nodes = getNodes(graph);
	//var table = []
	//console.log("-------");
	//var temp = JSON.parse(JSON.stringify(graph.edges[0]));
	//util._extend(temp, {flag: 0});
	//console.log(temp);
	//console.log(graph);
	var ret = doKruskal(graph, nodes.length)
	console.log(ret);
}

exports.kruskal = function(graph){
	var nodes = getNodes(graph);
	var ret = doKruskal(graph, nodes.length);
	return ret
}

var getNodes = function(graph){
	var array = [];
	for(var i = 0; i < graph.edges.length; i++){
		array.push(graph.edges[i].from);
		array.push(graph.edges[i].to);
	}
	var nodeSet = new Set(array);
	return Array.from(nodeSet);
}

var doKruskal = function(graph, nNodes){
	var ret = {edges : []}
	var kruskalTable = [];
	graph.edges.sort(function(a,b){
		return a.weight - b.weight;
	});
	console.log(graph);

	var counter = 0;
	var flag = 0;
	var iterator = 0;

	for(;iterator < graph.edges.length;){
		console.log("-->", graph.edges[iterator]);
		var cflag = nodeIsConnected(kruskalTable, graph.edges[iterator].from, graph.edges[iterator].to);
		if (cflag < 0) {
			var copy = JSON.parse(JSON.stringify(graph.edges[iterator]));
			util._extend(copy, {flag: flag});
			kruskalTable.push(copy);
			ret.edges.push(graph.edges[iterator])
			flag++
			counter++
		} else if (notCycle(kruskalTable, graph.edges[iterator].from, graph.edges[iterator].to)) {
			kruskalTable = updateFlag(kruskalTable, graph.edges[iterator].from, cflag);
			kruskalTable = updateFlag(kruskalTable, graph.edges[iterator].to, cflag);
			var copy = JSON.parse(JSON.stringify(graph.edges[iterator]));
			util._extend(copy, {flag: cflag});
			kruskalTable.push(copy);
			ret.edges.push(graph.edges[iterator]);
			counter++
		}
		iterator++
	}
	return ret;
}

var nodeIsConnected = function(table, node1, node2){
	var length = table.length
	var flags = []
	for (var i = 0; i < length; i++) {
		var row = table[i]
		if (row.from == node1 || row.to == node1 || row.from == node2 || row.to == node2) {
			flags.push(row.flag)
		}
	}
	if (flags.length > 0) {
		flags.sort()
		return flags[0];
	}
	return -1
}

var notCycle = function(table, node1, node2){
	var incidences = [];
	for (var i = 0; i < table.length; i++) {
		var row = table[i];
		if (row.from == node1 || row.to == node1 || row.from == node2 || row.to == node2) {
			incidences.push(row)
		}
	}

	//checking if incidences has 2 in the same flag
	var flagN1 = -1
	var flagN2 = -1
	for (var i = 0; i < incidences.length; i++) {
		var row = incidences[i]
		if (row.from == node1 || row.to == node1) {
			flagN1 = row.flag
		}
		if (row.from == node2 || row.to == node2) {
			flagN2 = row.flag
		}
	}
	return flagN1 != flagN2
}

var updateFlag = function(table, node, flag){
	for (var i = 0; i < table.length; i++) {
		if (table[i].from == node || table[i].to == node) {
			table[i].flag = flag
			console.log(table[i], " Flag updated to", flag)
		}
	}
	return table
}