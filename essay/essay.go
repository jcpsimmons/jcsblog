package essay

import (
	"embed"
	"fmt"
	"sort"
	"strings"

	"github.com/russross/blackfriday/v2"
)

//go:embed *.md
var essayFiles embed.FS

func FileNameToTitle(fileName string) string {
	humanReadableName := strings.Split(fileName, ".")[0]
	humanReadableName = strings.Split(humanReadableName, "_")[1]
	humanReadableName = strings.ReplaceAll(humanReadableName, "-", " ")
	return humanReadableName
}

func GetAllEssaysAsListItems() ([]string, error) {
	files, err := essayFiles.ReadDir(".")
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	var fileNames []string
	for _, file := range files {
		fileName := file.Name()
		humanReadableName := FileNameToTitle(fileName)
		if strings.HasSuffix(fileName, ".md") {
			// wrap in <a> and <li> tags
			htmlFileName := "<li><a href=\"/essay/" + fileName[:len(fileName)-3] + "\">" + humanReadableName + "</a></li>"
			fileNames = append(fileNames, htmlFileName)
		}
	}
	return fileNames, nil
}

func GetEssay(fileName string) ([]byte, error) {
	fmt.Println("fileName:", fileName)

	post, err := essayFiles.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return blackfriday.Run(post), nil
}