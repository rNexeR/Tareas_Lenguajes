var fs = require('fs');
var fse = require('fs-extra');
//var LineByLineReader = require('line-by-line');
const max = 5;

Array.prototype.subarray=function(start,end){
     if(!end){ end=0;} 
    return this.slice(start, end+1);
}

exports.orderEmails = function(filename, cb){
	//fs.rmdirSync('./files');
	//fs.mkdirSync('./files');
	console.log(odd(17));
	var cant = filterEmails(filename);
	var fileNumber = getDepth(cant);
	//var cant = 0;
	//var fileNumber = 0;	
	lastfile = './files/' + fileNumber.toString() + "_0.txt";
	/*mergeTwoFiles("./files/0_0.txt", "./files/0_1.txt", "./files/1_0.txt", function(err){
		if(err)
			cb(err, null);
		else{
			cb(null, lastfile);
		}
	});*/
	cb(null, lastfile);
	
}

var filterEmails = function(filename){
	var regx = new RegExp("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})");

	var filtered = './files/filtered_emails.txt';
	var emailsFile = fs.readFileSync(filename);
	var emails = emailsFile.toString().toLowerCase().split("\n");

	fs.writeFileSync(filtered, "");

	var sortedEmails = [];

	for (var i = 0; i < emails.length; i++) {
		var match = regx.exec(emails[i]);
		if(match != null){
			console.log(emails[i]);
			sortedEmails.push(emails[i]);
			fs.appendFileSync(filtered, emails[i] + "\n");
		}
	}
	//console.log(merge(1,1,1,1));
	//return 5;
	var cant = createSortedLeaves(sortedEmails);
	console.log("Create sorted Leaves cant: " + cant);
	//var cant = 5;
	console.log(cant);
	mergeLeaves(cant);
	return cant;
}

var odd = function(number){
	if (number == 1)
		return false;
	return number%2 != 0;
}

var getDepth = function(number){
	var number_of_cycles = 0;
	var temp = number;

	while(temp > 1){
		if(odd(temp))
			temp++;
		number_of_cycles++;
		temp = parseInt(temp/2);
	}
	return number_of_cycles;
}

var getNextNumberOfFiles = function(count){
	var temp = count;
	if(odd(temp))
		temp++;
	return parseInt(temp/2);
}

var merge = function(nfiles, ncreates, ccreated, depth){
	console.log("nfiles: %d ncreates: %d ccreated: %d depth: %d\n", nfiles, ncreates, ccreated, depth);
	
	if (ccreated < ncreates) {
		if (odd(nfiles) && ccreated+1 == ncreates) {
			//pressAnyKey()
			var fCopy = "./files/" + (depth-1) + "_" + (nfiles-1) + ".txt"
			var fResult = "./files/" + (depth) + "_" + (ccreated) + ".txt"
			
			fse.copySync(fCopy, fResult);


			console.log("copying %d_%d from %s\n", depth, ccreated, fCopy)
			//pressAnyKey()
		} else {
			var fResult = "./files/" + (depth) + "_" + (ccreated) + ".txt"
			var file1 = "./files/" + (depth-1) + "_" + (ccreated*2) + ".txt"
			var file2 = "./files/" + (depth-1) + "_" + (ccreated*2+1) + ".txt"
			console.log("creating %d_%d from %s and %s\n", depth, ccreated, file1, file2)
			mergeTwoFiles(file1, file2, fResult)
			merge(nfiles, ncreates, ccreated+1, depth)			
		}
	}
}

var mergeTwoFiles = function(filel, filer, fileResult, cb){
	var s_posl = 0;
	var s_posr = 0;

	console.log("merging: ", filel, filer, " to: ", fileResult);

	//var c_1 = fs.readFileSync(filel).toString()
	//var c_2 = fs.readFileSync(filer).toString()

	var emails_l = fs.readFileSync(filel).toString().split("\n");
	var emails_r = fs.readFileSync(filer).toString().split("\n");

	fs.writeFileSync(fileResult, "");

	//console.log("AR -> ", c_2);
	//console.log("AL -> ", c_1);

	//console.log("R ->\n", emails_r);
	//console.log("L ->\n", emails_l);

	//console.log("Leyendo linea por linea");
	while(true){
		var r_email = emails_r[s_posr];
		var l_email = emails_l[s_posl];

		console.log("R ->", r_email);
		console.log("L ->", l_email);

		if(l_email.length > 0 && r_email.length > 0){
			if (l_email < r_email) {
				//fmt.Printf("Writing: %s\n", l_email)
				fs.appendFileSync(fileResult, l_email + "\n");
				s_posl++;
			} else {
				//fmt.Printf("Writing: %s\n", r_email)
				fs.appendFileSync(fileResult, r_email + "\n");
				s_posr++;
			}
		}else{
			if (l_email.length > 0) {
				//fmt.Printf("Writing: %s\n", l_email)
				fs.appendFileSync(fileResult, l_email + "\n");
				s_posl++;
			}else if (r_email.length > 0) {
				//fmt.Printf("Writing: %s\n", r_email)
				fs.appendFileSync(fileResult, r_email + "\n");
				s_posr++;
			} else {
				break
			}
		}
	}
	
	/*
	scanl.on('line', function(line){
		scanl.pause();
		console.log("L--> ", line);
		email_l = line.toString();
		if(email_r != ""){
			if(email_l < email_r){
				fs.appendFileSync(fileResult, email_l + "\n");
			}else{
				fs.appendFileSync(fileResult, email_r + "\n");
			}
		}
		//while(email_r == "");
		
		//while(email_l == line);
		scanl.resume();
	});

	scanl.on('end', function(){
		scanl_finish = true;
	});

	scanr.on('line', function(line){
		scanr.pause();
		console.log("R--> ", line);
		email_r = line.toString();
		if(email_l != ""){
			if(email_l < email_r){
				fs.appendFileSync(fileResult, email_l + "\n");
			}else{
				fs.appendFileSync(fileResult, email_r + "\n");
			}
		}
		//while(email_l == "");
		//while(email_r == line);
		scanr.resume();
	});

	scanr.on('end', function(){
		scanr_finish = true;
	});
	//*/
	//cb(null);
}

var mergeLeaves = function(count){
	var depth = getDepth(count);
	var number_of_files = count;
	for (var i = depth; i > 0; i--) {
		var ncreates = getNextNumberOfFiles(number_of_files);
		merge(number_of_files, ncreates, 0, depth-i+1)
		number_of_files = ncreates
	}
}

var createSortedLeaves = function(emails){
	var cant = parseInt(emails.length / max);

	if (cant*max < emails.length) {
		cant++;
	}

	//fmt.Println("Cantidad de hojas a crear: " + strconv.Itoa(cant))

	for (var i = 0; i < cant; i++) {
		//fmt.Printf("Hoja actual: %d\n", i)
		var fn = "./files/0_" + (i) + ".txt";
	 	fs.writeFileSync(fn, "");

		var right = i*5 + 5
		if (right >= emails.length) {
			right = emails.length;
		}
		///*
		var emailSorted = emails.subarray(i*5, right-1)
		console.log("Emails sorted in ", fn, " : ",emailSorted);
		emailSorted.sort();
		//sort.Strings(emailSorted)

		//fmt.Print("Sorted emails to store: ")
		//fmt.Println(emailSorted)

		var toWrite = max;
		if (emailSorted.length < max) {
			toWrite = emailSorted.length;
		}

		for (var j = 0; j < toWrite; j++) {
			var store = emailSorted[j] + "\n";
			//fmt.Printf("Email to store: %s \n", store)
			fs.appendFileSync(fn, store);
			console.log(j, " - ", store);
		}
		//*/
	}
	//pressAnyKey()
	return cant
}

exports.test = function(variable, cb){
	var data = fs.readFileSync('./uploads/correos.txt');
	this.test2(data.toString());

	cb(null, variable);
}

exports.test2 = function(variable){
	console.log(variable);
}