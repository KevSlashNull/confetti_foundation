package foundation

import "os"

type BasePath string

const pathSeparator = string(os.PathSeparator)

// Get the path to the application "app" directory.
func (basePath BasePath) Path() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the base path of the Laravel installation.
func (basePath BasePath) BasePath() string {
	return string(basePath)
}

// Get the path to the bootstrap directory.
func (basePath BasePath) BootstrapPath() string {
	return string(basePath) + pathSeparator + "bootstrap"
}

// Get the path to the application configuration files.
func (basePath BasePath) ConfigPath() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the path to the database directory.
func (basePath BasePath) DatabasePath() string {
	return string(basePath) + pathSeparator + "database"
}

// Get the path to the language files.
func (basePath BasePath) LangPath() string {
	return basePath.ResourcePath() + pathSeparator + "lang"
}

// Get the path to the public / web directory.
func (basePath BasePath) PublicPath() string {
	return string(basePath) + pathSeparator + "public"
}

// Get the path to the storage directory.
func (basePath BasePath) StoragePath() string {
	return string(basePath) + pathSeparator + "storage"
}

// Get the path to the resources directory.
func (basePath BasePath) ResourcePath() string {
	return string(basePath) + pathSeparator + "resources"
}

// Get the path to the environment file.
func (basePath BasePath) EnvironmentFile() string {
	return string(basePath) + pathSeparator + ".env"
}
