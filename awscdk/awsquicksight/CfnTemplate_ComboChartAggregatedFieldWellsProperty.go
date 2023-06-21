package awsquicksight


// The aggregated field wells of a combo chart.
//
// Example:
//
//
type CfnTemplate_ComboChartAggregatedFieldWellsProperty struct {
	// The aggregated `BarValues` field well of a combo chart.
	BarValues interface{} `field:"optional" json:"barValues" yaml:"barValues"`
	// The aggregated category field wells of a combo chart.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The aggregated colors field well of a combo chart.
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The aggregated `LineValues` field well of a combo chart.
	LineValues interface{} `field:"optional" json:"lineValues" yaml:"lineValues"`
}

