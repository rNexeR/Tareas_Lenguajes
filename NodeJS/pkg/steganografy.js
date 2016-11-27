var fs = require("fs");
var bits = require('./bits')

exports.writeMessage = function(message, filename){
	var messageLenBits = intToBitsString(message.length, 32);
	var messageBits = stringToBitsString(message);
	console.log("messageLenBits: " + messageLenBits + " messageBits: " + messageBits);
	console.log("messageLen: " + parseInt(messageLenBits, 2) + " message: " + bitsStringToString(messageBits));
	
	writeBits(messageLenBits, 4*8, 0, filename);
	writeBits(messageBits, message.length*8, 4*8, filename);
}

exports.readMessage = function(filename, cb){
	var messageLenBits = readBits(4*8, 0, filename);
	var messageLen = parseInt(messageLenBits, 2);
	console.log("Message Len Bits: " + messageLenBits);
	console.log("RM Message Length: " + messageLen);

	var messageBits = readBits(messageLen*8, 4*8, filename);
	var message = bitsStringToString(messageBits);

	console.log("Message Bits: " + messageBits);
	console.log("RM Message: " + message);

	return message;
}

var getPixelSizeInBytes = function(filename){
	var fd = fs.openSync(filename, 'r');
	var buffer = new Buffer(2);
	fs.readSync(fd, buffer, 0, 2, 28);
	return buffer.readInt8(0);
}

var getDataOffset = function(filename){
	var fd = fs.openSync(filename, 'r');
	var buffer = new Buffer(2);
	fs.readSync(fd, buffer, 0, 2, 10);
	return buffer.readInt8(0);
}

var intToBitsString = function(n, size){
	var bits = n.toString(2);
	var retornar = "";
	for(var i = bits.length; i<size; i++){
		retornar += "0";
	}
	retornar += bits;
	return retornar;
}

var stringToBitsString = function(s){
	var bits = "";
	for(var i = 0; i < s.length; i++){
		bits += intToBitsString(s.charCodeAt(i), 8);
	}
	//console.log(bits.length);
	return bits;
}

var bitsStringToString = function(b){
	var bytesCount = b.length/8;
	var str = "";
	for(var i = 0; i < bytesCount; i++ ){
		var byte = "";
		for(var j = 0; j < 8; j++){
			byte += b.charAt(i*8 + j);
		}
		var character = parseInt(byte, 2);
		str += String.fromCharCode(character);
		//console.log(String.fromCharCode(character))
	}
	return str;
}

var writeBits = function(bitsToStore, length, offset, filename){
	var bytesPerPixel = getPixelSizeInBytes('./uploads/original.bmp');
	var dataOffset = getDataOffset('./uploads/original.bmp');
	console.log("ps: " + bytesPerPixel + " do: " + dataOffset);

	var whereToRW = dataOffset + offset;

	console.log("Whete to RW: " + whereToRW);

	file = fs.openSync(filename, 'r+');
	for(var i = 0; i < length; i++){
		fileOffset = whereToRW + i;
		lectura = new Buffer(bytesPerPixel);

		fs.readSync(file, lectura, 0, bytesPerPixel, fileOffset);

		//console.log("Byte before: " + lectura.readInt8(0));

		if(bitsToStore.charAt(i) == '1'){
			lectura.writeInt8(bits.setBit(lectura.readInt8(0)) ,0);
		}else{
			lectura.writeInt8(bits.clearBit(lectura.readInt8(0)) ,0);
		}

		//console.log("Byte after: " + lectura.readInt8(0));

		fs.writeSync(file, lectura, 0, bytesPerPixel, fileOffset);
	}
	fs.closeSync(file);
}

var readBits = function(length, offset, filename){
	var bytesPerPixel = getPixelSizeInBytes('./uploads/original.bmp');
	var dataOffset = getDataOffset('./uploads/original.bmp');

	var whereToRW = dataOffset + offset;
	var bitsRead = "";

	file = fs.openSync(filename, 'r');
	for(var i = 0; i < length; i++){
		lectura = new Buffer(bytesPerPixel);
		fileOffset = whereToRW + i;

		fs.readSync(file, lectura, 0, bytesPerPixel, fileOffset);

		if(bits.hasBit(lectura.readInt8(0), 0))
			bitsRead += "1";
		else
			bitsRead += "0";
	}
	fs.closeSync(file);
	return bitsRead;
}