package awsglue


// Defines how to extract values from HTTP responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   responseExtractionMappingProperty := &ResponseExtractionMappingProperty{
//   	ContentPath: jsii.String("contentPath"),
//   	HeaderKey: jsii.String("headerKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseextractionmapping.html
//
type CfnConnectionTypePropsMixin_ResponseExtractionMappingProperty struct {
	// A JSON path expression to extract a value from response body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseextractionmapping.html#cfn-glue-connectiontype-responseextractionmapping-contentpath
	//
	ContentPath *string `field:"optional" json:"contentPath" yaml:"contentPath"`
	// The name of an HTTP response header from which to extract the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-responseextractionmapping.html#cfn-glue-connectiontype-responseextractionmapping-headerkey
	//
	HeaderKey *string `field:"optional" json:"headerKey" yaml:"headerKey"`
}

