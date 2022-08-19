package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}


type Vlog struct{
	UserID string `json:"userid"`
	Logs []Log `json:"Logs"`
}

type Log struct{
	Time string  `json:"time"`
	GPS  string  `json:"gps"`
	Data []byte  `json:"data"`
	State string `json:"state"`
}

//Decentralcompany
func (s *SmartContract) AddUser(ctx contractapi.TransactionContextInterface, UserID string) error {
    //GetState
	var vlog = Vlog{UserID: UserID}

	vlogAsBytes, _ := json.Marshal(vlog)	

	return ctx.GetStub().PutState(UserID, vlogAsBytes)
}

func (s *SmartContract) QueryVlogUser(ctx contractapi.TransactionContextInterface, username string) (string, error) {

	UserAsBytes, err := ctx.GetStub().GetState(username)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if UserAsBytes == nil {
		return "", fmt.Errorf("%s does not exist", username)
	}

	// user := new(UserRating)
	// _ = json.Unmarshal(UserAsBytes, &user)
	
	return string(UserAsBytes[:]), nil	
}

//AddRating 을 킥보드 폐지로

// 사용시간 지속적으로 정보 등록
func (s *SmartContract) EnrollData(ctx contractapi.TransactionContextInterface, username string, time string, gps string, data string) error {
	
	// getState User 
	userAsBytes, err := ctx.GetStub().GetState(username)	

	if err != nil{
		return err
	} else if userAsBytes == nil{ // no State! error
		return fmt.Errorf("\"Error\":\"User does not exist: "+ username+"\"")
	}
	// state ok
	vlog := Vlog{}
	err = json.Unmarshal(userAsBytes, &vlog)
	if err != nil {
		return err
	}
	// create rate structure

	var log = Log{Time: time, GPS: gps, Data: 바이트변환(data)}

		// start 
	var log = Log{Time: time, GPS: gps, State: "1"}

		// finish 
	var log = Log{Time: time, GPS: gps, State: "2"}

	vlog.Logs=append(vlog.Logs,log)

	// update to User World state
	userAsBytes, err = json.Marshal(vlog);
	if err != nil {
		return fmt.Errorf("failed to Marshaling: %v", err)
	}	

	err = ctx.GetStub().PutState(username, userAsBytes)
	if err != nil {
		return fmt.Errorf("failed to AddRating: %v", err)
	}	
	return nil
}

func (s *SmartContract) ReviseUser(ctx contractapi.TransactionContextInterface, UserID string) error {
    //DelState

	return ctx.GetStub().DelState(UserID)
}



func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create teamate chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting teamate chaincode: %s", err.Error())
	}
}
