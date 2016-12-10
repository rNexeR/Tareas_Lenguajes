from PIL import Image
import bitarray
import bits

def test():
	'''
	message = "Hola"
	msgBits = stringToBits(message)
	print msgBits
	print bitsToString(msgBits)

	print 

	lenBits = intToBits(len(message))
	print lenBits

	print BitsToInt(lenBits)

	return
	'''

	writeMessage('./uploads/original.bmp', "Hola")
	message = readMessage('./uploads/original.bmp')
	print message

def stringToBits(text):
	ret = bitarray.bitarray()
	ret.fromstring(text)
	return ret

def bitsToString(bitss):
	return bitss.tostring()

def intToBits(entero):
	ret = bitarray.bitarray(32)
	for n in range(32):
		ret[32-n-1] = bits.HasBit(entero, n)
	return ret

def BitsToInt(bitss):
	ret = 0
	for n in range(32):
		if(bitss[32-n-1]):
			print "Set bit ", n
			ret += 2**n
	return ret

def writeMessage(filename, message):
	img = Image.open(filename)

	msgBits = stringToBits(message)
	lenBits = intToBits(len(message))

	print len(msgBits), msgBits
	print len(lenBits), lenBits

	for b in range(len(lenBits)):
		writeBit(lenBits[b], b, img)

	for b in range(32,len(msgBits)+32):
		writeBit(msgBits[b-32], b, img)

	img.save(filename)


def writeBit(bit, offset, img):
	width = img.width
	height = img.height

	y_offset = offset/width
	x_offset = (offset - y_offset*width)

	r,g,b = img.getpixel((x_offset, y_offset))

	if bit:
		b = bits.SetBit(b,0)
	else:
		b = bits.ClearBit(b,0)

	img.putpixel((x_offset, y_offset), (r,g,b))

	#print "height: ", height, " width: ", width
	#print "y_offset: ", y_offset, " x_offset: ", x_offset

def readMessage(filename):
	img = Image.open(filename)

	lengthBits = bitarray.bitarray()

	for n in range(32):
		lengthBits.append(readBit(img, n))

	print lengthBits

	length = BitsToInt(lengthBits)

	messageBits = bitarray.bitarray()

	for n in range(length*8):
		messageBits.append(readBit(img, 32+n))

	message = bitsToString(messageBits);

	print message

	return message


def readBit(img, offset):
	width = img.width
	height = img.height

	y_offset = offset/width
	x_offset = (offset - y_offset*width)

	r,g,b = img.getpixel((x_offset, y_offset))

	return bits.HasBit(b, 0)