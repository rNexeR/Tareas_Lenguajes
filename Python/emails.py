import re

def sayHello():
	return "hello"
'''
def orderEmails(filename):
	cant = filterEmails(filename)
	fileNumber = getDepth(cant)
	lastFile = "./files/" + str(fileNumber) + "_0.txt"
	return lastFile
'''

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

def test2(filename):
	file = open(filename, "a+")
	file.seek(0,0)
	#file.write("hola")

	texto = file.read()
	print "------", texto

	file.close()


def test():
	filterEmails("./uploads/correos.txt")
	test2("./uploads/correos.txt")
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