package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CloudTrail::EventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStore := awscdk.Aws_cloudtrail.NewCfnEventDataStore(this, jsii.String("MyCfnEventDataStore"), &cfnEventDataStoreProps{
//   	advancedEventSelectors: []interface{}{
//   		&advancedEventSelectorProperty{
//   			fieldSelectors: []interface{}{
//   				&advancedFieldSelectorProperty{
//   					field: jsii.String("field"),
//
//   					// the properties below are optional
//   					endsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					equalTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					notEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					notEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					notStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					startsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	multiRegionEnabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	organizationEnabled: jsii.Boolean(false),
//   	retentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	terminationProtectionEnabled: jsii.Boolean(false),
//   })
//
type CfnEventDataStore interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::CloudTrail::EventDataStore.AdvancedEventSelectors`.
	AdvancedEventSelectors() interface{}
	SetAdvancedEventSelectors(val interface{})
	AttrCreatedTimestamp() *string
	AttrEventDataStoreArn() *string
	AttrStatus() *string
	AttrUpdatedTimestamp() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::CloudTrail::EventDataStore.KmsKeyId`.
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
	// `AWS::CloudTrail::EventDataStore.MultiRegionEnabled`.
	MultiRegionEnabled() interface{}
	SetMultiRegionEnabled(val interface{})
	// `AWS::CloudTrail::EventDataStore.Name`.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// `AWS::CloudTrail::EventDataStore.OrganizationEnabled`.
	OrganizationEnabled() interface{}
	SetOrganizationEnabled(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::CloudTrail::EventDataStore.RetentionPeriod`.
	RetentionPeriod() *float64
	SetRetentionPeriod(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::CloudTrail::EventDataStore.Tags`.
	Tags() awscdk.TagManager
	// `AWS::CloudTrail::EventDataStore.TerminationProtectionEnabled`.
	TerminationProtectionEnabled() interface{}
	SetTerminationProtectionEnabled(val interface{})
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnEventDataStore
type jsiiProxy_CfnEventDataStore struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventDataStore) AdvancedEventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedEventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrCreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrEventDataStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEventDataStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrUpdatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) MultiRegionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiRegionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) OrganizationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) RetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) TerminationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStore(scope constructs.Construct, id *string, props *CfnEventDataStoreProps) CfnEventDataStore {
	_init_.Initialize()

	if err := validateNewCfnEventDataStoreParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventDataStore{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStore_Override(c CfnEventDataStore, scope constructs.Construct, id *string, props *CfnEventDataStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetAdvancedEventSelectors(val interface{}) {
	if err := j.validateSetAdvancedEventSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedEventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetMultiRegionEnabled(val interface{}) {
	if err := j.validateSetMultiRegionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiRegionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetOrganizationEnabled(val interface{}) {
	if err := j.validateSetOrganizationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore)SetTerminationProtectionEnabled(val interface{}) {
	if err := j.validateSetTerminationProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtectionEnabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventDataStore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventDataStore_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventDataStore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEventDataStore_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
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
func CfnEventDataStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventDataStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventDataStore_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudtrail.CfnEventDataStore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventDataStore) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEventDataStore) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnEventDataStore) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEventDataStore) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventDataStore) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventDataStore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEventDataStore) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

