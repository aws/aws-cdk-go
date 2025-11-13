package interfacesawslocation


// A reference to a RouteCalculator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeCalculatorReference := &RouteCalculatorReference{
//   	CalculatorName: jsii.String("calculatorName"),
//   	RouteCalculatorArn: jsii.String("routeCalculatorArn"),
//   }
//
type RouteCalculatorReference struct {
	// The CalculatorName of the RouteCalculator resource.
	CalculatorName *string `field:"required" json:"calculatorName" yaml:"calculatorName"`
	// The ARN of the RouteCalculator resource.
	RouteCalculatorArn *string `field:"required" json:"routeCalculatorArn" yaml:"routeCalculatorArn"`
}

