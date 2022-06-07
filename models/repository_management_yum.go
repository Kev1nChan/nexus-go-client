package models

type YUMSigning struct {
	Keypair    string `json:"keypair"`
	Passphrase string `json:"passphrase"`
}

type YUMAttributes struct {
	// Validate that all paths are RPMs or yum metadata
	// Enum: [PERMISSIVE STRICT]
	DeployPolicy string `json:"deployPolicy,omitempty"`

	// Specifies the repository depth where repodata folder(s) are created
	RepodataDepth int32 `json:"repodataDepth"`
}

type YUMHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage       `json:"storage"`
	YUM     *YUMAttributes `json:"yum"`
}

type YUMProxyRepository struct {
	Name          string         `json:"name"`
	Online        bool           `json:"online"`
	Storage       *Storage       `json:"storage"`
	Cleanup       *Cleanup       `json:"cleanup,omitempty"`
	Proxy         *Proxy         `json:"proxy"`
	NegativeCache *NegativeCache `json:"negativeCache"`
	HTTPClient    *HTTPClient    `json:"httpClient"`
	RoutingRule   string         `json:"routingRule,omitempty"`
	YumSigning    *YUMSigning    `json:"yumSigning"`
}

type YUMGroupRepository struct {
	Name       string     `json:"name"`
	Online     bool       `json:"online"`
	Storage    *Storage   `json:"storage"`
	Group      *Group     `json:"group"`
	YumSigning YUMSigning `json:"yumSigning"`
}
