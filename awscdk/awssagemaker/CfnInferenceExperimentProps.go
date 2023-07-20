package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInferenceExperiment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInferenceExperimentProps := &CfnInferenceExperimentProps{
//   	EndpointName: jsii.String("endpointName"),
//   	ModelVariants: []interface{}{
//   		&ModelVariantConfigProperty{
//   			InfrastructureConfig: &ModelInfrastructureConfigProperty{
//   				InfrastructureType: jsii.String("infrastructureType"),
//   				RealTimeInferenceConfig: &RealTimeInferenceConfigProperty{
//   					InstanceCount: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   				},
//   			},
//   			ModelName: jsii.String("modelName"),
//   			VariantName: jsii.String("variantName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DataStorageConfig: &DataStorageConfigProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		ContentType: &CaptureContentTypeHeaderProperty{
//   			CsvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			JsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	Description: jsii.String("description"),
//   	DesiredState: jsii.String("desiredState"),
//   	KmsKey: jsii.String("kmsKey"),
//   	Schedule: &InferenceExperimentScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	ShadowModeConfig: &ShadowModeConfigProperty{
//   		ShadowModelVariants: []interface{}{
//   			&ShadowModelVariantConfigProperty{
//   				SamplingPercentage: jsii.Number(123),
//   				ShadowModelVariantName: jsii.String("shadowModelVariantName"),
//   			},
//   		},
//   		SourceModelVariantName: jsii.String("sourceModelVariantName"),
//   	},
//   	StatusReason: jsii.String("statusReason"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html
//
type CfnInferenceExperimentProps struct {
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// An array of `ModelVariantConfigSummary` objects.
	//
	// There is one for each variant in the inference experiment. Each `ModelVariantConfigSummary` object in the array describes the infrastructure configuration for deploying the corresponding variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-modelvariants
	//
	ModelVariants interface{} `field:"required" json:"modelVariants" yaml:"modelVariants"`
	// The name of the inference experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of the inference experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The Amazon S3 location and configuration for storing inference request and response data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-datastorageconfig
	//
	DataStorageConfig interface{} `field:"optional" json:"dataStorageConfig" yaml:"dataStorageConfig"`
	// The description of the inference experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The desired state of the experiment after stopping. The possible states are the following:.
	//
	// - `Completed` : The experiment completed successfully
	// - `Cancelled` : The experiment was canceled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-desiredstate
	//
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The duration for which the inference experiment ran or will run.
	//
	// The maximum duration that you can set for an inference experiment is 30 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests.
	//
	// For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-shadowmodeconfig
	//
	ShadowModeConfig interface{} `field:"optional" json:"shadowModeConfig" yaml:"shadowModeConfig"`
	// The error message for the inference experiment status result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-statusreason
	//
	StatusReason *string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html#cfn-sagemaker-inferenceexperiment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

