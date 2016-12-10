import os
import json

class tableRow(object):
	def __init__(self, From, To, Weight, Flag):
		self.From = From
		self.To = To
		self.Weight = Weight
		self.Flag = Flag

class Graph(object):
    def __init__(self):
        self.edges = []
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, 
            sort_keys=True, indent=4)

class Edge(object):
    def __init__(self, From, To, Weight):
        self.From = From
        self.To = To
        self.Weight = Weight
    def __repr__(self):
        return repr((self.From, self.To, self.Weight))

def Kruskal(graph):
	nodes = getNodes(graph)
	ret = doKruskal(graph, len(nodes))

	return ret

def getNodes(graph):
	#print length
	nodes = set()
	for n in graph.edges:
		nodes.add(n.From)
		nodes.add(n.To)

	return nodes

def doKruskal(graph, length):
	ret = Graph()
	kruskalTable = []
	printGraph(graph)
	print
	graph.edges = sorted(graph.edges, key=lambda edge: edge.Weight)
	printGraph(graph)

	counter = 0
	iterator = 0
	flag = 0
	for node in graph.edges:
		print "--> From: ", node.From, " To: ", node.To, " Weight: ", node.Weight
		cflag = nodeIsConnected(kruskalTable, node.From, node.To)
		if cflag < 0:
			kruskalTable.append(tableRow(node.From, node.To, node.Weight, flag))
			ret.edges.append(node)
			flag += 1
			counter += 1
		elif notCycle(kruskalTable, node.From, node.To):
			kruskalTable = updateFlag(kruskalTable, node.From, cflag)
			kruskalTable = updateFlag(kruskalTable, node.To, cflag)
			kruskalTable.append(tableRow(node.From, node.To, node.Weight, cflag))
			ret.edges.append(node)
			print "\tAdded From: ", node.From, " To: ", node.To, " Weight: ", node.Weight
			counter += 1

	print
	printGraph(ret)

	return ret

def printGraph(graph):
	for node in graph.edges:
		print "From: ", node.From, " To: ", node.To, " Weight: ", node.Weight

def nodeIsConnected(table, node1, node2):
	length = len(table)
	flags = []
	for n in range(0, length):
		row = table[n]
		if row.From == node1 or row.To == node1 or row.From == node2 or row.To == node2 :
			flags.append(row.Flag)
	if len(flags) > 0:
		flags.sort()
		return flags[0]
	return -1

def notCycle(table, node1, node2):
	incidences = []
	for i in range(0, len(table)):
		row = table[i]
		if row.From == node1 or row.To == node1 or row.From == node2 or row.To == node2 :
			incidences.append(row)

	flagN1 = -1
	flagN2 = -1
	for i in range(0, len(incidences)):
		row = incidences[i]
		if row.From == node1 or row.To == node1 :
			flagN1 = row.Flag
		if row.From == node2 or row.To == node2 :
			flagN2 = row.Flag

	return flagN1 != flagN2	

def updateFlag(table, node, flag):
	for i in range(0, len(table)):
		if table[i].From == node or table[i].To == node :
			table[i].Flag = flag
			print "From: ", table[i].From, "To: ", table[i].To, "updated to ", flag

	return table