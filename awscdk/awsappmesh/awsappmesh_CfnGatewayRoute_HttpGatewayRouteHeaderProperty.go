package awsappmesh


// An object that represents the HTTP header in the gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteHeaderProperty := &httpGatewayRouteHeaderProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	invert: jsii.Boolean(false),
//   	match: &httpGatewayRouteHeaderMatchProperty{
//   		exact: jsii.String("exact"),
//   		prefix: jsii.String("prefix"),
//   		range: &gatewayRouteRangeMatchProperty{
//   			end: jsii.Number(123),
//   			start: jsii.Number(123),
//   		},
//   		regex: jsii.String("regex"),
//   		suffix: jsii.String("suffix"),
//   	},
//   }
//
type CfnGatewayRoute_HttpGatewayRouteHeaderProperty struct {
	// A name for the HTTP header in the gateway route that will be matched on.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// An object that represents the method and value to match with the header value sent in a request.
	//
	// Specify one match method.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

