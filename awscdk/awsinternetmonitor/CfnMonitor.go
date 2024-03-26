package awsinternetmonitor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinternetmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::InternetMonitor::Monitor` resource is an Internet Monitor resource type that contains information about how you create a monitor in Amazon CloudWatch Internet Monitor.
//
// A monitor in Internet Monitor provides visibility into performance and availability between your applications hosted on AWS and your end users, using a traffic profile that it creates based on the application resources that you add: Virtual Private Clouds (VPCs), Amazon CloudFront distributions, or WorkSpaces directories.
//
// Internet Monitor also alerts you to internet issues that impact your application in the city-networks (geographies and networks) where your end users use it. With Internet Monitor, you can quickly pinpoint the locations and providers that are affected, so that you can address the issue.
//
// For more information, see [Using Amazon CloudWatch Internet Monitor](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html) in the *Amazon CloudWatch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitor := awscdk.Aws_internetmonitor.NewCfnMonitor(this, jsii.String("MyCfnMonitor"), &CfnMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//
//   	// the properties below are optional
//   	HealthEventsConfig: &HealthEventsConfigProperty{
//   		AvailabilityLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		AvailabilityScoreThreshold: jsii.Number(123),
//   		PerformanceLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   			HealthScoreThreshold: jsii.Number(123),
//   			MinTrafficImpact: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		PerformanceScoreThreshold: jsii.Number(123),
//   	},
//   	IncludeLinkedAccounts: jsii.Boolean(false),
//   	InternetMeasurementsLogDelivery: &InternetMeasurementsLogDeliveryProperty{
//   		S3Config: &S3ConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   		},
//   	},
//   	LinkedAccountId: jsii.String("linkedAccountId"),
//   	MaxCityNetworksToMonitor: jsii.Number(123),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourcesToAdd: []*string{
//   		jsii.String("resourcesToAdd"),
//   	},
//   	ResourcesToRemove: []*string{
//   		jsii.String("resourcesToRemove"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPercentageToMonitor: jsii.Number(123),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-internetmonitor-monitor.html
//
type CfnMonitor interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The time when the monitor was created.
	AttrCreatedAt() *string
	// The last time that the monitor was modified.
	AttrModifiedAt() *string
	// The Amazon Resource Name (ARN) of the monitor.
	AttrMonitorArn() *string
	// The health of data processing for the monitor.
	//
	// For more information, see `ProcessingStatus` under [MonitorListMember](https://docs.aws.amazon.com/internet-monitor/latest/api/API_MonitorListMember.html) in the *Amazon CloudWatch Internet Monitor API Reference* .
	AttrProcessingStatus() *string
	// Additional information about the health of the data processing for the monitor.
	AttrProcessingStatusInfo() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A complex type with the configuration information that determines the threshold and other conditions for when Internet Monitor creates a health event for an overall performance or availability issue, across an application's geographies.
	HealthEventsConfig() interface{}
	SetHealthEventsConfig(val interface{})
	IncludeLinkedAccounts() interface{}
	SetIncludeLinkedAccounts(val interface{})
	// Publish internet measurements for a monitor for all city-networks (up to the 500,000 service limit) to another location, such as an Amazon S3 bucket.
	InternetMeasurementsLogDelivery() interface{}
	SetInternetMeasurementsLogDelivery(val interface{})
	LinkedAccountId() *string
	SetLinkedAccountId(val *string)
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
	// The maximum number of city-networks to monitor for your resources.
	MaxCityNetworksToMonitor() *float64
	SetMaxCityNetworksToMonitor(val *float64)
	// The name of the monitor.
	MonitorName() *string
	SetMonitorName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The resources that have been added for the monitor, listed by their Amazon Resource Names (ARNs).
	Resources() *[]*string
	SetResources(val *[]*string)
	// The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	ResourcesToAdd() *[]*string
	SetResourcesToAdd(val *[]*string)
	// The resources to remove from a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	ResourcesToRemove() *[]*string
	SetResourcesToRemove(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of a monitor.
	Status() *string
	SetStatus(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags for a monitor, listed as a set of *key:value* pairs.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The percentage of the internet-facing traffic for your application that you want to monitor.
	TrafficPercentageToMonitor() *float64
	SetTrafficPercentageToMonitor(val *float64)
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

// The jsii proxy struct for CfnMonitor
type jsiiProxy_CfnMonitor struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnMonitor) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) AttrMonitorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMonitorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) AttrProcessingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) AttrProcessingStatusInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingStatusInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) HealthEventsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthEventsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) IncludeLinkedAccounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeLinkedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) InternetMeasurementsLogDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetMeasurementsLogDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) LinkedAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) MaxCityNetworksToMonitor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCityNetworksToMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) MonitorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) ResourcesToAdd() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) ResourcesToRemove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) TrafficPercentageToMonitor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageToMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitor) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnMonitor(scope constructs.Construct, id *string, props *CfnMonitorProps) CfnMonitor {
	_init_.Initialize()

	if err := validateNewCfnMonitorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMonitor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnMonitor_Override(c CfnMonitor, scope constructs.Construct, id *string, props *CfnMonitorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMonitor)SetHealthEventsConfig(val interface{}) {
	if err := j.validateSetHealthEventsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthEventsConfig",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetIncludeLinkedAccounts(val interface{}) {
	if err := j.validateSetIncludeLinkedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeLinkedAccounts",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetInternetMeasurementsLogDelivery(val interface{}) {
	if err := j.validateSetInternetMeasurementsLogDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetMeasurementsLogDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetLinkedAccountId(val *string) {
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetMaxCityNetworksToMonitor(val *float64) {
	_jsii_.Set(
		j,
		"maxCityNetworksToMonitor",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetMonitorName(val *string) {
	if err := j.validateSetMonitorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorName",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetResources(val *[]*string) {
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetResourcesToAdd(val *[]*string) {
	_jsii_.Set(
		j,
		"resourcesToAdd",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetResourcesToRemove(val *[]*string) {
	_jsii_.Set(
		j,
		"resourcesToRemove",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnMonitor)SetTrafficPercentageToMonitor(val *float64) {
	_jsii_.Set(
		j,
		"trafficPercentageToMonitor",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMonitor_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitor_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnMonitor_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitor_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
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
func CfnMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMonitor_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_internetmonitor.CfnMonitor",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMonitor) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMonitor) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMonitor) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMonitor) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMonitor) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMonitor) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMonitor) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMonitor) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMonitor) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnMonitor) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnMonitor) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMonitor) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitor) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMonitor) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMonitor) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnMonitor) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnMonitor) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMonitor) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

