package awssagemaker


// The source type, ID, and URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sourceProperty := &SourceProperty{
//   	SourceId: jsii.String("sourceId"),
//   	SourceType: jsii.String("sourceType"),
//   	SourceUri: jsii.String("sourceUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html
//
type CfnContextPropsMixin_SourceProperty struct {
	// The ID of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourceid
	//
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The type of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// The URI of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourceuri
	//
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
}

