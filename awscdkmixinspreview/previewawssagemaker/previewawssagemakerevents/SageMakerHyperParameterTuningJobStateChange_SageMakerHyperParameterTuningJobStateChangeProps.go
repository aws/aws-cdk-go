package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerHyperParameterTuningJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerHyperParameterTuningJobStateChangeProps := &SageMakerHyperParameterTuningJobStateChangeProps{
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
//   	HyperParameterTuningJobArn: []*string{
//   		jsii.String("hyperParameterTuningJobArn"),
//   	},
//   	HyperParameterTuningJobName: []*string{
//   		jsii.String("hyperParameterTuningJobName"),
//   	},
//   	HyperParameterTuningJobStatus: []*string{
//   		jsii.String("hyperParameterTuningJobStatus"),
//   	},
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
//   	},
//   	ObjectiveStatusCounters: &ObjectiveStatusCounters{
//   		Failed: []*string{
//   			jsii.String("failed"),
//   		},
//   		Pending: []*string{
//   			jsii.String("pending"),
//   		},
//   		Succeeded: []*string{
//   			jsii.String("succeeded"),
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
//   	TrainingJobDefinition: &TrainingJobDefinition{
//   		AlgorithmSpecification: &AlgorithmSpecification{
//   			MetricDefinitions: []AlgorithmSpecificationItem{
//   				&AlgorithmSpecificationItem{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Regex: []*string{
//   						jsii.String("regex"),
//   					},
//   				},
//   			},
//   			TrainingImage: []*string{
//   				jsii.String("trainingImage"),
//   			},
//   			TrainingInputMode: []*string{
//   				jsii.String("trainingInputMode"),
//   			},
//   		},
//   		InputDataConfig: []TrainingJobDefinitionItem{
//   			&TrainingJobDefinitionItem{
//   				ChannelName: []*string{
//   					jsii.String("channelName"),
//   				},
//   				CompressionType: []*string{
//   					jsii.String("compressionType"),
//   				},
//   				ContentType: []*string{
//   					jsii.String("contentType"),
//   				},
//   				DataSource: &DataSource{
//   					S3DataSource: &S3DataSource{
//   						S3DataDistributionType: []*string{
//   							jsii.String("s3DataDistributionType"),
//   						},
//   						S3DataType: []*string{
//   							jsii.String("s3DataType"),
//   						},
//   						S3Uri: []*string{
//   							jsii.String("s3Uri"),
//   						},
//   					},
//   				},
//   				RecordWrapperType: []*string{
//   					jsii.String("recordWrapperType"),
//   				},
//   			},
//   		},
//   		OutputDataConfig: &OutputDataConfig{
//   			KmsKeyId: []*string{
//   				jsii.String("kmsKeyId"),
//   			},
//   			S3OutputPath: []*string{
//   				jsii.String("s3OutputPath"),
//   			},
//   		},
//   		ResourceConfig: &ResourceConfig{
//   			InstanceCount: []*string{
//   				jsii.String("instanceCount"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			VolumeKmsKeyId: []*string{
//   				jsii.String("volumeKmsKeyId"),
//   			},
//   			VolumeSizeInGb: []*string{
//   				jsii.String("volumeSizeInGb"),
//   			},
//   		},
//   		RoleArn: []*string{
//   			jsii.String("roleArn"),
//   		},
//   		StaticHyperParameters: &Tags{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   		StoppingCondition: &StoppingCondition{
//   			MaxRuntimeInSeconds: []*string{
//   				jsii.String("maxRuntimeInSeconds"),
//   			},
//   		},
//   		VpcConfig: &VpcConfig{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	TrainingJobStatusCounters: &TrainingJobStatusCounters{
//   		Completed: []*string{
//   			jsii.String("completed"),
//   		},
//   		InProgress: []*string{
//   			jsii.String("inProgress"),
//   		},
//   		NonRetryableError: []*string{
//   			jsii.String("nonRetryableError"),
//   		},
//   		RetryableError: []*string{
//   			jsii.String("retryableError"),
//   		},
//   		Stopped: []*string{
//   			jsii.String("stopped"),
//   		},
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_SageMakerHyperParameterTuningJobStateChangeProps struct {
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
	// HyperParameterTuningJobArn property.
	//
	// Specify an array of string values to match this event if the actual value of HyperParameterTuningJobArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HyperParameterTuningJobArn *[]*string `field:"optional" json:"hyperParameterTuningJobArn" yaml:"hyperParameterTuningJobArn"`
	// HyperParameterTuningJobName property.
	//
	// Specify an array of string values to match this event if the actual value of HyperParameterTuningJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HyperParameterTuningJobName *[]*string `field:"optional" json:"hyperParameterTuningJobName" yaml:"hyperParameterTuningJobName"`
	// HyperParameterTuningJobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of HyperParameterTuningJobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HyperParameterTuningJobStatus *[]*string `field:"optional" json:"hyperParameterTuningJobStatus" yaml:"hyperParameterTuningJobStatus"`
	// LastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of LastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// ObjectiveStatusCounters property.
	//
	// Specify an array of string values to match this event if the actual value of ObjectiveStatusCounters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectiveStatusCounters *SageMakerHyperParameterTuningJobStateChange_ObjectiveStatusCounters `field:"optional" json:"objectiveStatusCounters" yaml:"objectiveStatusCounters"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerHyperParameterTuningJobStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
	// TrainingJobDefinition property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingJobDefinition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobDefinition *SageMakerHyperParameterTuningJobStateChange_TrainingJobDefinition `field:"optional" json:"trainingJobDefinition" yaml:"trainingJobDefinition"`
	// TrainingJobStatusCounters property.
	//
	// Specify an array of string values to match this event if the actual value of TrainingJobStatusCounters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobStatusCounters *SageMakerHyperParameterTuningJobStateChange_TrainingJobStatusCounters `field:"optional" json:"trainingJobStatusCounters" yaml:"trainingJobStatusCounters"`
}

