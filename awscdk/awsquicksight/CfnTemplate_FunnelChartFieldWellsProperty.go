package awsquicksight


// The field well configuration of a `FunnelChartVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartfieldwells.html
//
type CfnTemplate_FunnelChartFieldWellsProperty struct {
	// The field well configuration of a `FunnelChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartfieldwells.html#cfn-quicksight-template-funnelchartfieldwells-funnelchartaggregatedfieldwells
	//
	FunnelChartAggregatedFieldWells interface{} `field:"optional" json:"funnelChartAggregatedFieldWells" yaml:"funnelChartAggregatedFieldWells"`
}

