package awslocation


// Properties for defining a `CfnRouteCalculator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteCalculatorProps := &cfnRouteCalculatorProps{
//   	calculatorName: jsii.String("calculatorName"),
//   	dataSource: jsii.String("dataSource"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	pricingPlan: jsii.String("pricingPlan"),
//   }
//
type CfnRouteCalculatorProps struct {
	// The name of the route calculator resource.
	//
	// Requirements:
	//
	// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique route calculator resource name.
	// - No spaces allowed. For example, `ExampleRouteCalculator` .
	CalculatorName *string `field:"required" json:"calculatorName" yaml:"calculatorName"`
	// Specifies the data provider of traffic and road network data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri`
	// - `Here`
	//
	// For more information about data providers, see the [Amazon Location Service data providers page](https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html) .
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The optional description for the route calculator resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
}

