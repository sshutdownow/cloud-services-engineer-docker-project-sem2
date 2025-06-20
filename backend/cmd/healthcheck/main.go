package main

import (
        "flag"
        "fmt"
        "net/http"
        "os"
        "time"
)

func main() {
        var url = flag.String("url", "http://localhost:8081/health", "HTTP URL for healthcheck")
        var verb = flag.Bool("verbose", false, "verbose output for healthcheck")

        flag.Parse()

        client := &http.Client{
                Timeout: 5 * time.Second,
        }
        resp, err := client.Get(*url)

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

        defer resp.Body.Close()

        // Process the response if successful
        if resp.StatusCode == http.StatusOK {
                if *verb == true {
                        fmt.Println("Health check successful!")
                }
                bodyBytes, err := io.ReadAll(resp.Body)
                if err != nil {
                        fmt.Println("io.ReadAll error:", err)
                        os.Exit(3)
                }
                os.Exit(0)
        }
        fmt.Printf("Health check failed with status: %d\n", resp.StatusCode)
        os.Exit(4)
}
