class tableRow(object):
	def __init__(self, From, To, Weight, Flag):
		self.From = From
		self.To = To
		self.Weight = Weight
		self.Flag = Flag

class Graph(object):
    def __init__(self):
        self.edges = []

class Edge(object):
    def __init__(self, From, To, Weight):
        self.From = From
        self.To = To
        self.Weight = Weight