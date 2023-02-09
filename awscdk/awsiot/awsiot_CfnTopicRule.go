package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoT::TopicRule`.
//
// Use the `AWS::IoT::TopicRule` resource to declare an AWS IoT rule. For information about working with AWS IoT rules, see [Rules for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicRule := awscdk.Aws_iot.NewCfnTopicRule(this, jsii.String("MyCfnTopicRule"), &cfnTopicRuleProps{
//   	topicRulePayload: &topicRulePayloadProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   					alarmName: jsii.String("alarmName"),
//   					roleArn: jsii.String("roleArn"),
//   					stateReason: jsii.String("stateReason"),
//   					stateValue: jsii.String("stateValue"),
//   				},
//   				cloudwatchLogs: &cloudwatchLogsActionProperty{
//   					logGroupName: jsii.String("logGroupName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   				},
//   				cloudwatchMetric: &cloudwatchMetricActionProperty{
//   					metricName: jsii.String("metricName"),
//   					metricNamespace: jsii.String("metricNamespace"),
//   					metricUnit: jsii.String("metricUnit"),
//   					metricValue: jsii.String("metricValue"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					metricTimestamp: jsii.String("metricTimestamp"),
//   				},
//   				dynamoDb: &dynamoDBActionProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2ActionProperty{
//   					putItem: &putItemInputProperty{
//   						tableName: jsii.String("tableName"),
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				elasticsearch: &elasticsearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				firehose: &firehoseActionProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					separator: jsii.String("separator"),
//   				},
//   				http: &httpActionProperty{
//   					url: jsii.String("url"),
//
//   					// the properties below are optional
//   					auth: &httpAuthorizationProperty{
//   						sigv4: &sigV4AuthorizationProperty{
//   							roleArn: jsii.String("roleArn"),
//   							serviceName: jsii.String("serviceName"),
//   							signingRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					confirmationUrl: jsii.String("confirmationUrl"),
//   					headers: []interface{}{
//   						&httpActionHeaderProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				iotAnalytics: &iotAnalyticsActionProperty{
//   					channelName: jsii.String("channelName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   				},
//   				iotEvents: &iotEventsActionProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					messageId: jsii.String("messageId"),
//   				},
//   				iotSiteWise: &iotSiteWiseActionProperty{
//   					putAssetPropertyValueEntries: []interface{}{
//   						&putAssetPropertyValueEntryProperty{
//   							propertyValues: []interface{}{
//   								&assetPropertyValueProperty{
//   									timestamp: &assetPropertyTimestampProperty{
//   										timeInSeconds: jsii.String("timeInSeconds"),
//
//   										// the properties below are optional
//   										offsetInNanos: jsii.String("offsetInNanos"),
//   									},
//   									value: &assetPropertyVariantProperty{
//   										booleanValue: jsii.String("booleanValue"),
//   										doubleValue: jsii.String("doubleValue"),
//   										integerValue: jsii.String("integerValue"),
//   										stringValue: jsii.String("stringValue"),
//   									},
//
//   									// the properties below are optional
//   									quality: jsii.String("quality"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				kafka: &kafkaActionProperty{
//   					clientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					destinationArn: jsii.String("destinationArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					key: jsii.String("key"),
//   					partition: jsii.String("partition"),
//   				},
//   				kinesis: &kinesisActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					streamName: jsii.String("streamName"),
//
//   					// the properties below are optional
//   					partitionKey: jsii.String("partitionKey"),
//   				},
//   				lambda: &lambdaActionProperty{
//   					functionArn: jsii.String("functionArn"),
//   				},
//   				location: &locationActionProperty{
//   					deviceId: jsii.String("deviceId"),
//   					latitude: jsii.String("latitude"),
//   					longitude: jsii.String("longitude"),
//   					roleArn: jsii.String("roleArn"),
//   					trackerName: jsii.String("trackerName"),
//
//   					// the properties below are optional
//   					timestamp: NewDate(),
//   				},
//   				openSearch: &openSearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				republish: &republishActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					headers: &republishActionHeadersProperty{
//   						contentType: jsii.String("contentType"),
//   						correlationData: jsii.String("correlationData"),
//   						messageExpiry: jsii.String("messageExpiry"),
//   						payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   						responseTopic: jsii.String("responseTopic"),
//   						userProperties: []interface{}{
//   							&userPropertyProperty{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					qos: jsii.Number(123),
//   				},
//   				s3: &s3ActionProperty{
//   					bucketName: jsii.String("bucketName"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   				sns: &snsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					messageFormat: jsii.String("messageFormat"),
//   				},
//   				sqs: &sqsActionProperty{
//   					queueUrl: jsii.String("queueUrl"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					useBase64: jsii.Boolean(false),
//   				},
//   				stepFunctions: &stepFunctionsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					stateMachineName: jsii.String("stateMachineName"),
//
//   					// the properties below are optional
//   					executionNamePrefix: jsii.String("executionNamePrefix"),
//   				},
//   				timestream: &timestreamActionProperty{
//   					databaseName: jsii.String("databaseName"),
//   					dimensions: []interface{}{
//   						&timestreamDimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					timestamp: &timestreamTimestampProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		sql: jsii.String("sql"),
//
//   		// the properties below are optional
//   		awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   		description: jsii.String("description"),
//   		errorAction: &actionProperty{
//   			cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   				alarmName: jsii.String("alarmName"),
//   				roleArn: jsii.String("roleArn"),
//   				stateReason: jsii.String("stateReason"),
//   				stateValue: jsii.String("stateValue"),
//   			},
//   			cloudwatchLogs: &cloudwatchLogsActionProperty{
//   				logGroupName: jsii.String("logGroupName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			cloudwatchMetric: &cloudwatchMetricActionProperty{
//   				metricName: jsii.String("metricName"),
//   				metricNamespace: jsii.String("metricNamespace"),
//   				metricUnit: jsii.String("metricUnit"),
//   				metricValue: jsii.String("metricValue"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				metricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			dynamoDb: &dynamoDBActionProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2ActionProperty{
//   				putItem: &putItemInputProperty{
//   					tableName: jsii.String("tableName"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			elasticsearch: &elasticsearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			firehose: &firehoseActionProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				separator: jsii.String("separator"),
//   			},
//   			http: &httpActionProperty{
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				auth: &httpAuthorizationProperty{
//   					sigv4: &sigV4AuthorizationProperty{
//   						roleArn: jsii.String("roleArn"),
//   						serviceName: jsii.String("serviceName"),
//   						signingRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				confirmationUrl: jsii.String("confirmationUrl"),
//   				headers: []interface{}{
//   					&httpActionHeaderProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			iotAnalytics: &iotAnalyticsActionProperty{
//   				channelName: jsii.String("channelName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			iotEvents: &iotEventsActionProperty{
//   				inputName: jsii.String("inputName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				messageId: jsii.String("messageId"),
//   			},
//   			iotSiteWise: &iotSiteWiseActionProperty{
//   				putAssetPropertyValueEntries: []interface{}{
//   					&putAssetPropertyValueEntryProperty{
//   						propertyValues: []interface{}{
//   							&assetPropertyValueProperty{
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kafka: &kafkaActionProperty{
//   				clientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				destinationArn: jsii.String("destinationArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				partition: jsii.String("partition"),
//   			},
//   			kinesis: &kinesisActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				partitionKey: jsii.String("partitionKey"),
//   			},
//   			lambda: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//   			},
//   			location: &locationActionProperty{
//   				deviceId: jsii.String("deviceId"),
//   				latitude: jsii.String("latitude"),
//   				longitude: jsii.String("longitude"),
//   				roleArn: jsii.String("roleArn"),
//   				trackerName: jsii.String("trackerName"),
//
//   				// the properties below are optional
//   				timestamp: NewDate(),
//   			},
//   			openSearch: &openSearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			republish: &republishActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				headers: &republishActionHeadersProperty{
//   					contentType: jsii.String("contentType"),
//   					correlationData: jsii.String("correlationData"),
//   					messageExpiry: jsii.String("messageExpiry"),
//   					payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   					responseTopic: jsii.String("responseTopic"),
//   					userProperties: []interface{}{
//   						&userPropertyProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				qos: jsii.Number(123),
//   			},
//   			s3: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//   				key: jsii.String("key"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			sns: &snsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				messageFormat: jsii.String("messageFormat"),
//   			},
//   			sqs: &sqsActionProperty{
//   				queueUrl: jsii.String("queueUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				useBase64: jsii.Boolean(false),
//   			},
//   			stepFunctions: &stepFunctionsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				stateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				executionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			timestream: &timestreamActionProperty{
//   				databaseName: jsii.String("databaseName"),
//   				dimensions: []interface{}{
//   					&timestreamDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				timestamp: &timestreamTimestampProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		ruleDisabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	ruleName: jsii.String("ruleName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnTopicRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name of the rule.
	RuleName() *string
	SetRuleName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags() awscdk.TagManager
	// The rule payload.
	TopicRulePayload() interface{}
	SetTopicRulePayload(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTopicRule
type jsiiProxy_CfnTopicRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopicRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) TopicRulePayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicRulePayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule(scope constructs.Construct, id *string, props *CfnTopicRuleProps) CfnTopicRule {
	_init_.Initialize()

	if err := validateNewCfnTopicRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule_Override(c CfnTopicRule, scope constructs.Construct, id *string, props *CfnTopicRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopicRule)SetRuleName(val *string) {
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_CfnTopicRule)SetTopicRulePayload(val interface{}) {
	if err := j.validateSetTopicRulePayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicRulePayload",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTopicRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTopicRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicRule) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopicRule) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopicRule) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopicRule) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTopicRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

