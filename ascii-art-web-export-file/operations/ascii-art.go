package operations

import (
	"os"
	"strings"
)

func ReadTextFile(file string) ([]string, error) {
	//Use os.readfile to read our banner file
	data, err := os.ReadFile(file)

	//Return an empty slice if there is an error
	if err != nil {
		return []string{}, err
	}

	cleanStr := strings.ReplaceAll(string(data), "\r", "")

	//Split data by new line
	dataSlice := strings.Split(cleanStr, "\n")

	return dataSlice, nil
}

func AsciiArt(inputSlice []string, dataSlice []string) (string, error) {

	var builder strings.Builder

	//Loop through input array
	for _, input := range inputSlice {

		if input == "" {
			builder.WriteString("\n")
			continue
		}

		//Loop 8 times to print all rowa of input
		for row := 1; row <= 8; row++ {

			//Now for each loop above print that row for each chs of input eg
			//row = 1 print row 1 of all chs of input
			for _, ch := range input {

				//Make sure ch is printable
				if ch >= 32 && ch <= 126 {

					//Calculate formular to get each rows
					print := ((int(ch) - 32) * 9) + row

					//Store string in buffer
					builder.WriteString(dataSlice[print])
				}

			}
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil

}
