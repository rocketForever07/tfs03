
package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	// "errors"
)



func GetOrderDetailById(id int64)(entity.OrderDetail,error){
	var orderDetail entity.OrderDetail

	result := client.DB.First(&orderDetail,id)

	if result.Error != nil {
		return  entity.OrderDetail{}, result.Error
	}

	return orderDetail, nil
}

func CreateOrderDetail(orderDetail entity.OrderDetail)(entity.OrderDetail,error){
	result := client.DB.Create(&orderDetail)

	if result.Error!=nil{
		return entity.OrderDetail{},result.Error
	}

	return orderDetail,nil
}


func UpdateOrderDetail(orderDetail entity.OrderDetail)(entity.OrderDetail,error){
	result := client.DB.Save(&orderDetail)

	if result.Error!=nil{
		return entity.OrderDetail{},result.Error
	}

	return orderDetail,nil
}

func DeleteOrderDetail(id int64) error{
	result := client.DB.Delete(entity.OrderDetail{},id)

	return result.Error
}