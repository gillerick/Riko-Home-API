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
}

var commands []Command

// Index Handler
func indexHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Welcome to Riko Home Automation!")
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
	command.ID = strconv.Itoa(rand.Intn(1000))
	commands = append(commands, command)
	json.NewEncoder(w).Encode(command)
}
func main() {
	commands = append(commands, Command{
		ID:        "1",
		Text:      "Turn on the kitchen lights"})

	commands = append(commands, Command{
		ID:        "2",
		Text:      "Turn Off the bedroom lights"})

	//Initialize the Mux Handler
	r := mux.NewRouter()

	//Define the endpoints
	r.HandleFunc("/commands/{id}", getCommand).Methods("GET")
	r.HandleFunc("/commands", createCommand).Methods("POST")
	r.HandleFunc("/commands", getCommands).Methods("GET")
	r.HandleFunc("/", indexHandler)

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
