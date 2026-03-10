package previewawssagemakerevents


// Type definition for TrainingJobDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trainingJobDefinition := &TrainingJobDefinition{
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		MetricDefinitions: []AlgorithmSpecificationItem{
//   			&AlgorithmSpecificationItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Regex: []*string{
//   					jsii.String("regex"),
//   				},
//   			},
//   		},
//   		TrainingImage: []*string{
//   			jsii.String("trainingImage"),
//   		},
//   		TrainingInputMode: []*string{
//   			jsii.String("trainingInputMode"),
//   		},
//   	},
//   	InputDataConfig: []TrainingJobDefinitionItem{
//   		&TrainingJobDefinitionItem{
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
//   	OutputDataConfig: &OutputDataConfig{
//   		KmsKeyId: []*string{
//   			jsii.String("kmsKeyId"),
//   		},
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
//   		VolumeKmsKeyId: []*string{
//   			jsii.String("volumeKmsKeyId"),
//   		},
//   		VolumeSizeInGb: []*string{
//   			jsii.String("volumeSizeInGb"),
//   		},
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   	StaticHyperParameters: &Tags{
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntimeInSeconds: []*string{
//   			jsii.String("maxRuntimeInSeconds"),
//   		},
//   	},
//   	VpcConfig: &VpcConfig{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_TrainingJobDefinition struct {
	// AlgorithmSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of AlgorithmSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmSpecification *SageMakerHyperParameterTuningJobStateChange_AlgorithmSpecification `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// InputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of InputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *[]*SageMakerHyperParameterTuningJobStateChange_TrainingJobDefinitionItem `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// OutputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of OutputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *SageMakerHyperParameterTuningJobStateChange_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// ResourceConfig property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceConfig *SageMakerHyperParameterTuningJobStateChange_ResourceConfig `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// RoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of RoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// StaticHyperParameters property.
	//
	// Specify an array of string values to match this event if the actual value of StaticHyperParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StaticHyperParameters *SageMakerHyperParameterTuningJobStateChange_Tags `field:"optional" json:"staticHyperParameters" yaml:"staticHyperParameters"`
	// StoppingCondition property.
	//
	// Specify an array of string values to match this event if the actual value of StoppingCondition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppingCondition *SageMakerHyperParameterTuningJobStateChange_StoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// VpcConfig property.
	//
	// Specify an array of string values to match this event if the actual value of VpcConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcConfig *SageMakerHyperParameterTuningJobStateChange_VpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

