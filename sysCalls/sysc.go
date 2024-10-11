package sysc

import (
	"os"
	"os/exec"
	"wallps/utils"
)

func SetWall(filePath string) {
	cmdPath := "file://" + filePath
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", cmdPath)

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// creates a shell script and schedule it's run
func ScheduleRun(timeStr string) {
	shellScript := utils.ProgramDirectory() + "/" + "run.sh"
	file := utils.Expect(os.Create(shellScript))
	file.WriteString("#!bin/bash\n" + utils.ProgramPath() + " a")
	err := exec.Command("at", timeStr, "-f", shellScript).Run()
	if err != nil {
		panic(err)
	}
}
