package awsquicksight


// The field wells of a tree map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnAnalysis_TreeMapFieldWellsProperty struct {
	// The aggregated field wells of a tree map.
	TreeMapAggregatedFieldWells interface{} `field:"optional" json:"treeMapAggregatedFieldWells" yaml:"treeMapAggregatedFieldWells"`
}

