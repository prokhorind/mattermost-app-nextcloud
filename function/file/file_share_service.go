package file

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type FileShareServiceImpl struct {
	Url   string
	Token string
}

func (s FileShareServiceImpl) GetSharesInfo(filePath string, shareType int32) (*FileShareModel, error) {
	shares, err := s.GetAllUserShares()

	if err != nil {
		return nil, err
	}

	for _, el := range shares.Data.Element {
		if el.Path == filePath {
			return &el, nil
		}
	}

	return s.CreateUserShare(filePath, shareType)
}

func (s FileShareServiceImpl) GetAllUserShares() (*SharedFilesResponseBody, error) {

	req, _ := http.NewRequest("GET", s.Url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	req.Header.Set("OCS-APIRequest", "true")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	xmlResp := SharedFilesResponseBody{}
	xml.NewDecoder(resp.Body).Decode(&xmlResp)

	return &xmlResp, err
}

func (s FileShareServiceImpl) CreateUserShare(filePath string, shareType int32) (*FileShareModel, error) {
	payload := FileShareRequestBody{filePath, shareType}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", s.Url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	req.Header.Set("OCS-APIRequest", "true")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	xmlResp := SharedFileResponseBody{}
	xml.NewDecoder(resp.Body).Decode(&xmlResp)

	return &xmlResp.Data, err
}
