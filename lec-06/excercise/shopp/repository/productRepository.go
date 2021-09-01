package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	// "errors"
)


func  GetAllProduct() ([]entity.Product,error){

	var products []entity.Product

	result := client.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil

}


func GetProductById(id int64)(entity.Product,error){
	var product entity.Product

	result := client.DB.First(&product,id)

	if result.Error != nil {
		return  entity.Product{}, result.Error
	}
	return product, nil
}

func CreateProduct(product entity.Product)(entity.Product,error){
	result := client.DB.Omit("OrderDetails","CartDetails").Create(&product)

	if result.Error!=nil{
		return entity.Product{},result.Error
	}

	return product,nil
}


func UpdateProduct(product entity.Product)(entity.Product,error){
	result := client.DB.Save(&product)

	if result.Error!=nil{
		return entity.Product{},result.Error
	}

	return product,nil
}

func DeleteProduct(id int64) error{
	result := client.DB.Delete(entity.Product{},id)

	return result.Error
}