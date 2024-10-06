package workspace

import (
	"errors"
	"fmt"
	"github.com/Ryeom/cosmos/log"
	"github.com/Ryeom/cosmos/psql"
	"gorm.io/gorm/clause"
	"time"

	"reflect"
)

type Workspace struct {
	Sno              uint   `json:"sno" gorm:"-:all;primary_key"`
	Name             string `json:"name" gorm:"column:name"`
	BusinessCategory string `json:"business_category" gorm:"column:business_category"` // 업종
	Address          string `json:"address" gorm:"column:address"`                     // 주소
	Grade            string `json:"grade" gorm:"column:grade"`                         // 등급
	OwnerId          string `json:"ownerId" gorm:"column:owner_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewWorkspace() Workspace {
	w := Workspace{}
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

func (w Workspace) Insert() error {
	db := psql.Postgresql.GetDB().
		Table(`cosmos.t_workspace`).
		Clauses(
			clause.Returning{
				Columns: []clause.Column{
					{Name: "sno"},
					{Name: "name"},
				},
			},
		).Create(&w)
	if db.Error != nil {
		log.Logger.Error(db.Error.Error())
		return db.Error
	}
	return nil
}
func (w Workspace) Update() error {
	db := psql.Postgresql.GetDB().
		Table(`cosmos.t_workspace`).
		Clauses(
			clause.Returning{
				Columns: []clause.Column{
					{Name: "sno"},
					{Name: "name"},
				},
			},
		).
		Where("sno = ?", w.Sno).
		UpdateColumns(&w)

	if db.Error != nil {
		log.Logger.Error(db.Error.Error())

		return db.Error
	}
	return nil
}

func (w Workspace) Delete() error {
	return errors.New("")
}

func Select() Workspace {
	w := Workspace{}

	return w
}

func MaxIndex() int64 {
	return 1
}
