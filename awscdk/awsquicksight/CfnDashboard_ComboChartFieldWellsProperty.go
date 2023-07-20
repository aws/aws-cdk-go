package awsquicksight


// The field wells of the visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartfieldwells.html
//
type CfnDashboard_ComboChartFieldWellsProperty struct {
	// The aggregated field wells of a combo chart.
	//
	// Combo charts only have aggregated field wells. Columns in a combo chart are aggregated by category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartfieldwells.html#cfn-quicksight-dashboard-combochartfieldwells-combochartaggregatedfieldwells
	//
	ComboChartAggregatedFieldWells interface{} `field:"optional" json:"comboChartAggregatedFieldWells" yaml:"comboChartAggregatedFieldWells"`
}

