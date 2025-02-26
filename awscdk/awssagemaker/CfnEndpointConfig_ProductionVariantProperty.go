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
//   	VariantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	AcceleratorType: jsii.String("acceleratorType"),
//   	ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   	EnableSsmAccess: jsii.Boolean(false),
//   	InitialInstanceCount: jsii.Number(123),
//   	InitialVariantWeight: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   	ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   		MaxInstanceCount: jsii.Number(123),
//   		MinInstanceCount: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	ModelName: jsii.String("modelName"),
//   	RoutingConfig: &RoutingConfigProperty{
//   		RoutingStrategy: jsii.String("routingStrategy"),
//   	},
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
	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting.
	//
	// For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests) .
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
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the `VariantWeight` to the sum of all `VariantWeight` values across all ProductionVariants. If unspecified, it defaults to 1.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-initialvariantweight
	//
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// The ML compute instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-managedinstancescaling
	//
	ManagedInstanceScaling interface{} `field:"optional" json:"managedInstanceScaling" yaml:"managedInstanceScaling"`
	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-modeldatadownloadtimeoutinseconds
	//
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-modelname
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-routingconfig
	//
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
	// The serverless configuration for an endpoint.
	//
	// Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-serverlessconfig
	//
	ServerlessConfig interface{} `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant.
	//
	// Currently only Amazon EBS gp2 storage volumes are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

