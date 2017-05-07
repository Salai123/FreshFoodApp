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

type Date struct{ time.Time }

// ============================================================================================================================
//  Response Definitions
// ============================================================================================================================
type Response struct {
	ObjectType string        `json:"Response"` //field for couchdb
                   Invno           string          `json:"Invoice No"`      //the fieldtags are needed to keep case from bouncing around
	 Item             string          `json:"Item"`
                   Retid          string           `json:"Retailer id "`     
                   PurchDate   string           `json:"Retailer Purchase Date"`       
	 Distid          string           `json:"Distributor Id"`     
                   DPurchDate   string           `json:"Distributor Purchase Date"`  
	 ExpDate      string           `jason:"Expiry Date"`	
                  Manid         string          `jason:"Manufacturer id"`
                  ManDate      string           `jason:"Manufacture Date"`
	Quality         string           `json:"Quality"`     	 
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
	Distid          string           `json:"distid"`     
	PurchDate   string           `json:"PurchDate"`
                  Retid        string              `json:"Retid"` 
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
	Ndays          int           `json:"ndays"`  
                  SellDate     string             `json:"selldate"`
                  Distid          string           `json:"Distid"`
}
// ============================================================================================================================
//  Distributor Definitions
// ============================================================================================================================
type Distributor struct {
	ObjectType string        `json:"docType"` //field for couchdb
                  Distid          string           `json:"distid"`	
	Item             string          `json:"item"`
                  Manid          string            `json:"manid"`
                  Retid          string           `json:"retid"` 
                  ExpDate      string           `jason:"expdate"`	    
	DPurchDate   string           `json:"dpurchdate"`  
                  SellDate   string           `json:"SellDate"`  
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
                   if function == "CreateDistributorDB" {
		return t.CreateDistributorDB(stub, args)
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
	return custAsBytes, nil
}
// CreateRetailerDB
func (t *SimpleChaincode) CreateRetailerDB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
                  var retailkey string
	

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var retail Retailer
                   retail.ObjectType  = "retailer_detail"
                   retail.Invno  = args[1]
                  retail.Retid = args[0]
                   retail.Item  = args[2]
                   retail.Distid = args[3]
                   retail.PurchDate = "2017-05-12 15:04:05"
                   
                retailkey = retail.Invno + retail.Item

                  retailAsBytes,_  :=  json.Marshal(retail) 

                   err = stub.PutState(retailkey, retailAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return retailAsBytes, nil
}
// CreateDistributorDB
func (t *SimpleChaincode) CreateDistributorDB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
                  var key string
                  
	var manud Manufacturer

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var dist Distributor
                   dist.ObjectType  = "Distributor_detail"
                   dist.Distid  = args[0]
                   dist.Item  = args[1]
                   dist.Manid = args[2]
                   dist.Retid = args[3]
                       
                   dist.DPurchDate = "2017-05-10 15:04:05"

              
                      key = dist.Manid + dist.Distid+dist.Item+dist.DPurchDate


                   valuezi,err := stub.GetState(key)
                   json.Unmarshal(valuezi,&manud)

                  tt,err := time.Parse("2017-01-31",manud.ManDate)

                 

                   et  := tt.AddDate(0,0,manud.Ndays)

                    dist.ExpDate = et.String()
             
                   
                     

                   

                    dist.SellDate = "2017-05-12 15:04:05"
                   

                  distAsBytes,_  :=  json.Marshal(dist) 

                   err = stub.PutState(dist.Distid+dist.Retid+dist.Item+ dist.SellDate, distAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return distAsBytes, nil
}

// CreateManDB
func (t *SimpleChaincode) CreateManDB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
                  var mankey string
	

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
                   
                   var manf Manufacturer
                   manf.ObjectType  = "Manufacturer_detail"
                   manf.Manid = args[0]
                   manf.Itemid = args[1]

                   manf.SellDate = "2017-05-10 15:04:05"

                    if manf.Itemid ==  "1" {
                   manf.Item  = "Plain Bread"
                   } else if manf.Itemid == "2" {
                    manf.Item = "Wheat Bread"
                   }else {
                   manf.Item = "Default"
                   }

                   

                    c :=  time.Now().Local()
                   manf.ManDate = c.String()
                   manf.Quality = args[2]
                   
                    if manf.Quality ==  "A" {
                   manf.Ndays  = 10
                   } else if manf.Quality == "B" {
                    manf.Ndays = 7
                   } else  {
                   manf.Ndays = 5
                   }
                   
                 manf.Distid = args[3]
                   
                mankey = manf.Manid + manf.Distid+manf.Item+manf.SellDate

                  manfAsBytes,_  :=  json.Marshal(manf) 

                   err = stub.PutState(mankey, manfAsBytes)

                   if err != nil {
			
                    return nil, err
	
                    }
	return manfAsBytes, nil
}
// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string,) ([]byte,error) {
	var key,key2,key3,key4,jsonResp string
	var err error                    
	if len(args) != 1 {
		return nil,errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

                   
                  var custr Customer
                  var retr Retailer 
                  var distr Distributor
                  var manfr Manufacturer
                    
         
                 var response Response
              
                   
          
	key = args[0]
                                     
                
	valuex,err := stub.GetState(key)
                   if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key+ "\"}"
		return nil, errors.New(jsonResp)
	}

                  json.Unmarshal(valuex,&custr)
 
                   key2 = custr.Invno+custr.Item

                  valuey,err := stub.GetState(key2)
                   
                   json.Unmarshal(valuey,&retr)
                    
                  key3=retr.Distid+retr.Retid+retr.Item+ retr.PurchDate
                   
                  valuez,err := stub.GetState(key3)
                   json.Unmarshal(valuez,&distr)
                   
                   key4 = distr.Manid + distr.Distid+distr.Item+distr.DPurchDate


                   valuezz,err := stub.GetState(key4)
                   json.Unmarshal(valuezz,&manfr)

                   response.Invno  =          custr.Invno
                  
	 response.Item         =     custr.Item
                   response.Retid          =   distr.Retid
                   response.PurchDate    =  retr.PurchDate
	 response.Distid              = distr.Distid
                   response.DPurchDate    = distr.DPurchDate
	 response.ExpDate       = distr.ExpDate
                  response.Manid         =  manfr.Manid
                  response.ManDate       = manfr.ManDate
	response.Quality             = manfr.Quality

                  
                    resAsBytes,_  :=  json.Marshal(response) 
                  
                  
                 

                   err = stub.PutState(response.Invno+response.ManDate, resAsBytes)
          
	
                          
	return resAsBytes,nil
                             
}

