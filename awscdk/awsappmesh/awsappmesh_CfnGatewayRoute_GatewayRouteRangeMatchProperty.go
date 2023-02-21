package awsappmesh


// An object that represents the range of values to match on.
//
// The first character of the range is included in the range, though the last character is not. For example, if the range specified were 1-100, only values 1-99 would be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteRangeMatchProperty := &gatewayRouteRangeMatchProperty{
//   	end: jsii.Number(123),
//   	start: jsii.Number(123),
//   }
//
type CfnGatewayRoute_GatewayRouteRangeMatchProperty struct {
	// The end of the range.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The start of the range.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

