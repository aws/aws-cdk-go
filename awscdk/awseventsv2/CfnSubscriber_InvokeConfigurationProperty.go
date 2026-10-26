package awseventsv2


// Configuration for how the subscriber invokes its target.
//
// Specify the target ARN, the IAM role, and, optionally, the target-specific parameters object that matches the target type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invokeConfigurationProperty := &InvokeConfigurationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	EventBusV2Parameters: &EventBusV2ParametersProperty{
//   		DeduplicationConfiguration: &DeduplicationConfigurationProperty{
//   			DeduplicationType: jsii.String("deduplicationType"),
//   		},
//   		Metadata: map[string]*string{
//   			"metadataKey": jsii.String("metadata"),
//   		},
//   		SystemMetadata: &EventBusV2SystemMetadataProperty{
//   			DeduplicationId: jsii.String("deduplicationId"),
//   			EventGroupId: jsii.String("eventGroupId"),
//   		},
//   	},
//   	HttpParameters: &HttpParametersProperty{
//   		HeaderParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   		PathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		QueryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	KinesisParameters: &KinesisParametersProperty{
//   		ExplicitHashKey: jsii.String("explicitHashKey"),
//   		PartitionKey: jsii.String("partitionKey"),
//   	},
//   	LambdaParameters: &LambdaParametersProperty{
//   		DurableExecutionName: jsii.String("durableExecutionName"),
//   		InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   		InvocationType: jsii.String("invocationType"),
//   		Qualifier: jsii.String("qualifier"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	SnsParameters: &SnsParametersProperty{
//   		MessageAttributes: map[string]interface{}{
//   			"messageAttributesKey": &MessageAttributeValueProperty{
//   				"dataType": jsii.String("dataType"),
//
//   				// the properties below are optional
//   				"binaryValue": jsii.String("binaryValue"),
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   		MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   		MessageGroupId: jsii.String("messageGroupId"),
//   		MessageStructure: jsii.String("messageStructure"),
//   		Subject: jsii.String("subject"),
//   	},
//   	SqsParameters: &SqsParametersProperty{
//   		DelaySeconds: jsii.String("delaySeconds"),
//   		MessageAttributes: map[string]interface{}{
//   			"messageAttributesKey": &MessageAttributeValueProperty{
//   				"dataType": jsii.String("dataType"),
//
//   				// the properties below are optional
//   				"binaryValue": jsii.String("binaryValue"),
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   		MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   		MessageGroupId: jsii.String("messageGroupId"),
//   		MessageSystemAttributes: map[string]interface{}{
//   			"messageSystemAttributesKey": &MessageAttributeValueProperty{
//   				"dataType": jsii.String("dataType"),
//
//   				// the properties below are optional
//   				"binaryValue": jsii.String("binaryValue"),
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	StepFunctionsParameters: &StepFunctionsParametersProperty{
//   		InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   		InvocationType: jsii.String("invocationType"),
//   		Name: jsii.String("name"),
//   		TraceHeader: jsii.String("traceHeader"),
//   	},
//   	UniversalTargetParameters: &UniversalTargetParametersProperty{
//   		Input: jsii.String("input"),
//
//   		// the properties below are optional
//   		InvocationTimeoutSeconds: jsii.String("invocationTimeoutSeconds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html
//
type CfnSubscriber_InvokeConfigurationProperty struct {
	// The ARN of the IAM role the service assumes to invoke the target.
	//
	// The role must belong to the same account as the subscriber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the target that the subscriber invokes.
	//
	// For universal service integration targets, use the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Parameters for forwarding events to another EventBridge event bus, used when TargetArn is an event bus ARN of the form arn:{partition}:events:{region}:{account}:event-busv2/{name}/{id}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-eventbusv2parameters
	//
	EventBusV2Parameters interface{} `field:"optional" json:"eventBusV2Parameters" yaml:"eventBusV2Parameters"`
	// Parameters for invoking an HTTP endpoint target, such as an Amazon API Gateway endpoint or an EventBridge API destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-httpparameters
	//
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Parameters for writing events to an Amazon Kinesis Data Streams target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-kinesisparameters
	//
	KinesisParameters interface{} `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// Parameters for invoking an AWS Lambda function target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-lambdaparameters
	//
	LambdaParameters interface{} `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// Parameters for publishing events to an Amazon SNS topic target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-snsparameters
	//
	SnsParameters interface{} `field:"optional" json:"snsParameters" yaml:"snsParameters"`
	// Parameters for sending events to an Amazon SQS queue target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-sqsparameters
	//
	SqsParameters interface{} `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
	// Parameters for starting an AWS Step Functions state machine execution target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-stepfunctionsparameters
	//
	StepFunctionsParameters interface{} `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
	// Parameters for invoking an AWS service API as a universal service integration target, used when TargetArn has the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-invokeconfiguration.html#cfn-eventsv2-subscriber-invokeconfiguration-universaltargetparameters
	//
	UniversalTargetParameters interface{} `field:"optional" json:"universalTargetParameters" yaml:"universalTargetParameters"`
}

