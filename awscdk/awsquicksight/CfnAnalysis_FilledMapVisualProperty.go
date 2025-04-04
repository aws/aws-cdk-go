package awsquicksight


// A filled map.
//
// For more information, see [Creating filled maps](https://docs.aws.amazon.com/quicksight/latest/user/filled-maps.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html
//
type CfnAnalysis_FilledMapVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration settings of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The column hierarchy that is used during drill-downs and drill-ups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-columnhierarchies
	//
	ColumnHierarchies interface{} `field:"optional" json:"columnHierarchies" yaml:"columnHierarchies"`
	// The conditional formatting of a `FilledMapVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-conditionalformatting
	//
	ConditionalFormatting interface{} `field:"optional" json:"conditionalFormatting" yaml:"conditionalFormatting"`
	// The subtitle that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapvisual.html#cfn-quicksight-analysis-filledmapvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

