package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const userAddressColumnList = `
id, user_id, province_id, city_id, area_id, province_str, city_str, area_str, address, is_default, 
is_del, create_time, update_time
`

func AddUserAddress(address *model.WechatMallUserAddressDO) error {
	sql := "INSERT INTO wechat_mall_user_address ( " + userAddressColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(address.UserId, address.ProvinceId, address.CityId, address.AreaId, address.ProvinceStr,
		address.CityStr, address.AreaStr, address.Address, address.IsDefault, 0, time.Now(), time.Now())
	return err
}

func ListUserAddress(userId, page, size int) (*[]model.WechatMallUserAddressDO, error) {
	sql := "SELECT " + userAddressColumnList + " FROM wechat_mall_user_address WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	addressList := []model.WechatMallUserAddressDO{}
	for rows.Next() {
		address := model.WechatMallUserAddressDO{}
		err := rows.Scan(&address.Id, &address.UserId, &address.ProvinceId, &address.CityId, &address.AreaId,
			&address.ProvinceStr, &address.CityStr, &address.AreaStr, &address.Address, &address.IsDefault,
			&address.Del, &address.CreateTime, &address.UpdateTime)
		if err != nil {
			return nil, err
		}
		addressList = append(addressList, address)
	}
	return &addressList, nil
}

func QueryUserAddressById(id int) (*model.WechatMallUserAddressDO, error) {
	sql := "SELECT " + userAddressColumnList + " FROM wechat_mall_user_address WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	address := model.WechatMallUserAddressDO{}
	if rows.Next() {
		err := rows.Scan(&address.Id, &address.UserId, &address.ProvinceId, &address.CityId, &address.AreaId,
			&address.ProvinceStr, &address.CityStr, &address.AreaStr, &address.Address, &address.IsDefault,
			&address.Del, &address.CreateTime, &address.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &address, nil
}

func UpdateUserAddress(address *model.WechatMallUserAddressDO) error {
	sql := `
UPDATE wechat_mall_user_address
SET user_id = ?, province_id = ?, city_id = ?, area_id = ?, province_str = ?, city_str = ?, area_str = ?, 
address = ?, is_default = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(address.UserId, address.ProvinceId, address.CityId, address.AreaId, address.ProvinceStr,
		address.CityStr, address.AreaStr, address.Address, address.IsDefault, 0, time.Now(), time.Now())
	return err
}
