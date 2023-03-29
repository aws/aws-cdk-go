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
type CfnInferenceExperimentProps struct {
	// `AWS::SageMaker::InferenceExperiment.EndpointName`.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// `AWS::SageMaker::InferenceExperiment.ModelVariants`.
	ModelVariants interface{} `field:"required" json:"modelVariants" yaml:"modelVariants"`
	// The name of the inference experiment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of the inference experiment.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `AWS::SageMaker::InferenceExperiment.DataStorageConfig`.
	DataStorageConfig interface{} `field:"optional" json:"dataStorageConfig" yaml:"dataStorageConfig"`
	// The description of the inference experiment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::SageMaker::InferenceExperiment.DesiredState`.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The duration for which the inference experiment ran or will run.
	//
	// The maximum duration that you can set for an inference experiment is 30 days.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// `AWS::SageMaker::InferenceExperiment.ShadowModeConfig`.
	ShadowModeConfig interface{} `field:"optional" json:"shadowModeConfig" yaml:"shadowModeConfig"`
	// The error message for the inference experiment status result.
	StatusReason *string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// `AWS::SageMaker::InferenceExperiment.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

