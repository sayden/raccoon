package parser

import "encoding/json"

type request struct {
	Zombiebook zombiebook  `json:"zombiebook"`
	Mansion    mansion `json:"mansion"`
}

//ParseRequest is called by the REST server to parse the JSON of a request
func ParseRequest(byt []byte) error {
	req := request{}
	err := json.Unmarshal(byt, &req)
	if err != nil {
		return err
	}

	generateJobs(&req.Mansion, generateZbookJob(req.Zombiebook))

	return nil
}