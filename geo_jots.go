package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mattbaird/elastigo/core"
	"net/http"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GeoJot struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type GeoJotJSON struct {
	GeoJot GeoJot `json:"geo_jot"`
}

type GeoJotsJSON struct {
	GeoJots []GeoJot `json:"geo_jots"`
}

var es_index = "geo_jots"
var es_index_type = "geo_jot"

func GeoJotsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		serveError(w, err)
		return
	}

	lat, err := strconv.ParseFloat(r.Form.Get("lat"), 64)
	if err != nil {
		serveError(w, err)
		return
	}

	lon, err := strconv.ParseFloat(r.Form.Get("lon"), 64)
	if err != nil {
		serveError(w, err)
		return
	}

	// Find all GeoJots that are close by
	qry := map[string]interface{}{
		"from": 0,
		"size": 50,
		"query": map[string]interface{}{
			"filtered": map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
				"filter": map[string]interface{}{
					"geo_distance": map[string]interface{}{
						"distance": "1km",
						"geo_jot.location": map[string]interface{}{
							"lat": lat,
							"lon": lon,
						},
					},
				},
			},
		},
	}

	resp, err := core.SearchRequest(true, es_index, es_index_type, qry, "")
	//resp, err := core.SearchUri(es_index, es_index_type, "name:*", "")

	if err != nil {
		serveError(w, err)
	}

	var jots []GeoJot
	if resp.Hits.Total > 0 {
		for _, value := range resp.Hits.Hits {
			var jot GeoJot
			if err := json.Unmarshal(value.Source, &jot); err != nil {
				serveError(w, err)
				return
			}
			jot.Id = value.Id
			jots = append(jots, jot)
		}
	}

	j, err := json.Marshal(GeoJotsJSON{GeoJots: jots})
	if err != nil {
		serveError(w, err)
	}
	w.Write(j)
}

func ShowGeoJotHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateGeoJotHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming geojot from the request body
	var geojotJSON GeoJotJSON
	err := json.NewDecoder(r.Body).Decode(&geojotJSON)
	if err != nil {
		serveError(w, err)
	}

	geojot := geojotJSON.GeoJot

	// add single go struct entity
	response, _ := core.Index(true, es_index, es_index_type, "", geojot)

	// Grab the Id for the client
	geojot.Id = response.Id

	// Serialize the modified geojot to JSON
	j, err := json.Marshal(GeoJotJSON{GeoJot: geojot})
	if err != nil {
		serveError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func UpdateGeoJotHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming geojot json
	var geojotJSON GeoJotJSON
	err := json.NewDecoder(r.Body).Decode(&geojotJSON)
	if err != nil {
		serveError(w, err)
	}

	// TODO find the GeoJot and update it in ES
	// Grab the geojot's id from the incoming url
	// vars := mux.Vars(r)
	// vars["id"]

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)
}

func DeleteGeoJotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := core.Delete(true, es_index, es_index_type, vars["id"], 1, "")
	if err != nil {
		serveError(w, err)
	}

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)

}
