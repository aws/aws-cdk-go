package awssns

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.BetweenCondition",
		reflect.TypeOf((*BetweenCondition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.CfnSubscription",
		reflect.TypeOf((*CfnSubscription)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "deliveryPolicy", GoGetter: "DeliveryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "filterPolicy", GoGetter: "FilterPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "filterPolicyScope", GoGetter: "FilterPolicyScope"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "rawMessageDelivery", GoGetter: "RawMessageDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicy", GoGetter: "RedrivePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "replayPolicy", GoGetter: "ReplayPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionRoleArn", GoGetter: "SubscriptionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicArn", GoGetter: "TopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnSubscriptionProps",
		reflect.TypeOf((*CfnSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.CfnTopic",
		reflect.TypeOf((*CfnTopic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "archivePolicy", GoGetter: "ArchivePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTopicArn", GoGetter: "AttrTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTopicName", GoGetter: "AttrTopicName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataProtectionPolicy", GoGetter: "DataProtectionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStatusLogging", GoGetter: "DeliveryStatusLogging"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "fifoTopic", GoGetter: "FifoTopic"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyId", GoGetter: "KmsMasterKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "signatureVersion", GoGetter: "SignatureVersion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subscription", GoGetter: "Subscription"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "topicName", GoGetter: "TopicName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tracingConfig", GoGetter: "TracingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTopic{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnTopic.LoggingConfigProperty",
		reflect.TypeOf((*CfnTopic_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnTopic.SubscriptionProperty",
		reflect.TypeOf((*CfnTopic_SubscriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.CfnTopicInlinePolicy",
		reflect.TypeOf((*CfnTopicInlinePolicy)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicArn", GoGetter: "TopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTopicInlinePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnTopicInlinePolicyProps",
		reflect.TypeOf((*CfnTopicInlinePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.CfnTopicPolicy",
		reflect.TypeOf((*CfnTopicPolicy)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topics", GoGetter: "Topics"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTopicPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnTopicPolicyProps",
		reflect.TypeOf((*CfnTopicPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.CfnTopicProps",
		reflect.TypeOf((*CfnTopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.Filter",
		reflect.TypeOf((*Filter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filterDoc", GoGetter: "FilterDoc"},
			_jsii_.MemberMethod{JsiiMethod: "isFilter", GoMethod: "IsFilter"},
			_jsii_.MemberMethod{JsiiMethod: "isPolicy", GoMethod: "IsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Filter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FilterOrPolicy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.FilterOrPolicy",
		reflect.TypeOf((*FilterOrPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isFilter", GoMethod: "IsFilter"},
			_jsii_.MemberMethod{JsiiMethod: "isPolicy", GoMethod: "IsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_FilterOrPolicy{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sns.FilterOrPolicyType",
		reflect.TypeOf((*FilterOrPolicyType)(nil)).Elem(),
		map[string]interface{}{
			"FILTER": FilterOrPolicyType_FILTER,
			"POLICY": FilterOrPolicyType_POLICY,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_sns.ITopic",
		reflect.TypeOf((*ITopic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleTarget", GoMethod: "BindAsNotificationRuleTarget"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesPublished", GoMethod: "MetricNumberOfMessagesPublished"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsDelivered", GoMethod: "MetricNumberOfNotificationsDelivered"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFailed", GoMethod: "MetricNumberOfNotificationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOut", GoMethod: "MetricNumberOfNotificationsFilteredOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutInvalidAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutInvalidAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutNoMessageAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutNoMessageAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishSize", GoMethod: "MetricPublishSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSMonthToDateSpentUSD", GoMethod: "MetricSMSMonthToDateSpentUSD"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSSuccessRate", GoMethod: "MetricSMSSuccessRate"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicArn", GoGetter: "TopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicName", GoGetter: "TopicName"},
		},
		func() interface{} {
			j := jsiiProxy_ITopic{}
			_jsii_.InitJsiiProxy(&j.Type__awscodestarnotificationsINotificationRuleTarget)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_sns.ITopicSubscription",
		reflect.TypeOf((*ITopicSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ITopicSubscription{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.LoggingConfig",
		reflect.TypeOf((*LoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sns.LoggingProtocol",
		reflect.TypeOf((*LoggingProtocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": LoggingProtocol_HTTP,
			"SQS": LoggingProtocol_SQS,
			"LAMBDA": LoggingProtocol_LAMBDA,
			"FIREHOSE": LoggingProtocol_FIREHOSE,
			"APPLICATION": LoggingProtocol_APPLICATION,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.NumericConditions",
		reflect.TypeOf((*NumericConditions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.Policy",
		reflect.TypeOf((*Policy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isFilter", GoMethod: "IsFilter"},
			_jsii_.MemberMethod{JsiiMethod: "isPolicy", GoMethod: "IsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "policyDoc", GoGetter: "PolicyDoc"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Policy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FilterOrPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.StringConditions",
		reflect.TypeOf((*StringConditions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.Subscription",
		reflect.TypeOf((*Subscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Subscription{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.SubscriptionFilter",
		reflect.TypeOf((*SubscriptionFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "conditions", GoGetter: "Conditions"},
		},
		func() interface{} {
			return &jsiiProxy_SubscriptionFilter{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.SubscriptionOptions",
		reflect.TypeOf((*SubscriptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.SubscriptionProps",
		reflect.TypeOf((*SubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_sns.SubscriptionProtocol",
		reflect.TypeOf((*SubscriptionProtocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": SubscriptionProtocol_HTTP,
			"HTTPS": SubscriptionProtocol_HTTPS,
			"EMAIL": SubscriptionProtocol_EMAIL,
			"EMAIL_JSON": SubscriptionProtocol_EMAIL_JSON,
			"SMS": SubscriptionProtocol_SMS,
			"SQS": SubscriptionProtocol_SQS,
			"APPLICATION": SubscriptionProtocol_APPLICATION,
			"LAMBDA": SubscriptionProtocol_LAMBDA,
			"FIREHOSE": SubscriptionProtocol_FIREHOSE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.Topic",
		reflect.TypeOf((*Topic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLoggingConfig", GoMethod: "AddLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleTarget", GoMethod: "BindAsNotificationRuleTarget"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberMethod{JsiiMethod: "createSSLPolicyDocument", GoMethod: "CreateSSLPolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "enforceSSL", GoGetter: "EnforceSSL"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesPublished", GoMethod: "MetricNumberOfMessagesPublished"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsDelivered", GoMethod: "MetricNumberOfNotificationsDelivered"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFailed", GoMethod: "MetricNumberOfNotificationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOut", GoMethod: "MetricNumberOfNotificationsFilteredOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutInvalidAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutInvalidAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutNoMessageAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutNoMessageAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishSize", GoMethod: "MetricPublishSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSMonthToDateSpentUSD", GoMethod: "MetricSMSMonthToDateSpentUSD"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSSuccessRate", GoMethod: "MetricSMSSuccessRate"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicArn", GoGetter: "TopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicName", GoGetter: "TopicName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Topic{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TopicBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.TopicBase",
		reflect.TypeOf((*TopicBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSubscription", GoMethod: "AddSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleTarget", GoMethod: "BindAsNotificationRuleTarget"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberMethod{JsiiMethod: "createSSLPolicyDocument", GoMethod: "CreateSSLPolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "enforceSSL", GoGetter: "EnforceSSL"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fifo", GoGetter: "Fifo"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublish", GoMethod: "GrantPublish"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfMessagesPublished", GoMethod: "MetricNumberOfMessagesPublished"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsDelivered", GoMethod: "MetricNumberOfNotificationsDelivered"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFailed", GoMethod: "MetricNumberOfNotificationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOut", GoMethod: "MetricNumberOfNotificationsFilteredOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutInvalidAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutInvalidAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfNotificationsFilteredOutNoMessageAttributes", GoMethod: "MetricNumberOfNotificationsFilteredOutNoMessageAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishSize", GoMethod: "MetricPublishSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSMonthToDateSpentUSD", GoMethod: "MetricSMSMonthToDateSpentUSD"},
			_jsii_.MemberMethod{JsiiMethod: "metricSMSSuccessRate", GoMethod: "MetricSMSSuccessRate"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicArn", GoGetter: "TopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "topicName", GoGetter: "TopicName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TopicBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITopic)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns.TopicPolicy",
		reflect.TypeOf((*TopicPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createSSLPolicyDocument", GoMethod: "CreateSSLPolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TopicPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.TopicPolicyProps",
		reflect.TypeOf((*TopicPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.TopicProps",
		reflect.TypeOf((*TopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns.TopicSubscriptionConfig",
		reflect.TypeOf((*TopicSubscriptionConfig)(nil)).Elem(),
	)
}
