
package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	// "errors"
)


func GetCartDetailById(id int64)(entity.CartDetail,error){
	var cartDetail entity.CartDetail

	result := client.DB.First(&cartDetail,id)

	if result.Error != nil {
		return  entity.CartDetail{}, result.Error
	}

	return cartDetail, nil
}

func CreateCartDetail(cartDetail entity.CartDetail)(entity.CartDetail,error){
	result := client.DB.Create(&cartDetail)

	if result.Error!=nil{
		return entity.CartDetail{},result.Error
	}

	return cartDetail,nil
}


func UpdateCartDetail(cartDetail entity.CartDetail)(entity.CartDetail,error){
	result := client.DB.Save(&cartDetail)

	if result.Error!=nil{
		return entity.CartDetail{},result.Error
	}

	return cartDetail,nil
}

func DeleteCartDetail(id int64) error{
	result := client.DB.Delete(entity.CartDetail{},id)

	return result.Error
}