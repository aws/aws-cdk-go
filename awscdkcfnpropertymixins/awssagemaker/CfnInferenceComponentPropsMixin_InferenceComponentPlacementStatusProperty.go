package awssagemaker


// The number of inference component copies currently placed on instances of a given type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceComponentPlacementStatusProperty := &InferenceComponentPlacementStatusProperty{
//   	CurrentCopyCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentplacementstatus.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentPlacementStatusProperty struct {
	// The number of copies for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentplacementstatus.html#cfn-sagemaker-inferencecomponent-inferencecomponentplacementstatus-currentcopycount
	//
	CurrentCopyCount *float64 `field:"optional" json:"currentCopyCount" yaml:"currentCopyCount"`
	// An ML compute instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentplacementstatus.html#cfn-sagemaker-inferencecomponent-inferencecomponentplacementstatus-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
}

