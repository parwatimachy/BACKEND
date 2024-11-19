package controller

import (
	"context"
	"fmt"
	"golang-database-user/model"
	"golang-database-user/service"
)

func DefaultChoose() {
	fmt.Println("Incorrect Number")
}

func CreateUser(userService service.UserService) {

	ctx := context.Background()

	var name, email, password, phoneNumber string

	fmt.Print("Masukkan Name: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&email)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	fmt.Print("Masukkan Phone Number: ")
	fmt.Scanln(&phoneNumber)

	user := model.MstUser{
		Name:        string (name),
		Email:       string (email),
		Password:    string (password),
		PhoneNumber: string (phoneNumber),
	}

	mstUser := userService.CreateUser(ctx, user)


	fmt.Println("SUCCESS: ")
	fmt.Println(mstUser)
}
func UpdateUser(userService service.UserService) {
	
	ctx := context.Background()

	var userId,name,email,password,phoneNumber string
	fmt.Print("Masukkan id user yang ingin di update: ")
	fmt.Scanln(&userId)

	fmt.Print("Masukkan Name: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&email)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	fmt.Print("Masukkan Phone Number: ")
	fmt.Scanln(&phoneNumber)

	user := model.MstUser{
		Name:        string (name),
		Email:       string (email),
		Password:    string (password),
		PhoneNumber: string (phoneNumber),
	}
	mstUser := userService.UpdateUser(ctx, user, userId)

	fmt.Println("SUCCESS: ")
	fmt.Println(mstUser)
}

func ReadUser(userService service.UserService) {
	ctx := context.Background()

	users, err := userService.ReadUsers(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nAll Users:")
	for _, mstUser := range users {
		fmt.Println("Id : ", mstUser.IdUser, "\nNama : ", mstUser.Name, "\nEmail : ", mstUser.Email, "\nNomor HP : ", mstUser.PhoneNumber)
		fmt.Println()
	}
}

func DeleteUser(userService service.UserService) {
	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan id user yang ingin di hapus: ")
	fmt.Scanln(&userId)

	err := userService.DeleteUser(ctx, userId)
	if err != nil {
		fmt.Println("Gagal menghapus user:", err)
	} else {
		fmt.Println("User berhasil dihapus")
	}
}

