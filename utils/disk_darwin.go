//go:build darwin

package utils

type App struct {
	DisplayIcon     string
	DisplayName     string
	DisplayVersion  string
	InstallLocation string
	Publisher       string
	UninstallString string
}

func QueryApps() []App {

}
