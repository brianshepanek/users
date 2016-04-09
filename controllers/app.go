package controllers

import (
    "net/http"
)


func AppIndex(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
