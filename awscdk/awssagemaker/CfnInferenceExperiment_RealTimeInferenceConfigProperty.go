package awssagemaker


// The infrastructure configuration for deploying the model to a real-time inference endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realTimeInferenceConfigProperty := &RealTimeInferenceConfigProperty{
//   	InstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-realtimeinferenceconfig.html
//
type CfnInferenceExperiment_RealTimeInferenceConfigProperty struct {
	// The number of instances of the type specified by `InstanceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-realtimeinferenceconfig.html#cfn-sagemaker-inferenceexperiment-realtimeinferenceconfig-instancecount
	//
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The instance type the model is deployed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-realtimeinferenceconfig.html#cfn-sagemaker-inferenceexperiment-realtimeinferenceconfig-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

