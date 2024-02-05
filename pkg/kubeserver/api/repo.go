package api

import "yunion.io/x/onecloud/pkg/apis"

type RepoType string

const (
	RepoTypeInternal RepoType = "internal"
	RepoTypeExternal RepoType = "external"
)

type RepoBackend string

const (
	RepoBackendCommon = "common"
	RepoBackendNexus  = "nexus"
)

type RepoCreateInput struct {
	apis.StatusInfrasResourceBaseCreateInput

	// Repo URL
	// required: true
	// example: http://mirror.azure.cn/kubernetes/charts
	Url string `json:"url"`

	// Repo username
	// required: false
	Username string `json:"username"`

	// Repo password
	Password string `json:"password"`

	// Repo type
	// enum: internal, external
	Type string `json:"type"`

	// Repo backend
	// enum: common, nexus
	Backend RepoBackend `json:"backend"`
}

type RepoListInput struct {
	apis.StatusInfrasResourceBaseListInput

	Type string `json:"type"`
	Url  string `json:"url"`
}

type RepoDetail struct {
	apis.StatusInfrasResourceBaseDetails

	Url          string `json:"url"`
	Type         string `json:"type"`
	ReleaseCount int    `json:"release_count"`
}

type RepoDownloadChartInput struct {
	ChartName string `json:"chart_name"`
	Version   string `json:"version"`
}
