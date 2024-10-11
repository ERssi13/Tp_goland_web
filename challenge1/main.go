package main

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)
type Etudiant struct {
    Prenom  string
    Nom     string
    Age     int
    Sexe    string
}
type Promotion struct {
    NomClasse   string
    Filiere     string
    Annee      string
    NbEtudiants int
    Etudiants   []Etudiant
}
func afficherPromotion(w http.ResponseWriter, r *http.Request) {
    infosPromotion := Promotion{
        NomClasse:   "B1 Informatique",
        Filiere:     "Informatique",
        Annee:      "Bachelor 1",
        NbEtudiants: 3,
        Etudiants: []Etudiant{
            {Prenom: "Maria", Nom: "Guash", Age: 42, Sexe: "F"},
            {Prenom: "Luther", Nom: "Martin", Age: 17, Sexe: "M"},
            {Prenom: "Charlie", Nom: "Hebdo", Age: 68, Sexe: "M"},
        },
    }
    tmplPath := filepath.Join("template", "temp.html")
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, infosPromotion)
    if err != nil {
        http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
    }
}

func main() {
    fs := http.FileServer(http.Dir("./image"))
    http.Handle("/image/", http.StripPrefix("/image/", fs))
    http.HandleFunc("/promo", afficherPromotion)
    log.Println("Serveur démarré sur http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
