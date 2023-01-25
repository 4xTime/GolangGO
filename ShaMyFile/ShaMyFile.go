package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func CreateFileForHashFileContent(PATH string, Dest_path string, Content string) {
	File, error := os.Create(Dest_path)
	if error != nil {
		fmt.Println(error)
	}

	defer File.Close()

	_, err2 := File.WriteString(Content)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Print("Done!")
}

func HashFileContent(PATH string, Dest_path string) {
	SHA := sha256.New()
	readFile, err := os.Open(PATH)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	TXT := ""
	for fileScanner.Scan() {
		SHA.Write([]byte(fileScanner.Text()))
		TXT += hex.EncodeToString(SHA.Sum(nil)) + "\n"
		SHA.Reset()
	}
	readFile.Close()
	CreateFileForHashFileContent(PATH, Dest_path, TXT)
}

func main() {
	//0 = Path to raw data,1 = Path to extract hash
	arg := os.Args[1:]
	HashFileContent(arg[0], arg[1])
}
