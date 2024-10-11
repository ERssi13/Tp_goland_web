package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "regexp"
)
type Utilisateur struct {
    Nom            string
    Prenom         string
    DateNaissance  string
    Sexe           string
}
func afficherFormulaire(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("template/form.html")
    if err != nil {
        http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
    }
}
func traiterFormulaire(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/user/form", http.StatusSeeOther)
        return
    }
    nom := r.FormValue("nom")
    prenom := r.FormValue("prenom")
    dateNaissance := r.FormValue("date_naissance")
    sexe := r.FormValue("sexe")
    if !validerNom(nom) || !validerNom(prenom) || !validerSexe(sexe) {
        http.Redirect(w, r, "/user/error", http.StatusSeeOther)
        return
    }
    utilisateur := Utilisateur{
        Nom:           nom,
        Prenom:        prenom,
        DateNaissance: dateNaissance,
        Sexe:          sexe,
    }
    http.Redirect(w, r, "/user/display?nom="+utilisateur.Nom+"&prenom="+utilisateur.Prenom+"&date_naissance="+utilisateur.DateNaissance+"&sexe="+utilisateur.Sexe, http.StatusSeeOther)
}
func afficherDonnees(w http.ResponseWriter, r *http.Request) {
    nom := r.URL.Query().Get("nom")
    prenom := r.URL.Query().Get("prenom")
    dateNaissance := r.URL.Query().Get("date_naissance")
    sexe := r.URL.Query().Get("sexe")

    if nom == "" || prenom == "" || dateNaissance == "" || sexe == "" {
        http.Error(w, "Veuillez renseigner toutes les informations personnelles.", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Nom: %s<br>Prénom: %s<br>Date de Naissance: %s<br>Sexe: %s", nom, prenom, dateNaissance, sexe)
}
func afficherErreur(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Données non valides. Veuillez vérifier vos informations et réessayer.")
}
func validerNom(nom string) bool {
    if len(nom) < 1 || len(nom) > 32 {
        return false
    }
    return regexp.MustCompile(`^[A-Za-zÀ-ÖØ-öø-ÿ]+$`).MatchString(nom)
}
func validerSexe(sexe string) bool {
    return sexe == "masculin" || sexe == "féminin" || sexe == "autre"
}

func main() {
    http.HandleFunc("/user/form", afficherFormulaire)
    http.HandleFunc("/user/treatment", traiterFormulaire)
    http.HandleFunc("/user/display", afficherDonnees)
    http.HandleFunc("/user/error", afficherErreur)
    log.Println("Serveur démarré sur http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
