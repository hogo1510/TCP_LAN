
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Gebruik: ./client <server_ip>")
        return
    }
    serverIP := os.Args[1]

    // Verbinding maken met de server op opgegeven IP-adres op poort 8080
    conn, err := net.Dial("tcp", serverIP+":8080")
    if err != nil {
        fmt.Println("Kan geen verbinding maken met server:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Verbonden met de server. Typ een bericht en druk op Enter om te verzenden.")

    // Luister naar invoer van de gebruiker en stuur berichten naar de server
    for {
        var msg string
        fmt.Scanln(&msg)

        // Stuur het bericht naar de server
        _, err := conn.Write([]byte(msg))
        if err != nil {
            fmt.Println("Fout bij het verzenden van gegevens naar server:", err)
            return
        }

        // Wacht op antwoord van de server
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Fout bij het ontvangen van gegevens van server:", err)
            return
        }
        fmt.Println("Ontvangen van server:", string(buf[:n]))
    }
}
