package awscloudwatch


// The position of the legend on a GraphWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.addWidgets(cloudwatch.NewGraphWidget(&graphWidgetProps{
//   	// ...
//
//   	legendPosition: cloudwatch.legendPosition_RIGHT,
//   }))
//
type LegendPosition string

const (
	// Legend appears below the graph (default).
	LegendPosition_BOTTOM LegendPosition = "BOTTOM"
	// Add shading above the annotation.
	LegendPosition_RIGHT LegendPosition = "RIGHT"
	// Add shading below the annotation.
	LegendPosition_HIDDEN LegendPosition = "HIDDEN"
)

