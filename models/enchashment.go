package models

type Enchashment struct {
	Id  		int 	`xorm:"id"`
	UID 		string 	`xorm:"UID"`
	Amount 		int64 	`xorm:"amount"`
	CreateTime  int64  	`xorm:"createTime"`
	UpdateTime  int64  	`xorm:"updateTime"`
	Operator 	string  `xorm:"operator"`
	State		int  	`xorm:"state`
	Memo		string 	`xorm:"memo"`
}

func AddEnchashment(enchashment *Enchashment) error{
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.InsertOne(enchashment)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}