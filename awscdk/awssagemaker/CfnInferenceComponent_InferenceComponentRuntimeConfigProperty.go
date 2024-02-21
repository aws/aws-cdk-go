package awssagemaker


// Runtime settings for a model that is deployed with an inference component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentRuntimeConfigProperty := &InferenceComponentRuntimeConfigProperty{
//   	CopyCount: jsii.Number(123),
//   	CurrentCopyCount: jsii.Number(123),
//   	DesiredCopyCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentruntimeconfig.html
//
type CfnInferenceComponent_InferenceComponentRuntimeConfigProperty struct {
	// The number of runtime copies of the model container to deploy with the inference component.
	//
	// Each copy can serve inference requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentruntimeconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentruntimeconfig-copycount
	//
	CopyCount *float64 `field:"optional" json:"copyCount" yaml:"copyCount"`
	// The number of runtime copies of the model container that are currently deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentruntimeconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentruntimeconfig-currentcopycount
	//
	CurrentCopyCount *float64 `field:"optional" json:"currentCopyCount" yaml:"currentCopyCount"`
	// The number of runtime copies of the model container that you requested to deploy with the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentruntimeconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentruntimeconfig-desiredcopycount
	//
	DesiredCopyCount *float64 `field:"optional" json:"desiredCopyCount" yaml:"desiredCopyCount"`
}

