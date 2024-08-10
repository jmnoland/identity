package model

import (
    "net/http"
)

var Actions = map[string]string{
   "Create": "Create",
   "Update": "Update",
   "Delete": "Delete",
}

var ResponseTypes = map[string]int{
    "EXCEPTION": http.StatusInternalServerError,
    "BADREQUEST": http.StatusBadRequest,
    "CREATED": http.StatusCreated,
    "NOTFOUND": http.StatusNotFound,
    "OK": http.StatusOK,
}

