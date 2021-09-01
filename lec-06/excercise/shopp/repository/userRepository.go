package repository

import(
	entity "shopp/models/entity"
	client "shopp/database"
	// "errors"
)


func GetUserById(id int)(entity.User,error){
	var user entity.User

	result := client.DB.First(&user,id)

	if result.Error != nil {
		return  entity.User{}, result.Error
	}
	return user, nil
}

func CreateUser(user entity.User)(entity.User,error){
	result := client.DB.Omit("Products").Create(&user)

	if result.Error!=nil{
		return entity.User{},result.Error
	}

	return user,nil
}

func UpdateUser(user entity.User)(entity.User,error){
	result := client.DB.Save(&user)

	if result.Error!=nil{
		return entity.User{},result.Error
	}

	return user,nil
}

func DeleteUser(id int) error{
	result := client.DB.Delete(entity.User{},id)

	return result.Error
}

