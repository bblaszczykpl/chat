package app

type Settings struct {
	Port int
}

type Application struct {
	Settings *Settings
}

func NewApplication() *Application {
	settings := &Settings{
		Port: 8080,
	}

	return &Application{
		Settings: settings,
	}
}
