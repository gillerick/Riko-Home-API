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
	params := mux.Vars(r)
	for _, item := range commands {
		if params["id"] == item.ID {
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
	fmt.Fprint(w, command)
}



//Command struct
type Command struct {
	ID        string  `json:"id"`
	Text      string  `json:"text"`
	Command   string  `json:"command"`
	TimeStamp string  `json:"timestamp"`
	Status    string  `json:"status"`
	Device    *Device `json:"device"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"status"`
}

var commands []Command
var devices []Device

func main() {

	//Initialize the Mux Handler
	r := mux.NewRouter()

	//Define the endpoints
	r.HandleFunc("/api/commands{id}", getCommand).Methods("GET")
	r.HandleFunc("/api/commands", createCommand).Methods("POST")
	r.HandleFunc("/api/commands", getCommands).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}




	devices = append(devices, Device{
		"1",
		"Kitchen Lights",
	})

	devices = append(devices, Device{
		"2",
		"Bedroom Lights",
	})

	devices = append(devices, Device{
		"3",
		"Lounge Room Lights",
	})

	devices = append(devices, Device{
		"4",
		"Washroom Lights",
	})

	devices = append(devices, Device{
		"5",
		"Buzzer",
	})

	commands = append(commands, Command{
		"1",
		"Turn on the kitchen lights",
		"on kitchen",
		"20200126",
		"ON",
		&Device{
			"1",
			"Kitchen Lights",
		},
	})

	commands = append(commands, Command{
		"2",
		"Turn Off the bedroom lights",
		"Off bedroom",
		"20200126",
		"OFF",
		&Device{
			"2",
			"Bedroom Lights",
		},
	})



}