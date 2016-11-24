/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , fileUpload = require('express-fileupload');

var app = express();
app.use(fileUpload());

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

app.get('/',function(req, res){
    res.json(200, "Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072" );
  }
);

app.post('/upload', function(req, res) {
    var sampleFile;
 
    if (!req.files) {
        res.status(500).send('No files were uploaded.');
        console.log("No files were uploaded.");
        return;
    }
 
    sampleFile = req.files.sampleFile;
    var uploadPath = __dirname + '/uploads/' + sampleFile.name;
    sampleFile.mv(uploadPath, function(err) {
        if (err) {
            res.status(500).send(err);
        }
        else {
          cloudinary.uploader.upload(uploadPath, function(result) { 
              res.json(200, result);
          });

          fs.unlink(uploadPath);
        }
    });
});


http.createServer(app).listen(app.get('port'), function(){
  console.log("Express server listening on port %s in %s mode.",  app.get('port'), app.settings.env);
});