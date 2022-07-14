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
// Creates a Data Store that can ingest and export FHIR formatted data.
//
// > Please note that when a user tries to do an Update operation via CloudFormation, changes to the Data Store name, Type Version, PreloadDataConfig, or SSEConfiguration will delete their existing Data Store for the stack and create a new one. This will lead to potential loss of data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFHIRDatastore := awscdk.Aws_healthlake.NewCfnFHIRDatastore(this, jsii.String("MyCfnFHIRDatastore"), &cfnFHIRDatastoreProps{
//   	datastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//
//   	// the properties below are optional
//   	datastoreName: jsii.String("datastoreName"),
//   	preloadDataConfig: &preloadDataConfigProperty{
//   		preloadDataType: jsii.String("preloadDataType"),
//   	},
//   	sseConfiguration: &sseConfigurationProperty{
//   		kmsEncryptionConfig: &kmsEncryptionConfigProperty{
//   			cmkType: jsii.String("cmkType"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFHIRDatastore interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	// The user generated name for the Data Store.
	DatastoreName() *string
	SetDatastoreName(val *string)
	// The FHIR version of the Data Store.
	//
	// The only supported version is R4.
	DatastoreTypeVersion() *string
	SetDatastoreTypeVersion(val *string)
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
	// The preloaded data configuration for the Data Store.
	//
	// Only data preloaded from Synthea is supported.
	PreloadDataConfig() interface{}
	SetPreloadDataConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The server-side encryption key configuration for a customer provided encryption key specified for creating a Data Store.
	SseConfiguration() interface{}
	SetSseConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFHIRDatastore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnFHIRDatastore) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The customer-managed-key(CMK) used when creating a Data Store.
//
// If a customer owned key is not specified, an Amazon owned key will be used for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kmsEncryptionConfigProperty := &kmsEncryptionConfigProperty{
//   	cmkType: jsii.String("cmkType"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFHIRDatastore_KmsEncryptionConfigProperty struct {
	// The type of customer-managed-key(CMK) used for encryption.
	//
	// The two types of supported CMKs are customer owned CMKs and Amazon owned CMKs. For more information on CMK types, see [KmsEncryptionConfig](https://docs.aws.amazon.com/healthlake/latest/APIReference/API_KmsEncryptionConfig.html#HealthLake-Type-KmsEncryptionConfig-CmkType) .
	CmkType *string `field:"required" json:"cmkType" yaml:"cmkType"`
	// The KMS encryption key id/alias used to encrypt the Data Store contents at rest.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Optional parameter to preload data upon creation of the Data Store.
//
// Currently, the only supported preloaded data is synthetic data generated from Synthea.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preloadDataConfigProperty := &preloadDataConfigProperty{
//   	preloadDataType: jsii.String("preloadDataType"),
//   }
//
type CfnFHIRDatastore_PreloadDataConfigProperty struct {
	// The type of preloaded data.
	//
	// Only Synthea preloaded data is supported.
	PreloadDataType *string `field:"required" json:"preloadDataType" yaml:"preloadDataType"`
}

// The server-side encryption key configuration for a customer provided encryption key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseConfigurationProperty := &sseConfigurationProperty{
//   	kmsEncryptionConfig: &kmsEncryptionConfigProperty{
//   		cmkType: jsii.String("cmkType"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
type CfnFHIRDatastore_SseConfigurationProperty struct {
	// The server-side encryption key configuration for a customer provided encryption key (CMK).
	KmsEncryptionConfig interface{} `field:"required" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
}

// Properties for defining a `CfnFHIRDatastore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFHIRDatastoreProps := &cfnFHIRDatastoreProps{
//   	datastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//
//   	// the properties below are optional
//   	datastoreName: jsii.String("datastoreName"),
//   	preloadDataConfig: &preloadDataConfigProperty{
//   		preloadDataType: jsii.String("preloadDataType"),
//   	},
//   	sseConfiguration: &sseConfigurationProperty{
//   		kmsEncryptionConfig: &kmsEncryptionConfigProperty{
//   			cmkType: jsii.String("cmkType"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFHIRDatastoreProps struct {
	// The FHIR version of the Data Store.
	//
	// The only supported version is R4.
	DatastoreTypeVersion *string `field:"required" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The user generated name for the Data Store.
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The preloaded data configuration for the Data Store.
	//
	// Only data preloaded from Synthea is supported.
	PreloadDataConfig interface{} `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The server-side encryption key configuration for a customer provided encryption key specified for creating a Data Store.
	SseConfiguration interface{} `field:"optional" json:"sseConfiguration" yaml:"sseConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

