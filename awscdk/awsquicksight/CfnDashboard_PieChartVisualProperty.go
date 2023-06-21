package awsquicksight


// A pie or donut chart.
//
// The `PieChartVisual` structure describes a visual that is a member of the pie chart family.
//
// The following charts can be described by using this structure:
//
// - Pie charts
// - Donut charts
//
// For more information, see [Using pie charts](https://docs.aws.amazon.com/quicksight/latest/user/pie-chart.html) in the *Amazon QuickSight User Guide* .
//
// For more information, see [Using donut charts](https://docs.aws.amazon.com/quicksight/latest/user/donut-chart.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
type CfnDashboard_PieChartVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of a pie chart.
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The column hierarchy that is used during drill-downs and drill-ups.
	ColumnHierarchies interface{} `field:"optional" json:"columnHierarchies" yaml:"columnHierarchies"`
	// The subtitle that is displayed on the visual.
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}

