package main

import (
    "fmt"
    "net"
    "flag"
)

func main() {

	var PORT string
	flag.StringVar(&PORT, "p","","Poort waar de server runt")
	flag.Parse()
	if PORT == ""{
	fmt.Println("geen poort gespecificeerd gebruik: ./LAN2.go -p <poort nummer>")
	return
	}
	address :=  fmt.Sprintf(":%s", PORT)

    //luister op alle netwerkinterfaces 
    listener, err := net.Listen("tcp", address)
    if err != nil {
        fmt.Println("Fout bij het starten van de server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server gestart. Wachtend op verbindingen...")

    //accepteer verbindingen van clients
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Fout bij het accepteren van verbinding:", err)
            continue
        }
        fmt.Println("Nieuwe verbinding geaccepteerd:", conn.RemoteAddr())

        //behandel de verbinding in aparte goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    //ontvang berichten van de client
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Fout bij het lezen van gegevens:", err)
            return
        }
        if n == 0 {
		fmt.Println("Verbinding met client verbroken!")
            return 
        }
        fmt.Printf("Ontvangen van %s: %s\n", conn.RemoteAddr(), string(buf[:n]))

        //terug sturen van bericht naar client
        _, err = conn.Write(buf[:n])
        if err != nil {
            fmt.Println("Fout bij het schrijven van gegevens:", err)
            return
        }
    }
}
