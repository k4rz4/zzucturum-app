package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"zzucturum-app/pkg/db"
	"zzucturum-app/pkg/models"
	"zzucturum-app/pkg/utils"
)

var log = logrus.New()
var timestamp int64
var msgStringFail = "Sorry, couldn't find data :("

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var counter []models.Counter

	slug, err := getPostFields(r)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteError(w, err)
		return
	}

	jsonString, err := json.Marshal(slug)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(jsonString, &counter); err != nil {
		log.Fatal(err)
	}

	for _, v := range counter {
		db.InsertCounter(v)
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteSuccess(w, fmt.Sprintf("Succesfull collect"))
}


func GetStatsByMinute (w http.ResponseWriter, r *http.Request) {
	stats, err := db.GetCounterStatsByMinutes()
	if err != nil {
		log.Fatal(err)
	}

	if len(stats) <= 0 {
		w.WriteHeader(http.StatusOK)
		utils.WriteSuccess(w, msgStringFail)
		return
	}

	statsMarshaled, err := json.Marshal(stats)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteSuccess(w,string(statsMarshaled))
}

func GetStatsByHour(w http.ResponseWriter, r *http.Request) {
	stats, err := db.GetCounterStatsByHour()
	if err != nil {
		log.Fatal(err)
	}

	if len(stats) <= 0 {
		w.WriteHeader(http.StatusOK)
		utils.WriteSuccess(w, msgStringFail)
		return
	}

	statsMarshaled, err := json.Marshal(stats)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteSuccess(w,string(statsMarshaled))

}

func getPostFields(r *http.Request) ([]map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	var m map[string]interface{}

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}

	mm := make([]map[string]interface{},0)
	log.Debug(m)

	if val, found := m["timestamp"]; found {
		timestamp, err = strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		delete(m, "timestamp")
	} else {
		return mm, errors.New("malformed request, no timestamp")
	}

	for k, v := range m {

		newMap := map[string]interface{} {
			"domain":           k,
			"numberOfRequests": v,
			"timestamp":        timestamp,
		}
		mm = append(mm, newMap)
	}
	return mm, nil
}