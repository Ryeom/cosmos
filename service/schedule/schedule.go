package schedule

import (
	"github.com/Ryeom/cosmos/mongo"
	"time"
)

type Schedule struct {
	Id          string   `json:"id" bson:"id"`
	Position    string   `json:"position" bson:"position"`       /* 일정 포지션 */
	Description string   `json:"description" bson:"description"` /* 상세메모 */
	Location    string   `json:"location" bson:"location"`       /* hall or 주방 ... */
	Attendees   []string `json:"attendees" bson:"attendees"`     /* 참여자 목록 */
	Reminder    bool     `json:"reminder" bson:"reminder"`       /* 알람여부 */
	Status      string   `json:"status" bson:"status"`           /* 일정 상태 : 예정, 진행중, 완료*/
	Recurring   string   `json:"recurring" bson:"recurring"`     /* 반복 주기 (매주월요일, 매달1일) */

	StartDateTime time.Time `json:"startDateTime" bson:"startDateTime"` /* 시작시간 */
	EndDateTime   time.Time `json:"endDateTime" bson:"endDateTime"`     /* 끝나는 시간 */
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
}

func (s Schedule) Insert() error {
	var err error
	d := mongo.ToBsonD(s)
	err = mongo.InsertOne("schedule", d)
	return err
}
