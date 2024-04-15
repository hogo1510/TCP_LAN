
package main

import (
    "fmt"
    "net"
)

func main() {
    // Luister op alle netwerkinterfaces op poort 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Fout bij het starten van de server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server gestart. Wachtend op verbindingen...")

    // Accepteer verbindingen van clients
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Fout bij het accepteren van verbinding:", err)
            continue
        }
        fmt.Println("Nieuwe verbinding geaccepteerd:", conn.RemoteAddr())

        // Behandel de verbinding in een aparte goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Ontvang berichten van de client en stuur ze terug
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Fout bij het lezen van gegevens:", err)
            return
        }
        if n == 0 {
            return // Verbinding verbroken door client
        }
        fmt.Printf("Ontvangen van %s: %s\n", conn.RemoteAddr(), string(buf[:n]))

        // Stuur het ontvangen bericht terug naar de client
        _, err = conn.Write(buf[:n])
        if err != nil {
            fmt.Println("Fout bij het schrijven van gegevens:", err)
            return
        }
    }
}
