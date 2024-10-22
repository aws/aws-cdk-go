package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Crawler` resource specifies an AWS Glue crawler.
//
// For more information, see [Cataloging Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) and [Crawler Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-crawling.html#aws-glue-api-crawler-crawling-Crawler) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCrawler := awscdk.Aws_glue.NewCfnCrawler(this, jsii.String("MyCfnCrawler"), &CfnCrawlerProps{
//   	Role: jsii.String("role"),
//   	Targets: &TargetsProperty{
//   		CatalogTargets: []interface{}{
//   			&CatalogTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				EventQueueArn: jsii.String("eventQueueArn"),
//   				Tables: []*string{
//   					jsii.String("tables"),
//   				},
//   			},
//   		},
//   		DeltaTargets: []interface{}{
//   			&DeltaTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				CreateNativeDeltaTable: jsii.Boolean(false),
//   				DeltaTables: []*string{
//   					jsii.String("deltaTables"),
//   				},
//   				WriteManifest: jsii.Boolean(false),
//   			},
//   		},
//   		DynamoDbTargets: []interface{}{
//   			&DynamoDBTargetProperty{
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		IcebergTargets: []interface{}{
//   			&IcebergTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				MaximumTraversalDepth: jsii.Number(123),
//   				Paths: []*string{
//   					jsii.String("paths"),
//   				},
//   			},
//   		},
//   		JdbcTargets: []interface{}{
//   			&JdbcTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				EnableAdditionalMetadata: []*string{
//   					jsii.String("enableAdditionalMetadata"),
//   				},
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		MongoDbTargets: []interface{}{
//   			&MongoDBTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		S3Targets: []interface{}{
//   			&S3TargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				EventQueueArn: jsii.String("eventQueueArn"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   				SampleSize: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Classifiers: []*string{
//   		jsii.String("classifiers"),
//   	},
//   	Configuration: jsii.String("configuration"),
//   	CrawlerSecurityConfiguration: jsii.String("crawlerSecurityConfiguration"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Description: jsii.String("description"),
//   	LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		UseLakeFormationCredentials: jsii.Boolean(false),
//   	},
//   	Name: jsii.String("name"),
//   	RecrawlPolicy: &RecrawlPolicyProperty{
//   		RecrawlBehavior: jsii.String("recrawlBehavior"),
//   	},
//   	Schedule: &ScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SchemaChangePolicy: &SchemaChangePolicyProperty{
//   		DeleteBehavior: jsii.String("deleteBehavior"),
//   		UpdateBehavior: jsii.String("updateBehavior"),
//   	},
//   	TablePrefix: jsii.String("tablePrefix"),
//   	Tags: tags,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-crawler.html
//
type CfnCrawler interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	Classifiers() *[]*string
	SetClassifiers(val *[]*string)
	// Crawler configuration information.
	Configuration() *string
	SetConfiguration(val *string)
	// The name of the `SecurityConfiguration` structure to be used by this crawler.
	CrawlerSecurityConfiguration() *string
	SetCrawlerSecurityConfiguration(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the database in which the crawler's output is stored.
	DatabaseName() *string
	SetDatabaseName(val *string)
	// A description of the crawler.
	Description() *string
	SetDescription(val *string)
	// Specifies whether the crawler should use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.
	LakeFormationConfiguration() interface{}
	SetLakeFormationConfiguration(val interface{})
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
	// The name of the crawler.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
	RecrawlPolicy() interface{}
	SetRecrawlPolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	Role() *string
	SetRole(val *string)
	// For scheduled crawlers, the schedule when the crawler runs.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The policy that specifies update and delete behaviors for the crawler.
	SchemaChangePolicy() interface{}
	SetSchemaChangePolicy(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The prefix added to the names of tables that are created.
	TablePrefix() *string
	SetTablePrefix(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags to use with this crawler.
	TagsRaw() interface{}
	SetTagsRaw(val interface{})
	// A collection of targets to crawl.
	Targets() interface{}
	SetTargets(val interface{})
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

// The jsii proxy struct for CfnCrawler
type jsiiProxy_CfnCrawler struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCrawler) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CrawlerSecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerSecurityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) LakeFormationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakeFormationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) RecrawlPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) SchemaChangePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) TagsRaw() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnCrawler(scope constructs.Construct, id *string, props *CfnCrawlerProps) CfnCrawler {
	_init_.Initialize()

	if err := validateNewCfnCrawlerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCrawler{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnCrawler_Override(c CfnCrawler, scope constructs.Construct, id *string, props *CfnCrawlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCrawler)SetClassifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetConfiguration(val *string) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetCrawlerSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"crawlerSecurityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetLakeFormationConfiguration(val interface{}) {
	if err := j.validateSetLakeFormationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lakeFormationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetRecrawlPolicy(val interface{}) {
	if err := j.validateSetRecrawlPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recrawlPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetSchemaChangePolicy(val interface{}) {
	if err := j.validateSetSchemaChangePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaChangePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetTablePrefix(val *string) {
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetTagsRaw(val interface{}) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler)SetTargets(val interface{}) {
	if err := j.validateSetTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCrawler_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCrawler_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCrawler_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCrawler_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
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
func CfnCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCrawler_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCrawler) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCrawler) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCrawler) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCrawler) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCrawler) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCrawler) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCrawler) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCrawler) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCrawler) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCrawler) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCrawler) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCrawler) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCrawler) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCrawler) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCrawler) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCrawler) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCrawler) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCrawler) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCrawler) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

