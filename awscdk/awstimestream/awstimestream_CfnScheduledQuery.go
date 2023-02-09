package awstimestream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Timestream::ScheduledQuery`.
//
// Create a scheduled query that will be run on your behalf at the configured schedule. Timestream assumes the execution role provided as part of the `ScheduledQueryExecutionRoleArn` parameter to run the query. You can use the `NotificationConfiguration` parameter to configure notification for your scheduled query operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledQuery := awscdk.Aws_timestream.NewCfnScheduledQuery(this, jsii.String("MyCfnScheduledQuery"), &cfnScheduledQueryProps{
//   	errorReportConfiguration: &errorReportConfigurationProperty{
//   		s3Configuration: &s3ConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			encryptionOption: jsii.String("encryptionOption"),
//   			objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   		},
//   	},
//   	notificationConfiguration: &notificationConfigurationProperty{
//   		snsConfiguration: &snsConfigurationProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   	},
//   	queryString: jsii.String("queryString"),
//   	scheduleConfiguration: &scheduleConfigurationProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	scheduledQueryExecutionRoleArn: jsii.String("scheduledQueryExecutionRoleArn"),
//
//   	// the properties below are optional
//   	clientToken: jsii.String("clientToken"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	scheduledQueryName: jsii.String("scheduledQueryName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetConfiguration: &targetConfigurationProperty{
//   		timestreamConfiguration: &timestreamConfigurationProperty{
//   			databaseName: jsii.String("databaseName"),
//   			dimensionMappings: []interface{}{
//   				&dimensionMappingProperty{
//   					dimensionValueType: jsii.String("dimensionValueType"),
//   					name: jsii.String("name"),
//   				},
//   			},
//   			tableName: jsii.String("tableName"),
//   			timeColumn: jsii.String("timeColumn"),
//
//   			// the properties below are optional
//   			measureNameColumn: jsii.String("measureNameColumn"),
//   			mixedMeasureMappings: []interface{}{
//   				&mixedMeasureMappingProperty{
//   					measureValueType: jsii.String("measureValueType"),
//
//   					// the properties below are optional
//   					measureName: jsii.String("measureName"),
//   					multiMeasureAttributeMappings: []interface{}{
//   						&multiMeasureAttributeMappingProperty{
//   							measureValueType: jsii.String("measureValueType"),
//   							sourceColumn: jsii.String("sourceColumn"),
//
//   							// the properties below are optional
//   							targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   						},
//   					},
//   					sourceColumn: jsii.String("sourceColumn"),
//   					targetMeasureName: jsii.String("targetMeasureName"),
//   				},
//   			},
//   			multiMeasureMappings: &multiMeasureMappingsProperty{
//   				multiMeasureAttributeMappings: []interface{}{
//   					&multiMeasureAttributeMappingProperty{
//   						measureValueType: jsii.String("measureValueType"),
//   						sourceColumn: jsii.String("sourceColumn"),
//
//   						// the properties below are optional
//   						targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				targetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   			},
//   		},
//   	},
//   })
//
type CfnScheduledQuery interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The `ARN` of the scheduled query.
	AttrArn() *string
	// The scheduled query error reporting configuration.
	AttrSqErrorReportConfiguration() *string
	// The KMS key used to encrypt the query resource, if a customer managed KMS key was provided.
	AttrSqKmsKeyId() *string
	// The scheduled query name.
	AttrSqName() *string
	// The scheduled query notification configuration.
	AttrSqNotificationConfiguration() *string
	// The scheduled query string..
	AttrSqQueryString() *string
	// The scheduled query schedule configuration.
	AttrSqScheduleConfiguration() *string
	// The ARN of the IAM role that will be used by Timestream to run the query.
	AttrSqScheduledQueryExecutionRoleArn() *string
	// The configuration for query output.
	AttrSqTargetConfiguration() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.
	//
	// Making multiple identical CreateScheduledQuery requests has the same effect as making a single request.
	//
	// - If CreateScheduledQuery is called without a `ClientToken` , the Query SDK generates a `ClientToken` on your behalf.
	// - After 8 hours, any request with the same `ClientToken` is treated as a new request.
	ClientToken() *string
	SetClientToken(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Configuration for error reporting.
	//
	// Error reports will be generated when a problem is encountered when writing the query results.
	ErrorReportConfiguration() interface{}
	SetErrorReportConfiguration(val interface{})
	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest.
	//
	// If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with *alias/*
	//
	// If ErrorReportConfiguration uses `SSE_KMS` as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Notification configuration for the scheduled query.
	//
	// A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
	NotificationConfiguration() interface{}
	SetNotificationConfiguration(val interface{})
	// The query string to run.
	//
	// Parameter names can be specified in the query string `@` character followed by an identifier. The named Parameter `@scheduled_runtime` is reserved and can be used in the query to get the time at which the query is scheduled to run.
	//
	// The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of `@scheduled_runtime` paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the `@scheduled_runtime` parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.
	QueryString() *string
	SetQueryString(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Schedule configuration.
	ScheduleConfiguration() interface{}
	SetScheduleConfiguration(val interface{})
	// The ARN for the IAM role that Timestream will assume when running the scheduled query.
	ScheduledQueryExecutionRoleArn() *string
	SetScheduledQueryExecutionRoleArn(val *string)
	// A name for the query.
	//
	// Scheduled query names must be unique within each Region.
	ScheduledQueryName() *string
	SetScheduledQueryName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of key-value pairs to label the scheduled query.
	Tags() awscdk.TagManager
	// Scheduled query target store configuration.
	TargetConfiguration() interface{}
	SetTargetConfiguration(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScheduledQuery
type jsiiProxy_CfnScheduledQuery struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScheduledQuery) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqErrorReportConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqErrorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqNotificationConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqNotificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqScheduleConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqScheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqScheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqTargetConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqTargetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ErrorReportConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) NotificationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduleConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduledQueryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) TargetConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Timestream::ScheduledQuery`.
func NewCfnScheduledQuery(scope awscdk.Construct, id *string, props *CfnScheduledQueryProps) CfnScheduledQuery {
	_init_.Initialize()

	if err := validateNewCfnScheduledQueryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScheduledQuery{}

	_jsii_.Create(
		"monocdk.aws_timestream.CfnScheduledQuery",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Timestream::ScheduledQuery`.
func NewCfnScheduledQuery_Override(c CfnScheduledQuery, scope awscdk.Construct, id *string, props *CfnScheduledQueryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_timestream.CfnScheduledQuery",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetClientToken(val *string) {
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetErrorReportConfiguration(val interface{}) {
	if err := j.validateSetErrorReportConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorReportConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetNotificationConfiguration(val interface{}) {
	if err := j.validateSetNotificationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetScheduleConfiguration(val interface{}) {
	if err := j.validateSetScheduleConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetScheduledQueryExecutionRoleArn(val *string) {
	if err := j.validateSetScheduledQueryExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledQueryExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetScheduledQueryName(val *string) {
	_jsii_.Set(
		j,
		"scheduledQueryName",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery)SetTargetConfiguration(val interface{}) {
	if err := j.validateSetTargetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnScheduledQuery_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledQuery_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_timestream.CfnScheduledQuery",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScheduledQuery_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledQuery_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_timestream.CfnScheduledQuery",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnScheduledQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledQuery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_timestream.CfnScheduledQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledQuery_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_timestream.CfnScheduledQuery",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnScheduledQuery) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScheduledQuery) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScheduledQuery) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnScheduledQuery) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

