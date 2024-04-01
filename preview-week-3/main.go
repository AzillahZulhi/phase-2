package main

import "preview-week-3/config"

func main() {
	config.DbConnect("root:@tcp(localhost:3306)/preview_week_3")
	defer config.DB.Close()
}
