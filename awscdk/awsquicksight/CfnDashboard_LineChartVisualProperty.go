package awsquicksight


// A line chart.
//
// For more information, see [Using line charts](https://docs.aws.amazon.com/quicksight/latest/user/line-charts.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html
//
type CfnDashboard_LineChartVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The column hierarchy that is used during drill-downs and drill-ups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-columnhierarchies
	//
	ColumnHierarchies interface{} `field:"optional" json:"columnHierarchies" yaml:"columnHierarchies"`
	// The subtitle that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

