package awscloudwatch


// The position of the legend on a GraphWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewGraphWidget(&GraphWidgetProps{
//   	// ...
//
//   	LegendPosition: cloudwatch.LegendPosition_RIGHT,
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

