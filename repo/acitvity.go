package repo

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Activity struct {
	ActivityId int       `orm:"auto" json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

const activitiesTableName string = "activities"

func (u *Activity) TableName() string {
	return activitiesTableName
}

func CreateActivity(o orm.Ormer, activity *Activity) error {
	_, err := o.Insert(activity)
	return err
}

func GetActivities(o orm.Ormer) ([]Activity, error) {
	activities := []Activity{}
	qs := o.QueryTable(activitiesTableName)

	_, err := qs.All(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func GetActivity(o orm.Ormer, id string) (*Activity, error) {
	activity := Activity{}

	err := o.QueryTable(activitiesTableName).
		Filter("activity_id", id).
		One(&activity)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func UpdateActivity(o orm.Ormer, activity *Activity, columns ...string) error {
	_, err := o.Update(activity, columns...)
	return err
}

func DeleteActivity(o orm.Ormer, id string) (isExist bool, err error) {
	idExist, err := o.QueryTable(activitiesTableName).
		Filter("activity_id", id).
		Delete()
	if err != nil {
		return false, err
	}

	return idExist != 0, nil
}
