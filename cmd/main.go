package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	mypkg "file-download-manager/pkg/myPackage"
)

var fs = []mypkg.Files{}
var downFs = []mypkg.downloadInfo{}

func main(){
	r:= mux.NewRouter()

	downFs = append(downFs,mypkg.downloadInfo{file_path:"xyz",download_path:"abc"})
	var newUrl = []string{"qwe"}
	fs = append(fs, mypkg.Files{ID:"1",url:newUrl,start_time:"st1",end_time:"et1",download_type:"SERIAL",file_info:downFs})

	//r.HandleFunc("/health", checkHealth()).Methods("GET")
	//r.HandleFunc("/downloads", downloadFile()).Methods("POST")
	r.HandleFunc("/downloads/{id}", downloadFileById()).Methods("GET")
	//r.HandleFunc("/files", getFiles()).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000",r))

}
