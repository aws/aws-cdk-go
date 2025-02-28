package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEndpointConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointConfigProps := &CfnEndpointConfigProps{
//   	ProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			VariantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//
//   				// the properties below are optional
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AsyncInferenceConfig: &AsyncInferenceConfigProperty{
//   		OutputConfig: &AsyncInferenceOutputConfigProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			NotificationConfig: &AsyncInferenceNotificationConfigProperty{
//   				ErrorTopic: jsii.String("errorTopic"),
//   				IncludeInferenceResponseIn: []*string{
//   					jsii.String("includeInferenceResponseIn"),
//   				},
//   				SuccessTopic: jsii.String("successTopic"),
//   			},
//   			S3FailurePath: jsii.String("s3FailurePath"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//
//   		// the properties below are optional
//   		ClientConfig: &AsyncInferenceClientConfigProperty{
//   			MaxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   	},
//   	DataCaptureConfig: &DataCaptureConfigProperty{
//   		CaptureOptions: []interface{}{
//   			&CaptureOptionProperty{
//   				CaptureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//   		InitialSamplingPercentage: jsii.Number(123),
//
//   		// the properties below are optional
//   		CaptureContentTypeHeader: &CaptureContentTypeHeaderProperty{
//   			CsvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			JsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		EnableCapture: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ExplainerConfig: &ExplainerConfigProperty{
//   		ClarifyExplainerConfig: &ClarifyExplainerConfigProperty{
//   			ShapConfig: &ClarifyShapConfigProperty{
//   				ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   					MimeType: jsii.String("mimeType"),
//   					ShapBaseline: jsii.String("shapBaseline"),
//   					ShapBaselineUri: jsii.String("shapBaselineUri"),
//   				},
//
//   				// the properties below are optional
//   				NumberOfSamples: jsii.Number(123),
//   				Seed: jsii.Number(123),
//   				TextConfig: &ClarifyTextConfigProperty{
//   					Granularity: jsii.String("granularity"),
//   					Language: jsii.String("language"),
//   				},
//   				UseLogit: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			EnableExplanations: jsii.String("enableExplanations"),
//   			InferenceConfig: &ClarifyInferenceConfigProperty{
//   				ContentTemplate: jsii.String("contentTemplate"),
//   				FeatureHeaders: []*string{
//   					jsii.String("featureHeaders"),
//   				},
//   				FeaturesAttribute: jsii.String("featuresAttribute"),
//   				FeatureTypes: []*string{
//   					jsii.String("featureTypes"),
//   				},
//   				LabelAttribute: jsii.String("labelAttribute"),
//   				LabelHeaders: []*string{
//   					jsii.String("labelHeaders"),
//   				},
//   				LabelIndex: jsii.Number(123),
//   				MaxPayloadInMb: jsii.Number(123),
//   				MaxRecordCount: jsii.Number(123),
//   				ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   				ProbabilityIndex: jsii.Number(123),
//   			},
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ShadowProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			VariantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//
//   				// the properties below are optional
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
//
type CfnEndpointConfigProps struct {
	// A list of `ProductionVariant` objects, one for each model that you want to host at this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-productionvariants
	//
	ProductionVariants interface{} `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-asyncinferenceconfig
	//
	AsyncInferenceConfig interface{} `field:"optional" json:"asyncInferenceConfig" yaml:"asyncInferenceConfig"`
	// Specifies how to capture endpoint data for model monitor.
	//
	// The data capture configuration applies to all production variants hosted at the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig
	//
	DataCaptureConfig interface{} `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-enablenetworkisolation
	//
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// The name of the endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-endpointconfigname
	//
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A parameter to activate explainers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-explainerconfig
	//
	ExplainerConfig interface{} `field:"optional" json:"explainerConfig" yaml:"explainerConfig"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Alias name: `alias/ExampleAlias`
	// - Alias name ARN: `arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias`
	//
	// The KMS key policy must grant permission to the IAM role that you specify in your `CreateEndpoint` , `UpdateEndpoint` requests. For more information, refer to the AWS Key Management Service section [Using Key Policies in AWS KMS](https://docs.aws.amazon.com//kms/latest/developerguide/key-policies.html)
	//
	// > Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a `KmsKeyId` when using an instance type with local storage. If any of the models that you specify in the `ProductionVariants` parameter use nitro-based instances with local storage, do not specify a value for the `KmsKeyId` parameter. If you specify a value for `KmsKeyId` when using any nitro-based instances with local storage, the call to `CreateEndpointConfig` fails.
	// >
	// > For a list of instance types that support local instance storage, see [Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes) .
	// >
	// > For more information about local instance storage encryption, see [SSD Instance Store Volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Array of `ProductionVariant` objects.
	//
	// There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on `ProductionVariants` . If you use this field, you can only specify one variant for `ProductionVariants` and one variant for `ShadowProductionVariants` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-shadowproductionvariants
	//
	ShadowProductionVariants interface{} `field:"optional" json:"shadowProductionVariants" yaml:"shadowProductionVariants"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

