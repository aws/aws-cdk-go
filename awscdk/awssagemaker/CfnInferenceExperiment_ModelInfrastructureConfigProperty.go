package awssagemaker


// The configuration for the infrastructure that the model will be deployed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInfrastructureConfigProperty := &ModelInfrastructureConfigProperty{
//   	InfrastructureType: jsii.String("infrastructureType"),
//   	RealTimeInferenceConfig: &RealTimeInferenceConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html
//
type CfnInferenceExperiment_ModelInfrastructureConfigProperty struct {
	// The inference option to which to deploy your model. Possible values are the following:.
	//
	// - `RealTime` : Deploy to real-time inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-infrastructuretype
	//
	InfrastructureType *string `field:"required" json:"infrastructureType" yaml:"infrastructureType"`
	// The infrastructure configuration for deploying the model to real-time inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-realtimeinferenceconfig
	//
	RealTimeInferenceConfig interface{} `field:"required" json:"realTimeInferenceConfig" yaml:"realTimeInferenceConfig"`
}

