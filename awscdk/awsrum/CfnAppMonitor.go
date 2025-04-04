package awsrum

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrum/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a CloudWatch RUM app monitor, which you can use to collect telemetry data from your application and send it to CloudWatch RUM.
//
// The data includes performance and reliability information such as page load time, client-side errors, and user behavior.
//
// After you create an app monitor, sign in to the CloudWatch RUM console to get the JavaScript code snippet to add to your web application. For more information, see [How do I find a code snippet that I've already generated?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppMonitor := awscdk.Aws_rum.NewCfnAppMonitor(this, jsii.String("MyCfnAppMonitor"), &CfnAppMonitorProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AppMonitorConfiguration: &AppMonitorConfigurationProperty{
//   		AllowCookies: jsii.Boolean(false),
//   		EnableXRay: jsii.Boolean(false),
//   		ExcludedPages: []*string{
//   			jsii.String("excludedPages"),
//   		},
//   		FavoritePages: []*string{
//   			jsii.String("favoritePages"),
//   		},
//   		GuestRoleArn: jsii.String("guestRoleArn"),
//   		IdentityPoolId: jsii.String("identityPoolId"),
//   		IncludedPages: []*string{
//   			jsii.String("includedPages"),
//   		},
//   		MetricDestinations: []interface{}{
//   			&MetricDestinationProperty{
//   				Destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				DestinationArn: jsii.String("destinationArn"),
//   				IamRoleArn: jsii.String("iamRoleArn"),
//   				MetricDefinitions: []interface{}{
//   					&MetricDefinitionProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						DimensionKeys: map[string]*string{
//   							"dimensionKeysKey": jsii.String("dimensionKeys"),
//   						},
//   						EventPattern: jsii.String("eventPattern"),
//   						Namespace: jsii.String("namespace"),
//   						UnitLabel: jsii.String("unitLabel"),
//   						ValueKey: jsii.String("valueKey"),
//   					},
//   				},
//   			},
//   		},
//   		SessionSampleRate: jsii.Number(123),
//   		Telemetries: []*string{
//   			jsii.String("telemetries"),
//   		},
//   	},
//   	CustomEvents: &CustomEventsProperty{
//   		Status: jsii.String("status"),
//   	},
//   	CwLogEnabled: jsii.Boolean(false),
//   	DeobfuscationConfiguration: &DeobfuscationConfigurationProperty{
//   		JavaScriptSourceMaps: &JavaScriptSourceMapsProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Domain: jsii.String("domain"),
//   	DomainList: []*string{
//   		jsii.String("domainList"),
//   	},
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: jsii.String("policyDocument"),
//
//   		// the properties below are optional
//   		PolicyRevisionId: jsii.String("policyRevisionId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html
//
type CfnAppMonitor interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// A structure that contains much of the configuration data for the app monitor.
	AppMonitorConfiguration() interface{}
	SetAppMonitorConfiguration(val interface{})
	// The ID of the app monitor, such as `123456ab-1234-4ca9-9d2f-a1b2c3456789` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies whether this app monitor allows the web client to define and send custom events.
	CustomEvents() interface{}
	SetCustomEvents(val interface{})
	// Data collected by CloudWatch RUM is kept by RUM for 30 days and then deleted.
	CwLogEnabled() interface{}
	SetCwLogEnabled(val interface{})
	// A structure that contains the configuration for how an app monitor can deobfuscate stack traces.
	DeobfuscationConfiguration() interface{}
	SetDeobfuscationConfiguration(val interface{})
	// The top-level internet domain name for which your application has administrative authority.
	Domain() *string
	SetDomain(val *string)
	// List the domain names for which your application has administrative authority.
	//
	// This parameter or the `Domain` parameter is required.
	DomainList() *[]*string
	SetDomainList(val *[]*string)
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
	// A name for the app monitor.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Use this structure to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it.
	ResourcePolicy() interface{}
	SetResourcePolicy(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Assigns one or more tags (key-value pairs) to the app monitor.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnAppMonitor
type jsiiProxy_CfnAppMonitor struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnAppMonitor) AppMonitorConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appMonitorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CustomEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CwLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cwLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) DeobfuscationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deobfuscationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) DomainList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) ResourcePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnAppMonitor(scope constructs.Construct, id *string, props *CfnAppMonitorProps) CfnAppMonitor {
	_init_.Initialize()

	if err := validateNewCfnAppMonitorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppMonitor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAppMonitor_Override(c CfnAppMonitor, scope constructs.Construct, id *string, props *CfnAppMonitorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetAppMonitorConfiguration(val interface{}) {
	if err := j.validateSetAppMonitorConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appMonitorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetCustomEvents(val interface{}) {
	if err := j.validateSetCustomEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEvents",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetCwLogEnabled(val interface{}) {
	if err := j.validateSetCwLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cwLogEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetDeobfuscationConfiguration(val interface{}) {
	if err := j.validateSetDeobfuscationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deobfuscationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetDomainList(val *[]*string) {
	_jsii_.Set(
		j,
		"domainList",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetResourcePolicy(val interface{}) {
	if err := j.validateSetResourcePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAppMonitor_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppMonitor_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAppMonitor_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppMonitor_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
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
func CfnAppMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppMonitor_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppMonitor) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAppMonitor) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAppMonitor) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAppMonitor) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnAppMonitor) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAppMonitor) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAppMonitor) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitor) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAppMonitor) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAppMonitor) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAppMonitor) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnAppMonitor) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitor) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

