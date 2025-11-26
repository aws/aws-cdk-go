package previewawsappmeshmixins


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpGatewayRouteMatchProperty := &HttpGatewayRouteMatchProperty{
//   	Headers: []interface{}{
//   		&HttpGatewayRouteHeaderProperty{
//   			Invert: jsii.Boolean(false),
//   			Match: &HttpGatewayRouteHeaderMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   				Range: &GatewayRouteRangeMatchProperty{
//   					End: jsii.Number(123),
//   					Start: jsii.Number(123),
//   				},
//   				Regex: jsii.String("regex"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Hostname: &GatewayRouteHostnameMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Suffix: jsii.String("suffix"),
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
//   			Match: &HttpQueryParameterMatchProperty{
//   				Exact: jsii.String("exact"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html
//
type CfnGatewayRoutePropsMixin_HttpGatewayRouteMatchProperty struct {
	// The client request headers to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The host name to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-hostname
	//
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The method to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The path to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-path
	//
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// The port number to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies the path to match requests with.
	//
	// This parameter must always start with `/` , which by itself matches all requests to the virtual service name. You can also match for path-based routing of requests. For example, if your virtual service name is `my-service.local` and you want the route to match requests to `my-service.local/metrics` , your prefix should be `/metrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The query parameter to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayroutematch.html#cfn-appmesh-gatewayroute-httpgatewayroutematch-queryparameters
	//
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}

