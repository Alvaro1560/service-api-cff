package upload_metadata

import "service-api-cff/pkg/indra/upload_metadata"

type RequestGenerate struct {
	InputData string `json:"input_data"`
	TypeInput string `json:"type_input"`
}

type RequestProcess struct {
	Metadata []upload_metadata.Metadata `json:"metadata"`
}

type ResponseUploadMetadata struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
	Msg   string `json:"msg"`
}
