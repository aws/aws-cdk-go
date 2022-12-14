package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Kendra::Index`.
//
// Specifies a new Amazon Kendra index. And index is a collection of documents and associated metadata that you want to search for relevant documents.
//
// Once the index is active you can add documents to your index using the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation or using one of the supported data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndex := awscdk.Aws_kendra.NewCfnIndex(this, jsii.String("MyCfnIndex"), &cfnIndexProps{
//   	edition: jsii.String("edition"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	capacityUnits: &capacityUnitsConfigurationProperty{
//   		queryCapacityUnits: jsii.Number(123),
//   		storageCapacityUnits: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	documentMetadataConfigurations: []interface{}{
//   		&documentMetadataConfigurationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			relevance: &relevanceProperty{
//   				duration: jsii.String("duration"),
//   				freshness: jsii.Boolean(false),
//   				importance: jsii.Number(123),
//   				rankOrder: jsii.String("rankOrder"),
//   				valueImportanceItems: []interface{}{
//   					&valueImportanceItemProperty{
//   						key: jsii.String("key"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			search: &searchProperty{
//   				displayable: jsii.Boolean(false),
//   				facetable: jsii.Boolean(false),
//   				searchable: jsii.Boolean(false),
//   				sortable: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	serverSideEncryptionConfiguration: &serverSideEncryptionConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userContextPolicy: jsii.String("userContextPolicy"),
//   	userTokenConfigurations: []interface{}{
//   		&userTokenConfigurationProperty{
//   			jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   				keyLocation: jsii.String("keyLocation"),
//
//   				// the properties below are optional
//   				claimRegex: jsii.String("claimRegex"),
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				issuer: jsii.String("issuer"),
//   				secretManagerArn: jsii.String("secretManagerArn"),
//   				url: jsii.String("url"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   })
//
type CfnIndex interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the index.
	//
	// For example: `arn:aws:kendra:us-west-2:111122223333:index/0123456789abcdef` .
	AttrArn() *string
	// The identifier for the index.
	//
	// For example: `f4aeaa10-8056-4b2c-a343-522ca0f41234` .
	AttrId() *string
	// Specifies capacity units configured for your index.
	//
	// You can add and remove capacity units to tune an index to your requirements. You can set capacity units only for Enterprise edition indexes.
	CapacityUnits() interface{}
	SetCapacityUnits(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the index.
	Description() *string
	SetDescription(val *string)
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	DocumentMetadataConfigurations() interface{}
	SetDocumentMetadataConfigurations(val interface{})
	// Indicates whether the index is a enterprise edition index or a developer edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	Edition() *string
	SetEdition(val *string)
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
	// The identifier of the index.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn() *string
	SetRoleArn(val *string)
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
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
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If there is an access control list, it is ignored. You can filter on user and group attributes.
	//
	// USER_TOKEN
	//
	// - Enables SSO and token-based user access control. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy() *string
	SetUserContextPolicy(val *string)
	// Defines the type of user token used for the index.
	UserTokenConfigurations() interface{}
	SetUserTokenConfigurations(val interface{})
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

// The jsii proxy struct for CfnIndex
type jsiiProxy_CfnIndex struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIndex) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CapacityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) DocumentMetadataConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentMetadataConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserContextPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userContextPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserTokenConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userTokenConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Index`.
func NewCfnIndex(scope constructs.Construct, id *string, props *CfnIndexProps) CfnIndex {
	_init_.Initialize()

	if err := validateNewCfnIndexParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIndex{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Index`.
func NewCfnIndex_Override(c CfnIndex, scope constructs.Construct, id *string, props *CfnIndexProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIndex)SetCapacityUnits(val interface{}) {
	if err := j.validateSetCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetDocumentMetadataConfigurations(val interface{}) {
	if err := j.validateSetDocumentMetadataConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentMetadataConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetServerSideEncryptionConfiguration(val interface{}) {
	if err := j.validateSetServerSideEncryptionConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetUserContextPolicy(val *string) {
	_jsii_.Set(
		j,
		"userContextPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnIndex)SetUserTokenConfigurations(val interface{}) {
	if err := j.validateSetUserTokenConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTokenConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnIndex_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIndex_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIndex_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnIndex_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
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
func CfnIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndex_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIndex) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIndex) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIndex) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIndex) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIndex) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIndex) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIndex) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIndex) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnIndex) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnIndex) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIndex) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIndex) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnIndex) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

