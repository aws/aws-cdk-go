package awsquicksight


// A sankey diagram.
//
// For more information, see [Using Sankey diagrams](https://docs.aws.amazon.com/quicksight/latest/user/sankey-diagram.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html
//
type CfnAnalysis_SankeyDiagramVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The subtitle that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramvisual.html#cfn-quicksight-analysis-sankeydiagramvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

