package main

import (
	"os"
	"strconv"
	"time"
	"wallps/data"
	sysc "wallps/sysCalls"
	"wallps/utils"
)

const SHELL_SCRIPT_NAME string = "run.sh"

func main() {
	args := os.Args
	if len(args) == 1 {
		simpleForcedExecution()
	} else if args[1] == "a" {
		simpleAutoExecution()
	} else if args[1] == "x" {
		cancelAutoExecution()
	} else if args[1] == "u" {
		completeExecution()
	}
}

func simpleForcedExecution() {
	progData := data.Data()
	progData.AutoSwitch = true
	data.WriteDataFile(progData)

	println("simpleExecution")
	ChangeImage()
	prepareNextExecution()
}

func simpleAutoExecution() {
	progData := data.Data()
	if !progData.AutoSwitch {
		return
	}
	ChangeImage()
	prepareNextExecution()
}

func cancelAutoExecution() {
	progData := data.Data()
	progData.AutoSwitch = false
	data.WriteDataFile(progData)
}

func completeExecution() {
	println("WALLPAPER SWITCHER!!!!!!!")
	println("select one of the options:")
	println("1) create new collection")
	println("2) switch collection")
	println("3) delete collection")
	println("4) configure colection")
}

// change the image of the current used collection
func ChangeImage() {
	progData := data.Data()                                                    // gets the program data
	imageIndex := progData.Collections[progData.CurrentCollection].NextImage() // switch the current image index
	data.WriteDataFile(progData)                                               // saves the changes made to the program data
	CollectionDirPath := progData.Collections[progData.CurrentCollection].Path // get the path of the collection directory
	dir := utils.Expect(os.ReadDir(CollectionDirPath))                         // reads the collections's directory
	imageName := dir[imageIndex].Name()                                        // gets the current image's name
	imagePath := CollectionDirPath + "/" + imageName                           // gets the current images'path
	sysc.SetWall(imagePath)                                                    // set the image as background
}

func prepareNextExecution() {
	progData := data.Data()
	changePeriod := time.Duration(progData.Collections[progData.CurrentCollection].ChangePeriod)

	now := time.Now()
	executionTime := now.Add(time.Minute * changePeriod)
	hour := executionTime.Hour()
	minute := executionTime.Minute()

	hourStr := strconv.FormatInt(int64(hour), 10)
	minuteStr := strconv.FormatInt(int64(minute), 10)
	if hour < 10 {
		hourStr = "0" + hourStr
	}
	if minute < 10 {
		minuteStr = "0" + minuteStr
	}
	timeStr := hourStr + ":" + minuteStr

	sysc.ScheduleRun(timeStr)
}
