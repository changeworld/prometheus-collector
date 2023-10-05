package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	mac := os.Getenv("MAC")

	if mac == "true" {
		// Checking if TokenConfig file exists
		_, err := os.Stat("/etc/mdsd.d/config-cache/metricsextension/TokenConfig.json")
		if os.IsNotExist(err) {
			azmonContainerStartTime, err := ioutil.ReadFile("/opt/microsoft/liveness/azmon-container-start-time")
			if err == nil {
				epochTimeNow := time.Now().Unix()
				azmonContainerStartTimeInt, _ := strconv.ParseInt(strings.TrimSpace(string(azmonContainerStartTime)), 10, 64)
				duration := epochTimeNow - azmonContainerStartTimeInt
				durationInMinutes := duration / 60

				if durationInMinutes%5 == 0 {
					fmt.Printf("%s No configuration present for the AKS resource\n", time.Now().Format("2006-01-02T15:04:05"))
				}

				if durationInMinutes > 15 {
					fmt.Println("No configuration present for the AKS resource")
					os.Exit(1)
				}
			}
		} else {
			// Check if ME is not running
			cmd := exec.Command("ps", "-ef")
			output, _ := cmd.Output()
			if !strings.Contains(string(output), "MetricsExt") {
				fmt.Println("Metrics Extension is not running (configuration exists)")
				os.Exit(1)
			}

			// Check if MDSD is not running
			cmd = exec.Command("ps", "-ef")
			output, _ = cmd.Output()
			if !strings.Contains(string(output), "mdsd") {
				fmt.Println("mdsd is not running (configuration exists)")
				os.Exit(1)
			}
		}

		// Adding liveness probe check for AMCS config update by MDSD
		if _, err := os.Stat("/opt/inotifyoutput-mdsd-config.txt"); err == nil {
			fmt.Println("inotifyoutput-mdsd-config.txt has been updated - mdsd config changed")
			os.Exit(1)
		}
	} else {
		// Non-MAC mode
		// Check if ME is not running
		cmd := exec.Command("ps", "-ef")
		output, _ := cmd.Output()
		if !strings.Contains(string(output), "MetricsExt") {
			fmt.Println("Metrics Extension is not running")
			os.Exit(1)
		}

		// Check if cert files have changed
		if _, err := os.Stat("/etc/config/settings/akv"); err == nil {
			if _, err := os.Stat("/opt/akv-copy/akv"); err == nil {
				cmd := exec.Command("diff", "-r", "-q", "/etc/config/settings/akv", "/opt/akv-copy/akv")
				output, _ := cmd.CombinedOutput()
				if len(output) > 0 {
					fmt.Println("A Metrics Account certificate has changed")
					os.Exit(1)
				}
			}
		}
	}

	// Check if otelcollector is running
	cmd := exec.Command("ps", "-ef")
	output, _ := cmd.Output()
	if !strings.Contains(string(output), "otelcollector") {
		fmt.Println("OpenTelemetryCollector is not running")
		os.Exit(1)
	}

	// Check if config changed
	if _, err := os.Stat("/opt/inotifyoutput.txt"); err == nil {
		fmt.Println("inotifyoutput.txt has been updated - config changed")
		os.Exit(1)
	}
}
