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
//   productionVariantProperty := &ProductionVariantProperty{
//   	InitialVariantWeight: jsii.Number(123),
//   	ModelName: jsii.String("modelName"),
//   	VariantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	AcceleratorType: jsii.String("acceleratorType"),
//   	ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   	EnableSsmAccess: jsii.Boolean(false),
//   	InitialInstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   	ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	ServerlessConfig: &ServerlessConfigProperty{
//   		MaxConcurrency: jsii.Number(123),
//   		MemorySizeInMb: jsii.Number(123),
//
//   		// the properties below are optional
//   		ProvisionedConcurrency: jsii.Number(123),
//   	},
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html
//
type CfnEndpointConfig_ProductionVariantProperty struct {
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the `VariantWeight` to the sum of all `VariantWeight` values across all ProductionVariants. If unspecified, it defaults to 1.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-initialvariantweight
	//
	InitialVariantWeight *float64 `field:"required" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-modelname
	//
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the production variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-variantname
	//
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	//
	// EI instances provide on-demand GPU computing for inference. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) . For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-acceleratortype
	//
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-containerstartuphealthchecktimeoutinseconds
	//
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// You can use this parameter to turn on native AWS Systems Manager (SSM) access for a production variant behind an endpoint.
	//
	// By default, SSM access is disabled for all production variants behind an endpoint. You can turn on or turn off SSM access for a production variant behind an existing endpoint by creating a new endpoint configuration and calling `UpdateEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-enablessmaccess
	//
	EnableSsmAccess interface{} `field:"optional" json:"enableSsmAccess" yaml:"enableSsmAccess"`
	// Number of instances to launch initially.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-initialinstancecount
	//
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// The ML compute instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-modeldatadownloadtimeoutinseconds
	//
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// The serverless configuration for an endpoint.
	//
	// Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-serverlessconfig
	//
	ServerlessConfig interface{} `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

