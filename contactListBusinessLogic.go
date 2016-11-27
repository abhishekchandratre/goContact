/* Author 1: Chaitanya Sri Krishna 
Contains the business Logic for the implementation of the Contact List.
It contains the following features:
1) Create File in a specified Path.
2) Insert the contact details into the CSV File.
3) Retrieves the data from CSV and forms the data for Front end.
4) Logical Deletion of the Record 
5) Update Record against a particular Field.
Low level Details:
1) Contains Implementations for the File Read, write.
2) Contains implementation for the indexing of the CSV File.
 */
package main

import(
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	s "strings"
	"log"
)

//  Create a file in a specified location.
func createFile(fullPath string){

	f,err:= os.Create(fullPath)
	check(err)
	 defer f.Close()

	dataString := "sep=|\n"
	n3, err1 := f.WriteString(dataString)
	 check(err1)
	fmt.Println("Succesffully written the file with length : ",n3)
}

// Write records into the given file.
func writeRecordIntoFile(dataString string, filename string){
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)

	defer f.Close()

	_, errWrite := f.WriteString(dataString);
	check(errWrite)
}

// Add Records into the give file.
func addRecordsToTheFile(name string, phoneNumber string,email string, address string, deleteFlag string){
	id := getRecordNumber("directory.csv")
	var index string = strconv.Itoa(id)
	var text string = index+"|"+name+"|"+phoneNumber+"|"+email+"|"+address+"|"+"0"+"\n"
	fmt.Println(text)
	writeRecordIntoFile(text,"directory.csv")
}

// Indexing Implementation which checks the current position of the record in CSV File.
func getRecordNumber(filename string) int {
	bytes,err := ioutil.ReadFile(filename)
	check(err)
	var data string = string(bytes)
	strarray := s.Split(data,"\n")
	id := len(strarray)-1
	return id
}

// Retrieve the CSV contact data and form a front end data.
func retrieveContactData(filename string) map[string][]string{
	bytes,err := ioutil.ReadFile(filename)
	check(err)
	var data string = string(bytes)

	records := s.Split(data,"\n")

	// Form the Front end data.
	dataToBeSent := make(map[string][]string)
	mapStr := make([]string, 1)
	count := 1;
	for i:=1; i<len(records);i++ {
		str:= records[i]
		newstr:= s.Split(str,"|")
		deleteFlag := newstr[len(newstr)-1]
		if(s.Compare(deleteFlag,"0")==0){

			finalstr:= s.Join(newstr,";")
			if(count==1){
				mapStr[0] = finalstr;
			} else{
				mapStr = append(mapStr,finalstr)
			}
			count+=1;
		}

	}
	dataToBeSent["contacts"] = mapStr
	fmt.Println("The data to be sent is :", dataToBeSent["contacts"])

	return dataToBeSent
}

// Logical deletion of the Contact Data.
func deleteRecord(filename string,id string){

	bytes,err := ioutil.ReadFile(filename)
	check(err)
	var data string = string(bytes)

	records := s.Split(data,"\n")

	createFile("directory.csv")
	// Form the Front end data.
	//count := 1;
	for i:=1; i<len(records);i++ {
		str:= records[i]
		newstr:= s.Split(str,"|")
		if(s.Compare(newstr[0],id)==0) {
			deleteVal := len(newstr)-1
			newstr[deleteVal] = "1"
		}

		finalstr:= s.Join(newstr,"|")
		finalstr = finalstr+"\n"
		writeRecordIntoFile(finalstr,filename)
	}
}

// Update Records against a specific Field given.
func updateRecord(filename string, id string, targetstr string, targetval string){

	bytes,err := ioutil.ReadFile(filename)
	check(err)
	var data string = string(bytes)

	records := s.Split(data,"\n")

	createFile("directory.csv")
	// Form the Front end data.
	//count := 1;
	for i:=1; i<len(records);i++ {
		str:= records[i]
		newstr:= s.Split(str,"|")
		if(s.Compare(newstr[0],id)==0) {
			//deleteVal := len(newstr)-1
			if(s.Compare(targetstr,"name") == 0 ){
				newstr[1] = targetval
			} else if(s.Compare(targetstr,"phoneNo") == 0 ){
				newstr[2] = targetval
			} else if(s.Compare(targetstr,"email")==0){
				newstr[3] = targetval
			} else if(s.Compare(targetstr,"text") == 0){
				newstr[4] = targetval
			}
		}

		finalstr:= s.Join(newstr,"|")
		finalstr = finalstr+"\n"
		writeRecordIntoFile(finalstr,filename)
	}
}

// Implementatation for Error Handling.
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}


// Main function implementation for the busines Logic.
func main() {
	//createFile("directory.csv")
	//addRecordsToTheFile("Chaitanya Lolla","9803187958","lollachaitanya@yahoo.com","UT","0")
	//retrieveContactData("directory.csv")
	//updateRecord("directory.csv","1","name","Chennu")
	deleteRecord("directory.csv","1")
}
