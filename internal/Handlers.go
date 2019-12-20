package internal

import (
	"encoding/json"
	"fmt"
	myPkg "file-download-manager/pkg/myPackage"
	mainHandler "file-download-manager/cmd/main"
)

//func checkHealth(){
//	fmt.Println("All good")
//}
//
//func downloadFile(){
//
//}


func downloadFileById(params []string){
	var downFs = []myPkg.downloadInfo{file_path:"xyz",download_path:"abc",}
	var newUrl = []string{"qwe"}
	//type FileS = myPkg.Files
	var FileS = []myPkg.Files{
		ID:"1",url:newUrl,start_time:"st1",end_time:"et1",download_type:"SERIAL",file_info:downFs,
	}
	for _,item := range FileS(){
		if item.ID == params["id"]{
			return item
		}]
	}
}