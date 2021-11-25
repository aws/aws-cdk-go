package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for adding an event selector.
//
// TODO: EXAMPLE
//
// Experimental.
type AddEventSelectorOptions struct {
	// An optional list of service event sources from which you do not want management events to be logged on your trail.
	// Experimental.
	ExcludeManagementEventSources *[]ManagementEventSources `json:"excludeManagementEventSources"`
	// Specifies whether the event selector includes management events for the trail.
	// Experimental.
	IncludeManagementEvents *bool `json:"includeManagementEvents"`
	// Specifies whether to log read-only events, write-only events, or all events.
	// Experimental.
	ReadWriteType ReadWriteType `json:"readWriteType"`
}

// A CloudFormation `AWS::CloudTrail::Trail`.
//
// TODO: EXAMPLE
//
type CfnTrail interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrSnsTopicArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CloudWatchLogsLogGroupArn() *string
	SetCloudWatchLogsLogGroupArn(val *string)
	CloudWatchLogsRoleArn() *string
	SetCloudWatchLogsRoleArn(val *string)
	CreationStack() *[]*string
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	EventSelectors() interface{}
	SetEventSelectors(val interface{})
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	InsightSelectors() interface{}
	SetInsightSelectors(val interface{})
	IsLogging() interface{}
	SetIsLogging(val interface{})
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TrailName() *string
	SetTrailName(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTrail
type jsiiProxy_CfnTrail struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTrail) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) AttrSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) InsightSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) TrailName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail(scope constructs.Construct, id *string, props *CfnTrailProps) CfnTrail {
	_init_.Initialize()

	j := jsiiProxy_CfnTrail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail_Override(c CfnTrail, scope constructs.Construct, id *string, props *CfnTrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTrail) SetCloudWatchLogsLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetCloudWatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetEnableLogFileValidation(val interface{}) {
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetEventSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"eventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIncludeGlobalServiceEvents(val interface{}) {
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetInsightSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"insightSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsLogging(val interface{}) {
	_jsii_.Set(
		j,
		"isLogging",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsMultiRegionTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsOrganizationTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetSnsTopicName(val *string) {
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetTrailName(val *string) {
	_jsii_.Set(
		j,
		"trailName",
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
func CfnTrail_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTrail_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnTrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrail_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnTrail) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnTrail) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnTrail) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnTrail) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnTrail) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnTrail) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnTrail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnTrail) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnTrail) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnTrail) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnTrail) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTrail) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnTrail) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnTrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnTrail) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnTrail_DataResourceProperty struct {
	// `CfnTrail.DataResourceProperty.Type`.
	Type *string `json:"type"`
	// `CfnTrail.DataResourceProperty.Values`.
	Values *[]*string `json:"values"`
}

// TODO: EXAMPLE
//
type CfnTrail_EventSelectorProperty struct {
	// `CfnTrail.EventSelectorProperty.DataResources`.
	DataResources interface{} `json:"dataResources"`
	// `CfnTrail.EventSelectorProperty.ExcludeManagementEventSources`.
	ExcludeManagementEventSources *[]*string `json:"excludeManagementEventSources"`
	// `CfnTrail.EventSelectorProperty.IncludeManagementEvents`.
	IncludeManagementEvents interface{} `json:"includeManagementEvents"`
	// `CfnTrail.EventSelectorProperty.ReadWriteType`.
	ReadWriteType *string `json:"readWriteType"`
}

// TODO: EXAMPLE
//
type CfnTrail_InsightSelectorProperty struct {
	// `CfnTrail.InsightSelectorProperty.InsightType`.
	InsightType *string `json:"insightType"`
}

// Properties for defining a `AWS::CloudTrail::Trail`.
//
// TODO: EXAMPLE
//
type CfnTrailProps struct {
	// `AWS::CloudTrail::Trail.CloudWatchLogsLogGroupArn`.
	CloudWatchLogsLogGroupArn *string `json:"cloudWatchLogsLogGroupArn"`
	// `AWS::CloudTrail::Trail.CloudWatchLogsRoleArn`.
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn"`
	// `AWS::CloudTrail::Trail.EnableLogFileValidation`.
	EnableLogFileValidation interface{} `json:"enableLogFileValidation"`
	// `AWS::CloudTrail::Trail.EventSelectors`.
	EventSelectors interface{} `json:"eventSelectors"`
	// `AWS::CloudTrail::Trail.IncludeGlobalServiceEvents`.
	IncludeGlobalServiceEvents interface{} `json:"includeGlobalServiceEvents"`
	// `AWS::CloudTrail::Trail.InsightSelectors`.
	InsightSelectors interface{} `json:"insightSelectors"`
	// `AWS::CloudTrail::Trail.IsLogging`.
	IsLogging interface{} `json:"isLogging"`
	// `AWS::CloudTrail::Trail.IsMultiRegionTrail`.
	IsMultiRegionTrail interface{} `json:"isMultiRegionTrail"`
	// `AWS::CloudTrail::Trail.IsOrganizationTrail`.
	IsOrganizationTrail interface{} `json:"isOrganizationTrail"`
	// `AWS::CloudTrail::Trail.KMSKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::CloudTrail::Trail.S3BucketName`.
	S3BucketName *string `json:"s3BucketName"`
	// `AWS::CloudTrail::Trail.S3KeyPrefix`.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
	// `AWS::CloudTrail::Trail.SnsTopicName`.
	SnsTopicName *string `json:"snsTopicName"`
	// `AWS::CloudTrail::Trail.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::CloudTrail::Trail.TrailName`.
	TrailName *string `json:"trailName"`
}

// Resource type for a data event.
// Experimental.
type DataResourceType string

const (
	DataResourceType_LAMBDA_FUNCTION DataResourceType = "LAMBDA_FUNCTION"
	DataResourceType_S3_OBJECT DataResourceType = "S3_OBJECT"
)

// Types of management event sources that can be excluded.
// Experimental.
type ManagementEventSources string

const (
	ManagementEventSources_KMS ManagementEventSources = "KMS"
	ManagementEventSources_RDS_DATA_API ManagementEventSources = "RDS_DATA_API"
)

// Types of events that CloudTrail can log.
//
// TODO: EXAMPLE
//
// Experimental.
type ReadWriteType string

const (
	ReadWriteType_READ_ONLY ReadWriteType = "READ_ONLY"
	ReadWriteType_WRITE_ONLY ReadWriteType = "WRITE_ONLY"
	ReadWriteType_ALL ReadWriteType = "ALL"
	ReadWriteType_NONE ReadWriteType = "NONE"
)

// Selecting an S3 bucket and an optional prefix to be logged for data events.
//
// TODO: EXAMPLE
//
// Experimental.
type S3EventSelector struct {
	// S3 bucket.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// Data events for objects whose key matches this prefix will be logged.
	// Experimental.
	ObjectPrefix *string `json:"objectPrefix"`
}

// Cloud trail allows you to log events that happen in your AWS account For example:.
//
// import { CloudTrail } from '@aws-cdk/aws-cloudtrail'
//
// const cloudTrail = new CloudTrail(this, 'MyTrail');
//
// NOTE the above example creates an UNENCRYPTED bucket by default,
// If you are required to use an Encrypted bucket you can supply a preconfigured bucket
// via TrailProps
//
// TODO: EXAMPLE
//
// Experimental.
type Trail interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	LogGroup() awslogs.ILogGroup
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	TrailArn() *string
	TrailSnsTopicArn() *string
	AddEventSelector(dataResourceType DataResourceType, dataResourceValues *[]*string, options *AddEventSelectorOptions)
	AddLambdaEventSelector(handlers *[]awslambda.IFunction, options *AddEventSelectorOptions)
	AddS3EventSelector(s3Selector *[]*S3EventSelector, options *AddEventSelectorOptions)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LogAllLambdaDataEvents(options *AddEventSelectorOptions)
	LogAllS3DataEvents(options *AddEventSelectorOptions)
	ToString() *string
}

// The jsii proxy struct for Trail
type jsiiProxy_Trail struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Trail) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) TrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) TrailSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailSnsTopicArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewTrail(scope constructs.Construct, id *string, props *TrailProps) Trail {
	_init_.Initialize()

	j := jsiiProxy_Trail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTrail_Override(t Trail, scope constructs.Construct, id *string, props *TrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Trail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Trail_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Create an event rule for when an event is recorded by any Trail in the account.
//
// Note that the event doesn't necessarily have to come from this Trail, it can
// be captured from any one.
//
// Be sure to filter the event further down using an event pattern.
// Experimental.
func Trail_OnEvent(scope constructs.Construct, id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	_init_.Initialize()

	var returns awsevents.Rule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		"onEvent",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
//
// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
//
// This method adds an Event Selector for filtering events that match either S3 or Lambda function operations.
//
// Data events: These events provide insight into the resource operations performed on or within a resource.
// These are also known as data plane operations.
// Experimental.
func (t *jsiiProxy_Trail) AddEventSelector(dataResourceType DataResourceType, dataResourceValues *[]*string, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addEventSelector",
		[]interface{}{dataResourceType, dataResourceValues, options},
	)
}

// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
//
// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
//
// This method adds a Lambda Data Event Selector for filtering events that match Lambda function operations.
//
// Data events: These events provide insight into the resource operations performed on or within a resource.
// These are also known as data plane operations.
// Experimental.
func (t *jsiiProxy_Trail) AddLambdaEventSelector(handlers *[]awslambda.IFunction, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addLambdaEventSelector",
		[]interface{}{handlers, options},
	)
}

// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
//
// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
//
// This method adds an S3 Data Event Selector for filtering events that match S3 operations.
//
// Data events: These events provide insight into the resource operations performed on or within a resource.
// These are also known as data plane operations.
// Experimental.
func (t *jsiiProxy_Trail) AddS3EventSelector(s3Selector *[]*S3EventSelector, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addS3EventSelector",
		[]interface{}{s3Selector, options},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (t *jsiiProxy_Trail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (t *jsiiProxy_Trail) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (t *jsiiProxy_Trail) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (t *jsiiProxy_Trail) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Log all Lamda data events for all lambda functions the account.
// See: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
//
// Experimental.
func (t *jsiiProxy_Trail) LogAllLambdaDataEvents(options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"logAllLambdaDataEvents",
		[]interface{}{options},
	)
}

// Log all S3 data events for all objects for all buckets in the account.
// See: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
//
// Experimental.
func (t *jsiiProxy_Trail) LogAllS3DataEvents(options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"logAllS3DataEvents",
		[]interface{}{options},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_Trail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AWS CloudTrail trail.
//
// TODO: EXAMPLE
//
// Experimental.
type TrailProps struct {
	// The Amazon S3 bucket.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// Log Group to which CloudTrail to push logs to.
	//
	// Ignored if sendToCloudWatchLogs is set to false.
	// Experimental.
	CloudWatchLogGroup awslogs.ILogGroup `json:"cloudWatchLogGroup"`
	// How long to retain logs in CloudWatchLogs.
	//
	// Ignored if sendToCloudWatchLogs is false or if cloudWatchLogGroup is set.
	// Experimental.
	CloudWatchLogsRetention awslogs.RetentionDays `json:"cloudWatchLogsRetention"`
	// To determine whether a log file was modified, deleted, or unchanged after CloudTrail delivered it, you can use CloudTrail log file integrity validation.
	//
	// This feature is built using industry standard algorithms: SHA-256 for hashing and SHA-256 with RSA for digital signing.
	// This makes it computationally infeasible to modify, delete or forge CloudTrail log files without detection.
	// You can use the AWS CLI to validate the files in the location where CloudTrail delivered them.
	// Experimental.
	EnableFileValidation *bool `json:"enableFileValidation"`
	// The AWS Key Management Service (AWS KMS) key ID that you want to use to encrypt CloudTrail logs.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// For most services, events are recorded in the region where the action occurred.
	//
	// For global services such as AWS Identity and Access Management (IAM), AWS STS, Amazon CloudFront, and Route 53,
	// events are delivered to any trail that includes global services, and are logged as occurring in US East (N. Virginia) Region.
	// Experimental.
	IncludeGlobalServiceEvents *bool `json:"includeGlobalServiceEvents"`
	// Whether or not this trail delivers log files from multiple regions to a single S3 bucket for a single account.
	// Experimental.
	IsMultiRegionTrail *bool `json:"isMultiRegionTrail"`
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method sets the management configuration for this trail.
	//
	// Management events provide insight into management operations that are performed on resources in your AWS account.
	// These are also known as control plane operations.
	// Management events can also include non-API events that occur in your account.
	// For example, when a user logs in to your account, CloudTrail logs the ConsoleLogin event.
	// Experimental.
	ManagementEvents ReadWriteType `json:"managementEvents"`
	// An Amazon S3 object key prefix that precedes the name of all log files.
	// Experimental.
	S3KeyPrefix *string `json:"s3KeyPrefix"`
	// If CloudTrail pushes logs to CloudWatch Logs in addition to S3.
	//
	// Disabled for cost out of the box.
	// Experimental.
	SendToCloudWatchLogs *bool `json:"sendToCloudWatchLogs"`
	// SNS topic that is notified when new log files are published.
	// Experimental.
	SnsTopic awssns.ITopic `json:"snsTopic"`
	// The name of the trail.
	//
	// We recommend customers do not set an explicit name.
	// Experimental.
	TrailName *string `json:"trailName"`
}

