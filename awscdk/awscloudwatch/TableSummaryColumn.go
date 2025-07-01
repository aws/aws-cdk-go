package awscloudwatch


// Standard table summary columns.
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
type TableSummaryColumn string

const (
	// Minimum of all data points.
	TableSummaryColumn_MINIMUM TableSummaryColumn = "MINIMUM"
	// Maximum of all data points.
	TableSummaryColumn_MAXIMUM TableSummaryColumn = "MAXIMUM"
	// Sum of all data points.
	TableSummaryColumn_SUM TableSummaryColumn = "SUM"
	// Average of all data points.
	TableSummaryColumn_AVERAGE TableSummaryColumn = "AVERAGE"
)

