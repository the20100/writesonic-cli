package api

// ContentResult is returned by most content endpoints.
type ContentResult struct {
	Text string `json:"text"`
}

// LandingPage is returned by the landing-pages endpoint.
type LandingPage struct {
	Title                string `json:"title"`
	Subtitle             string `json:"subtitle"`
	MainFeatureTitle     string `json:"main_feature_title"`
	MainFeatureSubtitle  string `json:"main_feature_subtitle"`
	Feature1Title        string `json:"feature_1_title"`
	Feature1Subtitle     string `json:"feature_1_subtitle"`
	Feature2Title        string `json:"feature_2_title"`
	Feature2Subtitle     string `json:"feature_2_subtitle"`
	Feature3Title        string `json:"feature_3_title"`
	Feature3Subtitle     string `json:"feature_3_subtitle"`
	CTA                  string `json:"cta"`
	Button               string `json:"button"`
}

// ValidationError is returned on HTTP 422.
type ValidationError struct {
	Detail []struct {
		Loc  []interface{} `json:"loc"`
		Msg  string        `json:"msg"`
		Type string        `json:"type"`
	} `json:"detail"`
}

func (v *ValidationError) Error() string {
	if len(v.Detail) == 0 {
		return "validation error"
	}
	return v.Detail[0].Msg
}
