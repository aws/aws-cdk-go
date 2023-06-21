package awsquicksight


// The aggregated field wells of a heat map.
//
// Example:
//
//
type CfnTemplate_HeatMapAggregatedFieldWellsProperty struct {
	// The columns field well of a heat map.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well of a heat map.
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well of a heat map.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

