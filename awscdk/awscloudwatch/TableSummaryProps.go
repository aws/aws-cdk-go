package awscloudwatch


// Properties for TableWidget's summary columns.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTableWidget(&TableWidgetProps{
//   	// ...
//
//   	Summary: &TableSummaryProps{
//   		Columns: []tableSummaryColumn{
//   			cloudwatch.*tableSummaryColumn_AVERAGE,
//   		},
//   		HideNonSummaryColumns: jsii.Boolean(true),
//   		Sticky: jsii.Boolean(true),
//   	},
//   }))
//
type TableSummaryProps struct {
	// Summary columns.
	// Default: - No summary columns will be shown.
	//
	Columns *[]TableSummaryColumn `field:"optional" json:"columns" yaml:"columns"`
	// Prevent the columns of datapoints from being displayed, so that only the label and summary columns are displayed.
	// Default: - false.
	//
	HideNonSummaryColumns *bool `field:"optional" json:"hideNonSummaryColumns" yaml:"hideNonSummaryColumns"`
	// Make the summary columns sticky, so that they remain in view while scrolling.
	// Default: - false.
	//
	Sticky *bool `field:"optional" json:"sticky" yaml:"sticky"`
}

