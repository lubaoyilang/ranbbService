package models


/*
Tid	int	主键
TaoBaoAccount	String	刷单所用淘宝账号
CreateTime	long	创建时间
UpdateTime	long	更新时间
Memo	String	描述
 */

type TaobaoAccount struct {
	Tid int `xorm:"tid pk int(11) unique autoincr "`
	UID string `xorm:"'UID' index notnull varchar(35)"`
	TaoBaoAccount string  `xorm:"'taobaoAccount' index notnull unique varchar(35)"`
	CreateTime int64 `xorm:"createTime default 0"`
	UpdateTime int64 `xorm:"updateTime default 0"`
	Memo string  `xorm:"memo blob"`
}


func GetTaobaoAccList(UID string) ([]TaobaoAccount,int,error){
	sess := Engine.NewSession()
	defer sess.Close()
	accs := make([]TaobaoAccount,0)
	err := sess.Where("UID = ?",UID).Find(&accs)
	return accs,len(accs),err
}

func UpdateTaobaoAcc(acc *TaobaoAccount) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.Where("tid = ?",acc.Tid).Update(acc)
	if err != nil {
		sess.Rollback()
		return err
	}
	return nil
}


func DeleteTaobaoAcc(tid int) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.Exec(`delete from taobao_account where tid = ?`,tid)
	if err != nil {
		sess.Rollback()
		return err
	}
	return nil
}

func AddTaobaoAcc(acc * TaobaoAccount) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.InsertOne(acc)
	if err != nil {
		sess.Rollback()
		return err
	}
	return nil
}