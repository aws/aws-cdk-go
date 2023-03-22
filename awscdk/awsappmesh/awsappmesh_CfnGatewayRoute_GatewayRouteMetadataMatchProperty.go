package awsappmesh


// An object representing the method header to be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteMetadataMatchProperty := &gatewayRouteMetadataMatchProperty{
//   	exact: jsii.String("exact"),
//   	prefix: jsii.String("prefix"),
//   	range: &gatewayRouteRangeMatchProperty{
//   		end: jsii.Number(123),
//   		start: jsii.Number(123),
//   	},
//   	regex: jsii.String("regex"),
//   	suffix: jsii.String("suffix"),
//   }
//
type CfnGatewayRoute_GatewayRouteMetadataMatchProperty struct {
	// The exact method header to be matched on.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The specified beginning characters of the method header to be matched on.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An object that represents the range of values to match on.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The regex used to match the method header.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The specified ending characters of the method header to match on.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

