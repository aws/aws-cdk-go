package awscdklocationalpha


// Properties for a route calculator.
//
// Example:
//   location.NewRouteCalculator(this, jsii.String("RouteCalculator"), &RouteCalculatorProps{
//   	RouteCalculatorName: jsii.String("MyRouteCalculator"),
//   	 // optional, defaults to a generated name
//   	DataSource: location.DataSource_ESRI,
//   })
//
// Experimental.
type RouteCalculatorProps struct {
	// Data source for the route calculator.
	// Experimental.
	DataSource DataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// A description for the route calculator.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the route calculator.
	//
	// Must be between 1 and 100 characters and contain only alphanumeric characters,
	// hyphens, periods and underscores.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	RouteCalculatorName *string `field:"optional" json:"routeCalculatorName" yaml:"routeCalculatorName"`
}

