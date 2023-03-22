package awssagemaker


// Contains information about the deployment options of a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelVariantConfigProperty := &modelVariantConfigProperty{
//   	infrastructureConfig: &modelInfrastructureConfigProperty{
//   		infrastructureType: jsii.String("infrastructureType"),
//   		realTimeInferenceConfig: &realTimeInferenceConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   		},
//   	},
//   	modelName: jsii.String("modelName"),
//   	variantName: jsii.String("variantName"),
//   }
//
type CfnInferenceExperiment_ModelVariantConfigProperty struct {
	// The configuration for the infrastructure that the model will be deployed to.
	InfrastructureConfig interface{} `field:"required" json:"infrastructureConfig" yaml:"infrastructureConfig"`
	// The name of the Amazon SageMaker Model entity.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the variant.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
}

