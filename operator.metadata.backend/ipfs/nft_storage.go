package ipfs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NftStorageCli struct {
	token    string
	endpoint string
	hc       *http.Client
}

type PinStatus string

// List of PinStatus
const (
	QUEUED_PinStatus  PinStatus = "queued"
	PINNING_PinStatus PinStatus = "pinning"
	PINNED_PinStatus  PinStatus = "pinned"
	FAILED_PinStatus  PinStatus = "failed"
)

type Pin struct {
	Cid     string       `json:"cid,omitempty"`
	Name    string       `json:"name,omitempty"`
	Meta    *interface{} `json:"meta,omitempty"`
	Status  *PinStatus   `json:"status,omitempty"`
	Created *time.Time   `json:"created,omitempty"`
	Size    float64      `json:"size,omitempty"`
}

type FilesInner struct {
	Name  string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
}
type Deal struct {
	BatchRootCid string `json:"batchRootCid,omitempty"`
	// This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ.
	LastChange string `json:"lastChange"`
	// Miner ID
	Miner string `json:"miner,omitempty"`
	// Filecoin network for this Deal
	Network string `json:"network,omitempty"`
	// Piece CID string
	PieceCid string `json:"pieceCid,omitempty"`
	// Deal status
	Status string `json:"status"`
	// Deal status description.
	StatusText string `json:"statusText,omitempty"`
	// Identifier for the deal stored on chain.
	ChainDealID float64 `json:"chainDealID,omitempty"`
	// This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ.
	DealActivation string `json:"dealActivation,omitempty"`
	// This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ.
	DealExpiration string `json:"dealExpiration,omitempty"`
}
type Nft struct {
	Cid string `json:"cid,omitempty"`
	// Size in bytes of the NFT data.
	Size    float64    `json:"size,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	// MIME type of the uploaded file or 'directory' when uploading multiple files.
	Type_ string `json:"type,omitempty"`
	// Name of the JWT token used to create this NFT.
	Scope string        `json:"scope,omitempty"`
	Pin   *Pin          `json:"pin,omitempty"`
	Files *[]FilesInner `json:"files,omitempty"`
	Deals []Deal        `json:"deals,omitempty"`
}
type UploadResponse struct {
	Ok    bool `json:"ok,omitempty"`
	Value *Nft `json:"value,omitempty"`
}

type GetResponse struct {
	Ok    bool `json:"ok,omitempty"`
	Value *Nft `json:"value,omitempty"`
}

type ErrorResponseError struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
}
type ErrorResponse struct {
	Ok    bool                `json:"ok,omitempty"`
	Error *ErrorResponseError `json:"error,omitempty"`
}

// type NftStorageCli interface {
// 	Upload(context.Context, io.Reader, ...interface{}) (*UploadResponse, error)
// 	Status(context.Context, string) (*GetResponse, error)
// }

func NewNftStorageClient(endpoint, token string) *NftStorageCli {
	if len(endpoint) == 0 {
		endpoint = "https://api.nft.storage"
	}
	return &NftStorageCli{
		token:    token,
		endpoint: endpoint,
		hc:       &http.Client{},
	}
}

// Uoload implements NftStorageCli
func (nsc *NftStorageCli) Upload(ctx context.Context, file io.Reader, options ...interface{}) (*UploadResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", nsc.endpoint+"/upload", file)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "*/*")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", nsc.token))
	res, err := nsc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response status: %d", res.StatusCode)
	}
	d := json.NewDecoder(res.Body)
	var out UploadResponse
	err = d.Decode(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (nsc *NftStorageCli) Status(ctx context.Context, cid string) (*GetResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", nsc.endpoint+fmt.Sprintf("/%s", cid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", nsc.token))
	res, err := nsc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response status: %d", res.StatusCode)
	}
	d := json.NewDecoder(res.Body)
	var out GetResponse
	err = d.Decode(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
