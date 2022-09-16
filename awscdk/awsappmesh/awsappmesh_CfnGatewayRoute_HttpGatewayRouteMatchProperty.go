package awsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteMatchProperty := &httpGatewayRouteMatchProperty{
//   	headers: []interface{}{
//   		&httpGatewayRouteHeaderProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			invert: jsii.Boolean(false),
//   			match: &httpGatewayRouteHeaderMatchProperty{
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   				range: &gatewayRouteRangeMatchProperty{
//   					end: jsii.Number(123),
//   					start: jsii.Number(123),
//   				},
//   				regex: jsii.String("regex"),
//   				suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	hostname: &gatewayRouteHostnameMatchProperty{
//   		exact: jsii.String("exact"),
//   		suffix: jsii.String("suffix"),
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
//   }
//
type CfnGatewayRoute_HttpGatewayRouteMatchProperty struct {
	// The client request headers to match on.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The host name to match on.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The method to match on.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The path to match on.
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// `CfnGatewayRoute.HttpGatewayRouteMatchProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The query parameter to match on.
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}

