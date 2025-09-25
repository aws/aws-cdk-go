package awsquicksight


// The field well configuration of a waterfall visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartaggregatedfieldwells.html
//
type CfnTemplate_WaterfallChartAggregatedFieldWellsProperty struct {
	// The breakdown field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartaggregatedfieldwells.html#cfn-quicksight-template-waterfallchartaggregatedfieldwells-breakdowns
	//
	Breakdowns interface{} `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// The category field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartaggregatedfieldwells.html#cfn-quicksight-template-waterfallchartaggregatedfieldwells-categories
	//
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// The value field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartaggregatedfieldwells.html#cfn-quicksight-template-waterfallchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

