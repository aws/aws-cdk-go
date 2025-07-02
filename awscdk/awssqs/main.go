package awssqs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.CfnQueue",
		reflect.TypeOf((*CfnQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrQueueName", GoGetter: "AttrQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "attrQueueUrl", GoGetter: "AttrQueueUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deduplicationScope", GoGetter: "DeduplicationScope"},
			_jsii_.MemberProperty{JsiiProperty: "delaySeconds", GoGetter: "DelaySeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fifoQueue", GoGetter: "FifoQueue"},
			_jsii_.MemberProperty{JsiiProperty: "fifoThroughputLimit", GoGetter: "FifoThroughputLimit"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsDataKeyReusePeriodSeconds", GoGetter: "KmsDataKeyReusePeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyId", GoGetter: "KmsMasterKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maximumMessageSize", GoGetter: "MaximumMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "messageRetentionPeriod", GoGetter: "MessageRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "receiveMessageWaitTimeSeconds", GoGetter: "ReceiveMessageWaitTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "redriveAllowPolicy", GoGetter: "RedriveAllowPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicy", GoGetter: "RedrivePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sqsManagedSseEnabled", GoGetter: "SqsManagedSseEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityTimeout", GoGetter: "VisibilityTimeout"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.CfnQueueInlinePolicy",
		reflect.TypeOf((*CfnQueueInlinePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "queue", GoGetter: "Queue"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueInlinePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.CfnQueueInlinePolicyProps",
		reflect.TypeOf((*CfnQueueInlinePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.CfnQueuePolicy",
		reflect.TypeOf((*CfnQueuePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "queues", GoGetter: "Queues"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.CfnQueuePolicyProps",
		reflect.TypeOf((*CfnQueuePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.CfnQueueProps",
		reflect.TypeOf((*CfnQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.DeadLetterQueue",
		reflect.TypeOf((*DeadLetterQueue)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sqs.DeduplicationScope",
		reflect.TypeOf((*DeduplicationScope)(nil)).Elem(),
		map[string]interface{}{
			"MESSAGE_GROUP": DeduplicationScope_MESSAGE_GROUP,
			"QUEUE": DeduplicationScope_QUEUE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sqs.FifoThroughputLimit",
		reflect.TypeOf((*FifoThroughputLimit)(nil)).Elem(),
		map[string]interface{}{
			"PER_QUEUE": FifoThroughputLimit_PER_QUEUE,
			"PER_MESSAGE_GROUP_ID": FifoThroughputLimit_PER_MESSAGE_GROUP_ID,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_sqs.IQueue",
		reflect.TypeOf((*IQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionMasterKey", GoGetter: "EncryptionMasterKey"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionType", GoGetter: "EncryptionType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConsumeMessages", GoMethod: "GrantConsumeMessages"},
			_jsii_.MemberMethod{JsiiMethod: "grantPurge", GoMethod: "GrantPurge"},
			_jsii_.MemberMethod{JsiiMethod: "grantSendMessages", GoMethod: "GrantSendMessages"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateAgeOfOldestMessage", GoMethod: "MetricApproximateAgeOfOldestMessage"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesDelayed", GoMethod: "MetricApproximateNumberOfMessagesDelayed"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesNotVisible", GoMethod: "MetricApproximateNumberOfMessagesNotVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesVisible", GoMethod: "MetricApproximateNumberOfMessagesVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfEmptyReceives", GoMethod: "MetricNumberOfEmptyReceives"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesDeleted", GoMethod: "MetricNumberOfMessagesDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesReceived", GoMethod: "MetricNumberOfMessagesReceived"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesSent", GoMethod: "MetricNumberOfMessagesSent"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentMessageSize", GoMethod: "MetricSentMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "queueArn", GoGetter: "QueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "queueUrl", GoGetter: "QueueUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.Queue",
		reflect.TypeOf((*Queue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionMasterKey", GoGetter: "EncryptionMasterKey"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionType", GoGetter: "EncryptionType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConsumeMessages", GoMethod: "GrantConsumeMessages"},
			_jsii_.MemberMethod{JsiiMethod: "grantPurge", GoMethod: "GrantPurge"},
			_jsii_.MemberMethod{JsiiMethod: "grantSendMessages", GoMethod: "GrantSendMessages"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateAgeOfOldestMessage", GoMethod: "MetricApproximateAgeOfOldestMessage"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesDelayed", GoMethod: "MetricApproximateNumberOfMessagesDelayed"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesNotVisible", GoMethod: "MetricApproximateNumberOfMessagesNotVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesVisible", GoMethod: "MetricApproximateNumberOfMessagesVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfEmptyReceives", GoMethod: "MetricNumberOfEmptyReceives"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesDeleted", GoMethod: "MetricNumberOfMessagesDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesReceived", GoMethod: "MetricNumberOfMessagesReceived"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesSent", GoMethod: "MetricNumberOfMessagesSent"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentMessageSize", GoMethod: "MetricSentMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queueArn", GoGetter: "QueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "queueUrl", GoGetter: "QueueUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Queue{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_QueueBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.QueueAttributes",
		reflect.TypeOf((*QueueAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.QueueBase",
		reflect.TypeOf((*QueueBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionMasterKey", GoGetter: "EncryptionMasterKey"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionType", GoGetter: "EncryptionType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConsumeMessages", GoMethod: "GrantConsumeMessages"},
			_jsii_.MemberMethod{JsiiMethod: "grantPurge", GoMethod: "GrantPurge"},
			_jsii_.MemberMethod{JsiiMethod: "grantSendMessages", GoMethod: "GrantSendMessages"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateAgeOfOldestMessage", GoMethod: "MetricApproximateAgeOfOldestMessage"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesDelayed", GoMethod: "MetricApproximateNumberOfMessagesDelayed"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesNotVisible", GoMethod: "MetricApproximateNumberOfMessagesNotVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricApproximateNumberOfMessagesVisible", GoMethod: "MetricApproximateNumberOfMessagesVisible"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfEmptyReceives", GoMethod: "MetricNumberOfEmptyReceives"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesDeleted", GoMethod: "MetricNumberOfMessagesDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesReceived", GoMethod: "MetricNumberOfMessagesReceived"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesSent", GoMethod: "MetricNumberOfMessagesSent"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentMessageSize", GoMethod: "MetricSentMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queueArn", GoGetter: "QueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "queueUrl", GoGetter: "QueueUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_QueueBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IQueue)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sqs.QueueEncryption",
		reflect.TypeOf((*QueueEncryption)(nil)).Elem(),
		map[string]interface{}{
			"UNENCRYPTED": QueueEncryption_UNENCRYPTED,
			"KMS_MANAGED": QueueEncryption_KMS_MANAGED,
			"KMS": QueueEncryption_KMS,
			"SQS_MANAGED": QueueEncryption_SQS_MANAGED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sqs.QueuePolicy",
		reflect.TypeOf((*QueuePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queuePolicyId", GoGetter: "QueuePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_QueuePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.QueuePolicyProps",
		reflect.TypeOf((*QueuePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.QueueProps",
		reflect.TypeOf((*QueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sqs.RedriveAllowPolicy",
		reflect.TypeOf((*RedriveAllowPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sqs.RedrivePermission",
		reflect.TypeOf((*RedrivePermission)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW_ALL": RedrivePermission_ALLOW_ALL,
			"DENY_ALL": RedrivePermission_DENY_ALL,
			"BY_QUEUE": RedrivePermission_BY_QUEUE,
		},
	)
}
