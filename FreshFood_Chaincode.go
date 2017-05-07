package main

import (
                 "time"
	"errors"
	"fmt"
                  "encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
              

// SimpleChaincode example simple Chaincode implementation

type SimpleChaincode struct {
}

// ============================================================================================================================
//  Customer Definitions
// ============================================================================================================================
type Customer struct {
	ObjectType string        `json:"docType"` //field for couchdb
	Invno           string          `json:"invno"`      //the fieldtags are needed to keep case from bouncing around
	Item             string          `json:"item"`
	Quantity      string               `json:"quantity"`    //size in mm of marble
	 Cost           string                `json:"cost"`  
}
// ============================================================================================================================
//  Retailer Definitions
// ============================================================================================================================
type Retailer struct {
	ObjectType string        `json:"docType"` //field for couchdb
	Invno           string          `json:"invno"`      //the fieldtags are needed to keep case from bouncing around
	Item             string          `json:"item"`
                  Itemid          string          `json:"itemid"`
	Distid          string           `json:"distid"`     
	PurchDate   string           `json:"PurchDate"`  
}

// ============================================================================================================================
//  Manufacture Definitions
// ============================================================================================================================
type Manufacturer struct {
	ObjectType string        `json:"docType"` //field for couchdb
                  Manid         string          `jason:"manid"`	
	Item             string          `json:"item"`
                  Itemid          string          `json:"itemid"`
                  ManDate      string           `jason:"mandate"`
	Quality         string           `json:"quality"`     
	Ndays          string           `json:"ndays"`  
}
// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}
// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil,err
	}

	return nil, nil
}
// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
                  }
	if function == "write" {
		return t.write(stub, args)
                  }
                  if function == "CreateRetailerDB" {
		return t.CreateRetailerDB(stub, args)
                  }
	
                  if function == "CreateManDB" {
		return t.CreateManDB(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}
// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	fmt.Println("running write()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var cust Customer
                   cust.ObjectType  = "cust_type"
                   cust.Invno  = args[0]
                   cust.Item  = args[1]
                   cust.Quantity = args[2]
                   cust.Cost =args[3]
                    
                     
                    
                   
                  custAsBytes,_  :=  json.Marshal(cust) 
                  
                  
                 

                   err = stub.PutState(cust.Invno, custAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return nil, nil
}
// CreateRetailerDB
func (t *SimpleChaincode) CreateRetailerDB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
                  var retailkey string
	

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var retail Retailer
                   retail.ObjectType  = "retailer_detail"
                   retail.Invno  = args[0]
                   retail.Item  = args[1]
                   retail.Itemid = args[2]
                   retail.Distid = args[3]
                   retail.PurchDate = args[4]
                   
                retailkey = retail.Invno + retail.Item

                  retailAsBytes,_  :=  json.Marshal(retail) 

                   err = stub.PutState(retailkey, retailAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return nil, nil
}

// CreateManDB
func (t *SimpleChaincode) CreateManDB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
                  var mankey string
	

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var manf Manufacturer
                   manf.ObjectType  = "Manufacturer_detail"
                   manf.Manid = "M001"
                   manf.Itemid = args[0]

                    if manf.Itemid ==  "1" {
                   manf.Item  = "Plain Bread"
                   } else if manf.Itemid == "2" {
                    manf.Item = "Wheat Bread"
                   }else {
                   manf.Item = "Default"
                   }

                   c := time.Now().Local()

                   manf.ManDate =  c.String()
                   manf.Quality = args[1]
                   
                    if manf.Quality==  "A" {
                   manf.Ndays  = "10"
                   } else if manf.Quality == "B" {
                    manf.Ndays = "7"
                   } else  {
                   manf.Ndays = "5"
                   }
                   
                 
                   
                mankey = manf.Manid + manf.Item

                  manfAsBytes,_  :=  json.Marshal(manf) 

                   err = stub.PutState(mankey, manfAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return nil, nil
}
// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string,) ([]byte,error) {
	var key,jsonResp string
	var err error                    
	if len(args) != 1 {
		return nil,errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

                   
                  
          
	key = args[0]
                                     
                
	valuex,err := stub.GetState(key)

                   

               
          
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key+ "\"}"
		return nil, errors.New(jsonResp)
	}
                          
	return valuex,nil
                             
}
