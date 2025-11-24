package mixinsawspipes


// The parameters required to set up a target for your pipe.
//
// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeTargetParametersProperty := &PipeTargetParametersProperty{
//   	BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   		ArrayProperties: &BatchArrayPropertiesProperty{
//   			Size: jsii.Number(123),
//   		},
//   		ContainerOverrides: &BatchContainerOverridesProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Environment: []interface{}{
//   				&BatchEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			ResourceRequirements: []interface{}{
//   				&BatchResourceRequirementProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		DependsOn: []interface{}{
//   			&BatchJobDependencyProperty{
//   				JobId: jsii.String("jobId"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		JobDefinition: jsii.String("jobDefinition"),
//   		JobName: jsii.String("jobName"),
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		RetryStrategy: &BatchRetryStrategyProperty{
//   			Attempts: jsii.Number(123),
//   		},
//   	},
//   	CloudWatchLogsParameters: &PipeTargetCloudWatchLogsParametersProperty{
//   		LogStreamName: jsii.String("logStreamName"),
//   		Timestamp: jsii.String("timestamp"),
//   	},
//   	EcsTaskParameters: &PipeTargetEcsTaskParametersProperty{
//   		CapacityProviderStrategy: []interface{}{
//   			&CapacityProviderStrategyItemProperty{
//   				Base: jsii.Number(123),
//   				CapacityProvider: jsii.String("capacityProvider"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		EnableEcsManagedTags: jsii.Boolean(false),
//   		EnableExecuteCommand: jsii.Boolean(false),
//   		Group: jsii.String("group"),
//   		LaunchType: jsii.String("launchType"),
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		Overrides: &EcsTaskOverrideProperty{
//   			ContainerOverrides: []interface{}{
//   				&EcsContainerOverrideProperty{
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Cpu: jsii.Number(123),
//   					Environment: []interface{}{
//   						&EcsEnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					EnvironmentFiles: []interface{}{
//   						&EcsEnvironmentFileProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Memory: jsii.Number(123),
//   					MemoryReservation: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					ResourceRequirements: []interface{}{
//   						&EcsResourceRequirementProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			Cpu: jsii.String("cpu"),
//   			EphemeralStorage: &EcsEphemeralStorageProperty{
//   				SizeInGiB: jsii.Number(123),
//   			},
//   			ExecutionRoleArn: jsii.String("executionRoleArn"),
//   			InferenceAcceleratorOverrides: []interface{}{
//   				&EcsInferenceAcceleratorOverrideProperty{
//   					DeviceName: jsii.String("deviceName"),
//   					DeviceType: jsii.String("deviceType"),
//   				},
//   			},
//   			Memory: jsii.String("memory"),
//   			TaskRoleArn: jsii.String("taskRoleArn"),
//   		},
//   		PlacementConstraints: []interface{}{
//   			&PlacementConstraintProperty{
//   				Expression: jsii.String("expression"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlacementStrategy: []interface{}{
//   			&PlacementStrategyProperty{
//   				Field: jsii.String("field"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlatformVersion: jsii.String("platformVersion"),
//   		PropagateTags: jsii.String("propagateTags"),
//   		ReferenceId: jsii.String("referenceId"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TaskCount: jsii.Number(123),
//   		TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   	},
//   	EventBridgeEventBusParameters: &PipeTargetEventBridgeEventBusParametersProperty{
//   		DetailType: jsii.String("detailType"),
//   		EndpointId: jsii.String("endpointId"),
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Source: jsii.String("source"),
//   		Time: jsii.String("time"),
//   	},
//   	HttpParameters: &PipeTargetHttpParametersProperty{
//   		HeaderParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		PathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		QueryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	InputTemplate: jsii.String("inputTemplate"),
//   	KinesisStreamParameters: &PipeTargetKinesisStreamParametersProperty{
//   		PartitionKey: jsii.String("partitionKey"),
//   	},
//   	LambdaFunctionParameters: &PipeTargetLambdaFunctionParametersProperty{
//   		InvocationType: jsii.String("invocationType"),
//   	},
//   	RedshiftDataParameters: &PipeTargetRedshiftDataParametersProperty{
//   		Database: jsii.String("database"),
//   		DbUser: jsii.String("dbUser"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		Sqls: []*string{
//   			jsii.String("sqls"),
//   		},
//   		StatementName: jsii.String("statementName"),
//   		WithEvent: jsii.Boolean(false),
//   	},
//   	SageMakerPipelineParameters: &PipeTargetSageMakerPipelineParametersProperty{
//   		PipelineParameterList: []interface{}{
//   			&SageMakerPipelineParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SqsQueueParameters: &PipeTargetSqsQueueParametersProperty{
//   		MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   		MessageGroupId: jsii.String("messageGroupId"),
//   	},
//   	StepFunctionStateMachineParameters: &PipeTargetStateMachineParametersProperty{
//   		InvocationType: jsii.String("invocationType"),
//   	},
//   	TimestreamParameters: &PipeTargetTimestreamParametersProperty{
//   		DimensionMappings: []interface{}{
//   			&DimensionMappingProperty{
//   				DimensionName: jsii.String("dimensionName"),
//   				DimensionValue: jsii.String("dimensionValue"),
//   				DimensionValueType: jsii.String("dimensionValueType"),
//   			},
//   		},
//   		EpochTimeUnit: jsii.String("epochTimeUnit"),
//   		MultiMeasureMappings: []interface{}{
//   			&MultiMeasureMappingProperty{
//   				MultiMeasureAttributeMappings: []interface{}{
//   					&MultiMeasureAttributeMappingProperty{
//   						MeasureValue: jsii.String("measureValue"),
//   						MeasureValueType: jsii.String("measureValueType"),
//   						MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   					},
//   				},
//   				MultiMeasureName: jsii.String("multiMeasureName"),
//   			},
//   		},
//   		SingleMeasureMappings: []interface{}{
//   			&SingleMeasureMappingProperty{
//   				MeasureName: jsii.String("measureName"),
//   				MeasureValue: jsii.String("measureValue"),
//   				MeasureValueType: jsii.String("measureValueType"),
//   			},
//   		},
//   		TimeFieldType: jsii.String("timeFieldType"),
//   		TimestampFormat: jsii.String("timestampFormat"),
//   		TimeValue: jsii.String("timeValue"),
//   		VersionValue: jsii.String("versionValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html
//
type CfnPipePropsMixin_PipeTargetParametersProperty struct {
	// The parameters for using an AWS Batch job as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-batchjobparameters
	//
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// The parameters for using an CloudWatch Logs log stream as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-cloudwatchlogsparameters
	//
	CloudWatchLogsParameters interface{} `field:"optional" json:"cloudWatchLogsParameters" yaml:"cloudWatchLogsParameters"`
	// The parameters for using an Amazon ECS task as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-ecstaskparameters
	//
	EcsTaskParameters interface{} `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// The parameters for using an EventBridge event bus as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-eventbridgeeventbusparameters
	//
	EventBridgeEventBusParameters interface{} `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
	// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-httpparameters
	//
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	//
	// To remove an input template, specify an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// The parameters for using a Kinesis stream as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-kinesisstreamparameters
	//
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// The parameters for using a Lambda function as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-lambdafunctionparameters
	//
	LambdaFunctionParameters interface{} `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-redshiftdataparameters
	//
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The parameters for using a SageMaker AI pipeline as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-sagemakerpipelineparameters
	//
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// The parameters for using a Amazon SQS stream as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-sqsqueueparameters
	//
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// The parameters for using a Step Functions state machine as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-stepfunctionstatemachineparameters
	//
	StepFunctionStateMachineParameters interface{} `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
	// The parameters for using a Timestream for LiveAnalytics table as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-timestreamparameters
	//
	TimestreamParameters interface{} `field:"optional" json:"timestreamParameters" yaml:"timestreamParameters"`
}

