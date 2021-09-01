
package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	// "errors"
)

func GetAllPayment()([]entity.Payment,error){
	var payments []entity.Payment

	result := client.DB.First(&payments)

	if result.Error != nil {
		return  []entity.Payment{}, result.Error
	}

	return payments, nil
}


func GetPaymentById(id int64)(entity.Payment,error){
	var payment entity.Payment

	result := client.DB.First(&payment,id)

	if result.Error != nil {
		return  entity.Payment{}, result.Error
	}

	return payment, nil
}

func CreatePayment(payment entity.Payment)(entity.Payment,error){
	result := client.DB.Create(&payment)

	if result.Error!=nil{
		return entity.Payment{},result.Error
	}

	return payment,nil
}


func UpdatePayment(payment entity.Payment)(entity.Payment,error){
	result := client.DB.Save(&payment)

	if result.Error!=nil{
		return entity.Payment{},result.Error
	}

	return payment,nil
}
