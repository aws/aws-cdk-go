package awsappmesh


// An object that represents the match metadata for the route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteMetadataProperty := &GrpcRouteMetadataProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Invert: jsii.Boolean(false),
//   	Match: &GrpcRouteMetadataMatchMethodProperty{
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   		Range: &MatchRangeProperty{
//   			End: jsii.Number(123),
//   			Start: jsii.Number(123),
//   		},
//   		Regex: jsii.String("regex"),
//   		Suffix: jsii.String("suffix"),
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

