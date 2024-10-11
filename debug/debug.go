package testes

import "wallps/data"

func MinimumDataFileCreation() {

	colA := data.Collection{
		Nome:         "dark",
		Tipo:         "s",
		Path:         "/home/rmxv/tudo/images/wallpapers/dark",
		ChangePeriod: 2,
	}

	colB := data.Collection{
		Nome:         "light",
		Tipo:         "s",
		Path:         "/home/rmxv/tudo/images/wallpapers/light",
		ChangePeriod: 2,
	}

	var mimimumData data.ProgramData

	mimimumData.Collections = append(mimimumData.Collections, colA, colB)
	mimimumData.AutoSwitch = true

	data.WriteDataFile(mimimumData)
}
