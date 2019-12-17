package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"nama"`
	Address string `json:"alamat"`
	Age     string `json:"umur"`
}

type UserData struct {
	Success bool         `json:"success"`
	User    []UserDetail `json:"user"`
}

type UserDetail struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	// biodata := []User{{"Muhammad", "jakarta", "17"}, {"Rijalul", "jakarta", "27"}, {"Haq", "jakarta", "28"}}
	// var jsonData, err = json.Marshal(biodata)

	// if err != nil {
	// 	fmt.Println("Error nih ", err.Error)
	// 	return
	// }

	outputs := `{"success":true,"data":[{"id":"1","name":"user 1"},{"id":"2","name":"user 2"},{"id":"3","name":"user 3"}]}`
	var dataDetail = []byte(outputs)

	var data UserData
	//var detail UserDetail

	var err = json.Unmarshal(dataDetail, &data)

	if err != nil {
		fmt.Println("Erronr Unmarshalling json " + err.Error())
		return
	}

	// var errlagi = json.Unmarshal([]byte(data.User), &detail)
	// if errlagi != nil {
	// 	fmt.Println("Erronr Unmarshalling json " + errlagi.Error())
	// 	return
	// }

	if data.Success == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	for _, det := range data.User {
		fmt.Println("id_", det.Id, "#name_", det.Name)
	}

	// var x, y int
	// fmt.Scan(&x)
	// fmt.Scan(&y)
	// //Untuk memasukkan input pada command line
	// scanner := bufio.NewScanner(os.Stdin)
	// for {

	// 	fmt.Print("Enter Text: ")
	// 	scanner.Scan()

	// 	text := scanner.Text()

	// 	if len(text) == 0 {
	// 		break
	// 	} else {
	// 		operator := scanner.Text()
	// 		resultOperation := Operation(x, y, operator)
	// 		fmt.Println("Hasil "+operator+" :", resultOperation)
	// 	}
	// }

	// a := []string{"foo1", "Bar1", "foo2", "bar2"}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println("perulangan for biasa : ", a[i])
	// }

	// for _, avalue := range a {
	// 	fmt.Println("perulangan dengan for range : ", avalue)
	// }

	//operator := "jumlah"
	// resultAdd := Add(x, y)
	// resultMin := Min(x, y)
	// resultMultiple := Multiple(x, y)
	// resultDivision := Division(x, y)

	// fmt.Println("Hasil Jumlah : ", resultAdd)
	// fmt.Println("Hasil Kurang : ", resultMin)
	// fmt.Println("Hasil Kali   : ", resultMultiple)
	// fmt.Println("Hasil Bagi   : ", resultDivision)

	// resultOperation := Operation(x, y, operator)
	// fmt.Println("Hasil "+operator+" :", resultOperation)

}

func Add(x int, y int) int {
	result := x + y
	return result
}

func Min(x int, y int) int {
	result := x - y
	return result
}

func Multiple(x int, y int) int {
	result := x * y
	return result
}

func Division(x int, y int) float32 {

	result := float32(x) / float32(y)
	return result
}

func Operation(x int, y int, operator string) float32 {
	var result float32
	switch operator {
	case "jumlah":
		result = float32(Add(x, y))
	case "kurang":
		result = float32(Min(x, y))
	case "kali":
		result = float32(Multiple(x, y))
	case "bagi":
		result = float32(Division(x, y))
	}

	return result
}
