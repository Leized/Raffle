package dao

import (
	"raffle/model"
	"raffle/utils"
)

//CheckAward 查询奖品
func CheckAward(award string) (*model.Data, error) {
	sqlStr := "select * from datas where award = ? "
	row := utils.Db.QueryRow(sqlStr, award)
	data := &model.Data{}
	row.Scan(&data.ID, &data.Member, &data.Award)
	return data, nil
}

//CheckIdAndAward 查询id和奖品
func CheckIdAndAward(id int, award string) (*model.Data, error) {
	sqlStr := "select * from datas where id = ? and award = ? "
	row := utils.Db.QueryRow(sqlStr, id, award)
	data := &model.Data{}
	row.Scan(&data.ID, &data.Member, &data.Award)
	return data, nil
}

//CheckId 查询id
func CheckId(id int) (*model.Data, error) {
	sqlStr := "select * from datas where id = ? "
	row := utils.Db.QueryRow(sqlStr, id)
	data := &model.Data{}
	row.Scan(&data.ID, &data.Member, &data.Award)
	return data, nil
}

//SaveData 向数据库插入奖品信息
func SaveData(id int, award string) error {
	sqlStr := "update datas set award = ? where id = ?"
	_, err := utils.Db.Exec(sqlStr, award, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteData 清除奖品数据
func DeleteData(id int) (*model.Data, error) {
	sqlStr := "update datas set award='' where  id = ?"
	row := utils.Db.QueryRow(sqlStr, id)
	data := &model.Data{}
	row.Scan(&data.ID, &data.Member, &data.Award)
	return data, nil
}
