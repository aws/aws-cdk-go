package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactSourceTypeProperty := &ArtifactSourceTypeProperty{
//   	SourceIdType: jsii.String("sourceIdType"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsourcetype.html
//
type CfnArtifact_ArtifactSourceTypeProperty struct {
	// The type of ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsourcetype.html#cfn-sagemaker-artifact-artifactsourcetype-sourceidtype
	//
	SourceIdType *string `field:"required" json:"sourceIdType" yaml:"sourceIdType"`
	// The ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-artifact-artifactsourcetype.html#cfn-sagemaker-artifact-artifactsourcetype-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

