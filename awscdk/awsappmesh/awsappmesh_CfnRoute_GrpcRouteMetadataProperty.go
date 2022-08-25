package awsappmesh


// An object that represents the match metadata for the route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteMetadataProperty := &grpcRouteMetadataProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	invert: jsii.Boolean(false),
//   	match: &grpcRouteMetadataMatchMethodProperty{
//   		exact: jsii.String("exact"),
//   		prefix: jsii.String("prefix"),
//   		range: &matchRangeProperty{
//   			end: jsii.Number(123),
//   			start: jsii.Number(123),
//   		},
//   		regex: jsii.String("regex"),
//   		suffix: jsii.String("suffix"),
//   	},
//   }
//
type CfnRoute_GrpcRouteMetadataProperty struct {
	// The name of the route.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// An object that represents the data to match from the request.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

