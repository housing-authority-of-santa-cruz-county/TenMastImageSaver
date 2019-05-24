package main

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strings"
)

var (
    fileName    string
    fullUrlFile string
)

func main() {

    fullUrlFile = "http://hasctd1/members/GetPriPreview.aspx?Width=750&PriGuid=e8ec7a41-228c-4df4-9ddc-af4c2963263f&ext=.png"

    // Build fileName from fullPath
    buildFileName()

    // Create blank file
    file := createFile()

    // Put content on file
    putFile(file, httpClient())

}

func putFile(file *os.File, client *http.Client) {

  var username string = "admin"
  var passwd string = "Tendocs!"

  //req, err := http.NewRequest("GET", "http://hasctd1/members/GetPriPreview.aspx?Width=750&PriGuid=e8ec7a41-228c-4df4-9ddc-af4c2963263f", nil)
  urlData := url.Values{}
  urlData.Set("ctl00$ContentPlaceHolder1$LoginControl$UserName", username)
  urlData.Set("ctl00$ContentPlaceHolder1$LoginControl$Password", passwd)
  urlData.Set("ctl00_ContentPlaceHolder1_LoginControl_UserName", username)
  urlData.Set("ctl00_ContentPlaceHolder1_LoginControl_Password", passwd)
  urlData.Set("__EVENTVALIDATION","/wEWBQLM3LOQDQLnwJbtCwKq+5mgAQKg262YDwL5j6nUAQVQJ3szmQbob5ITJOCn8uERtK54");
  urlData.Set("__VIEWSTATEGENERATOR","CA0B0334")
  resp, err := http.PostForm("http://hasctd1/Default.aspx?ReturnUrl=%2fmembers%2fGetPriPreview.aspx%3fWidth%3d750%26PriGuid%3de8ec7a41-228c-4df4-9ddc-af4c2963263f&amp;Width=750&amp;PriGuid=e8ec7a41-228c-4df4-9ddc-af4c2963263f", urlData)

  //req, err := http.NewRequest("POST", "", nil)
  //req.SetBasicAuth(username, passwd)
  //resp, err := client.Do(req)


    //resp, err := client.Get(fullUrlFile)

    checkError(err)

    defer resp.Body.Close()

    size, err := io.Copy(file, resp.Body)

    defer file.Close()

    checkError(err)

    fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func buildFileName() {
    fileUrl, err := url.Parse(fullUrlFile)
    checkError(err)

    path := fileUrl.Path
    segments := strings.Split(path, "/")

    fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {

    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }

    return &client
}

func createFile() *os.File {
    file, err := os.Create(fileName)

    checkError(err)
    return file
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
