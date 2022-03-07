package main

import (
	"codeup.aliyun.com/go-online/public/permission-service/lib"
	"codeup.aliyun.com/go-online/public/permission-service/pb/authorization"
	"context"
	"fmt"
	gormadapter "github.com/dengh57/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

const tableRuleName = "policies"

type authService struct {
	authorization.UnimplementedAuthorizationServer
}

func (*authService) AddModel(ctx context.Context, req *authorization.AddModelReq) (*authorization.AddModelRes, error) {
	DB := lib.GetDB()

	// 向model表中需要添加的数据
	modelData := &lib.Model{
		RequestDefinition: req.RequestDefinition,
		PolicyDefinition:  req.PolicyDefinition,
		RoleDefinition:    req.RoleDefinition,
		PolicyEffect:      req.PolicyEffect,
		Matchers:          req.Matchers}
	// 添加记录
	result := DB.Create(modelData)
	//DB.Table(tableRuleName).AutoMigrate(&lib.Rule{})

	return &authorization.AddModelRes{ModelId: modelData.ID}, result.Error
}

func (*authService) AddPolicy(ctx context.Context, req *authorization.AddPolicyReq) (*authorization.AddPolicyRes, error) {

	DB := lib.GetDB()
	res := make([]lib.Rule, len(req.Policies))
	for i, v := range req.Policies {
		//v.Data = append(v.Data, strconv.FormatInt(req.ModelId, 10))
		for len(v.Data) < 11 {
			v.Data = append(v.Data, "")
		}
		res[i] = lib.Rule{
			Ptype:   v.Data[0],
			V0:      v.Data[1],
			V1:      v.Data[2],
			V2:      v.Data[3],
			V3:      v.Data[4],
			V4:      v.Data[5],
			V5:      v.Data[6],
			V6:      v.Data[7],
			V7:      v.Data[8],
			V8:      v.Data[9],
			V9:      v.Data[10],
			ModelID: strconv.FormatInt(req.ModelId, 10),
		}
	}
	DB.Table(tableRuleName).Create(&res)
	ans := make([]int64, len(req.Policies))
	for i, v := range res {
		ans[i] = v.ID
	}
	return &authorization.AddPolicyRes{PolicyId: ans}, nil
}

func (*authService) DeletePolicy(ctx context.Context, req *authorization.DeletePolicyReq) (*authorization.Empty, error) {
	DB := lib.GetDB()
	DB.Unscoped().Table(tableRuleName).Delete(&lib.Rule{}, req.PolicyId)

	return &authorization.Empty{}, nil
}

func (*authService) Authorization(ctx context.Context, req *authorization.AuthReq) (*authorization.AuthRes, error) {
	DB := lib.GetDB()
	e, _ := lib.GetEnforce(req.ModelId, DB)
	id := strconv.FormatInt(req.ModelId, 10)
	fmt.Println(id)
	m := req.Filter
	filter := gormadapter.Filter{}
	if _, ok := m["Ptype"]; ok {
		filter.Ptype = m["Ptype"].Data
	}
	if _, ok := m["V0"]; ok {
		filter.V0 = m["V0"].Data
	}
	if _, ok := m["V1"]; ok {
		filter.V1 = m["V1"].Data
	}
	for k, v := range m {
		switch k {
		case "Ptype":
			filter.Ptype = v.Data
		case "V0":
			filter.V0 = v.Data
		case "V1":
			filter.V1 = v.Data
		case "V2":
			filter.V2 = v.Data
		case "V3":
			filter.V3 = v.Data
		case "V4":
			filter.V4 = v.Data
		case "V5":
			filter.V5 = v.Data
		case "V6":
			filter.V6 = v.Data
		case "V7":
			filter.V7 = v.Data
		case "V8":
			filter.V8 = v.Data
		case "V9":
			filter.V9 = v.Data
		}
	}
	e.LoadFilteredPolicy(filter)
	//e.LoadPolicy()
	l := len(req.Request)
	fmt.Println(l)
	request := make([][]string, l)
	ans := make([]bool, l)
	for i, temp := range req.Request {
		request[i] = temp.Data
		res := make([]interface{}, len(request[i]))
		for i, v := range request[i] {
			res[i] = v
		}
		fmt.Println(res)
		ans[i], _ = e.Enforce(res...)
	}

	return &authorization.AuthRes{
		Result: ans,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	s := grpc.NewServer()
	authorization.RegisterAuthorizationServer(s, &authService{})
	s.Serve(l)
}
