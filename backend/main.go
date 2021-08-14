package main

import (
	"fmt"
)

func main() {
	a := "string"
	fmt.Println(a)
}

// // --------
// // model↓
// // --------

// // User ...
// type User struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// // --------
// // router↓
// // --------

// // InitRouting ...
// func InitRouting(e *echo.Echo, u *User) {
// 	e.POST("user", u.CreateUser)
// 	e.GET("user/:id", u.GetUser)
// }

// // --------
// // handler↓
// // --------

// // CreateUser ...
// func (u *User) CreateUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	user = User{
// 		ID:   1,
// 		Name: user.Name,
// 		Age:  user.Age,
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// // GetUser ...
// func (u *User) GetUser(c echo.Context) error {
// 	user := User{}

// 	if err := c.Bind(&user); err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	// Getメソッドのイメージ
// 	if id == 1 {
// 		user = User{
// 			ID:   1,
// 			Name: "Tom",
// 			Age:  29,
// 		}
// 	} else if id == 2 {
// 		user = User{
// 			ID:   2,
// 			Name: "Bob",
// 			Age:  35,
// 		}

// 	} else {
// 		return c.JSON(http.StatusOK, "Not Found")
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// // --------
// // main.go↓
// // --------

// func main() {
// 	e := echo.New()

// 	u := new(User)
// 	InitRouting(e, u)

// 	e.Start(":9115")
// }
