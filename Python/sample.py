from bottle import route, run, get, post, request
from emails import sayHello
from emails import test
import json

class ALGO(object):
    def __init__(self, messages):
        self.messages = messages

@post('/orderEmails')
def emails():
    return "Working on that!"

@post('/hideMessage')
def hide():
    return "Working on that!"

@post('/discoverMessage')
def discover():
    return "Working on that!"

@post('/kruskal')
def kruskal():
    #request body = {"messages" : [ {"message":"hola"}, {"message" : "hola"}]}
    for m in request.json["messages"]:
        print m
    algo = ALGO(**request.json)
    print algo.messages
    #print request.json.message
    return "Working on that!"

@get('/')
def index():
    return "Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072"

@get('/hello')
def hello():
    return sayHello()

test()

run(host='localhost', port=8080, debug=True)