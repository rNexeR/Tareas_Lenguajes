from bottle import route, run, get, post, request, static_file
from emails import orderEmails
import json
import os
from kruskal import Kruskal
from steganography import writeMessage, readMessage

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

def upload(req, filename):
    if os.path.isfile(filename):
        os.remove(filename)
    file = req.files.get('file')
    file.save(filename)

def download(filename):
    return static_file(filename, root='', download=filename)

@post('/orderEmails')
def emails():
    filename = "./uploads/test.txt"
    upload(request, filename)
    fileRet = orderEmails(filename)
    return download(fileRet)

@post('/hideMessage')       
def hide():
    filename = "./uploads/img.bmp"
    upload(request, filename)
    message = request.forms.get('message')
    writeMessage(filename, message)

    return download(filename)

@post('/discoverMessage')
def discover():
    filename = "./uploads/img.bmp"
    upload(request, filename)
    message = readMessage(filename)
    return message

@post('/kruskal')
def kruskal():
    #request body = {"messages" : [ {"message":"hola"}, {"message" : "hola"}]}
    graph = Graph()
    for m in request.json["edges"]:
        From = m["from"]
        To = m["to"]
        Weight = m["weight"]
        graph.edges.append(Edge(From, To, Weight))

    ret = Kruskal(graph)

    return ret.toJSON()

@get('/')
def index():
    return "Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072"


run(host='localhost', port=3000, debug=True)