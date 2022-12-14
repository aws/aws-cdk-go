package awssagemaker


// Specifies a model that you want to host and the resources to deploy for hosting it.
//
// If you are deploying multiple models, tell Amazon SageMaker how to distribute traffic among the models by specifying the `InitialVariantWeight` objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   productionVariantProperty := &productionVariantProperty{
//   	initialVariantWeight: jsii.Number(123),
//   	modelName: jsii.String("modelName"),
//   	variantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	acceleratorType: jsii.String("acceleratorType"),
//   	containerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   	initialInstanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	modelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	serverlessConfig: &serverlessConfigProperty{
//   		maxConcurrency: jsii.Number(123),
//   		memorySizeInMb: jsii.Number(123),
//   	},
//   	volumeSizeInGb: jsii.Number(123),
//   }
//
type CfnEndpointConfig_ProductionVariantProperty struct {
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the `VariantWeight` to the sum of all `VariantWeight` values across all ProductionVariants. If unspecified, it defaults to 1.0.
	InitialVariantWeight *float64 `field:"required" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the production variant.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	//
	// EI instances provide on-demand GPU computing for inference. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) . For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// `CfnEndpointConfig.ProductionVariantProperty.ContainerStartupHealthCheckTimeoutInSeconds`.
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// Number of instances to launch initially.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// The ML compute instance type.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// `CfnEndpointConfig.ProductionVariantProperty.ModelDataDownloadTimeoutInSeconds`.
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// The serverless configuration for an endpoint.
	//
	// Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	ServerlessConfig interface{} `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// `CfnEndpointConfig.ProductionVariantProperty.VolumeSizeInGB`.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

