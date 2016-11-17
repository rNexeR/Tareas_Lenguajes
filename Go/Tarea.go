package main

import (
    "github.com/go-martini/martini"
    //"github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"
    "net/http"
    "io"
    "io/ioutil"
    "os"
    "fmt"
    "regexp"
    "strconv"
    "bufio"
    "sort"
    s "strings"
    "./emails"
    //"github.com/martini-contrib/cors"
    //"fmt"
)

const max int = 5

func OrderEmails(filename string) string {
    FilterEmails(filename)
    return "hola"
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func FilterEmails(filename string) string {
    selected := "./files/filtered_emails.txt"
    emails, err := ioutil.ReadFile(filename)
    emailsString := string(emails)
    check(err)
    fileW, err := os.Create(selected)
    check(err)
    defer fileW.Close()

    emailsString = s.ToLower(emailsString)

    rg, err := regexp.Compile("([a-z]+)@([a-z]+)\\.([a-z]+)")
    check(err)

    emailsSelected := rg.FindAllString(emailsString, -1)
    fmt.Println(len(emailsSelected))
    
    for _, email := range emailsSelected {
        fileW.WriteString(email + "\n")
    }

    cant := createSortedLeaves(emailsSelected)
    fmt.Println(cant)
    mergeLeaves(cant)
    return selected
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
        //fmt.Println("\tTemp: " + strconv.Itoa(temp))
    }
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
            fmt.Printf("copying %d_%d\n", depth, ccreated)
        }else{
            fmt.Printf("creating %d_%d\n", depth, ccreated)
            merge(nfiles, ncreates, ccreated + 1, depth)
        }
    }
}

func seekForever() {
    file, err := os.Open("./uploads/emails.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lectura := scanner.Text()
        fmt.Println(lectura)
        cpos, _ := file.Seek(0, 1)
        fmt.Printf("cpos: %d\n", cpos)
        file.Seek(cpos - int64(len(lectura) + 1), 0)
        scanner = bufio.NewScanner(file)
        //fmt.Println("\t" + scanner.Text())
        //break

    }
}

///*
func mergeLeaves(count int) {
    depth := getDepth(count)
    number_of_files := count
    for i:= depth; i > 0; i-- {
        ncreates := getNextNumberOfFiles(number_of_files)
        merge(number_of_files, ncreates, 0, depth - i + 1)
        number_of_files = ncreates
    }

    /*
    filename := "./files/" + cnumber + "_"
    _, err := os.Stat(filename)
    if err == nil{
        _, err := os.Create(filename)
        check(err)
    }
    sortedFile, err := os.Open(filename)
    check(err)
    leaveFile, err := os.Open(strconv.Itoa(currentFile))
    check(err)
    

    if currentFile == 0{
        io.Copy(sortedFile, leaveFile)
    }else{

    }

    leaveFile.Close()
    sortedFile.Close()

    if currentFile + 1 < count{
        //recursion
    }
    */
}
//*/

func createSortedLeaves(emails []string) int {
    cant := len(emails) / max

    if cant*max < len(emails){
        cant++
    }

    fmt.Println(cant)

    for i:=0; i < cant; i++ {
        f, err:= os.Create("./files/0_" + strconv.Itoa(i) + ".txt")
        check(err)

        right := i*5 +5;
        if right >= len(emails){
            right = len(emails)
        }

        emailSorted := emails[i*5:right]
        sort.Strings(emailSorted)

        toWrite := max
        if (len(emailSorted) < max){
            toWrite = len(emailSorted)
        }

        for j:=0; j < toWrite; j++ {
            f.WriteString(emailSorted[j] + "\n")
        }
        
        f.Close()
    }

    return cant
}


func main(){
    fmt.Println("Con 5")
    mergeLeaves(5);

    seekForever()

    emails.SayHello()
    /*
    fmt.Println("Con 5")
    mergeLeaves(5);
    fmt.Println("Con 28")
    mergeLeaves(28);
    fmt.Println("Con 29")
    mergeLeaves(29);
    fmt.Println("Con 100")
    mergeLeaves(100);
    */

    m := martini.Classic();

    m.Use(render.Renderer())

    m.Get("/",  func(r render.Render) {
        r.JSON(http.StatusOK, "It works")
    })

    m.Post("/orderEmails", func(r *http.Request, res render.Render){
        err := r.ParseMultipartForm(100000)
        if err != nil {
            res.JSON(http.StatusInternalServerError, err.Error())
        }
        files := r.MultipartForm.File["files"]
        file := files[0]
        fOpen, err := file.Open()
        defer fOpen.Close()
        if err != nil {
            res.JSON(http.StatusInternalServerError, err.Error())
        }

        dst, err := os.Create("./uploads/emails.txt")
        defer dst.Close()
        if err != nil {
            res.JSON(http.StatusInternalServerError, err.Error())
        }

        if _, err := io.Copy(dst, fOpen); err != nil {
                res.JSON(http.StatusInternalServerError, err.Error())
        }

        //call function to order emails
        OrderEmails("./uploads/emails.txt")
        res.JSON(http.StatusOK, "Sorting")

        })

    m.Get("/download", func(r render.Render) {
        f, err := ioutil.ReadFile("./uploads/71.jpg");
        if err != nil {
            r.JSON(http.StatusInternalServerError, err.Error())
        }
        r.Data(http.StatusOK, f);
    })

    m.Run()
}