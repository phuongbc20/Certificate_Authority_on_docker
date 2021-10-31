package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", Action)
	http.ListenAndServe(":8080", nil)
}

//--------------------------------------

func Action(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := r.FormFile("File")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	name := handler.Filename[0 : len(handler.Filename)-4]
	Type := handler.Filename[len(handler.Filename)-4 : len(handler.Filename)]
	tempFile, err := ioutil.TempFile("Uploads", name+"*"+Type)

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	//------------Sign for CSR-------------//

	nameCSR := tempFile.Name()[8 : len(tempFile.Name())-4]
	cmd := exec.Command("./sign.sh", "-f", nameCSR)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Sign CSR fail")
		return
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println(err)
	}

	//------------Response to Client-------------//

	FileName := "pki/issued/" + nameCSR + ".crt"
	fileCrt, err := os.Open(FileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileCrt.Close()
	io.Copy(w, fileCrt)
	err = os.Remove(FileName)
	if err != nil {
		fmt.Println(err)
	}
}
