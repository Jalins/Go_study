package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"log"
)

type CC struct {
}

type User struct {
	Name   string            `json:"name"`
	Id     string            `json:"id"`
	Assets map[string]string `json:"asserts"`
}

type Asset struct {
	Name     string            `json:"name"`
	Id       string            `json:"id"`
	MetaData map[string]string `json:"meta_data"`
}

type AssetChangeHistory struct {
	AssetId      string `json:"asset_id"`
	OriginOwner  string `json:"origin_owner"`
	CurrentOwner string `json:"current_owner"`
}

func constructUserKey(userId string) string {
	return fmt.Sprintf("user_%s", userId)
}

func constructAssetKey(assetId string) string {
	return fmt.Sprintf("asset_%s", assetId)
}

func constructAssetHistoryKey(originUserId, assetId, currentUserId string) string {
	return fmt.Sprintf("history_%s_%s_%s", originUserId, assetId, currentUserId)
}

func userRegister(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 检查参数个数
	if len(args) != 2 {
		return shim.Error("not enough args!")
	}
	// 验证参数的正确性
	name := args[0]
	id := args[1]
	if name == "" || id == "" {
		return shim.Error("invalid args!")
	}

	// 验证数据是否已经存在
	userBytes, err := stub.GetState(constructUserKey(id))
	if err == nil && len(userBytes) != 0 {
		return shim.Error("user already exist!")
	}

	// 写入状态
	user := &User{
		Name:   name,
		Id:     id,
		Assets: make(map[string]string),
	}

	if userBytes, err = json.Marshal(user); err != nil {
		return shim.Error(fmt.Sprintf("marshal user error %s", err))
	}

	if err = stub.PutState(constructUserKey(id), userBytes); err != nil {
		return shim.Error(fmt.Sprintf("put user error %s", err))
	}

	// 成功返回
	return shim.Success(nil)
}

func userDestroy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 检查参数个数
	if len(args) != 1 {
		return shim.Error("not enough args!")
	}

	id := args[0]
	if id == "" {
		return shim.Error("invalid args!")
	}

	userBytes, err := stub.GetState(constructUserKey(id))
	if err != nil || len(userBytes) == 0 {
		return shim.Error("user not found!")
	}
	return shim.Success(nil)
}

func assetEnroll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("not enough args!")
	}

	name := args[0]
	id := args[1]
	metadata := args[2]
	owner := args[3]
	if name == "" || id == "" || owner == "" {
		return shim.Error("invalid args!")
	}

	userBytes, err := stub.GetState(constructAssetKey(id))
	if err == nil && len(userBytes) != 0 {
		return shim.Error("user already exist!")
	}

	assetBytes, err := stub.GetState(constructUserKey(owner))
	if err == nil && len(assetBytes) != 0 {
		return shim.Error("asset already exist!")
	}
	return shim.Success(nil)
}

func assetExchange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("not enough args!")
	}

	originOwner := args[0]
	id := args[1]
	currentOwner := args[2]
	if originOwner == "" || id == "" || currentOwner == "" {
		return shim.Error("invalid args!")
	}

	originOwnerBytes, err := stub.GetState(constructUserKey(originOwner))
	if err != nil || len(originOwnerBytes) == 0 {
		return shim.Error("originOwner not found!")
	}

	assetBytes, err := stub.GetState(constructAssetKey(id))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found!")
	}

	currentOwnerBytes, err := stub.GetState(constructUserKey(currentOwner))
	if err == nil || len(currentOwnerBytes) != 0 {
		return shim.Error("currentUser not found!")
	}
	return shim.Success(nil)
}

func userQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("not enough args!")
	}

	userId := args[0]
	if userId == "" {
		return shim.Error("invalid args!")
	}

	userBytes, err := stub.GetState(constructUserKey(userId))
	if err != nil || len(userBytes) == 0 {
		return shim.Error("user not found!")
	}
	return shim.Success(userBytes)
}

func assetQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("not enough args!")
	}

	assetId := args[0]
	if assetId == "" {
		return shim.Error("invalid args!")
	}

	assetBytes, err := stub.GetState(constructUserKey(assetId))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found!")
	}
	return shim.Success(assetBytes)
}

func queryAssetExchangeHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 && len(args) != 1 {
		return shim.Error("not enough args!")
	}

	assetId := args[0]
	queryType := args[1]
	if assetId == "" {
		return shim.Error("invalid args!")
	}

	assetBytes, err := stub.GetState(constructUserKey(assetId))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found!")
	}
	return shim.Success(assetBytes)
}

func (c CC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("successful"))
}

func (c CC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funName, args := stub.GetFunctionAndParameters()

	switch funName {
	case "userRegister":
		return userRegister(stub, args)
	case "userDestroy":
		return userDestroy(stub, args)
	case "assetEnroll":
		return assetEnroll(stub, args)
	case "assetExchange":
		return assetExchange(stub, args)
	case "userQuery":
		return userQuery(stub, args)
	case "assetQuery":
		return assetQuery(stub, args)
	case "queryAssetExchangeHistory":
		return queryAssetExchangeHistory(stub, args)

	}
	return shim.Success([]byte("successful"))
}

func main() {
	err := shim.Start(new(CC))
	if err != nil {
		log.Fatal(err)
	}
}
