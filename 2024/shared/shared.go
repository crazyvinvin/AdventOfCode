package shared

import "os"

func ReadInput(path string) string {
	data, err := os.ReadFile(path + "/input.txt")

	if err != nil {
		//No need to handle the error
		panic(err)
	}

	content := string(data)

	return content
}
