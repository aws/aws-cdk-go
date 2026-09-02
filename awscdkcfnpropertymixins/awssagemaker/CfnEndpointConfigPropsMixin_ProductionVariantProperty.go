package awssagemaker


// Specifies a model that you want to host and the resources to deploy for hosting it.
//
// If you are deploying multiple models, tell Amazon SageMaker how to distribute traffic among the models by specifying the `InitialVariantWeight` objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   productionVariantProperty := &ProductionVariantProperty{
//   	AcceleratorType: jsii.String("acceleratorType"),
//   	CapacityReservationConfig: &CapacityReservationConfigProperty{
//   		CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		MlReservationArn: jsii.String("mlReservationArn"),
//   	},
//   	ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   	CoreDumpConfig: &CoreDumpConfigProperty{
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EnableSsmAccess: jsii.Boolean(false),
//   	InferenceAmiVersion: jsii.String("inferenceAmiVersion"),
//   	InitialInstanceCount: jsii.Number(123),
//   	InitialVariantWeight: jsii.Number(123),
//   	InstancePools: []interface{}{
//   		&InstancePoolProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			ModelNameOverride: jsii.String("modelNameOverride"),
//   			Priority: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   		MaxInstanceCount: jsii.Number(123),
//   		MinInstanceCount: jsii.Number(123),
//   		ScaleInPolicy: &ScaleInPolicyProperty{
//   			CooldownInMinutes: jsii.Number(123),
//   			MaximumStepSize: jsii.Number(123),
//   			Strategy: jsii.String("strategy"),
//   		},
//   		Status: jsii.String("status"),
//   	},
//   	ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	ModelName: jsii.String("modelName"),
//   	RoutingConfig: &RoutingConfigProperty{
//   		PrefixAwareRoutingConfig: &PrefixAwareRoutingConfigProperty{
//   			ConcurrencyThreshold: jsii.Number(123),
//   			PrefixLength: jsii.Number(123),
//   		},
//   		RoutingStrategy: jsii.String("routingStrategy"),
//   	},
//   	ServerlessConfig: &ServerlessConfigProperty{
//   		MaxConcurrency: jsii.Number(123),
//   		MemorySizeInMb: jsii.Number(123),
//   		ProvisionedConcurrency: jsii.Number(123),
//   	},
//   	VariantInstanceProvisionTimeoutInSeconds: jsii.Number(123),
//   	VariantName: jsii.String("variantName"),
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html
//
type CfnEndpointConfigPropsMixin_ProductionVariantProperty struct {
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	//
	// EI instances provide on-demand GPU computing for inference. For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) . For more information, see [Using Elastic Inference in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-acceleratortype
	//
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// Settings for the capacity reservation for the compute instances that SageMaker AI reserves for an endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-capacityreservationconfig
	//
	CapacityReservationConfig interface{} `field:"optional" json:"capacityReservationConfig" yaml:"capacityReservationConfig"`
	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting.
	//
	// For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-containerstartuphealthchecktimeoutinseconds
	//
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// Specifies where SageMaker writes core dumps from the model container when the process crashes, and how it encrypts them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-coredumpconfig
	//
	CoreDumpConfig interface{} `field:"optional" json:"coreDumpConfig" yaml:"coreDumpConfig"`
	// You can use this parameter to turn on native AWS Systems Manager (SSM) access for a production variant behind an endpoint.
	//
	// By default, SSM access is disabled for all production variants behind an endpoint. You can turn on or turn off SSM access for a production variant behind an existing endpoint by creating a new endpoint configuration and calling `UpdateEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-enablessmaccess
	//
	EnableSsmAccess interface{} `field:"optional" json:"enableSsmAccess" yaml:"enableSsmAccess"`
	// Specifies an option from a collection of preconfigured Amazon Machine Image (AMI) images.
	//
	// Each image is configured by AWS with a set of software and driver versions. AWS optimizes these configurations for different machine learning workloads. By selecting an AMI version, you can ensure that your inference environment is compatible with specific software requirements, such as CUDA driver versions, Linux kernel versions, or AWS Neuron driver versions
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-inferenceamiversion
	//
	InferenceAmiVersion *string `field:"optional" json:"inferenceAmiVersion" yaml:"inferenceAmiVersion"`
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
	// A list of instance pools for the production variant.
	//
	// Each instance pool specifies an instance type and its priority for provisioning. Use instance pools to configure heterogeneous endpoints that deploy models across multiple instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-instancepools
	//
	InstancePools interface{} `field:"optional" json:"instancePools" yaml:"instancePools"`
	// The ML compute instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Settings that control the range in the number of instances that the endpoint provisions as it scales up or down to accommodate traffic.
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
	// Settings that control how the endpoint routes incoming traffic to the instances that the endpoint hosts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-routingconfig
	//
	RoutingConfig interface{} `field:"optional" json:"routingConfig" yaml:"routingConfig"`
	// The serverless configuration for an endpoint.
	//
	// Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-serverlessconfig
	//
	ServerlessConfig interface{} `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// The timeout value, in seconds, for provisioning instances for the production variant.
	//
	// When SageMaker encounters an insufficient capacity error while provisioning instances, it retries with the next instance pool (if configured) or waits until the timeout expires. This timeout applies only to capacity provisioning and does not include the time for model download or container startup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-variantinstanceprovisiontimeoutinseconds
	//
	VariantInstanceProvisionTimeoutInSeconds *float64 `field:"optional" json:"variantInstanceProvisionTimeoutInSeconds" yaml:"variantInstanceProvisionTimeoutInSeconds"`
	// The name of the production variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-variantname
	//
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant.
	//
	// Currently only Amazon EBS gp2 storage volumes are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html#cfn-sagemaker-endpointconfig-productionvariant-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

