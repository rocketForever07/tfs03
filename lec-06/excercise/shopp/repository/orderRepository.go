
package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	"log"
	// "errors"
)


func  GetAllOrder() ([]entity.Order,error){

	var orders []entity.Order

	result := client.DB.Preload("OrderDetails").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil

}


func GetOrderById(id int64)(entity.Order,error){
	var order entity.Order

	result := client.DB.Preload("OrderDetails").First(&order,id)

	if result.Error != nil {
		return  entity.Order{}, result.Error
	}

	log.Println("---------------")
	log.Println(order.OrderDetails)
	return order, nil
}

func CreateOrder(order entity.Order)(entity.Order,error){
	result := client.DB.Omit("OrderDetails").Create(&order)

	if result.Error!=nil{
		return entity.Order{},result.Error
	}

	return order,nil
}

func UpdateOrder(order entity.Order)(entity.Order,error){
	result := client.DB.Save(&order)

	if result.Error!=nil{
		return entity.Order{},result.Error
	}

	return order,nil
}

func DeleteOrder(id int64) error{
	result := client.DB.Delete(entity.Order{},id)

	return result.Error
}