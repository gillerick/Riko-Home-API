package main

import ("fmt"
"net/http"
"github.com/gorilla/mux"
"encoding/json"
"os"
"log"
"strconv"
"math/rand"
)


//Command struct
type Command struct {
	ID        string  `json:"id"`
	Text      string  `json:"text"`
	Command   string  `json:"command"`
	TimeStamp string  `json:"timestamp"`
	Device    *Device `json:"device"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"status"`
	Status string `json:"status"`
}

var commands []Command

// Index Handler
func indexHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

//index handler that fetches all commands
func getCommands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commands)
}


// indexHandler that fetches commands
func getCommand(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range commands {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Command{})
}

// indexHandler that fetches commands
func createCommand(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var command Command
	_ = json.NewDecoder(r.Body).Decode(&command)
	command.ID = strconv.Itoa(rand.Intn(100))
	commands = append(commands, command)
	json.NewEncoder(w).Encode(command)
}

// Update Commands
func updateCommands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range commands {
		if item.ID == params["id"] {
			commands = append(commands[:index], commands[index+1:]...)
			var command Command
			_ = json.NewDecoder(r.Body).Decode(&command)
			command.ID = params["id"]
			commands = append(commands, command)
			json.NewEncoder(w).Encode(command)
			return
		}
	}
}

func main() {



	commands = append(commands, Command{
		ID:        "1",
		Text:      "Turn on the kitchen lights",
		Command:   "on kitchen",
		TimeStamp: "20200126",
		Device: &Device{
			ID:     "1",
			Name:   "Kitchen Lights",
			Status: "ON"}})

	commands = append(commands, Command{
		ID:        "2",
		Text:      "Turn Off the bedroom lights",
		Command:   "Off bedroom",
		TimeStamp: "20200126",
		Device: &Device{
			ID:     "2",
			Name:   "Bedroom Lights",
			Status: "OFF"}})



	//Initialize the Mux Handler
	r := mux.NewRouter()

	//Define the endpoints
	r.HandleFunc("/commands/{id}", getCommand).Methods("GET")
	r.HandleFunc("/commands", createCommand).Methods("POST")
	r.HandleFunc("/commands", getCommands).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
