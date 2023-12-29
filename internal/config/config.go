package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

func LoadData() (interface {}, error){
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf("load env data is not working properly %v", err)
	}

	dataFileName := os.Getenv("DATA_FILE_NAME")
	dataFolderName := os.Getenv("DATA_FOLDER")

	dataFilePath := "/home/adithya/thesis/karma/karma-backend/internal/data/data.json"

	file, errorOpeningFile := os.OpenFile(dataFilePath, os.O_RDWR, 0644)
	
	if errorOpeningFile != nil {
		return nil,fmt.Errorf(
			"error opening data.json file %s %s %v",
			dataFileName,
			dataFolderName,
			errorOpeningFile,
		)
	}

	defer file.Close()

	content, errReadingFile := io.ReadAll(file)
	if errReadingFile != nil {
		// If an error occurs reading the file, return an error message
		return nil, fmt.Errorf("error reading JSON file content: %v", errReadingFile)
	}

	var jsonData interface{}

	errUnmarshalingJson := json.Unmarshal(content, &jsonData)

	if errUnmarshalingJson != nil {
		return nil,fmt.Errorf("error unmarshaling json %v", errUnmarshalingJson)
	}

	return jsonData,nil
}