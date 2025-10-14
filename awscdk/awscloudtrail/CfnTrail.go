package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrail := awscdk.Aws_cloudtrail.NewCfnTrail(this, jsii.String("MyCfnTrail"), &CfnTrailProps{
//   	IsLogging: jsii.Boolean(false),
//   	S3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	AdvancedEventSelectors: []interface{}{
//   		&AdvancedEventSelectorProperty{
//   			FieldSelectors: []interface{}{
//   				&AdvancedFieldSelectorProperty{
//   					Field: jsii.String("field"),
//
//   					// the properties below are optional
//   					EndsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					EqualTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					NotEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					NotEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					NotStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					StartsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	CloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	EnableLogFileValidation: jsii.Boolean(false),
//   	EventSelectors: []interface{}{
//   		&EventSelectorProperty{
//   			DataResources: []interface{}{
//   				&DataResourceProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			ExcludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			IncludeManagementEvents: jsii.Boolean(false),
//   			ReadWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	IncludeGlobalServiceEvents: jsii.Boolean(false),
//   	InsightSelectors: []interface{}{
//   		&InsightSelectorProperty{
//   			InsightType: jsii.String("insightType"),
//   		},
//   	},
//   	IsMultiRegionTrail: jsii.Boolean(false),
//   	IsOrganizationTrail: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	SnsTopicName: jsii.String("snsTopicName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrailName: jsii.String("trailName"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
//
type CfnTrail interface {
	awscdk.CfnResource
	ITrailRef
	awscdk.IInspectable
	awscdk.ITaggable
	// Specifies the settings for advanced event selectors.
	AdvancedEventSelectors() interface{}
	SetAdvancedEventSelectors(val interface{})
	// `Ref` returns the ARN of the CloudTrail trail, such as `arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail` .
	AttrArn() *string
	// `Ref` returns the ARN of the Amazon SNS topic that's associated with the CloudTrail trail, such as `arn:aws:sns:us-east-2:123456789012:mySNSTopic` .
	AttrSnsTopicArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.
	CloudWatchLogsLogGroupArn() *string
	SetCloudWatchLogsLogGroupArn(val *string)
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	CloudWatchLogsRoleArn() *string
	SetCloudWatchLogsRoleArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether log file validation is enabled.
	//
	// The default is false.
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	// Use event selectors to further specify the management and data event settings for your trail.
	EventSelectors() interface{}
	SetEventSelectors(val interface{})
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	// A JSON string that contains the Insights types you want to log on a trail.
	InsightSelectors() interface{}
	SetInsightSelectors(val interface{})
	// Whether the CloudTrail trail is currently logging AWS API calls.
	IsLogging() interface{}
	SetIsLogging(val interface{})
	// Specifies whether the trail applies only to the current Region or to all Regions.
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations , or only for the current AWS account .
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	// Specifies the AWS KMS key ID to use to encrypt the logs and digest files delivered by CloudTrail.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	S3BucketName() *string
	SetS3BucketName(val *string)
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	// Specifies the name or ARN of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A custom set of tags (key-value pairs) for this trail.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Specifies the name of the trail.
	//
	// The name must meet the following requirements:.
	TrailName() *string
	SetTrailName(val *string)
	// A reference to a Trail resource.
	TrailRef() *TrailReference
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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

// The jsii proxy struct for CfnTrail
type jsiiProxy_CfnTrail struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ITrailRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnTrail) AdvancedEventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedEventSelectors",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnTrail) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnTrail) TrailRef() *TrailReference {
	var returns *TrailReference
	_jsii_.Get(
		j,
		"trailRef",
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

func (j *jsiiProxy_CfnTrail) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnTrail(scope constructs.Construct, id *string, props *CfnTrailProps) CfnTrail {
	_init_.Initialize()

	if err := validateNewCfnTrailParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnTrail_Override(c CfnTrail, scope constructs.Construct, id *string, props *CfnTrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTrail)SetAdvancedEventSelectors(val interface{}) {
	if err := j.validateSetAdvancedEventSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedEventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetCloudWatchLogsLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetCloudWatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetEnableLogFileValidation(val interface{}) {
	if err := j.validateSetEnableLogFileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetEventSelectors(val interface{}) {
	if err := j.validateSetEventSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIncludeGlobalServiceEvents(val interface{}) {
	if err := j.validateSetIncludeGlobalServiceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetInsightSelectors(val interface{}) {
	if err := j.validateSetInsightSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsLogging(val interface{}) {
	if err := j.validateSetIsLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLogging",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsMultiRegionTrail(val interface{}) {
	if err := j.validateSetIsMultiRegionTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsOrganizationTrail(val interface{}) {
	if err := j.validateSetIsOrganizationTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetSnsTopicName(val *string) {
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetTrailName(val *string) {
	_jsii_.Set(
		j,
		"trailName",
		val,
	)
}

// Creates a new ITrailRef from an ARN.
func CfnTrail_FromTrailArn(scope constructs.Construct, id *string, arn *string) ITrailRef {
	_init_.Initialize()

	if err := validateCfnTrail_FromTrailArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns ITrailRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"fromTrailArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new ITrailRef from a trailName.
func CfnTrail_FromTrailName(scope constructs.Construct, id *string, trailName *string) ITrailRef {
	_init_.Initialize()

	if err := validateCfnTrail_FromTrailNameParameters(scope, id, trailName); err != nil {
		panic(err)
	}
	var returns ITrailRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"fromTrailName",
		[]interface{}{scope, id, trailName},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTrail_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTrail_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		"isCfnResource",
		[]interface{}{x},
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
func CfnTrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnTrail) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTrail) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTrail) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTrail) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTrail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTrail) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTrail) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTrail) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTrail) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTrail) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTrail) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTrail) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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

func (c *jsiiProxy_CfnTrail) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

