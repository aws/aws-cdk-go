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
//   cfnEndpointConfigProps := &cfnEndpointConfigProps{
//   	productionVariants: []interface{}{
//   		&productionVariantProperty{
//   			initialVariantWeight: jsii.Number(123),
//   			modelName: jsii.String("modelName"),
//   			variantName: jsii.String("variantName"),
//
//   			// the properties below are optional
//   			acceleratorType: jsii.String("acceleratorType"),
//   			containerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			initialInstanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			modelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			serverlessConfig: &serverlessConfigProperty{
//   				maxConcurrency: jsii.Number(123),
//   				memorySizeInMb: jsii.Number(123),
//   			},
//   			volumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	asyncInferenceConfig: &asyncInferenceConfigProperty{
//   		outputConfig: &asyncInferenceOutputConfigProperty{
//   			s3OutputPath: jsii.String("s3OutputPath"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			notificationConfig: &asyncInferenceNotificationConfigProperty{
//   				errorTopic: jsii.String("errorTopic"),
//   				successTopic: jsii.String("successTopic"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		clientConfig: &asyncInferenceClientConfigProperty{
//   			maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   	},
//   	dataCaptureConfig: &dataCaptureConfigProperty{
//   		captureOptions: []interface{}{
//   			&captureOptionProperty{
//   				captureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		destinationS3Uri: jsii.String("destinationS3Uri"),
//   		initialSamplingPercentage: jsii.Number(123),
//
//   		// the properties below are optional
//   		captureContentTypeHeader: &captureContentTypeHeaderProperty{
//   			csvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			jsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		enableCapture: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	endpointConfigName: jsii.String("endpointConfigName"),
//   	explainerConfig: &explainerConfigProperty{
//   		clarifyExplainerConfig: &clarifyExplainerConfigProperty{
//   			shapConfig: &clarifyShapConfigProperty{
//   				shapBaselineConfig: &clarifyShapBaselineConfigProperty{
//   					mimeType: jsii.String("mimeType"),
//   					shapBaseline: jsii.String("shapBaseline"),
//   					shapBaselineUri: jsii.String("shapBaselineUri"),
//   				},
//
//   				// the properties below are optional
//   				numberOfSamples: jsii.Number(123),
//   				seed: jsii.Number(123),
//   				textConfig: &clarifyTextConfigProperty{
//   					granularity: jsii.String("granularity"),
//   					language: jsii.String("language"),
//   				},
//   				useLogit: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			enableExplanations: jsii.String("enableExplanations"),
//   			inferenceConfig: &clarifyInferenceConfigProperty{
//   				contentTemplate: jsii.String("contentTemplate"),
//   				featureHeaders: []*string{
//   					jsii.String("featureHeaders"),
//   				},
//   				featuresAttribute: jsii.String("featuresAttribute"),
//   				featureTypes: []*string{
//   					jsii.String("featureTypes"),
//   				},
//   				labelAttribute: jsii.String("labelAttribute"),
//   				labelHeaders: []*string{
//   					jsii.String("labelHeaders"),
//   				},
//   				labelIndex: jsii.Number(123),
//   				maxPayloadInMb: jsii.Number(123),
//   				maxRecordCount: jsii.Number(123),
//   				probabilityAttribute: jsii.String("probabilityAttribute"),
//   				probabilityIndex: jsii.Number(123),
//   			},
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEndpointConfigProps struct {
	// A list of `ProductionVariant` objects, one for each model that you want to host at this endpoint.
	ProductionVariants interface{} `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig interface{} `field:"optional" json:"asyncInferenceConfig" yaml:"asyncInferenceConfig"`
	// Specifies how to capture endpoint data for model monitor.
	//
	// The data capture configuration applies to all production variants hosted at the endpoint.
	DataCaptureConfig interface{} `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// The name of the endpoint configuration.
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// `AWS::SageMaker::EndpointConfig.ExplainerConfig`.
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
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

