package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zzucturum-app/pkg/db"
	"zzucturum-app/pkg/models"
	"zzucturum-app/pkg/utils"
)

func AddSocialMedia(w http.ResponseWriter, r *http.Request) {
	var socialMedia []models.SocialMedia

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

	if err := json.Unmarshal(jsonString, &socialMedia); err != nil {
		log.Fatal(err)
	}

	for _, v := range socialMedia {
		db.InsertCounter(v)
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteSuccess(w, fmt.Sprintf("Succesfull collect"))
}

func EditSocialMedia(w http.ResponseWriter, r *http.Request) {
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
	utils.WriteSuccess(w, string(statsMarshaled))
}

func GetSocialMedia(w http.ResponseWriter, r *http.Request) {
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
	utils.WriteSuccess(w, string(statsMarshaled))

}
