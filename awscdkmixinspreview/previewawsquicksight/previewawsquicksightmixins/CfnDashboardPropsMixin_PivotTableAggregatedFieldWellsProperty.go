package previewawsquicksightmixins


// The aggregated field well for the pivot table.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_PivotTableAggregatedFieldWellsProperty struct {
	// The columns field well for a pivot table.
	//
	// Values are grouped by columns fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableaggregatedfieldwells.html#cfn-quicksight-dashboard-pivottableaggregatedfieldwells-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well for a pivot table.
	//
	// Values are grouped by rows fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableaggregatedfieldwells.html#cfn-quicksight-dashboard-pivottableaggregatedfieldwells-rows
	//
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well for a pivot table.
	//
	// Values are aggregated based on rows and columns fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableaggregatedfieldwells.html#cfn-quicksight-dashboard-pivottableaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

