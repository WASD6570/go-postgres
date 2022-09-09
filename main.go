package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/wasd6570/go-postgres/db"
	"github.com/wasd6570/go-postgres/routes"
)

func get_workdir() string {
	path_dir, path_err := os.Getwd()
	if path_err != nil {
		log.Fatal(path_err)
	}
	return path_dir
}

func load_env_vars() error {

	work_dir := get_workdir()

	pathErr := godotenv.Load(filepath.Join(work_dir, ".env"))

	return pathErr
}

func main() {
	err := load_env_vars()
	if err != nil {
		log.Fatalf("error loading env vars, error: %e", err)
	}

	_, conn_err := db.Connection()

	if conn_err != nil {
		log.Fatalf("error connecting to the db, error: %e", conn_err)
	}

	log.Printf("connected to the db")

	// conn.Exec("DROP TABLE IF EXISTS auths cascade;")
	// conn.Exec("DROP TABLE IF EXISTS users cascade;")

	// conn.AutoMigrate(models.User{}, models.Auth{})

	routes.Init_server()

}
