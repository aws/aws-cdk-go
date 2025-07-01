package awsappmesh


// An object that represents the requirements for a route to match HTTP requests for a virtual router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRouteMatchProperty := &HttpRouteMatchProperty{
//   	Headers: []interface{}{
//   		&HttpRouteHeaderProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Invert: jsii.Boolean(false),
//   			Match: &HeaderMatchMethodProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   				Range: &MatchRangeProperty{
//   					End: jsii.Number(123),
//   					Start: jsii.Number(123),
//   				},
//   				Regex: jsii.String("regex"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	Method: jsii.String("method"),
//   	Path: &HttpPathMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Regex: jsii.String("regex"),
//   	},
//   	Port: jsii.Number(123),
//   	Prefix: jsii.String("prefix"),
//   	QueryParameters: []interface{}{
//   		&QueryParameterProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Match: &HttpQueryParameterMatchProperty{
//   				Exact: jsii.String("exact"),
//   			},
//   		},
//   	},
//   	Scheme: jsii.String("scheme"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html
//
type CfnRoute_HttpRouteMatchProperty struct {
	// The client request headers to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The client request method to match on.
	//
	// Specify only one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The client request path to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-path
	//
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// The port number to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The client request query parameters to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-queryparameters
	//
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The client request scheme to match on.
	//
	// Specify only one. Applicable only for HTTP2 routes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproutematch.html#cfn-appmesh-route-httproutematch-scheme
	//
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

