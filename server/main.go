package main

import (
    "./controllers"
    "log"
    "os"
)

func main() {
    server := controllers.NewServer()

    go func() {
        log.Printf("Server running on port " + os.Getenv("SERVER_PORT") + "...")
        err := server.Server.ListenAndServe()
        if err != nil {
            log.Printf("Error: %v", err)
        }
    }()

    server.WaitShutdown()
    log.Printf("Server shuted down")
}
