package main

import (
	"encoding/json"
	"fmt"
	"log"

	"errors"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"

	"os"
	"strconv"
	"strings"

	terminal "golang.org/x/term"
)

type RoadmapJsonData struct {
	Categories []struct {
		Category    string `json:"CATEGORY"`
		Description string `json:"DESCRIPTION"`
		Topics      []struct {
			Topic       string              `json:"TOPIC"`
			Description string              `json:"DESCRIPTION"`
			Subtopics   []map[string]string `json:"SUBTOPICS"`
		} `json:"TOPICS"`
	} `json:"CATEGORIES"`
}

// File paths
var roadmapJsonFile string = "/usr/local/etc/roadmap-sec-engineer/roadmap.json"
var learningStateFile string = "/usr/local/etc/roadmap-sec-engineer/learning-state.state"

func main() {

	if len(os.Args) > 1 {

		switch os.Args[1] {

		case "--done":

			done()

			os.Exit(0)

		case "--custom":

			if len(os.Args) < 3 {
				log.Fatal("Error: You need to parse a number to change the progress!")

				os.Exit(1)
			}

			num, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Panic("Error: Invalid integer!")
				os.Exit(1)
			}

			custom(num)

		case "--help":
			help()

		case "-h":
			help()

		}
	}

	show_default()

	// // Access the data
	// fmt.Println(roadmapData.Categories[1].Category)

}

// Shows the help menu (Documentation)
func help() {

	fmt.Println("\033[0;38;5;183;1;4mROADMAP-SECURITY-ENGINEER")
	fmt.Println()
	fmt.Println("\033[0;38;5;159;1mNAME\033[0;38;5;231m")
	fmt.Println("\t roadmap-security-engineer [OPTION]")
	fmt.Println()
	fmt.Println("\033[0;38;5;159;1mDESCRIPTION\033[0;38;5;231m")
	fmt.Println("\t Shows a roadmap to be a successful Cyber Security Engineer. By learning one topic per day.")
	fmt.Println()
	fmt.Println("\t \033[0;38;5;159;1m-h, --help\033[0;38;5;231m")
	fmt.Println("\t\t Show this help/documentation page.")
	fmt.Println()
	fmt.Println("\t \033[0;38;5;159;1m --done\033[0;38;5;231m")
	fmt.Println("\t\t Increment the progress. (Assume you have learnt current topic.)")
	fmt.Println()
	fmt.Println("\t \033[0;38;5;159;1m --custom\033[0;38;5;231m <level>")
	fmt.Println("\t\t Change the progress to desired level.")
	fmt.Println()
	fmt.Println("\033[0;38;5;159;1mAUTHOR\033[0;38;5;231m")
	fmt.Println("\t Written by A. Alham Azhar")
	fmt.Println()

	os.Exit(0)

}

// Show from saved progress
func show_default() {

	// check if the 'roadmap.json' file exist
	if !fileExists(roadmapJsonFile) {
		log.Fatal("Error: 'roadmap.json' file does not exist!")
	}

	// retrieve data from the 'roadmap.json' file.
	roadmapData, err := loadJsonData(roadmapJsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	progress, err := getLearningProgress()
	if err != nil {
		log.Println(err)
		log.Println("Showing the 1st day of learning instructions")

		custom(0)
	}

	// const (
	// 	defaultMargin     = 4
	// 	defaultListIndent = 2
	// )

	width := terminalWidth()

	customStylingConfig := ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				// BlockPrefix: "\n",
				// BlockSuffix: "\n",
				Color: stringPtr("231"),
			},
			Margin: uintPtr(0),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr("229"),
				Italic: boolPtr(true),
				Faint:  boolPtr(true),
			},
			Indent: uintPtr(2),
			// IndentToken: stringPtr("│"),
		},
		Paragraph: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr("229"),
				Italic: boolPtr(true),
			},
			Indent: uintPtr(2),
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				// BlockSuffix: "\n",
				Color: stringPtr("183"),
				Bold:  boolPtr(true),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "🎯 ",
				// Suffix: " 🞘🞘",
				// Color:  stringPtr("38"),
			},
			Indent: uintPtr(1),
		},
		Strikethrough: ansi.StylePrimitive{
			CrossedOut: boolPtr(true),
		},
		Emph: ansi.StylePrimitive{
			Color:  stringPtr("222"),
			Italic: boolPtr(true),
		},
		Strong: ansi.StylePrimitive{
			Bold:  boolPtr(true),
			Color: stringPtr("222"),
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "📑 ",
			Color:       stringPtr("159"),
			// Indent: uintPtr(4),
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
		},
		Task: ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         "[✓] ",
			Unticked:       "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:     stringPtr("5"),
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{
			Color: stringPtr("30"),
			Bold:  boolPtr(true),
		},
		Image: ansi.StylePrimitive{
			Color:     stringPtr("212"),
			Underline: boolPtr(true),
		},
		ImageText: ansi.StylePrimitive{
			Color:  stringPtr("243"),
			Format: "Image: {{.text}} →",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix:     "'",
				BlockSuffix:     "'",
				Prefix:          "",
				Suffix:          "",
				Color:           stringPtr("120"),
				BackgroundColor: stringPtr("236"),
				Italic:          boolPtr(true),
			},
		},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{},
			},
			CenterSeparator: stringPtr("┼"),
			ColumnSeparator: stringPtr("│"),
			RowSeparator:    stringPtr("─"),
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: "\n🠶 ",
		},
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dracula"),
		glamour.WithEmoji(),
		glamour.WithPreservedNewLines(),
		glamour.WithWordWrap(width),
		glamour.WithStyles(customStylingConfig),
	)
	if err != nil {
		fmt.Println("Error: formatting the text!")
	}

	tempVar := *progress
	// Progress := *progress
	var Category, CategoryDescription, Topic, TopicDescription, SubTopic, SubTopicDescription string

	for i := 0; i < len(roadmapData.Categories); i++ {
		for j := 0; j < len(roadmapData.Categories[i].Topics); j++ {
			if len(roadmapData.Categories[i].Topics[j].Subtopics) > tempVar {

				a := roadmapData.Categories[i].Topics[j].Subtopics[tempVar]

				for key, value := range a {

					SubTopic = key
					SubTopicDescription = value

				}

				Category = roadmapData.Categories[i].Category
				CategoryDescription = roadmapData.Categories[i].Description
				Topic = roadmapData.Categories[i].Topics[j].Topic
				TopicDescription = roadmapData.Categories[i].Topics[j].Description

				RawResult := CategoryDescription +
					"\n  ## " + Topic +
					"\n > " + TopicDescription

				halfWidth := strings.Repeat(" ", (width-30)/2)
				fmt.Println(halfWidth, "\033[0;38;5;159m────── \033[0;38;5;210;1m✨ Day", *progress, "✨ \033[0;38;5;159m──────")

				halfWidth = strings.Repeat("─", (width-len(Category)-8)/2)
				result, _ := r.Render(RawResult)

				fmt.Print("\033[0;38;5;159;1m", halfWidth, " 🚀 \033[0;38;5;183;1m", Category, " 🚀 \033[0;38;5;159;1m", halfWidth)
				fmt.Println(result)

				halfWidth = strings.Repeat(" ", (width-len(SubTopic)-7)/2)
				fmt.Print(halfWidth, "\033[0;38;5;183;1;4m📌 ", SubTopic, "\033[0;m")

				result, _ = r.Render(SubTopicDescription)
				fmt.Println(result)

				fmt.Println(strings.Repeat("\033[0;38;5;159m─", width))

				os.Exit(0)

			}

			tempVar -= len(roadmapData.Categories[i].Topics[j].Subtopics)

		}
	}
	// Show some final instruction or something.
	// In here we can consider the learning state is higher than we have to teach
	os.Exit(1)

}

func terminalWidth() int {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 100
	}
	return width
}

func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
func boolPtr(b bool) *bool       { return &b }

// by reading the learning-state.state file get the progress of current learning
func getLearningProgress() (*int, error) {

	// check if the 'learning-state.state' exist
	if fileExists(learningStateFile) {

		// read the learning-state.state file
		learningStateData, err := os.ReadFile(learningStateFile)
		if err != nil {
			return nil, errors.New("error: Reading the 'learning-state.state' file")

		}

		// convert string to int
		strProgress := strings.TrimSpace(string(learningStateData))
		progress, _ := strconv.Atoi(strProgress)

		return &progress, nil

	}

	return nil, errors.New("the 'learning-state.state' file does not exist")

}

// Increment the progress
func done() {

	currentProgress, err := getLearningProgress()
	if err != nil {
		newProgress := 0
		writeState(newProgress)
	}

	fmt.Println("🔥 \033[0;38;5;212;1mCongrats🎉 \033[0;38;5;120mon Your journey to become a successful Security Engineer\033[0;38;5;212;1m!")
	fmt.Println("🔥 \033[0;38;5;212mYou have progressed to Day\033[38;5;120m", *currentProgress+1)
	fmt.Println()

	writeState(*currentProgress + 1)

	show_default()

	os.Exit(0)

}

// Change for the custom progress
func custom(progress int) {

	err := writeState(progress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	show_default()

	os.Exit(0)

}

// Writes the given integer into the state file
func writeState(state int) error {

	strState := strconv.Itoa(state)

	data := []byte(strState)

	err := os.WriteFile(learningStateFile, data, 0644)
	if err != nil {
		return errors.New("error Writing the progress into learning-state file")
	}

	return nil

}

// Check if the given file exist
func fileExists(file string) bool {
	_, err := os.Stat(file)

	return err == nil

}

// Retrieve data from the JSON file
func loadJsonData(filename string) (*RoadmapJsonData, error) {

	// Open json file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// Close the json file
	defer file.Close()

	// Read the json file content
	var roadmapData RoadmapJsonData

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&roadmapData)
	if err != nil {
		return nil, err
	}

	return &roadmapData, nil

}
