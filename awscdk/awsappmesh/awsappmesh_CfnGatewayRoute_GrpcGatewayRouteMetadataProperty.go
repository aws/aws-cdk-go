package awsappmesh


// An object representing the metadata of the gateway route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcGatewayRouteMetadataProperty := &grpcGatewayRouteMetadataProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	invert: jsii.Boolean(false),
//   	match: &gatewayRouteMetadataMatchProperty{
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
type CfnGatewayRoute_GrpcGatewayRouteMetadataProperty struct {
	// A name for the gateway route metadata.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// The criteria for determining a metadata match.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

