package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	// employee1 := employee{
	// 	employeeID:   "101",
	// 	employeeName: "anawat",
	// 	phone:        "086765xxxx",
	// }
	// fmt.Println("employee1 = ", employee1)

	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeID:   "101",
	// 	employeeName: "anawat",
	// 	phone:        "086765xxxx",
	// }
	// employeeList[1] = employee{
	// 	employeeID:   "102",
	// 	employeeName: "boat",
	// 	phone:        "086765xxxx",
	// }
	// employeeList[2] = employee{
	// 	employeeID:   "103",
	// 	employeeName: "mon",
	// 	phone:        "086765xxxx",
	// }
	// fmt.Println("employeeList = ", employeeList)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "anawat",
		phone:        "086765xxxx",
	}

	employee2 := employee{
		employeeID:   "102",
		employeeName: "boat",
		phone:        "086765xxxx",
	}

	employee3 := employee{
		employeeID:   "103",
		employeeName: "mon",
		phone:        "086765xxxx",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)
	fmt.Println("employeeList = ", employeeList)

}
