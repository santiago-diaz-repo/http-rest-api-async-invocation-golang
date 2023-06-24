package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"http-rest-api-async-invocation-golang/models"
	"http-rest-api-async-invocation-golang/repository"
	"io/ioutil"
	"net/http"
	"time"
)

func HandleRequests(){
	http.HandleFunc("/uuid-body",handlerUuidBody)
	http.HandleFunc("/location-header", handlerLocation)
	http.HandleFunc("/check",handlerCheck)
	http.HandleFunc("/push", handlerPush)
	http.ListenAndServe(":8088",nil)
}

func handlerUuidBody(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json"{
		http.Error(w, "Only Json", http.StatusUnsupportedMediaType)
	}

	uuid := uuid.New().String()

	go asyncProcessPooling(uuid)

	res := models.Response{
		Uuid: uuid,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(res)

}

func handlerLocation(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST"{
		http.Error(w,"Only POST", http.StatusMethodNotAllowed)
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json"{
		http.Error(w,"Only Json", http.StatusUnsupportedMediaType)
	}

	uuid := uuid.New().String()
	go asyncProcessPooling(uuid)

	w.Header().Set("Location", "http://"+r.Host+"/check?uuid="+uuid)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func handlerCheck(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w,"Only POST", http.StatusMethodNotAllowed)
	}

	uuid := r.URL.Query()["uuid"]
	if len(uuid[0]) < 1 {
		http.Error(w,"uuid is missing", http.StatusBadRequest)
	}

	id := repository.NewRepository()[uuid[0]]
	res := models.ResponseAsync{
		State: id,
	}

	w.Header().Set("Content-Type", "application/json")
	if !id {
		w.WriteHeader(http.StatusConflict)
	}else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(res)
}

func handlerPush(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w,"Only POST", http.StatusMethodNotAllowed)
		return
	}

	if contentType := r.Header.Get("Content-Type") ; contentType != "application/json"{
		http.Error(w, "Only json", http.StatusUnsupportedMediaType)
		return
	}

	uuid := uuid.New().String()
	var req models.RequestPush

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error reading Body" + err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error reading Body" + err.Error(), http.StatusInternalServerError)
		return
	}

	go asyncProcessPush(req.UrlCallback, uuid)
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
}

func asyncProcessPooling(uuid string)  {
	time.Sleep(5 * time.Second)
	fmt.Println("Processing uuid " + uuid)
	repository.NewRepository()[uuid] = true
}

func asyncProcessPush(url string, uuid string){
	time.Sleep(5 * time.Second)
	fmt.Printf("Invoking callback %s for UUID %s", url, uuid)
}