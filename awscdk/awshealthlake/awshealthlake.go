package awshealthlake

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awshealthlake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::HealthLake::FHIRDatastore`.
//
// TODO: EXAMPLE
//
type CfnFHIRDatastore interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDatastoreArn() *string
	AttrDatastoreEndpoint() *string
	AttrDatastoreId() *string
	AttrDatastoreStatus() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatastoreName() *string
	SetDatastoreName(val *string)
	DatastoreTypeVersion() *string
	SetDatastoreTypeVersion(val *string)
	LogicalId() *string
	Node() constructs.Node
	PreloadDataConfig() interface{}
	SetPreloadDataConfig(val interface{})
	Ref() *string
	SseConfiguration() interface{}
	SetSseConfiguration(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnFHIRDatastore
type jsiiProxy_CfnFHIRDatastore struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrDatastoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatastoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrDatastoreEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatastoreEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrDatastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrDatastoreStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatastoreStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) DatastoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) DatastoreTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) PreloadDataConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) SseConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::HealthLake::FHIRDatastore`.
func NewCfnFHIRDatastore(scope constructs.Construct, id *string, props *CfnFHIRDatastoreProps) CfnFHIRDatastore {
	_init_.Initialize()

	j := jsiiProxy_CfnFHIRDatastore{}

	_jsii_.Create(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::HealthLake::FHIRDatastore`.
func NewCfnFHIRDatastore_Override(c CfnFHIRDatastore, scope constructs.Construct, id *string, props *CfnFHIRDatastoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore) SetDatastoreName(val *string) {
	_jsii_.Set(
		j,
		"datastoreName",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore) SetDatastoreTypeVersion(val *string) {
	_jsii_.Set(
		j,
		"datastoreTypeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore) SetPreloadDataConfig(val interface{}) {
	_jsii_.Set(
		j,
		"preloadDataConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore) SetSseConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"sseConfiguration",
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
func CfnFHIRDatastore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFHIRDatastore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
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
func CfnFHIRDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFHIRDatastore_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFHIRDatastore) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFHIRDatastore) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFHIRDatastore) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFHIRDatastore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFHIRDatastore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFHIRDatastore) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFHIRDatastore) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFHIRDatastore) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFHIRDatastore) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFHIRDatastore) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFHIRDatastore) ToString() *string {
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
func (c *jsiiProxy_CfnFHIRDatastore) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnFHIRDatastore_KmsEncryptionConfigProperty struct {
	// `CfnFHIRDatastore.KmsEncryptionConfigProperty.CmkType`.
	CmkType *string `json:"cmkType"`
	// `CfnFHIRDatastore.KmsEncryptionConfigProperty.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
}

// TODO: EXAMPLE
//
type CfnFHIRDatastore_PreloadDataConfigProperty struct {
	// `CfnFHIRDatastore.PreloadDataConfigProperty.PreloadDataType`.
	PreloadDataType *string `json:"preloadDataType"`
}

// TODO: EXAMPLE
//
type CfnFHIRDatastore_SseConfigurationProperty struct {
	// `CfnFHIRDatastore.SseConfigurationProperty.KmsEncryptionConfig`.
	KmsEncryptionConfig interface{} `json:"kmsEncryptionConfig"`
}

// Properties for defining a `AWS::HealthLake::FHIRDatastore`.
//
// TODO: EXAMPLE
//
type CfnFHIRDatastoreProps struct {
	// `AWS::HealthLake::FHIRDatastore.DatastoreName`.
	DatastoreName *string `json:"datastoreName"`
	// `AWS::HealthLake::FHIRDatastore.DatastoreTypeVersion`.
	DatastoreTypeVersion *string `json:"datastoreTypeVersion"`
	// `AWS::HealthLake::FHIRDatastore.PreloadDataConfig`.
	PreloadDataConfig interface{} `json:"preloadDataConfig"`
	// `AWS::HealthLake::FHIRDatastore.SseConfiguration`.
	SseConfiguration interface{} `json:"sseConfiguration"`
	// `AWS::HealthLake::FHIRDatastore.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

