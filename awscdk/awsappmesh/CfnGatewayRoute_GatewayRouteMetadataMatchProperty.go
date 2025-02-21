package awsappmesh


// An object representing the method header to be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteMetadataMatchProperty := &GatewayRouteMetadataMatchProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html
//
type CfnGatewayRoute_GatewayRouteMetadataMatchProperty struct {
	// The exact method header to be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html#cfn-appmesh-gatewayroute-gatewayroutemetadatamatch-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The specified beginning characters of the method header to be matched on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html#cfn-appmesh-gatewayroute-gatewayroutemetadatamatch-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html#cfn-appmesh-gatewayroute-gatewayroutemetadatamatch-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The regex used to match the method header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html#cfn-appmesh-gatewayroute-gatewayroutemetadatamatch-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The specified ending characters of the method header to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.html#cfn-appmesh-gatewayroute-gatewayroutemetadatamatch-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

