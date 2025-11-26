package previewawsappmeshmixins


// An object that represents the match metadata for the route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   grpcRouteMetadataProperty := &GrpcRouteMetadataProperty{
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
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutemetadata.html
//
type CfnRoutePropsMixin_GrpcRouteMetadataProperty struct {
	// Specify `True` to match anything except the match criteria.
	//
	// The default value is `False` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutemetadata.html#cfn-appmesh-route-grpcroutemetadata-invert
	//
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// An object that represents the data to match from the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutemetadata.html#cfn-appmesh-route-grpcroutemetadata-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// The name of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutemetadata.html#cfn-appmesh-route-grpcroutemetadata-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

