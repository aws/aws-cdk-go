package awsquicksight


// An insight visual.
//
// For more information, see [Working with insights](https://docs.aws.amazon.com/quicksight/latest/user/computational-insights.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//
//
type CfnAnalysis_InsightVisualProperty struct {
	// The dataset that is used in the insight visual.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of an insight visual.
	InsightConfiguration interface{} `field:"optional" json:"insightConfiguration" yaml:"insightConfiguration"`
	// The subtitle that is displayed on the visual.
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}

