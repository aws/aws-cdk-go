package awsquicksight


// The field wells for a pivot table visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnAnalysis_PivotTableFieldWellsProperty struct {
	// The aggregated field well for the pivot table.
	PivotTableAggregatedFieldWells interface{} `field:"optional" json:"pivotTableAggregatedFieldWells" yaml:"pivotTableAggregatedFieldWells"`
}

