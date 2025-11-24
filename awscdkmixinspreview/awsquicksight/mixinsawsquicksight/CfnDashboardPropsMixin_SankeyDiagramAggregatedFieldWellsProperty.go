package mixinsawsquicksight


// The field well configuration of a sankey diagram.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_SankeyDiagramAggregatedFieldWellsProperty struct {
	// The destination field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-dashboard-sankeydiagramaggregatedfieldwells-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-dashboard-sankeydiagramaggregatedfieldwells-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The weight field wells of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramaggregatedfieldwells.html#cfn-quicksight-dashboard-sankeydiagramaggregatedfieldwells-weight
	//
	Weight interface{} `field:"optional" json:"weight" yaml:"weight"`
}

