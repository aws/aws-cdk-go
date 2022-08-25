package awsappmesh


// Base options for all gateway route specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonGatewayRouteSpecOptions := &commonGatewayRouteSpecOptions{
//   	priority: jsii.Number(123),
//   }
//
type CommonGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

