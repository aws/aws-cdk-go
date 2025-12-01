package previewawssagemakerevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   requestParameters := &RequestParameters{
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		TrainingImage: []*string{
//   			jsii.String("trainingImage"),
//   		},
//   		TrainingInputMode: []*string{
//   			jsii.String("trainingInputMode"),
//   		},
//   	},
//   	EnableInterContainerTrafficEncryption: []*string{
//   		jsii.String("enableInterContainerTrafficEncryption"),
//   	},
//   	EnableManagedSpotTraining: []*string{
//   		jsii.String("enableManagedSpotTraining"),
//   	},
//   	EnableNetworkIsolation: []*string{
//   		jsii.String("enableNetworkIsolation"),
//   	},
//   	EndpointConfigName: []*string{
//   		jsii.String("endpointConfigName"),
//   	},
//   	EndpointName: []*string{
//   		jsii.String("endpointName"),
//   	},
//   	ExecutionRoleArn: []*string{
//   		jsii.String("executionRoleArn"),
//   	},
//   	HyperParameters: &HyperParameters{
//   		EvalMetric: []*string{
//   			jsii.String("evalMetric"),
//   		},
//   		NumRound: []*string{
//   			jsii.String("numRound"),
//   		},
//   		Objective: []*string{
//   			jsii.String("objective"),
//   		},
//   	},
//   	InputDataConfig: []RequestParametersItem2{
//   		&RequestParametersItem2{
//   			ChannelName: []*string{
//   				jsii.String("channelName"),
//   			},
//   			ContentType: []*string{
//   				jsii.String("contentType"),
//   			},
//   			DataSource: &DataSource1{
//   				S3DataSource: &S3DataSource1{
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
//   		},
//   	},
//   	ModelName: []*string{
//   		jsii.String("modelName"),
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		RemoveJobNameFromS3OutputPath: []*string{
//   			jsii.String("removeJobNameFromS3OutputPath"),
//   		},
//   		S3OutputPath: []*string{
//   			jsii.String("s3OutputPath"),
//   		},
//   	},
//   	PrimaryContainer: &PrimaryContainer{
//   		ContainerHostname: []*string{
//   			jsii.String("containerHostname"),
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		ModelDataUrl: []*string{
//   			jsii.String("modelDataUrl"),
//   		},
//   	},
//   	ProductionVariants: []RequestParametersItem{
//   		&RequestParametersItem{
//   			InitialInstanceCount: []*string{
//   				jsii.String("initialInstanceCount"),
//   			},
//   			InitialVariantWeight: []*string{
//   				jsii.String("initialVariantWeight"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			ModelName: []*string{
//   				jsii.String("modelName"),
//   			},
//   			VariantName: []*string{
//   				jsii.String("variantName"),
//   			},
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
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntimeInSeconds: []*string{
//   			jsii.String("maxRuntimeInSeconds"),
//   		},
//   	},
//   	Tags: []interface{}{
//   		tags,
//   	},
//   	TrainingJobName: []*string{
//   		jsii.String("trainingJobName"),
//   	},
//   	TransformInput: &TransformInput{
//   		CompressionType: []*string{
//   			jsii.String("compressionType"),
//   		},
//   		ContentType: []*string{
//   			jsii.String("contentType"),
//   		},
//   		DataSource: &DataSource{
//   			S3DataSource: &S3DataSource{
//   				S3DataType: []*string{
//   					jsii.String("s3DataType"),
//   				},
//   				S3Uri: []*string{
//   					jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   	},
//   	TransformJobName: []*string{
//   		jsii.String("transformJobName"),
//   	},
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: []*string{
//   			jsii.String("s3OutputPath"),
//   		},
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: []*string{
//   			jsii.String("instanceCount"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// algorithmSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of algorithmSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmSpecification *ModelEvents_AWSAPICallViaCloudTrail_AlgorithmSpecification `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// enableInterContainerTrafficEncryption property.
	//
	// Specify an array of string values to match this event if the actual value of enableInterContainerTrafficEncryption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnableInterContainerTrafficEncryption *[]*string `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// enableManagedSpotTraining property.
	//
	// Specify an array of string values to match this event if the actual value of enableManagedSpotTraining is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnableManagedSpotTraining *[]*string `field:"optional" json:"enableManagedSpotTraining" yaml:"enableManagedSpotTraining"`
	// enableNetworkIsolation property.
	//
	// Specify an array of string values to match this event if the actual value of enableNetworkIsolation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnableNetworkIsolation *[]*string `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// endpointConfigName property.
	//
	// Specify an array of string values to match this event if the actual value of endpointConfigName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointConfigName *[]*string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// endpointName property.
	//
	// Specify an array of string values to match this event if the actual value of endpointName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointName *[]*string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// executionRoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of executionRoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionRoleArn *[]*string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// hyperParameters property.
	//
	// Specify an array of string values to match this event if the actual value of hyperParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HyperParameters *ModelEvents_AWSAPICallViaCloudTrail_HyperParameters `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// inputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of inputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputDataConfig *[]*ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem2 `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// modelName property.
	//
	// Specify an array of string values to match this event if the actual value of modelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelName *[]*string `field:"optional" json:"modelName" yaml:"modelName"`
	// outputDataConfig property.
	//
	// Specify an array of string values to match this event if the actual value of outputDataConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDataConfig *ModelEvents_AWSAPICallViaCloudTrail_OutputDataConfig `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// primaryContainer property.
	//
	// Specify an array of string values to match this event if the actual value of primaryContainer is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrimaryContainer *ModelEvents_AWSAPICallViaCloudTrail_PrimaryContainer `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// productionVariants property.
	//
	// Specify an array of string values to match this event if the actual value of productionVariants is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductionVariants *[]*ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"productionVariants" yaml:"productionVariants"`
	// resourceConfig property.
	//
	// Specify an array of string values to match this event if the actual value of resourceConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceConfig *ModelEvents_AWSAPICallViaCloudTrail_ResourceConfig `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// roleArn property.
	//
	// Specify an array of string values to match this event if the actual value of roleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// stoppingCondition property.
	//
	// Specify an array of string values to match this event if the actual value of stoppingCondition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppingCondition *ModelEvents_AWSAPICallViaCloudTrail_StoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
	// trainingJobName property.
	//
	// Specify an array of string values to match this event if the actual value of trainingJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TrainingJobName *[]*string `field:"optional" json:"trainingJobName" yaml:"trainingJobName"`
	// transformInput property.
	//
	// Specify an array of string values to match this event if the actual value of transformInput is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformInput *ModelEvents_AWSAPICallViaCloudTrail_TransformInput `field:"optional" json:"transformInput" yaml:"transformInput"`
	// transformJobName property.
	//
	// Specify an array of string values to match this event if the actual value of transformJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformJobName *[]*string `field:"optional" json:"transformJobName" yaml:"transformJobName"`
	// transformOutput property.
	//
	// Specify an array of string values to match this event if the actual value of transformOutput is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformOutput *ModelEvents_AWSAPICallViaCloudTrail_TransformOutput `field:"optional" json:"transformOutput" yaml:"transformOutput"`
	// transformResources property.
	//
	// Specify an array of string values to match this event if the actual value of transformResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformResources *ModelEvents_AWSAPICallViaCloudTrail_TransformResources `field:"optional" json:"transformResources" yaml:"transformResources"`
}

