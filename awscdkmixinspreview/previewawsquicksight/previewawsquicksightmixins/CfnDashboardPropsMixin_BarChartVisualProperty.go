package previewawsquicksightmixins


// A bar chart.
//
// The `BarChartVisual` structure describes a visual that is a member of the bar chart family. The following charts can be described using this structure:
//
// - Horizontal bar chart
// - Vertical bar chart
// - Horizontal stacked bar chart
// - Vertical stacked bar chart
// - Horizontal stacked 100% bar chart
// - Vertical stacked 100% bar chart
//
// For more information, see [Using bar charts](https://docs.aws.amazon.com/quicksight/latest/user/bar-charts.html) in the *Amazon Quick Suite User Guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html
//
type CfnDashboardPropsMixin_BarChartVisualProperty struct {
	// The list of custom actions that are configured for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration settings of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The column hierarchy that is used during drill-downs and drill-ups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-columnhierarchies
	//
	ColumnHierarchies interface{} `field:"optional" json:"columnHierarchies" yaml:"columnHierarchies"`
	// The subtitle that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartvisual.html#cfn-quicksight-dashboard-barchartvisual-visualid
	//
	VisualId *string `field:"optional" json:"visualId" yaml:"visualId"`
}

