package awsquicksight


// The field well configuration of a KPI visual.
//
// Example:
//
//
type CfnAnalysis_KPIFieldWellsProperty struct {
	// The target value field wells of a KPI visual.
	TargetValues interface{} `field:"optional" json:"targetValues" yaml:"targetValues"`
	// The trend group field wells of a KPI visual.
	TrendGroups interface{} `field:"optional" json:"trendGroups" yaml:"trendGroups"`
	// The value field wells of a KPI visual.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

