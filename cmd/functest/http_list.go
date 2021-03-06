// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/profile.proto
package functest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"

	pb "github.com/gomeet-examples/svc-profile/pb"
)

func TestHttpList(config FunctionalTestConfig) (failures []TestFailure) {
	client, serverAddr, proto, err := httpClient(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "List/HTTP", Message: fmt.Sprintf("HTTP client initialization error (%v)", err)})
		return failures
	}

	var testCaseResults []*TestCaseResult
	reqs, extras, err := testGetListRequest(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "List/HTTP", Message: fmt.Sprintf("HTTP testGetListRequest error (%v)", err)})
		return failures
	}

	for _, req := range reqs {
		url := fmt.Sprintf("%s://%s/api/v1/list", proto, serverAddr)

		// Proto to JSON
		ma := jsonpb.Marshaler{}
		sMsg, err := ma.MarshalToString(req)
		if err != nil {
			testCaseResults = append(
				testCaseResults,
				&TestCaseResult{
					req,
					nil,
					fmt.Errorf("List/HTTP POST error to marshalling the message with %s (%v) - %v", url, err, req),
				},
			)
			continue
		}

		data := bytes.NewBufferString(sMsg)

		// construct HTTP request
		httpReq, err := http.NewRequest("POST", url, data)
		if err != nil {
			testCaseResults = append(
				testCaseResults,
				&TestCaseResult{
					req,
					nil,
					fmt.Errorf("List/HTTP POST error to construct the http request with %s (%v) - %v", url, err, req),
				},
			)
			continue
		}
		httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.JsonWebToken))
		httpReq.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(httpReq)
		if err != nil {
			testCaseResults = append(
				testCaseResults,
				&TestCaseResult{
					req,
					nil,
					fmt.Errorf("List/HTTP POST error on %s (%v) - %v", url, err, req),
				},
			)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			testCaseResults = append(
				testCaseResults,
				&TestCaseResult{
					req,
					nil,
					fmt.Errorf("List/HTTP POST error on %s (%v) - %v - readAll body", url, err, req),
				},
			)
			continue
		}

		var httpError HttpError
		err = json.Unmarshal(body, &httpError)
		if err == nil && (httpError.Code != 0 || httpError.Error != "") {
			testCaseResults = append(
				testCaseResults,
				&TestCaseResult{
					req,
					nil,
					fmt.Errorf("List/HTTP POST error on %s (Code: %d, Error: %s) - %v", url, httpError.Code, httpError.Error, req),
				},
			)
			continue
		}

		res := &pb.ProfileList{}
		err = jsonpb.UnmarshalString(string(body), res)
		testCaseResults = append(testCaseResults, &TestCaseResult{req, res, err})
	}

	return testListResponse(config, FUNCTEST_HTTP, testCaseResults, extras)
}
