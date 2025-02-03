package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "unknown"
	}

	err = tmpl.Execute(w, PageData{Hostname: hostname})
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	slow := r.URL.Query().Get("slow") == "true"

	if slow {
		time.Sleep(2 * time.Second)
	}

	products := []Product{
		{"PlayStation 1", "https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/PSX-Console-wController.png/1600px-PSX-Console-wController.png"},
		{"PlayStation 2", "https://upload.wikimedia.org/wikipedia/commons/thumb/5/58/PS2-Fat-Console-Set.png/987px-PS2-Fat-Console-Set.png"},
		{"PlayStation 3", "https://banner2.cleanpng.com/20180417/ase/kisspng-playstation-3-playstation-2-playstation-4-grand-th-sony-playstation-5ad62609071490.750687421523983881029.jpg"},
		{"PlayStation 4", "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/PS4-Console-wDS4.png/1600px-PS4-Console-wDS4.png"},
		{"PlayStation 5", "https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Black_and_white_Playstation_5_base_edition_with_controller.png/1024px-Black_and_white_Playstation_5_base_edition_with_controller.png"},
		{"PSP", "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/PSP-1000.png/1600px-PSP-1000.png"},
		{"PS Vita", "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/PlayStation-Vita-1101-FL.png/1600px-PlayStation-Vita-1101-FL.png"},
	}
	rand.Shuffle(len(products), func(i, j int) { products[i], products[j] = products[j], products[i] })

	elapsed := time.Since(start).Seconds()
	searchDuration.WithLabelValues("simple").Observe(elapsed)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
