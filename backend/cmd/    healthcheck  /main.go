package main

import (
        "flag"
        "fmt"
        "net/http"
        "os"
        "time"
)

func main() {
        var port = flag.Int("port", 8081, "HTTP port for healthcheck")
        var verb = flag.Bool("verbose", false, "verbose output for healthcheck")

        flag.Parse()

        client := &http.Client{
                Timeout: 5 * time.Second,
        }
        resp, err := client.Get(fmt.Sprintf("http://localhost:%d/health", *port))
        defer resp.Body.Close()
        if err != nil {
                // Handle error, including potential timeout errors
                if os.IsTimeout(err) {
                        // This is a timeout error
                        fmt.Println("Health check timed out:", err)
                        os.Exit(1)
                } else {
                        // Other network or HTTP errors
                        fmt.Println("Health check error:", err)
                }
                os.Exit(2)

        }

        // Process the response if successful
        if resp.StatusCode == http.StatusOK {
                if *verb == true {
                        fmt.Println("Health check successful!")
                }
        } else {
                fmt.Printf("Health check failed with status: %d\n", resp.StatusCode)
                os.Exit(3)
        }
        os.Exit(0)
}
