import (
	"encoding/json"
        "fmt"
	"log"
	"net/http"
        "gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/moezilla/antispam-api/config"
	. "github.com/moezilla/antispam-api/db"
	. "github.com/moezilla/antispam-api/type"
    }

func init() {
	fmt.Println("Server Starting...")
}

var config = Config{}
var db = MoviesDB{}

func respondWithError(scammer http.ResponseWriter, code int, msg string) {
	respondWithJSON(scammer, code, map[string]string{"error": msg})
}

func respondWithJSON(scammer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	scammer.Header().Set("Content-Type", "application/json")
	scammer.WriteHeader(code)
	scammer.Write(response)
}

func FindScammer(scammer http.ResponseWriter, router *http.Request) {
	params := mux.Vars(router)
	antispam, err := db.FindById(params["id"])
	if err != nil {
		respondWithError(scammer, http.StatusBadRequest, "Who is this user")
		return
	}
	respondWithJson(scammer, http.StatusOK, antispam)
}

func BanScammer(scammer http.ResponseWriter, router *http.Request) {
	defer router.Body.Close()
	var antispam Antispam
	if err := json.NewDecoder(router.Body).Decode(&antispam); err != nil {
		respondWithError(scammer, http.StatusBadRequest, "Failed Request")
		return
	}
	antispam.ID = bson.NewObjectId()
	if err := db.Insert(antispam); err != nil {
		respondWithError(scammer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(scammer, http.StatusCreated, antispam)
}

func UpdateBanScammer(scammer http.ResponseWriter, router *http.Request) {
	defer router.Body.Close()
	var antispam Antispam
	if err := json.NewDecoder(router.Body).Decode(&antispam); err != nil {
		respondWithError(scammer, http.StatusBadRequest, "Failed Scammer")
		return
	}
	if err := db.Update(antispam); err != nil {
		respondWithError(scammer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(scammer, http.StatusOK, map[string]string{"success"})
}

func UnbanScammer(scammer http.ResponseWriter, router *http.Request) {
	defer router.Body.Close()
	var antispam Antispam
	if err := json.NewDecoder(router.Body).Decode(&antispam); err != nil {
		respondWithError(scammer, http.StatusBadRequest, "Failed Request")
		return
	}
	if err := db.Delete(antispam); err != nil {
		respondWithError(scammer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(scammer, http.StatusOK, map[string]string{"success"})
}
 
func init() {
	config.Read()

	db.Server = config.Server
	fmt.Println("Server Name:", config.Server)
	db.Database = config.Database
	fmt.Println("Database:", config.Database)
	db.Connect()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/banscammer", BanScammer).Methods("POST")
	router.HandleFunc("/banscammer", UpdateBanScammer).Methods("PUT")
	router.HandleFunc("/unban", UnbanScammer).Methods("DELETE")
	router.HandleFunc("/findscammer/{id}", FindScammer).Methods("GET")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
