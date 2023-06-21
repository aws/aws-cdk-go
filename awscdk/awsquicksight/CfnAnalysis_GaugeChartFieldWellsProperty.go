package awsquicksight


// The field well configuration of a `GaugeChartVisual` .
//
// Example:
//
//
type CfnAnalysis_GaugeChartFieldWellsProperty struct {
	// The target value field wells of a `GaugeChartVisual` .
	TargetValues interface{} `field:"optional" json:"targetValues" yaml:"targetValues"`
	// The value field wells of a `GaugeChartVisual` .
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

