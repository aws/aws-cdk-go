package awsquicksight


// The field wells of the visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnTemplate_ComboChartFieldWellsProperty struct {
	// The aggregated field wells of a combo chart.
	//
	// Combo charts only have aggregated field wells. Columns in a combo chart are aggregated by category.
	ComboChartAggregatedFieldWells interface{} `field:"optional" json:"comboChartAggregatedFieldWells" yaml:"comboChartAggregatedFieldWells"`
}

