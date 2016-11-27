/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , emails = require('./pkg/emails')
  , uploads = require('./pkg/uploads')
  , steganografy = require('./pkg/steganografy')
  , fs = require('fs')
  , kruskal = require('./pkg/kruskal')
  , fileUpload = require('express-fileupload');

var app = express();
app.use(fileUpload());

/*
app.configure(function(){
  app.set('port', process.env.PORT || 8000);
  app.use(express.favicon());
  app.use(express.methodOverride());
  app.use(app.router);
  app.use(express.static(__dirname + '/public'));
});

app.configure('development', function(){
  app.use(express.errorHandler());
});
*/

app.get('/',function(req, res){
    res.status(200).json("Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072" );
  }
);

app.post('/orderEmails', function(req, res){
  //Eliminar directorio, descargar archivo y division de enteros pendientes
  var filename = "./uploads/emails.txt"
  uploads.upload(req, filename, function(err){
    if(err)
      res.status(500).json(err)
    else{
      emails.orderEmails(filename, function(err, data){
        if(err)
          res.status(500).json(err)
        else
          res.download(data);
      });
    }
  })
})

app.post('/hideMessage', function(req, res){
  var message = req.body.mensaje;
  console.log(message);
  var filename = "./uploads/img.bmp";
  uploads.upload(req, filename, function(err){
    if(err)
      res.status(500).json(err)
    else{
      steganografy.writeMessage(message, filename);
      res.download(filename);
    }
  });
})

app.post('/discoverMessage', function(req, res){
  var filename = "./uploads/img.bmp";
  uploads.upload(req, filename, function(err){
    if(err)
      res.status(500).json("Error uploading file")
    else{
      var message = steganografy.readMessage(filename);
      res.status(200).json(message);
    }
  })
})

app.post('/kruskal', function(req, res){
  res.status(200).json("Working in that");
})

kruskal.test();

http.createServer(app).listen(8000, function(){
  console.log("Express server listening on port %s in %s mode.",  8000, app.settings.env);
});