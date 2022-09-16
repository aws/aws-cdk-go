package awsappmesh


// An object that represents the requirements for a route to match HTTP requests for a virtual router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRouteMatchProperty := &httpRouteMatchProperty{
//   	headers: []interface{}{
//   		&httpRouteHeaderProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			invert: jsii.Boolean(false),
//   			match: &headerMatchMethodProperty{
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   				range: &matchRangeProperty{
//   					end: jsii.Number(123),
//   					start: jsii.Number(123),
//   				},
//   				regex: jsii.String("regex"),
//   				suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	method: jsii.String("method"),
//   	path: &httpPathMatchProperty{
//   		exact: jsii.String("exact"),
//   		regex: jsii.String("regex"),
//   	},
//   	port: jsii.Number(123),
//   	prefix: jsii.String("prefix"),
//   	queryParameters: []interface{}{
//   		&queryParameterProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			match: &httpQueryParameterMatchProperty{
//   				exact: jsii.String("exact"),
//   			},
//   		},
//   	},
//   	scheme: jsii.String("scheme"),
//   }
//
type CfnRoute_HttpRouteMatchProperty struct {
	// The client request headers to match on.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The client request method to match on.
	//
	// Specify only one.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The client request path to match on.
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// `CfnRoute.HttpRouteMatchProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The client request query parameters to match on.
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The client request scheme to match on.
	//
	// Specify only one. Applicable only for HTTP2 routes.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

