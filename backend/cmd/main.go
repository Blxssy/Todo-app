package main

import (
	"github.com/Blxssy/Todo-app/backend/internal/model/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/Blxssy/Todo-app/backend/internal/controller"
	httphandler "github.com/Blxssy/Todo-app/backend/internal/handler/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	staticDir := "./frontend/static"

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.New(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	if err = db.CreateTables(); err != nil {
		log.Println(err)
		return
	}

	router := mux.NewRouter()

	//repo := memory.New()
	ctrl := controller.New(db)
	h := httphandler.New(ctrl)

	router.HandleFunc("/todos", h.Home).Methods("GET")
	router.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/complete", h.CompleteTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", h.HandleDeleteTodo).Methods("DELETE")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	http.ListenAndServe(":8080", router)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
