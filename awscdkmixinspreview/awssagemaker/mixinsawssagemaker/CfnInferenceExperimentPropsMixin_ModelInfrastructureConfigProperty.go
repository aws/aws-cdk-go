package mixinsawssagemaker


// The configuration for the infrastructure that the model will be deployed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnInferenceExperimentPropsMixin_ModelInfrastructureConfigProperty struct {
	// The inference option to which to deploy your model. Possible values are the following:.
	//
	// - `RealTime` : Deploy to real-time inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-infrastructuretype
	//
	InfrastructureType *string `field:"optional" json:"infrastructureType" yaml:"infrastructureType"`
	// The infrastructure configuration for deploying the model to real-time inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-realtimeinferenceconfig
	//
	RealTimeInferenceConfig interface{} `field:"optional" json:"realTimeInferenceConfig" yaml:"realTimeInferenceConfig"`
}

