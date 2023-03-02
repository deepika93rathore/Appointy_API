package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Driver represents a driver with their ID, current location, and the time of the last update.
type Driver struct {
	ID         int
	Location   Location
	LastUpdate time.Time
}

// Location represents a geographic location with a latitude, longitude, and accuracy in meters.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  float64 `json:"accuracy"`
}

// Drivers represents a collection of drivers with a lock to ensure thread safety.
type Drivers struct {
	sync.RWMutex
	Data map[int]*Driver
}

// NewDrivers creates a new instance of Drivers.
func NewDrivers() *Drivers {
	return &Drivers{
		Data: make(map[int]*Driver),
	}
}

// UpdateLocation updates a driver's location.
func (d *Drivers) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/drivers/"):])
	if err != nil || id < 1 || id > 50000 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var loc Location
	err = json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if loc.Latitude < -90 || loc.Latitude > 90 {
		http.Error(w, "Latitude should be between +/- 90", http.StatusUnprocessableEntity)
		return
	}

	if loc.Longitude < -180 || loc.Longitude > 180 {
		http.Error(w, "Longitude should be between +/- 180", http.StatusUnprocessableEntity)
		return
	}

	d.Lock()
	defer d.Unlock()

	driver, ok := d.Data[id]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	driver.Location = loc
	driver.LastUpdate = time.Now()

	w.WriteHeader(http.StatusOK)
}

// FindDrivers finds drivers within a certain radius of a given location.
func (d *Drivers) FindDrivers(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")
	limitStr := r.URL.Query().Get("limit")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil || lat < -90 || lat > 90 {
		http.Error(w, "Latitude should be between +/- 90", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil || lon < -180 || lon > 180 {
		http.Error(w, "Longitude should be between +/- 180", http.StatusBadRequest)
		return
	}

	radius := 500.0
	if radiusStr != "" {
		radius, err = strconv.ParseFloat(radiusStr, 64)
		if err != nil || radius < 0 {
			http.Error(w, "Radius should be a positive number", http.StatusBadRequest)
			return
		}
	}

	limit := 10
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			http.Error(w, "Limit should be a positive integer", http.StatusBadRequest)
			return
		}
	}

	d.R
