package config

type WorkConfig struct {
	Works []*Work `json:"works"`
}

type Work struct {
	Name                        string   `json:"name"`
	MaxRequestPerSecond         int      `json:"max_request_per_second"`
	ConcurrentRequests          int      `json:"concurrent_requests"`
	ConcurrentRequestsPerDomain int      `json:"concurrent_requests_per_domain"`
	ConcurrentRequestsPerIp     int      `json:"concurrent_requests_per_ip"`
	Type                        string   `json:"type"`
	Seeds                       []string `json:"seeds"`
	Rules                       []Rule   `json:"rules"`
}

type Rule struct {
	Type       string           `json:"type"`
	UrnPattern []string         `json:"urn_pattern"`
	Pipelines  []PipelineConfig `json:"pipelines"`
	Follow     Follow           `json:"follow"`
}

type PipelineConfig struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

type Follow struct {
	AllowDomains []string     `json:"allow_domains"`
	DenyDomains  []string     `json:"deny_domains"`
	AllowRules   []FollowRule `json:"allow_rules"`
	DenyRules    []FollowRule `json:"deny_rules"`
}

type FollowRule struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
