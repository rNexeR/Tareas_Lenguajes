var fs = require('fs');

exports.upload = function(req, filename, cb){
	var file;
 
    if (!req.files) {
        cb(new Error("No files were uploaded"))
    }
 
    file = req.files.file;
    if(!file){
    	cb(new Error("Cannot find file in files param"))
    }else{
	    var uploadPath = filename;
	    file.mv(uploadPath, function(err) {
	        if (err) {
	            cb(new Error('Cannot copy file'));
	        }
	        else {
	        	cb(null);
	          	//fs.unlink(uploadPath);
	        }
	    });
	}
}