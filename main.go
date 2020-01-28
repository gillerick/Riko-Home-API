package main

import ("fmt"
"net/http"
"github.com/gorilla/mux"
"encoding/json"
"os"
"log"
"strconv"
"math/rand"
"regexp"
)


//Transcript struct
type Transcript struct {
	ID int  `json:"id"`
	Text  string  `json:"text"`
}

//Command struct
type Command struct {
	ID int `json:"id"`
	Command int `json:"command"`
	DeviceID int `json:"device_id"`
}

var commands []Command
var transcripts []Transcript

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


func getCommand(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range commands {
		if strconv.Itoa(item.ID) == params["device_id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Command{})
}


// indexHandler that fetches transcripts
func getTranscript(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range transcripts {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Transcript{})
}

// indexHandler that fetches commands
func createTranscript(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transcript Transcript
	_ = json.NewDecoder(r.Body).Decode(&transcript)
	transcript.ID = rand.Intn(rand.Intn(10000))
	transcripts = append(transcripts, transcript)
	json.NewEncoder(w).Encode(transcript)


	//Patterns to check out for keywords in transcribed text
	pattern1 := regexp.MustCompile(`on.*bedroom|bedroom.*on`)
	pattern2 := regexp.MustCompile(`off.*bedroom|bedroom.*off`)
	pattern3 := regexp.MustCompile(`on.*sitting|sitting.*on|lounge.*on|on.*lounge`)
	pattern4 := regexp.MustCompile(`off.*sitting|sitting.*off|lounge.*off|off.*lounge`)
	pattern5 := regexp.MustCompile(`on.*washroom|washroom.*on| on.*bathroom|on.*bathroom`)
	pattern6 := regexp.MustCompile(`off.*washroom|washroom.*off| off.*bathroom|off.*bathroom`)
	pattern7 := regexp.MustCompile(`on.*kitchen|kitchen.*on`)
	pattern8 := regexp.MustCompile(`off.*kitchen|kitchen.*off`)
	pattern9 := regexp.MustCompile(`on.*buzzer|buzzer.*on|on.*alarm|alarm.*on`)
	pattern10 := regexp.MustCompile(`play.*music|music.*play`)
	pattern11 := regexp.MustCompile(`lights.*all.*on|all.*lights.*on`)
	pattern12 := regexp.MustCompile(`lights.*all.*off|all.*lights.*off`)


	if pattern1.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:1})
	}

	if pattern2.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:1})
	}

	if pattern3.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:2})
	}

	if pattern4.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:2})
	}

	if pattern5.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:3})
	}

	if pattern6.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:3})
	}

	if pattern7.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:4})
	}

	if pattern8.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:4})
	}

	if pattern9.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:5})
	}

	if pattern10.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:5})
	}

	if pattern11.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:1, DeviceID:0})
	}

	if pattern12.MatchString(transcript.Text){
		delPreviousCommands()
		commands = append(commands, Command{ID: transcript.ID, Command:0, DeviceID:0})
	}
}

//Function that deletes previous commands
func delPreviousCommands() {
	commands = nil
}


func main() {
	//Hard coded transcript for testing
	transcripts = append(transcripts, Transcript{ ID: 1, Text: "Turn on the kitchen lights"})

	//Initialize the Mux Handler
	r := mux.NewRouter()

	//Define the endpoints
	r.HandleFunc("/transcripts/{id}", getTranscript).Methods("GET")
	r.HandleFunc("/transcripts", createTranscript).Methods("POST")
	r.HandleFunc("/commands/{id}", getCommand).Methods("GET")
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
