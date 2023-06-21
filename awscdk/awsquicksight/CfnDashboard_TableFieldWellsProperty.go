package awsquicksight


// The field wells for a table visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnDashboard_TableFieldWellsProperty struct {
	// The aggregated field well for the table.
	TableAggregatedFieldWells interface{} `field:"optional" json:"tableAggregatedFieldWells" yaml:"tableAggregatedFieldWells"`
	// The unaggregated field well for the table.
	TableUnaggregatedFieldWells interface{} `field:"optional" json:"tableUnaggregatedFieldWells" yaml:"tableUnaggregatedFieldWells"`
}

