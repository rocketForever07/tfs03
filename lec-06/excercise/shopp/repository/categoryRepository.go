package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	"log"
	// "errors"
)


func  GetAllCategory() ([]entity.Category,error){

	var categories []entity.Category

	result := client.DB.Preload("Products").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil

}


func GetCategoryById(id int64)(entity.Category,error){
	var category entity.Category

	result := client.DB.Preload("Products").First(&category,id)

	if result.Error != nil {
		return  entity.Category{}, result.Error
	}

	log.Println("---------------")
	log.Println(category.Products)
	return category, nil
}

func CreateCategory(category entity.Category)(entity.Category,error){
	result := client.DB.Omit("Products").Create(&category)

	if result.Error!=nil{
		return entity.Category{},result.Error
	}

	return category,nil
}


func UpdateCategory(category entity.Category)(entity.Category,error){
	result := client.DB.Save(&category)

	if result.Error!=nil{
		return entity.Category{},result.Error
	}

	return category,nil
}

func DeleteCategory(id int64) error{
	result := client.DB.Delete(entity.Category{},id)

	return result.Error
}