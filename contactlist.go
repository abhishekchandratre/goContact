/* Author 1: Chaitanya Sri Krishna 
   Author 2: Nikhil Jonnalagadda
1. Contains the Backend implementation for retrieving the directory records of the Contact List.

 */
package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"strconv"
	"strings"
)
var lineCount int = 0
func main() {

		/*file,err := os.Open("contactlist.csv")
		checkError("Cannot read file",err)

		// Get a CSV Reader Object.
		reader := csv.NewReader(file)
		reader.Comma = ';'
		records,err := reader.ReadAll()	
		checkError("Read Error",err)
		// Finding the number of records.
		lineCount := 0
		for _, record := range records {
			fmt.Println("Records are :",record)
			lineCount+=1;
		}


		id := lineCount+1;


		fmt.Println("The number of records in CSV File is : ", lineCount)
		fmt.Println("The pointer in the CSV File is : ", id)

		file.Close()

		//writeRecordsIntoDirectory(id,"name")
		writeRecordsIntoDirectory(id,"Tejaswi","9803181751","abhi@yahoo.com","Ashford UT","0") */

		//sendDataToFrontEnd()
		//defer file.Close()
		deleteRecords("1")

}


func checkError(message string, err error){
	if err!= nil {
		log.Fatal(message,err)
	}

}



func writeRecordsIntoDirectory(id int, name string, phoneNumber string,email string, address string, deleteFlag string)  {

	file,err := os.Open("contactlist.csv")
	checkError("Cannot read file",err)

		// Get a CSV Reader Object.
	reader := csv.NewReader(file)
	reader.Comma = ','

	records,err:= reader.ReadAll()
	checkError("Read Error", err)
	
	file.Close()

	var err1 = os.Remove("contactlist.csv")
	checkError("Delete Error", err1)

	csvfile, err2 := os.Create("contactlist.csv")
	checkError("Creation error",err2)
	writer := csv.NewWriter(csvfile)

	var index string = strconv.Itoa(id)
	data := [][]string{{index,name,phoneNumber,email,address,deleteFlag}}

	for _, record := range records {
		fmt.Println("ENtered")
		fmt.Println("Value of teh DAta is :",record)
		err:= writer.Write(record)
		checkError("Cannot Write into the file",err)
	}

	for _,value := range data {
		fmt.Println("New Data")
		err:= writer.Write(value)
		checkError("Cannot Write into the file",err)
	}

	 writer.Flush()
	 csvfile.Close()

}

/* Get all the Contacts from the CSV File */
func sendDataToFrontEnd(){

	file,err := os.Open("contactlist.csv")
	checkError("Cannot read file",err)

	// Get a CSV Reader Object.
	reader := csv.NewReader(file)
	reader.Comma = ','

	records,err:= reader.ReadAll()
	dataToBeSent := make(map[string][]string)
	s := make([]string, 1)

	count := 1;
	//Convert Record Objects to Strings
	for _, record := range records {
		
		fmt.Println("Value of teh data is :",record)
		var recordString string= ""
		for i :=0 ; i < len(record); i++ {
			var str string = record[i]
			if i == len(record) -1  {
				recordString+= str
			}else{
				recordString+=str
				recordString+=";"
			}

			

		}

		if(count == 1){
			s[0] = recordString;
		} else{
			s = append(s,recordString)
		}
		
		count+=1;
		
	}

	dataToBeSent["contacts"] = s

	fmt.Println("The data is :", dataToBeSent["contacts"])
	
}

func deleteRecords(id string){

	file,err := os.Open("contactlist.csv")
	checkError("Cannot read file",err)

		// Get a CSV Reader Object.
	reader := csv.NewReader(file)
	reader.Comma = ','

	records,err:= reader.ReadAll()
	checkError("Read Error", err)
	
	file.Close()

	var err1 = os.Remove("contactlist.csv")
	checkError("Delete Error", err1)

	csvfile, err2 := os.Create("contactlist.csv")
	checkError("Creation error",err2)
	writer := csv.NewWriter(csvfile)

	for _, record := range records {
		fmt.Println("Value of the Data is :",record)

		for i:=0; i<len(record); i++ {
			var val int = strings.Compare(record[0],id)
			if(val == 0){
				record[5] = "1"
			}
		}

		err:= writer.Write(record)
		checkError("Cannot Write into the file",err)
	}

	 writer.Flush()
	 csvfile.Close()


}