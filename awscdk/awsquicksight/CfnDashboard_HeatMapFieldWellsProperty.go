package awsquicksight


// The field well configuration of a heat map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnDashboard_HeatMapFieldWellsProperty struct {
	// The aggregated field wells of a heat map.
	HeatMapAggregatedFieldWells interface{} `field:"optional" json:"heatMapAggregatedFieldWells" yaml:"heatMapAggregatedFieldWells"`
}

