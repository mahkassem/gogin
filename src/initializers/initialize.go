package initializers

func Initialize(migrateFile bool) {
	HandleCommandArguments(migrateFile)
	LoadEnv()
	LoadConfig()
	ConnectDatabase()
}
