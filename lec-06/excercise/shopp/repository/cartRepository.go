
package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	"log"
	// "errors"
)


func  GetAllCart() ([]entity.Cart,error){

	var carts []entity.Cart

	result := client.DB.Preload("CartDetails").Find(&carts)
	if result.Error != nil {
		return nil, result.Error
	}
	return carts, nil

}


func GetCartById(id int64)(entity.Cart,error){
	var cart entity.Cart

	result := client.DB.Preload("CartDetails").First(&cart,id)

	if result.Error != nil {
		return  entity.Cart{}, result.Error
	}

	log.Println("---------------")
	log.Println(cart.CartDetails)
	return cart, nil
}

func CreateCart(cart entity.Cart)(entity.Cart,error){
	result := client.DB.Omit("CartDetails").Create(&cart)

	if result.Error!=nil{
		return entity.Cart{},result.Error
	}

	return cart,nil
}

func UpdateCart(cart entity.Cart)(entity.Cart,error){
	result := client.DB.Save(&cart)

	if result.Error!=nil{
		return entity.Cart{},result.Error
	}

	return cart,nil
}

func DeleteCart(id int64) error{
	result := client.DB.Delete(entity.Cart{},id)

	return result.Error
}

func FindCartByUserId(userId int64) ([]entity.Cart,error){


	var carts []entity.Cart
	result := client.DB.Where("name = ?", "jinzhu").Find(&carts)

	if result.Error !=nil{
		return []entity.Cart{},result.Error
	}

	return carts,nil


}