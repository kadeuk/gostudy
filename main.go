package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("data.csv")
	if err != nil {
		panic(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	var resultLines []string

	for i, line := range lines {
			if strings.Contains(line, "[이미지파일]") {
				continue
			}
		
			newLine := fmt.Sprintf("%d번째 줄: %s", i+1, line)

			resultLines = append(resultLines, newLine)
	}
		finalResult := strings.Join(resultLines, "\n")

		os.WriteFile("result.csv", []byte(finalResult), 0644)

		fmt.Println("처리완료 rusult.csv를 확인하세요")
	
}

