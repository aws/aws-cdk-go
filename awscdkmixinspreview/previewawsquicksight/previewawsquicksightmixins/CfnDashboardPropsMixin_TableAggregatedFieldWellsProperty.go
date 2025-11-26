package previewawsquicksightmixins


// The aggregated field well for the table.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_TableAggregatedFieldWellsProperty struct {
	// The group by field well for a pivot table.
	//
	// Values are grouped by group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableaggregatedfieldwells.html#cfn-quicksight-dashboard-tableaggregatedfieldwells-groupby
	//
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// The values field well for a pivot table.
	//
	// Values are aggregated based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableaggregatedfieldwells.html#cfn-quicksight-dashboard-tableaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

