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

var fullPath string = ""

/* Function to check if the file exists or not.
 1. This function returns true if the file already exits.
 2. Else it will create a file in the given path and returns false.
*/
func isFileExists(filePath string) bool {

	var isExists bool = false
	_,err := os.Stat(filePath)
	if err!= nil {
		fmt.Println("File doesn't exists in the given directory. A new file will be created.")
		fullPath = filePath;
		createFile()
		// Create The file.
	} else {
		fmt.Println("File already Exists.")
		isExists = true
		fullPath = filePath;
	}

	return isExists
}


//  Create a file in the Location which has been set in the global path.
func createFile(){

	f,err:= os.Create(fullPath)
	check(err)
	 defer f.Close()

	dataString := "sep=|\n"
	n3, err1 := f.WriteString(dataString)
	 check(err1)
	fmt.Println("Succesffully written the file with length : ",n3)
}


// Write records into the given file specified in the global path.
func writeRecordIntoFile(dataString string){
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)

	defer f.Close()

	_, errWrite := f.WriteString(dataString);
	check(errWrite)
}

// Add Records into the give file specified in the global path.
func addRecordsToTheFile(name string, phoneNumber string,email string, address string, deleteFlag string){
	id := getRecordNumber(fullPath)
	var index string = strconv.Itoa(id)
	var text string = index+"|"+name+"|"+phoneNumber+"|"+email+"|"+address+"|"+"0"+"\n"
	fmt.Println(text)
	writeRecordIntoFile(text)
}

//Helper Function.
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
func retrieveContactData() map[string][]string{
	bytes,err := ioutil.ReadFile(fullPath)
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
		//deleteFlag := newstr[len(newstr)-1]
		//TODO: No need for the condition.
		//if(s.Compare(deleteFlag,"0")==0){

			finalstr:= s.Join(newstr,";")
			if(count==1){
				mapStr[0] = finalstr;
			} else{
				mapStr = append(mapStr,finalstr)
			}
			count+=1;
		//}

	}
	dataToBeSent["contacts"] = mapStr
	fmt.Println("The data to be sent is :", dataToBeSent["contacts"])

	return dataToBeSent
}

// Logical deletion of the Contact Data.
func deleteRecord(id string){

	bytes,err := ioutil.ReadFile(fullPath)
	check(err)
	var data string = string(bytes)

	records := s.Split(data,"\n")

	createFile()
	// Form the Front end data.
	//count := 1;
	for i:=1; i<len(records)-1;i++ {
		str:= records[i]
		newstr:= s.Split(str,"|")
		if(s.Compare(newstr[0],id)==0) {
			deleteVal := len(newstr)-1
			newstr[deleteVal] = "1"
		}

		finalstr:= s.Join(newstr,"|")
		finalstr = finalstr+"\n"
		writeRecordIntoFile(finalstr)
	}
}

// Update Records against a specific Field given.
func updateRecord(id string, targetstr string, targetval string){

	bytes,err := ioutil.ReadFile(fullPath)
	check(err)
	var data string = string(bytes)

	records := s.Split(data,"\n")
	fmt.Println("The records are: ",records)
	fmt.Println("The length of the records are : ", len(records))
	createFile()
	// Form the Front end data.
	//count := 1;
	for i:=1; i<len(records)-1;i++ {
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
		fmt.Println(finalstr)
		finalstr = finalstr+"\n"
		writeRecordIntoFile(finalstr)
	}
}

// Implementatation for Error Handling.
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}


// Main function implementation for the busines Logic.
//func main() {
	 //isFileExists("newdir1.csv")
	//createFile("directory.csv")
	//addRecordsToTheFile("Chaitanya","9803187958","chaitanya@gmail.com","UT","0")
	//addRecordsToTheFile("Nikhil","9803187951","nikhil@gmail.com","UT","0")
	//addRecordsToTheFile("Abhishek","9803187952","abhi@gmail.com","Harrisburg","0")
	//addRecordsToTheFile("Teja","9803187953","teja@gmail.com","Ashford","0")
	//retrieveContactData()
	//updateRecord("1","email","abhishek@gmail.com")
	//deleteRecord("5")
//}


