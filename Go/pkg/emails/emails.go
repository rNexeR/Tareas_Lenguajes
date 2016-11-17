package emails

import (
   //"io"
    //"io/ioutil"
    "io"
    "io/ioutil"
    "os"
    "fmt"
    "regexp"
    "strconv"
    "bufio"
    "sort"
    s "strings"
    //"os"
)

const max int = 5

func OrderEmails(filename string) string {
    cant := FilterEmails(filename)
    fileNumber := getDepth(cant)
    lastfile := "./files/" + strconv.Itoa(fileNumber) + "_0.txt"
    return lastfile
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func pressAnyKey(){
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func FilterEmails(filename string) int {
    selected := "./files/filtered_emails.txt"
    emails, err := ioutil.ReadFile(filename)
    emailsString := string(emails)
    check(err)
    fileW, err := os.Create(selected)
    check(err)
    defer fileW.Close()

    emailsString = s.ToLower(emailsString)

    rg, err := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
    check(err)

    emailsSelected := rg.FindAllString(emailsString, -1)
    //fmt.Println(len(emailsSelected))
    
    for _, email := range emailsSelected {
        fileW.WriteString(email + "\n")
    }

    cant := createSortedLeaves(emailsSelected)
    fmt.Println(cant)
    MergeLeaves(cant)
    return cant
}

func odd(number int) bool {
    if number == 1 {
        return false
    }
    return number % 2 != 0
}

func getDepth(number int) int {
    number_of_cycles := 0
    temp := number

    for temp > 1 {
        if odd(temp){
            temp++
        }
        number_of_cycles++
        temp = temp / 2
    }
    //fmt.Printf("With %d the depth is: %d\n", number, number_of_cycles)
    return number_of_cycles
}

func getNextNumberOfFiles(count int) int {
    if odd(count){
        count++
    }
    return count/2
}

func merge(nfiles, ncreates, ccreated, depth int) {
    fmt.Printf("nfiles: %d ncreates: %d ccreated: %d depth: %d\n", nfiles, ncreates, ccreated, depth )

    if ccreated < ncreates{
        if odd(nfiles) && ccreated +1  == ncreates {
        	//pressAnyKey()
            fCopy := "./files/" + strconv.Itoa(depth -1) + "_" + strconv.Itoa(nfiles-1) + ".txt"
            fResult := "./files/" + strconv.Itoa(depth) + "_" + strconv.Itoa(ccreated) + ".txt"
            dst, _ := os.Create(fResult)
            defer dst.Close()
            src, _ := os.Open(fCopy)
            defer src.Close()
            _, err := io.Copy(dst, src)
            check(err)
            //fmt.Printf("copying %d_%d from %s\n", depth, ccreated, fCopy)
            //pressAnyKey()
        }else{
            fResult := "./files/" + strconv.Itoa(depth) + "_" + strconv.Itoa(ccreated) + ".txt"
            file1 := "./files/" + strconv.Itoa(depth-1) + "_" + strconv.Itoa(ccreated*2) + ".txt"
            file2 := "./files/" + strconv.Itoa(depth-1) + "_" + strconv.Itoa(ccreated*2 + 1) + ".txt"
            //fmt.Printf("creating %d_%d from %s and %s\n", depth, ccreated, file1, file2)
            mergeTwoFiles(file1, file2, fResult)
            merge(nfiles, ncreates, ccreated + 1, depth)
        }
    }
}

func mergeTwoFiles(file1, file2, fileResult string){
	fOpenl,_ := os.Open(file1)
	defer fOpenl.Close()
	fOpenr, _ := os.Open(file2)
	defer fOpenr.Close()
	fResult, _ := os.Create(fileResult)
	defer fResult.Close()

	scanl := bufio.NewScanner(fOpenl)
	scanr := bufio.NewScanner(fOpenr)

	s_posl := 0
	s_posr := 0

	for{
		scanl.Scan()
		scanr.Scan()

		l_email := scanl.Text()
		r_email := scanr.Text()

		//fmt.Printf("To Compare: %s with %s\n", l_email, r_email)

		if(len(l_email) >0 && len(r_email) > 0) {
			if(l_email < r_email) {
				//fmt.Printf("Writing: %s\n", l_email)
				fResult.WriteString(l_email + "\n")
				s_posl += len(l_email) + 1
				fOpenr.Seek(int64(s_posr), 0)
				scanr = bufio.NewScanner(fOpenr)
			}else{
				//fmt.Printf("Writing: %s\n", r_email)
				fResult.WriteString(r_email + "\n")
				s_posr += len(r_email) + 1
				fOpenl.Seek(int64(s_posl), 0)
				scanl = bufio.NewScanner(fOpenl)
			}
		}else{
			if(len(l_email) > 0){
				//fmt.Printf("Writing: %s\n", l_email)
				fResult.WriteString(l_email + "\n")
			}else if(len(r_email) > 0){
				//fmt.Printf("Writing: %s\n", r_email)
				fResult.WriteString(r_email + "\n")
			}else{
				break
			}
		}
	}
}

///*
func MergeLeaves(count int) {
    depth := getDepth(count)
    number_of_files := count
    for i:= depth; i > 0; i-- {
        ncreates := getNextNumberOfFiles(number_of_files)
        merge(number_of_files, ncreates, 0, depth - i + 1)
        number_of_files = ncreates
    }
}
//*/

func createSortedLeaves(emails []string) int {
    cant := len(emails) / max

    if cant*max < len(emails){
        cant++
    }

    //fmt.Println("Cantidad de hojas a crear: " + strconv.Itoa(cant))

    for i:=0; i < cant; i++ {
    	//fmt.Printf("Hoja actual: %d\n", i)
        f, err:= os.Create("./files/0_" + strconv.Itoa(i) + ".txt")
        check(err)

        right := i*5 +5;
        if right >= len(emails){
            right = len(emails)
        }

        emailSorted := emails[i*5:right]
        sort.Strings(emailSorted)

        //fmt.Print("Sorted emails to store: ")
        //fmt.Println(emailSorted)

        toWrite := max
        if (len(emailSorted) < max){
            toWrite = len(emailSorted)
        }

        for j:=0; j < toWrite; j++ {
        	store := emailSorted[j] + "\n"
        	//fmt.Printf("Email to store: %s \n", store)
            _, err := f.WriteString(store)
            check(err)
            //fmt.Println(n)
        }
        
        f.Close()
    }
    //pressAnyKey()
    return cant
}
