package main

import (
	"fmt"
	"io/ioutil"
	"os"

	hb "github.com/iuh88/hybrik-sdk-go"
)

type Config struct {
	hb.Config
}

func main() {

	url := os.Getenv("HYBRIK_URL")
	complianceDate := os.Getenv("HYBRIK_COMPLIANCE_DATE")
	oapiKey := os.Getenv("HYBRIK_OAPI_KEY")
	oapiSecret := os.Getenv("HYBRIK_OAPI_SECRET")
	authKey := os.Getenv("HYBRIK_AUTH_KEY")
	authSecret := os.Getenv("HYBRIK_AUTH_SECRET")
	oapiURL := os.Getenv("HYBRIK_BASE_URL")

	c, err := hb.NewClient(hb.Config{
		URL:            url,
		ComplianceDate: complianceDate,
		OAPIKey:        oapiKey,
		OAPISecret:     oapiSecret,
		AuthKey:        authKey,
		AuthSecret:     authSecret,
		OAPIURL:        oapiURL,
	})

	if err != nil {
		fmt.Println("Error calling NewClient:", err)
		return
	}

	data, err := ioutil.ReadFile("example.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	jobJSON := string(data)

	jobID, err := c.QueueJob(jobJSON)
	if err != nil {
		fmt.Println("Error calling connect:", err)
		return
	}

	fmt.Println("jobID: %w", jobID)

	jobInfo, err := c.GetJobInfo(jobID)
	if err != nil {
		fmt.Println("Error calling GetJobInfo:", err)
		return
	}

	fmt.Print(jobInfo)
}
