
package main

import (
    "fmt"
    "net"
    "os"
    "flag"
)

func main() {
    if len(os.Args) != 3 {
       
    }
    //serverIP, serverPOORT := os.Args[1], os.Args[2]
    var serverIP string
    var serverPORT string
    flag.StringVar(&serverIP, "i","", "Zet hier de server IP")
    flag.StringVar(&serverPORT, "p","", "Zet hier de server Poort")

    flag.Parse()

    if serverIP == "" || serverPORT == ""{
        fmt.Println("Gebruik: ./client -i <server_ip> -p <server_poort>")
        return
    }
    fmt.Println("IP : ",serverIP)
    fmt.Println("PORT : ",serverPORT)

	//samenvoegen
	address := fmt.Sprintf("%s:%s", serverIP, serverPORT)
    // Verbinding maken met de server
    conn, err := net.Dial("tcp", address)
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
