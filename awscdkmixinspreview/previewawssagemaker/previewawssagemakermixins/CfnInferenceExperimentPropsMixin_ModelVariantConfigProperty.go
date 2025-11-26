package previewawssagemakermixins


// Contains information about the deployment options of a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelVariantConfigProperty := &ModelVariantConfigProperty{
//   	InfrastructureConfig: &ModelInfrastructureConfigProperty{
//   		InfrastructureType: jsii.String("infrastructureType"),
//   		RealTimeInferenceConfig: &RealTimeInferenceConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   		},
//   	},
//   	ModelName: jsii.String("modelName"),
//   	VariantName: jsii.String("variantName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelvariantconfig.html
//
type CfnInferenceExperimentPropsMixin_ModelVariantConfigProperty struct {
	// The configuration for the infrastructure that the model will be deployed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelvariantconfig.html#cfn-sagemaker-inferenceexperiment-modelvariantconfig-infrastructureconfig
	//
	InfrastructureConfig interface{} `field:"optional" json:"infrastructureConfig" yaml:"infrastructureConfig"`
	// The name of the Amazon SageMaker Model entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelvariantconfig.html#cfn-sagemaker-inferenceexperiment-modelvariantconfig-modelname
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The name of the variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelvariantconfig.html#cfn-sagemaker-inferenceexperiment-modelvariantconfig-variantname
	//
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
}

