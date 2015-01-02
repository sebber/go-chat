package main

import (
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
  "github.com/sebber/go-chat/core/entities"
)


var(
  messageRepo entities.Messages
)


/**
 * The handlers
 */

func GetMessages(w rest.ResponseWriter, req *rest.Request) {
  w.WriteJson(messageRepo.GetAll())
}

func PostMessage(w rest.ResponseWriter, req *rest.Request) {
  message := entities.Message{}

  if err := req.DecodeJsonPayload(&message); err != nil {
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  messageRepo.Put(message)

  w.WriteJson(message)
}


/**
 *  Startingpoint
 */

func main() {
  messageRepo.Store = []*entities.Message{}

  apiHandler := rest.ResourceHandler{
    EnableRelaxedContentType: true,
  }

  err := apiHandler.SetRoutes(
      &rest.Route{"GET", "/messages", GetMessages},
      &rest.Route{"POST", "/messages", PostMessage},
  )
  if err != nil {
      log.Fatal(err)
  }

  http.Handle("/api/", http.StripPrefix("/api", &apiHandler))

  http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir("public/"))))

  log.Fatal(http.ListenAndServe(":8080", nil))
}
