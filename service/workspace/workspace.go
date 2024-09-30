package workspace

import (
	"fmt"
	"github.com/Ryeom/cosmos/mongo"
	"github.com/Ryeom/cosmos/util"
	//"go.mongodb.org/mongo-driver/v2/bson"
	"reflect"
)

type Workspace struct {
	Index            int64    `json:"index,omitempty" bson:"index"`
	Id               string   `json:"id" bson:"id"`
	Name             string   `json:"name" bson:"name"`
	BusinessCategory string   `json:"businessCategory" bson:"businessCategory"` // 업종
	Address          string   `json:"address" bson:"address"`                   // 주소
	Grade            string   `json:"grade" bson:"grade"`                       // 등급
	InService        []string `json:"inService" bson:"inService"`               // 사용중인 서비스
}

func NewWorkspace() Workspace {
	w := Workspace{
		Id: util.GetUUID(),
	}
	return w
}

func (w Workspace) PrintDataInfo() {
	e := reflect.ValueOf(&w).Elem()
	fieldNum := e.NumField()
	for i := 0; i < fieldNum; i++ {
		v := e.Field(i)
		t := e.Type().Field(i)
		fmt.Printf("[Name: %s] Type: %s | Value: %v\n",
			t.Name, t.Type, v.Interface())

	}
}

func (w Workspace) Save() error {
	var err error
	d := mongo.ToBsonD(w)
	err = mongo.InsertOne("workspace", d)
	return err
}

func MaxIndex() int64 {
	return 1
}
