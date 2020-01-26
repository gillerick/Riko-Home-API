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
type Transcript struct {
	ID int  `json:"id"`
	Text  string  `json:"text"`
}

type Command struct {
	ID int `json:"id"`
	Command string `json:"command"`
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

//index handler that fetches all transcripts
func getTranscripts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transcripts)
}


// indexHandler that fetches commands
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
	transcript.ID = rand.Intn(rand.Intn(100))
	transcripts = append(transcripts, transcript)
	json.NewEncoder(w).Encode(transcript)
}
func main() {
	transcripts = append(transcripts, Transcript{ ID: 1, Text: "Turn on the kitchen lights"})
	transcripts = append(transcripts, Transcript{ ID: 2, Text: "Turn Off the bedroom lights"})

	//Initialize the Mux Handler
	r := mux.NewRouter()

	//Define the endpoints
	r.HandleFunc("/transcript/{id}", getTranscript).Methods("GET")
	r.HandleFunc("/transcript", createTranscript).Methods("POST")
	r.HandleFunc("/transcripts", getTranscripts).Methods("GET")
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
