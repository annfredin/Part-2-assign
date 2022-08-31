package lib

import (
	"io/ioutil"
	"os"
)


func CreateFile(name string) error {
    file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    return file.Close()
}

func WriteFile(fileName, data string) error{
	// Open file using READ & WRITE permission.
    var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
    if err != nil{
        return err
    }
    defer file.Close()

	file.WriteString(data)

	return nil
}

func ReadFile(fileName string) (string, error){
	bytesData, err := ioutil.ReadFile(fileName)
   if err != nil{
	return "", err
   } 

   return string(bytesData), nil
}