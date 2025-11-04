package awshealthlake

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awshealthlake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Data Store that can ingest and export FHIR formatted data.
//
// > Please note that when a user tries to do an Update operation via CloudFormation, changes to the Data Store name, Type Version, PreloadDataConfig, or SSEConfiguration will delete their existing Data Store for the stack and create a new one. This will lead to potential loss of data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFHIRDatastore := awscdk.Aws_healthlake.NewCfnFHIRDatastore(this, jsii.String("MyCfnFHIRDatastore"), &CfnFHIRDatastoreProps{
//   	DatastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//
//   	// the properties below are optional
//   	DatastoreName: jsii.String("datastoreName"),
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		AuthorizationStrategy: jsii.String("authorizationStrategy"),
//
//   		// the properties below are optional
//   		FineGrainedAuthorizationEnabled: jsii.Boolean(false),
//   		IdpLambdaArn: jsii.String("idpLambdaArn"),
//   		Metadata: jsii.String("metadata"),
//   	},
//   	PreloadDataConfig: &PreloadDataConfigProperty{
//   		PreloadDataType: jsii.String("preloadDataType"),
//   	},
//   	SseConfiguration: &SseConfigurationProperty{
//   		KmsEncryptionConfig: &KmsEncryptionConfigProperty{
//   			CmkType: jsii.String("cmkType"),
//
//   			// the properties below are optional
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html
//
type CfnFHIRDatastore interface {
	awscdk.CfnResource
	IFHIRDatastoreRef
	awscdk.IInspectable
	awscdk.ITaggable
	// The time that a Data Store was created.
	AttrCreatedAt() awscdk.IResolvable
	AttrCreatedAtNanos() *float64
	AttrCreatedAtSeconds() *string
	// The Data Store ARN is generated during the creation of the Data Store and can be found in the output from the initial Data Store creation request.
	AttrDatastoreArn() *string
	// The endpoint for the created Data Store.
	AttrDatastoreEndpoint() *string
	// The Amazon generated Data Store id.
	//
	// This id is in the output from the initial Data Store creation call.
	AttrDatastoreId() *string
	// The status of the FHIR Data Store.
	//
	// Possible statuses are ‘CREATING’, ‘ACTIVE’, ‘DELETING’, ‘DELETED’.
	AttrDatastoreStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The data store name (user-generated).
	DatastoreName() *string
	SetDatastoreName(val *string)
	// The FHIR release version supported by the data store.
	DatastoreTypeVersion() *string
	SetDatastoreTypeVersion(val *string)
	Env() *awscdk.ResourceEnvironment
	// A reference to a FHIRDatastore resource.
	FhirDatastoreRef() *FHIRDatastoreReference
	// The identity provider configuration selected when the data store was created.
	IdentityProviderConfiguration() interface{}
	SetIdentityProviderConfiguration(val interface{})
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
	// The preloaded Synthea data configuration for the data store.
	PreloadDataConfig() interface{}
	SetPreloadDataConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The server-side encryption key configuration for a customer-provided encryption key specified for creating a data store.
	SseConfiguration() interface{}
	SetSseConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An array of key-value pairs to apply to this resource.
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

// The jsii proxy struct for CfnFHIRDatastore
type jsiiProxy_CfnFHIRDatastore struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IFHIRDatastoreRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrCreatedAt() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrCreatedAtNanos() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreatedAtNanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) AttrCreatedAtSeconds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAtSeconds",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnFHIRDatastore) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) FhirDatastoreRef() *FHIRDatastoreReference {
	var returns *FHIRDatastoreReference
	_jsii_.Get(
		j,
		"fhirDatastoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFHIRDatastore) IdentityProviderConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderConfiguration",
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

func (j *jsiiProxy_CfnFHIRDatastore) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnFHIRDatastore) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnFHIRDatastore(scope constructs.Construct, id *string, props *CfnFHIRDatastoreProps) CfnFHIRDatastore {
	_init_.Initialize()

	if err := validateNewCfnFHIRDatastoreParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFHIRDatastore{}

	_jsii_.Create(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnFHIRDatastore_Override(c CfnFHIRDatastore, scope constructs.Construct, id *string, props *CfnFHIRDatastoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetDatastoreName(val *string) {
	_jsii_.Set(
		j,
		"datastoreName",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetDatastoreTypeVersion(val *string) {
	if err := j.validateSetDatastoreTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreTypeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetIdentityProviderConfiguration(val interface{}) {
	if err := j.validateSetIdentityProviderConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetPreloadDataConfig(val interface{}) {
	if err := j.validateSetPreloadDataConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preloadDataConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetSseConfiguration(val interface{}) {
	if err := j.validateSetSseConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sseConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFHIRDatastore)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Creates a new IFHIRDatastoreRef from a datastoreId.
func CfnFHIRDatastore_FromDatastoreId(scope constructs.Construct, id *string, datastoreId *string) IFHIRDatastoreRef {
	_init_.Initialize()

	if err := validateCfnFHIRDatastore_FromDatastoreIdParameters(scope, id, datastoreId); err != nil {
		panic(err)
	}
	var returns IFHIRDatastoreRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		"fromDatastoreId",
		[]interface{}{scope, id, datastoreId},
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
func CfnFHIRDatastore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFHIRDatastore_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnFHIRDatastore_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFHIRDatastore_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_healthlake.CfnFHIRDatastore",
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
func CfnFHIRDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFHIRDatastore_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnFHIRDatastore) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFHIRDatastore) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFHIRDatastore) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFHIRDatastore) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFHIRDatastore) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFHIRDatastore) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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

func (c *jsiiProxy_CfnFHIRDatastore) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

