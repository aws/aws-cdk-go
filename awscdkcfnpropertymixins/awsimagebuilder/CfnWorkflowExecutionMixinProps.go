package awsimagebuilder


// Properties for CfnWorkflowExecutionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkflowExecutionMixinProps := &CfnWorkflowExecutionMixinProps{
//   	ImageBuildVersionArn: jsii.String("imageBuildVersionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflowexecution.html
//
type CfnWorkflowExecutionMixinProps struct {
	// The Amazon Resource Name (ARN) of the image resource build version that the specified runtime instance of the workflow created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflowexecution.html#cfn-imagebuilder-workflowexecution-imagebuildversionarn
	//
	ImageBuildVersionArn *string `field:"optional" json:"imageBuildVersionArn" yaml:"imageBuildVersionArn"`
}

