package awsappmesh


// An object that represents the method and value to match with the header value sent in a request.
//
// Specify one match method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteHeaderMatchProperty := &HttpGatewayRouteHeaderMatchProperty{
//   	Exact: jsii.String("exact"),
//   	Prefix: jsii.String("prefix"),
//   	Range: &GatewayRouteRangeMatchProperty{
//   		End: jsii.Number(123),
//   		Start: jsii.Number(123),
//   	},
//   	Regex: jsii.String("regex"),
//   	Suffix: jsii.String("suffix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html
//
type CfnGatewayRoute_HttpGatewayRouteHeaderMatchProperty struct {
	// The value sent by the client must match the specified value exactly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html#cfn-appmesh-gatewayroute-httpgatewayrouteheadermatch-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The value sent by the client must begin with the specified characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html#cfn-appmesh-gatewayroute-httpgatewayrouteheadermatch-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html#cfn-appmesh-gatewayroute-httpgatewayrouteheadermatch-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The value sent by the client must include the specified characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html#cfn-appmesh-gatewayroute-httpgatewayrouteheadermatch-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The value sent by the client must end with the specified characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-httpgatewayrouteheadermatch.html#cfn-appmesh-gatewayroute-httpgatewayrouteheadermatch-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

