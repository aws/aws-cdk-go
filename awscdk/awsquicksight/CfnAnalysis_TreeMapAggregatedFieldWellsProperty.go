package awsquicksight


// Aggregated field wells of a tree map.
//
// Example:
//
//
type CfnAnalysis_TreeMapAggregatedFieldWellsProperty struct {
	// The color field well of a tree map.
	//
	// Values are grouped by aggregations based on group by fields.
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The group by field well of a tree map.
	//
	// Values are grouped based on group by fields.
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// The size field well of a tree map.
	//
	// Values are aggregated based on group by fields.
	Sizes interface{} `field:"optional" json:"sizes" yaml:"sizes"`
}

