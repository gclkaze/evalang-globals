package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gclkaze/evalang-globals/globals/evalogger"
)

func PerformPostRequest(data map[string]string, addr string, endpoint string, logger *evalogger.ILogger, h http.Header) (bodyString string, err error) {

	// Convert to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", addr+"/"+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	//	// Set headers
	//	req.Header.Set("Content-Type", "application/json")
	///	req.Header.Set("Authorization", "Bearer YOUR_TOKEN_HERE") // Replace with your actual token
	//	req.Header.Set("Programid", "test")                       // Optional custom header

	req.Header = h
	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString = string(body)
	/*	fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(body))
	*/
	return bodyString, nil
}

func PerformPostRequestWithContent(data map[string]string, addr string, endpoint string, logger *evalogger.ILogger, h http.Header, content *string) (bodyString string, err error) {

	if content == nil {
		return "", fmt.Errorf("the content is null")
	}
	bb := bytes.NewBufferString(*content)

	// Create a new POST request
	req, err := http.NewRequest("POST", addr+"/"+endpoint, bb)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "text/plain")
	//	// Set headers
	//	req.Header.Set("Content-Type", "application/json")
	///	req.Header.Set("Authorization", "Bearer YOUR_TOKEN_HERE") // Replace with your actual token
	//	req.Header.Set("Programid", "test")                       // Optional custom header

	req.Header = h
	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return bodyString, fmt.Errorf("an error has occurred")
	}

	bodyString = string(body)
	return bodyString, nil
}

func PerformGetRequestWithFileAttachment(addr string, endpoint string, h http.Header, saveAs string) (res bool, err error) {

	req, err := http.NewRequest("GET", addr+"/"+endpoint, nil)
	if err != nil {
		return false, err
	}

	req.Header = h

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Status NOT OK %s", strconv.Itoa(resp.StatusCode))
	}

	// Create the file locally
	outFile, err := os.Create(saveAs)
	if err != nil {
		return false, err
	}
	defer outFile.Close()

	// Copy the response body to the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return false, err
	}

	/*println("File downloaded successfully.")*/
	return true, nil

}
