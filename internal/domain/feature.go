package domain

type Feature struct {
	CvId      string `json:"cvId"`
	FeatureId int    `json:"featureId"`
	Value     string `json:"value"`
}

type FullfieldFeature struct {
	FeatureName  string  `json:"featureName"`
	PriorityName string  `json:"priorityName"`
	Coefficient  float32 `json:"coefficient"`
}
