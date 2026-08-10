package awssagemaker


// The source type, ID, and URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	SourceUri: jsii.String("sourceUri"),
//
//   	// the properties below are optional
//   	SourceId: jsii.String("sourceId"),
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html
//
type CfnContext_SourceProperty struct {
	// The URI of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourceuri
	//
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// The ID of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourceid
	//
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The type of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-source.html#cfn-sagemaker-context-source-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

