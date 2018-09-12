import (
	"fmt"
	"io/ioutil"
)

func ReadFromFile(filepath string) string {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in readFromFile(), Error Info: ", err1)
		}
	}()

	strData := ""

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File read error: ", err)
		return ""
	}

	strData = string(data)
	return strData
}
