package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetParametersProperty := &PipeTargetParametersProperty{
//   	BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   		JobDefinition: jsii.String("jobDefinition"),
//   		JobName: jsii.String("jobName"),
//
//   		// the properties below are optional
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
//   		TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		CapacityProviderStrategy: []interface{}{
//   			&CapacityProviderStrategyItemProperty{
//   				CapacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				Base: jsii.Number(123),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		EnableEcsManagedTags: jsii.Boolean(false),
//   		EnableExecuteCommand: jsii.Boolean(false),
//   		Group: jsii.String("group"),
//   		LaunchType: jsii.String("launchType"),
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
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
//   		Tags: []cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TaskCount: jsii.Number(123),
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
//   		Sqls: []*string{
//   			jsii.String("sqls"),
//   		},
//
//   		// the properties below are optional
//   		DbUser: jsii.String("dbUser"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html
//
type CfnPipe_PipeTargetParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-batchjobparameters
	//
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-cloudwatchlogsparameters
	//
	CloudWatchLogsParameters interface{} `field:"optional" json:"cloudWatchLogsParameters" yaml:"cloudWatchLogsParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-ecstaskparameters
	//
	EcsTaskParameters interface{} `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-eventbridgeeventbusparameters
	//
	EventBridgeEventBusParameters interface{} `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-kinesisstreamparameters
	//
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-lambdafunctionparameters
	//
	LambdaFunctionParameters interface{} `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-redshiftdataparameters
	//
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-sagemakerpipelineparameters
	//
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-sqsqueueparameters
	//
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-stepfunctionstatemachineparameters
	//
	StepFunctionStateMachineParameters interface{} `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

