import re
from shutil import copyfile

max = 5
'''
def orderEmails(filename):
	cant = filterEmails(filename)
	fileNumber = getDepth(cant)
	lastFile = "./files/" + str(fileNumber) + "_0.txt"
	return lastFile
'''

def orderEmails(filename):
	cant = filterEmails(filename)
	fileNumber = getDepth(cant)
	lastFile = './files/' + str(fileNumber) + "_0.txt";
	return lastFile

def filterEmails(filename):
	selected = "./files/filtered_emails.txt"
	emailsF = open(filename, "r")
	emails = emailsF.read()
	emails = emails.lower()

	emails = emails.split("\n")
	print emails

	emailsFiltered = []
	pattern = "(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})"
	regexp = re.compile(pattern)

	for email in emails:
		if regexp.match(email):
			emailsFiltered.append(email)

	print emailsFiltered

	fileW = open(selected, "w+")

	for n in emailsFiltered:
		fileW.write(n + "\n")

	fileW.close()
	emailsF.close()

	cant = createSortedLeaves(emailsFiltered)
	print "Create sorted leaves returned cant: ", cant
	mergeLeaves(cant)
	return cant

def odd(number):
	if number == 1:
		return False
	return number%2 != 0

def getDepth(number):
	number_of_cycles = 0
	temp = number

	while temp > 1:
		if odd(temp):
			temp += 1
		number_of_cycles += 1
		temp = temp//2

	return number_of_cycles

def getNextNumberOfFiles(count):
	temp = count
	if odd(temp):
		temp += 1
	return temp//2

def merge(nfiles, ncreates, ccreated, depth):
	print "nfiles: ", nfiles, " ncreates: ", ncreates, " ccreated: ", ccreated, " depth: ", depth

	if ccreated < ncreates:
		if odd(nfiles) and ccreated + 1 == ncreates:
			fCopy = "./files/" + str(depth-1) + "_" + str(nfiles-1) + ".txt"
			fResult = "./files/" + str(depth) + "_" + str(ccreated) + ".txt"

			copyfile(fCopy, fResult)

			print "copying ", depth, "_", ccreated, " from ", fCopy

		else:
			fResult = "./files/" + str(depth) + "_" + str(ccreated) + ".txt"
			file1 = "./files/" + str(depth-1) + "_" + str(ccreated*2) + ".txt"
			file2 = "./files/" + str(depth-1) + "_" + str(ccreated*2+1) + ".txt"

			print "creating ", depth, "_", ccreated, " from ", file1, " and ", file2

			mergeTwoFiles(file1, file2, fResult)
			merge(nfiles, ncreates, ccreated+1, depth)

def mergeTwoFiles(filel, filer, fileResult):

	print "merging: ", filel, " ", filer, " to: ", fileResult
	fOpenl = open(filel, "r")
	fOpenr = open(filer, "r")
	fResult = open(fileResult, "w")

	l_email = fOpenl.readline()
	r_email = fOpenr.readline()

	while(True):
		if len(l_email) > 0 and len(r_email) > 0:
			if l_email < r_email:
				fResult.write(l_email)
				l_email = fOpenl.readline()
			else:
				fResult.write(r_email)
				r_email = fOpenr.readline()
		else:
			if len(l_email) > 0:
				fResult.write(l_email)
				l_email = fOpenl.readline()
			elif len(r_email) > 0:
				fResult.write(r_email)
				r_email = fOpenr.readline()
			else:
				break

def mergeLeaves(count):
	depth = getDepth(count)
	number_of_files = count
	i = depth
	while(i > 0):
		ncreates = getNextNumberOfFiles(number_of_files)
		merge(number_of_files, ncreates, 0, depth-i+1)
		number_of_files = ncreates
		i -= 1

def createSortedLeaves(emails):
	cant = len(emails) / max
	if cant*max < len(emails):
		cant += 1

	for n in range(0, cant):
		filename = "./files/0_" + str(n) + ".txt"
		file = open(filename, "w")

		right = n*5 + 5
		if right >= len(emails):
			right = len(emails)

		emailSorted = emails[n*5 : right]
		emailSorted.sort()

		toWrite = max
		if len(emailSorted) < max:
			toWrite = len(emailSorted)

		for m in range(0, toWrite):
			store = emailSorted[m] + "\n"
			file.write(store)

		file.close()
	return cant

def test2(filename):
	file = open(filename, "a+")
	file.seek(0,0)
	#file.write("hola")

	texto = file.read()
	print "------", texto

	file.close()


def test():
	orderEmails("./uploads/correos.txt")
	#test2("./uploads/correos.txt")
	return
	filename = "./uploads/correos2.txt"
	file = open(filename, "w+")
	#file.truncate()
	print file.readline()
	pattern = "(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})"
	regexp = re.compile(pattern)
	array = []
	with open(filename, "r") as f:
	  for line in f:
	    if re.search(pattern, line) :
	    	print "match: ", line
	    	array.append(line)
	    else:
	    	print "_don't match: ", line
	#print array
	file.close()