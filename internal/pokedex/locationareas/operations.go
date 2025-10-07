package locationareas

import "fmt"

func (l LocationAreaList) PrintLocationAreaResultsName() {
	for _, v := range l.Results {
		fmt.Println(v.Name)
	}
}
