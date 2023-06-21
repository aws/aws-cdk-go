package awsquicksight


// A histogram.
//
// For more information, see [Using histograms](https://docs.aws.amazon.com/quicksight/latest/user/histogram-charts.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
type CfnTemplate_HistogramVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration for a `HistogramVisual` .
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The subtitle that is displayed on the visual.
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}

