package inits

func SyncDB() {
	DB.AutoMigrate()
}
