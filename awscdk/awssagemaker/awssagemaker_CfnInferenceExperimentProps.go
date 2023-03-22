package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnInferenceExperiment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInferenceExperimentProps := &cfnInferenceExperimentProps{
//   	endpointName: jsii.String("endpointName"),
//   	modelVariants: []interface{}{
//   		&modelVariantConfigProperty{
//   			infrastructureConfig: &modelInfrastructureConfigProperty{
//   				infrastructureType: jsii.String("infrastructureType"),
//   				realTimeInferenceConfig: &realTimeInferenceConfigProperty{
//   					instanceCount: jsii.Number(123),
//   					instanceType: jsii.String("instanceType"),
//   				},
//   			},
//   			modelName: jsii.String("modelName"),
//   			variantName: jsii.String("variantName"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	dataStorageConfig: &dataStorageConfigProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		contentType: &captureContentTypeHeaderProperty{
//   			csvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			jsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	description: jsii.String("description"),
//   	desiredState: jsii.String("desiredState"),
//   	kmsKey: jsii.String("kmsKey"),
//   	schedule: &inferenceExperimentScheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		startTime: jsii.String("startTime"),
//   	},
//   	shadowModeConfig: &shadowModeConfigProperty{
//   		shadowModelVariants: []interface{}{
//   			&shadowModelVariantConfigProperty{
//   				samplingPercentage: jsii.Number(123),
//   				shadowModelVariantName: jsii.String("shadowModelVariantName"),
//   			},
//   		},
//   		sourceModelVariantName: jsii.String("sourceModelVariantName"),
//   	},
//   	statusReason: jsii.String("statusReason"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

