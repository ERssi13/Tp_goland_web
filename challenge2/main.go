package main

import (
    "html/template"
    "log"
    "net/http"
    "sync"
)

var compteurVues int
var mutex sync.Mutex
type VueData struct {
    Compteur int
    Pair     bool
}

func afficherChange(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    compteurVues++
    compteur := compteurVues
    mutex.Unlock()
    pair := (compteur % 2) == 0
    data := VueData{
        Compteur: compteur,
        Pair:     pair,
    }
    tmpl, err := template.ParseFiles("template/change.html")
    if err != nil {
        http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
    }
}

func main() {
    compteurVues = 0
    http.HandleFunc("/change", afficherChange)
    log.Println("Serveur démarré sur http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
