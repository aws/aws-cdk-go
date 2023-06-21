package awsquicksight


// The aggregated field well for the pivot table.
//
// Example:
//
//
type CfnAnalysis_PivotTableAggregatedFieldWellsProperty struct {
	// The columns field well for a pivot table.
	//
	// Values are grouped by columns fields.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well for a pivot table.
	//
	// Values are grouped by rows fields.
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well for a pivot table.
	//
	// Values are aggregated based on rows and columns fields.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

