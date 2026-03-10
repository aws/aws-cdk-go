package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerTrainingJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTrainingJobStateChangeProps := &SageMakerTrainingJobStateChangeProps{
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		TrainingImage: []*string{
//   			jsii.String("trainingImage"),
//   		},
//   		TrainingInputMode: []*string{
//   			jsii.String("trainingInputMode"),
//   		},
//   	},
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	HyperParameters: []*string{
//   		jsii.String("hyperParameters"),
//   	},
//   	InputDataConfig: []SageMakerTrainingJobStateChangeItem{
//   		&SageMakerTrainingJobStateChangeItem{
//   			ChannelName: []*string{
//   				jsii.String("channelName"),
//   			},
//   			CompressionType: []*string{
//   				jsii.String("compressionType"),
//   			},
//   			ContentType: []*string{
//   				jsii.String("contentType"),
//   			},
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3DataDistributionType: []*string{
//   						jsii.String("s3DataDistributionType"),
//   					},
//   					S3DataType: []*string{
//   						jsii.String("s3DataType"),
//   					},
//   					S3Uri: []*string{
//   						jsii.String("s3Uri"),
//   					},
//   				},
//   			},
//   			RecordWrapperType: []*string{
//   				jsii.String("recordWrapperType"),
//   			},
//   		},
//   	},
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
//   	},
//   	ModelArtifacts: &ModelArtifacts{
//   		S3ModelArtifacts: []*string{
//   			jsii.String("s3ModelArtifacts"),
//   		},
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3OutputPath: []*string{
//   			jsii.String("s3OutputPath"),
//   		},
//   	},
//   	ResourceConfig: &ResourceConfig{
//   		InstanceCount: []*string{
//   			jsii.String("instanceCount"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		VolumeSizeInGb: []*string{
//   			jsii.String("volumeSizeInGb"),
//   		},
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   	SecondaryStatus: []*string{
//   		jsii.String("secondaryStatus"),
//   	},
//   	SecondaryStatusTransitions: []SageMakerTrainingJobStateChangeItem1{
//   		&SageMakerTrainingJobStateChangeItem1{
//   			StartTime: []*string{
//   				jsii.String("startTime"),
//   			},
//   			Status: []*string{
//   				jsii.String("status"),
//   			},
//   			StatusMessage: []*string{
//   				jsii.String("statusMessage"),
//   			},
//   		},
//   	},
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntimeInSeconds: []*string{
//   			jsii.String("maxRuntimeInSeconds"),
//   		},
//   	},
//   	Tags: []Tags{
//   		&Tags{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	TrainingEndTime: []*string{
//   		jsii.String("trainingEndTime"),
//   	},
//   	TrainingJobArn: []*string{
//   		jsii.String("trainingJobArn"),
//   	},
//   	TrainingJobName: []*string{
//   		jsii.String("trainingJobName"),
//   	},
//   	TrainingJobStatus: []*string{
//   		jsii.String("trainingJobStatus"),
//   	},
//   	TrainingStartTime: []*string{
//   		jsii.String("trainingStartTime"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeProps struct {
	// AlgorithmSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of AlgorithmSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmSpecification *SageMakerTrainingJobStateChange_AlgorithmSpecification `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// CreationTime property.
	//
	// Specify an array of string values to match this event if the actual value of CreationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// HyperParameters property.
	//
	// Specify an array of string values to match this event if the actual value of HyperParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HyperParameters *[]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// InputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of InputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *[]*SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeItem `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// LastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of LastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// ModelArtifacts property.
	//
	// Specify an array of string values to match this event if the actual value of ModelArtifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelArtifacts *SageMakerTrainingJobStateChange_ModelArtifacts `field:"optional" json:"modelArtifacts" yaml:"modelArtifacts"`
	// OutputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of OutputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *SageMakerTrainingJobStateChange_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// ResourceConfig property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceConfig *SageMakerTrainingJobStateChange_ResourceConfig `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// RoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of RoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// SecondaryStatus property.
	//
	// Specify an array of string values to match this event if the actual value of SecondaryStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecondaryStatus *[]*string `field:"optional" json:"secondaryStatus" yaml:"secondaryStatus"`
	// SecondaryStatusTransitions property.
	//
	// Specify an array of string values to match this event if the actual value of SecondaryStatusTransitions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecondaryStatusTransitions *[]*SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeItem1 `field:"optional" json:"secondaryStatusTransitions" yaml:"secondaryStatusTransitions"`
	// StoppingCondition property.
	//
	// Specify an array of string values to match this event if the actual value of StoppingCondition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppingCondition *SageMakerTrainingJobStateChange_StoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerTrainingJobStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
	// TrainingEndTime property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingEndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingEndTime *[]*string `field:"optional" json:"trainingEndTime" yaml:"trainingEndTime"`
	// TrainingJobArn property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingJobArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobArn *[]*string `field:"optional" json:"trainingJobArn" yaml:"trainingJobArn"`
	// TrainingJobName property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobName *[]*string `field:"optional" json:"trainingJobName" yaml:"trainingJobName"`
	// TrainingJobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingJobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobStatus *[]*string `field:"optional" json:"trainingJobStatus" yaml:"trainingJobStatus"`
	// TrainingStartTime property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingStartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingStartTime *[]*string `field:"optional" json:"trainingStartTime" yaml:"trainingStartTime"`
}

