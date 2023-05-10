package awsappmesh


// Base options for all route specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSpecOptionsBase := &RouteSpecOptionsBase{
//   	Priority: jsii.Number(123),
//   }
//
// Experimental.
type RouteSpecOptionsBase struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

