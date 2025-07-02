package awsquicksight


// The field well configuration of a sankey diagram.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramaggregatedfieldwells.html
//
type CfnTemplate_SankeyDiagramAggregatedFieldWellsProperty struct {
	// The destination field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-template-sankeydiagramaggregatedfieldwells-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-template-sankeydiagramaggregatedfieldwells-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The weight field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-template-sankeydiagramaggregatedfieldwells-weight
	//
	Weight interface{} `field:"optional" json:"weight" yaml:"weight"`
}

