package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionSourceProperty := &ActionSourceProperty{
//   	SourceUri: jsii.String("sourceUri"),
//
//   	// the properties below are optional
//   	SourceId: jsii.String("sourceId"),
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-action-actionsource.html
//
type CfnAction_ActionSourceProperty struct {
	// The URI of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-action-actionsource.html#cfn-sagemaker-action-actionsource-sourceuri
	//
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// The ID of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-action-actionsource.html#cfn-sagemaker-action-actionsource-sourceid
	//
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The type of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-action-actionsource.html#cfn-sagemaker-action-actionsource-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

