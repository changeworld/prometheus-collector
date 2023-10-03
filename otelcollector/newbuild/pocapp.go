package pocapp

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
)

func main() {
    // Start otelcollector.exe and metricsextension.exe processes
    startProcess("otelcollector.exe")
    startProcess("metricsextension.exe")

    // Expose a health endpoint for liveness probe
    http.HandleFunc("/health", healthHandler)
    http.ListenAndServe(":8080", nil)
}

func startProcess(processName string) {
    cmd := exec.Command(processName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Start()
    if err != nil {
        fmt.Printf("Error starting %s: %v\n", processName, err)
    }
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    otelCollectorRunning := isProcessRunning("otelcollector.exe")
    metricsExtensionRunning := isProcessRunning("metricsextension.exe")

    if otelCollectorRunning && metricsExtensionRunning {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Both otelcollector.exe and metricsextension.exe are running.")
    } else {
        w.WriteHeader(http.StatusServiceUnavailable)
        fmt.Fprintln(w, "One or both of the processes are not running.")
    }
}

func isProcessRunning(processName string) bool {
    cmd := exec.Command("pgrep", processName)
    err := cmd.Run()
    return err == nil
}
