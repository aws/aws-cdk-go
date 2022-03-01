package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to add a secret attachment to a secret.
//
// TODO: EXAMPLE
//
type AttachedSecretOptions struct {
	// The target to attach the secret to.
	Target ISecretAttachmentTarget `json:"target" yaml:"target"`
}

// The type of service or database that's being associated with the secret.
type AttachmentTargetType string

const (
	AttachmentTargetType_RDS_DB_INSTANCE AttachmentTargetType = "RDS_DB_INSTANCE"
	AttachmentTargetType_RDS_DB_CLUSTER AttachmentTargetType = "RDS_DB_CLUSTER"
	AttachmentTargetType_RDS_DB_PROXY AttachmentTargetType = "RDS_DB_PROXY"
	AttachmentTargetType_REDSHIFT_CLUSTER AttachmentTargetType = "REDSHIFT_CLUSTER"
	AttachmentTargetType_DOCDB_DB_INSTANCE AttachmentTargetType = "DOCDB_DB_INSTANCE"
	AttachmentTargetType_DOCDB_DB_CLUSTER AttachmentTargetType = "DOCDB_DB_CLUSTER"
)

// A CloudFormation `AWS::SecretsManager::ResourcePolicy`.
//
// Attaches a resource-based permission policy to a secret. A resource-based policy is optional. For more information, see [Authentication and access control for Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
//
// For information about attaching a policy in the console, see [Attach a permissions policy to a secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html) .
//
// *Required permissions:* `secretsmanager:PutResourcePolicy` . For more information, see [IAM policy actions for Secrets Manager](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awssecretsmanager.html#awssecretsmanager-actions-as-permissions) and [Authentication and access control in Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html) .
//
// TODO: EXAMPLE
//
type CfnResourcePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	BlockPublicPolicy() interface{}
	SetBlockPublicPolicy(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ResourcePolicy() interface{}
	SetResourcePolicy(val interface{})
	SecretId() *string
	SetSecretId(val *string)
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnResourcePolicy
type jsiiProxy_CfnResourcePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourcePolicy) BlockPublicPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) ResourcePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SecretsManager::ResourcePolicy`.
func NewCfnResourcePolicy(scope constructs.Construct, id *string, props *CfnResourcePolicyProps) CfnResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResourcePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SecretsManager::ResourcePolicy`.
func NewCfnResourcePolicy_Override(c CfnResourcePolicy, scope constructs.Construct, id *string, props *CfnResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetBlockPublicPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetResourcePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"resourcePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetSecretId(val *string) {
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourcePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourcePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
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
func CfnResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.CfnResourcePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnResourcePolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnResourcePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnResourcePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnResourcePolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnResourcePolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnResourcePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnResourcePolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourcePolicy`.
//
// TODO: EXAMPLE
//
type CfnResourcePolicyProps struct {
	// A JSON-formatted string for an AWS resource-based policy.
	//
	// For example policies, see [Permissions policy examples](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html) .
	ResourcePolicy interface{} `json:"resourcePolicy" yaml:"resourcePolicy"`
	// The ARN or name of the secret to attach the resource-based policy.
	//
	// For an ARN, we recommend that you specify a complete ARN rather than a partial ARN.
	SecretId *string `json:"secretId" yaml:"secretId"`
	// Specifies whether to block resource-based policies that allow broad access to the secret.
	//
	// By default, Secrets Manager blocks policies that allow broad access, for example those that use a wildcard for the principal.
	BlockPublicPolicy interface{} `json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
}

// A CloudFormation `AWS::SecretsManager::RotationSchedule`.
//
// Configures rotation for a secret. You must already configure the secret with the details of the database or service. If you define both the secret and the database or service in an AWS CloudFormation template, then define the [AWS::SecretsManager::SecretTargetAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html) resource to populate the secret with the connection details of the database or service before you attempt to configure rotation.
//
// TODO: EXAMPLE
//
type CfnRotationSchedule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HostedRotationLambda() interface{}
	SetHostedRotationLambda(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RotateImmediatelyOnUpdate() interface{}
	SetRotateImmediatelyOnUpdate(val interface{})
	RotationLambdaArn() *string
	SetRotationLambdaArn(val *string)
	RotationRules() interface{}
	SetRotationRules(val interface{})
	SecretId() *string
	SetSecretId(val *string)
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnRotationSchedule
type jsiiProxy_CfnRotationSchedule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRotationSchedule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) HostedRotationLambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostedRotationLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) RotateImmediatelyOnUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateImmediatelyOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) RotationLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) RotationRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotationRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationSchedule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SecretsManager::RotationSchedule`.
func NewCfnRotationSchedule(scope constructs.Construct, id *string, props *CfnRotationScheduleProps) CfnRotationSchedule {
	_init_.Initialize()

	j := jsiiProxy_CfnRotationSchedule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SecretsManager::RotationSchedule`.
func NewCfnRotationSchedule_Override(c CfnRotationSchedule, scope constructs.Construct, id *string, props *CfnRotationScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRotationSchedule) SetHostedRotationLambda(val interface{}) {
	_jsii_.Set(
		j,
		"hostedRotationLambda",
		val,
	)
}

func (j *jsiiProxy_CfnRotationSchedule) SetRotateImmediatelyOnUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"rotateImmediatelyOnUpdate",
		val,
	)
}

func (j *jsiiProxy_CfnRotationSchedule) SetRotationLambdaArn(val *string) {
	_jsii_.Set(
		j,
		"rotationLambdaArn",
		val,
	)
}

func (j *jsiiProxy_CfnRotationSchedule) SetRotationRules(val interface{}) {
	_jsii_.Set(
		j,
		"rotationRules",
		val,
	)
}

func (j *jsiiProxy_CfnRotationSchedule) SetSecretId(val *string) {
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRotationSchedule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRotationSchedule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
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
func CfnRotationSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRotationSchedule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.CfnRotationSchedule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRotationSchedule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRotationSchedule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRotationSchedule) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnRotationSchedule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRotationSchedule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRotationSchedule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRotationSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRotationSchedule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRotationSchedule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRotationSchedule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRotationSchedule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRotationSchedule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRotationSchedule) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRotationSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRotationSchedule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies that you want to create a hosted Lambda rotation function.
//
// To use these values, you must specify `Transform: AWS::SecretsManager-2020-07-23` at the beginning of the CloudFormation template.
//
// TODO: EXAMPLE
//
type CfnRotationSchedule_HostedRotationLambdaProperty struct {
	// The type of rotation template to use. For more information, see [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) .
	//
	// You can specify one of the following `RotationTypes` :
	//
	// - MySQLSingleUser
	// - MySQLMultiUser
	// - PostgreSQLSingleUser
	// - PostgreSQLMultiUser
	// - OracleSingleUser
	// - OracleMultiUser
	// - MariaDBSingleUser
	// - MariaDBMultiUser
	// - SQLServerSingleUser
	// - SQLServerMultiUser
	// - RedshiftSingleUser
	// - RedshiftMultiUser
	// - MongoDBSingleUser
	// - MongoDBMultiUser
	RotationType *string `json:"rotationType" yaml:"rotationType"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the secret.
	//
	// If you don't specify this value, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the secret that contains elevated credentials.
	//
	// The Lambda rotation function uses this secret for the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) .
	MasterSecretArn *string `json:"masterSecretArn" yaml:"masterSecretArn"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the elevated secret if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . If you don't specify this value and you use the alternating users strategy, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	MasterSecretKmsKeyArn *string `json:"masterSecretKmsKeyArn" yaml:"masterSecretKmsKeyArn"`
	// The name of the Lambda rotation function.
	RotationLambdaName *string `json:"rotationLambdaName" yaml:"rotationLambdaName"`
	// The ARN of the secret that contains elevated credentials.
	//
	// The Lambda rotation function uses this secret for the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) .
	SuperuserSecretArn *string `json:"superuserSecretArn" yaml:"superuserSecretArn"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the elevated secret if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . If you don't specify this value and you use the alternating users strategy, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	SuperuserSecretKmsKeyArn *string `json:"superuserSecretKmsKeyArn" yaml:"superuserSecretKmsKeyArn"`
	// A comma-separated list of security group IDs applied to the target database.
	//
	// The templates applies the same security groups as on the Lambda rotation function that is created as part of this stack.
	VpcSecurityGroupIds *string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A comma separated list of VPC subnet IDs of the target database network.
	//
	// The Lambda rotation function is in the same subnet group.
	VpcSubnetIds *string `json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

// A structure that defines the rotation configuration for the secret.
//
// TODO: EXAMPLE
//
type CfnRotationSchedule_RotationRulesProperty struct {
	// The number of days between automatic scheduled rotations of the secret.
	//
	// You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.
	//
	// In `DescribeSecret` and `ListSecrets` , this value is calculated from the rotation schedule after every successful rotation. In `RotateSecret` , you can set the rotation schedule in `RotationRules` with `AutomaticallyAfterDays` or `ScheduleExpression` , but not both.
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
	// The length of the rotation window in hours, for example `3h` for a three hour window.
	//
	// Secrets Manager rotates your secret at any time during this window. The window must not go into the next UTC day. If you don't specify this value, the window automatically ends at the end of the UTC day. The window begins according to the `ScheduleExpression` . For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) .
	Duration *string `json:"duration" yaml:"duration"`
	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret.
	//
	// Secrets Manager rotation schedules use UTC time zone.
	//
	// Secrets Manager `rate()` expressions represent the interval in days that you want to rotate your secret, for example `rate(10 days)` . If you use a `rate()` expression, the rotation window opens at midnight, and Secrets Manager rotates your secret any time that day after midnight. You can set a `Duration` to shorten the rotation window.
	//
	// You can use a `cron()` expression to create rotation schedules that are more detailed than a rotation interval. For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) . If you use a `cron()` expression, Secrets Manager rotates your secret any time during that day after the window opens. For example, `cron(0 8 1 * ? *)` represents a rotation window that occurs on the first day of every month beginning at 8:00 AM UTC. Secrets Manager rotates the secret any time that day after 8:00 AM. You can set a `Duration` to shorten the rotation window.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
}

// Properties for defining a `CfnRotationSchedule`.
//
// TODO: EXAMPLE
//
type CfnRotationScheduleProps struct {
	// The ARN or name of the secret to rotate.
	//
	// To reference a secret also created in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	SecretId *string `json:"secretId" yaml:"secretId"`
	// To use these values, you must specify `Transform: AWS::SecretsManager-2020-07-23` at the beginning of the CloudFormation template.
	//
	// When you enter valid values for `RotationSchedule.HostedRotationLambda` , Secrets Manager launches a Lambda that performs rotation on the secret specified in the `secret-id` property. The template creates a Lambda as part of a nested stack within the current stack.
	HostedRotationLambda interface{} `json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	//
	// The rotation schedule is defined in `RotationRules` .
	//
	// If you don't immediately rotate the secret, Secrets Manager tests the rotation configuration by running the [`testSecret` step](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) of the Lambda rotation function. The test creates an `AWSPENDING` version of the secret and then removes it.
	//
	// If you don't specify this value, then by default, Secrets Manager rotates the secret immediately.
	RotateImmediatelyOnUpdate interface{} `json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The ARN of the Lambda function that can rotate the secret.
	//
	// If you don't specify this parameter, then the secret must already have the ARN of a Lambda function configured.
	//
	// To reference a Lambda function also created in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the function's logical ID.
	RotationLambdaArn *string `json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	RotationRules interface{} `json:"rotationRules" yaml:"rotationRules"`
}

// A CloudFormation `AWS::SecretsManager::Secret`.
//
// Creates a new secret. A *secret* is a set of credentials, such as a user name and password, that you store in an encrypted form in Secrets Manager. The secret also includes the connection information to access a database or other service, which Secrets Manager doesn't encrypt. A secret in Secrets Manager consists of both the protected secret data and the important information needed to manage the secret.
//
// For information about creating a secret in the console, see [Create a secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html) .
//
// For information about creating a secret using the CLI or SDK, see [CreateSecret](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html) .
//
// To specify the encrypted value for the secret, you must include either the `GenerateSecretString` or the `SecretString` property, but not both. We recommend that you use the `GenerateSecretString` property to generate a random password as shown in the examples. You can't generate a secret with a `SecretBinary` secret value using AWS CloudFormation .
//
// > Do not create a dynamic reference using a backslash `(\)` as the final value. AWS CloudFormation cannot resolve those references, which causes a resource failure.
//
// TODO: EXAMPLE
//
type CfnSecret interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	GenerateSecretString() interface{}
	SetGenerateSecretString(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	ReplicaRegions() interface{}
	SetReplicaRegions(val interface{})
	SecretString() *string
	SetSecretString(val *string)
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

// The jsii proxy struct for CfnSecret
type jsiiProxy_CfnSecret struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecret) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) GenerateSecretString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecretString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) ReplicaRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) SecretString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecret) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SecretsManager::Secret`.
func NewCfnSecret(scope constructs.Construct, id *string, props *CfnSecretProps) CfnSecret {
	_init_.Initialize()

	j := jsiiProxy_CfnSecret{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SecretsManager::Secret`.
func NewCfnSecret_Override(c CfnSecret, scope constructs.Construct, id *string, props *CfnSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecret) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnSecret) SetGenerateSecretString(val interface{}) {
	_jsii_.Set(
		j,
		"generateSecretString",
		val,
	)
}

func (j *jsiiProxy_CfnSecret) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnSecret) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSecret) SetReplicaRegions(val interface{}) {
	_jsii_.Set(
		j,
		"replicaRegions",
		val,
	)
}

func (j *jsiiProxy_CfnSecret) SetSecretString(val *string) {
	_jsii_.Set(
		j,
		"secretString",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSecret_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSecret_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
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
func CfnSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecret_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.CfnSecret",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSecret) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSecret) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSecret) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnSecret) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSecret) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSecret) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSecret) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSecret) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSecret) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSecret) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecret) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSecret) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecret) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Generates a random password.
//
// We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
//
// *Required permissions:* `secretsmanager:GetRandomPassword` . For more information, see [IAM policy actions for Secrets Manager](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awssecretsmanager.html#awssecretsmanager-actions-as-permissions) and [Authentication and access control in Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html) .
//
// TODO: EXAMPLE
//
type CfnSecret_GenerateSecretStringProperty struct {
	// A string of the characters that you don't want in the password.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies whether to exclude lowercase letters from the password.
	//
	// If you don't include this switch, the password can contain lowercase letters.
	ExcludeLowercase interface{} `json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies whether to exclude numbers from the password.
	//
	// If you don't include this switch, the password can contain numbers.
	ExcludeNumbers interface{} `json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies whether to exclude the following punctuation characters from the password: `!
	//
	// " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~` . If you don't include this switch, the password can contain punctuation.
	ExcludePunctuation interface{} `json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies whether to exclude uppercase letters from the password.
	//
	// If you don't include this switch, the password can contain uppercase letters.
	ExcludeUppercase interface{} `json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name for the key/value pair, where the value is the generated password.
	//
	// This pair is added to the JSON structure specified by the `SecretStringTemplate` parameter. If you specify this parameter, then you must also specify `SecretStringTemplate` .
	GenerateStringKey *string `json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies whether to include the space character.
	//
	// If you include this switch, the password can contain space characters.
	IncludeSpace interface{} `json:"includeSpace" yaml:"includeSpace"`
	// The length of the password.
	//
	// If you don't include this parameter, the default length is 32 characters.
	PasswordLength *float64 `json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation.
	//
	// If you don't include this switch, the password contains at least one of every character type.
	RequireEachIncludedType interface{} `json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A template that the generated string must match.
	SecretStringTemplate *string `json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

// A custom type that specifies a `Region` and the `KmsKeyId` for a replica secret.
//
// TODO: EXAMPLE
//
type CfnSecret_ReplicaRegionProperty struct {
	// `CfnSecret.ReplicaRegionProperty.Region`.
	Region *string `json:"region" yaml:"region"`
	// The ARN, key ID, or alias of the KMS key to encrypt the secret.
	//
	// If you don't include this field, Secrets Manager uses `aws/secretsmanager` .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Properties for defining a `CfnSecret`.
//
// TODO: EXAMPLE
//
type CfnSecretProps struct {
	// The description of the secret.
	Description *string `json:"description" yaml:"description"`
	// A structure that specifies how to generate a password to encrypt and store in the secret.
	//
	// Either `GenerateSecretString` or `SecretString` must have a value, but not both. They cannot both be empty.
	//
	// We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	GenerateSecretString interface{} `json:"generateSecretString" yaml:"generateSecretString"`
	// The ARN, key ID, or alias of the AWS KMS key that Secrets Manager uses to encrypt the secret value in the secret.
	//
	// To use a AWS KMS key in a different account, use the key ARN or the alias ARN.
	//
	// If you don't specify this value, then Secrets Manager uses the key `aws/secretsmanager` . If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//
	// If the secret is in a different AWS account from the credentials calling the API, then you can't use `aws/secretsmanager` to encrypt the secret, and you must create and use a customer managed AWS KMS key.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the new secret.
	//
	// The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//
	// Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	Name *string `json:"name" yaml:"name"`
	// A custom type that specifies a `Region` and the `KmsKeyId` for a replica secret.
	ReplicaRegions interface{} `json:"replicaRegions" yaml:"replicaRegions"`
	// The text to encrypt and store in the secret.
	//
	// We recommend you use a JSON structure of key/value pairs for your secret value.
	//
	// Either `GenerateSecretString` or `SecretString` must have a value, but not both. They cannot both be empty. We recommend that you use the `GenerateSecretString` property to generate a random password.
	SecretString *string `json:"secretString" yaml:"secretString"`
	// A list of tags to attach to the secret.
	//
	// Each tag is a key and value pair of strings in a JSON text string, for example:
	//
	// `[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]`
	//
	// Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//
	// If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an `Access Denied` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2) .
	//
	// For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json) . If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//
	// The following restrictions apply to tags:
	//
	// - Maximum number of tags per secret: 50
	// - Maximum key length: 127 Unicode characters in UTF-8
	// - Maximum value length: 255 Unicode characters in UTF-8
	// - Tag keys and values are case sensitive.
	// - Do not use the `aws:` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	// - If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::SecretsManager::SecretTargetAttachment`.
//
// The `AWS::SecretsManager::SecretTargetAttachment` resource completes the final link between a Secrets Manager secret and the associated database. This is required because each has a dependency on the other. No matter which one you create first, the other doesn't exist yet. To resolve this, you must create the resources in the following order:
//
// - Define the secret without referencing the service or database. You can't reference the service or database because it doesn't exist yet. The secret must contain a user name and password.
// - Next, define the service or database. Include the reference to the secret to use stored credentials to define the database admin user and password.
// - Finally, define a `SecretTargetAttachment` resource type to finish configuring the secret with the required database engine type and the connection details of the service or database. The rotation function requires the details, if you attach one later by defining a [AWS::SecretsManager::RotationSchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html) resource type.
//
// TODO: EXAMPLE
//
type CfnSecretTargetAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SecretId() *string
	SetSecretId(val *string)
	Stack() awscdk.Stack
	TargetId() *string
	SetTargetId(val *string)
	TargetType() *string
	SetTargetType(val *string)
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

// The jsii proxy struct for CfnSecretTargetAttachment
type jsiiProxy_CfnSecretTargetAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecretTargetAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecretTargetAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SecretsManager::SecretTargetAttachment`.
func NewCfnSecretTargetAttachment(scope constructs.Construct, id *string, props *CfnSecretTargetAttachmentProps) CfnSecretTargetAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnSecretTargetAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SecretsManager::SecretTargetAttachment`.
func NewCfnSecretTargetAttachment_Override(c CfnSecretTargetAttachment, scope constructs.Construct, id *string, props *CfnSecretTargetAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecretTargetAttachment) SetSecretId(val *string) {
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

func (j *jsiiProxy_CfnSecretTargetAttachment) SetTargetId(val *string) {
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

func (j *jsiiProxy_CfnSecretTargetAttachment) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSecretTargetAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSecretTargetAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
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
func CfnSecretTargetAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecretTargetAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.CfnSecretTargetAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSecretTargetAttachment) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnSecretTargetAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSecretTargetAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSecretTargetAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSecretTargetAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecretTargetAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSecretTargetAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecretTargetAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSecretTargetAttachment`.
//
// TODO: EXAMPLE
//
type CfnSecretTargetAttachmentProps struct {
	// The ARN or name of the secret.
	//
	// To reference a secret also created in this template, use the see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	SecretId *string `json:"secretId" yaml:"secretId"`
	// The ID of the database or cluster.
	TargetId *string `json:"targetId" yaml:"targetId"`
	// A string that defines the type of service or database associated with the secret.
	//
	// This value instructs Secrets Manager how to update the secret with the details of the service or database. This value must be one of the following:
	//
	// - AWS::RDS::DBInstance
	// - AWS::RDS::DBCluster
	// - AWS::Redshift::Cluster
	// - AWS::DocDB::DBInstance
	// - AWS::DocDB::DBCluster
	TargetType *string `json:"targetType" yaml:"targetType"`
}

// A hosted rotation.
//
// TODO: EXAMPLE
//
type HostedRotation interface {
	awsec2.IConnectable
	Connections() awsec2.Connections
	Bind(secret ISecret, scope constructs.Construct) *CfnRotationSchedule_HostedRotationLambdaProperty
}

// The jsii proxy struct for HostedRotation
type jsiiProxy_HostedRotation struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_HostedRotation) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// MariaDB Multi User.
func HostedRotation_MariaDbMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mariaDbMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MariaDB Single User.
func HostedRotation_MariaDbSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mariaDbSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MongoDB Multi User.
func HostedRotation_MongoDbMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mongoDbMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MongoDB Single User.
func HostedRotation_MongoDbSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mongoDbSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MySQL Multi User.
func HostedRotation_MysqlMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mysqlMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MySQL Single User.
func HostedRotation_MysqlSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mysqlSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Oracle Multi User.
func HostedRotation_OracleMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"oracleMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Oracle Single User.
func HostedRotation_OracleSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"oracleSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// PostgreSQL Multi User.
func HostedRotation_PostgreSqlMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"postgreSqlMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// PostgreSQL Single User.
func HostedRotation_PostgreSqlSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"postgreSqlSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Redshift Multi User.
func HostedRotation_RedshiftMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"redshiftMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Redshift Single User.
func HostedRotation_RedshiftSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"redshiftSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// SQL Server Multi User.
func HostedRotation_SqlServerMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"sqlServerMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// SQL Server Single User.
func HostedRotation_SqlServerSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"sqlServerSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Binds this hosted rotation to a secret.
func (h *jsiiProxy_HostedRotation) Bind(secret ISecret, scope constructs.Construct) *CfnRotationSchedule_HostedRotationLambdaProperty {
	var returns *CfnRotationSchedule_HostedRotationLambdaProperty

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{secret, scope},
		&returns,
	)

	return returns
}

// Hosted rotation type.
//
// TODO: EXAMPLE
//
type HostedRotationType interface {
	IsMultiUser() *bool
	Name() *string
}

// The jsii proxy struct for HostedRotationType
type jsiiProxy_HostedRotationType struct {
	_ byte // padding
}

func (j *jsiiProxy_HostedRotationType) IsMultiUser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isMultiUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedRotationType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func HostedRotationType_MARIADB_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MARIADB_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MARIADB_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MARIADB_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MONGODB_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MONGODB_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MONGODB_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MONGODB_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MYSQL_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MYSQL_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MYSQL_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MYSQL_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_ORACLE_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"ORACLE_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_ORACLE_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"ORACLE_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_POSTGRESQL_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"POSTGRESQL_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_POSTGRESQL_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"POSTGRESQL_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_REDSHIFT_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"REDSHIFT_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_REDSHIFT_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"REDSHIFT_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_SQLSERVER_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"SQLSERVER_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_SQLSERVER_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"SQLSERVER_SINGLE_USER",
		&returns,
	)
	return returns
}

// A secret in AWS Secrets Manager.
type ISecret interface {
	awscdk.IResource
	// Adds a rotation schedule to the secret.
	AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Attach a target to this secret.
	//
	// Returns: An attached secret
	Attach(target ISecretAttachmentTarget) ISecret
	// Denies the `DeleteSecret` action to all principals within the current account.
	DenyAccountRootDelete()
	// Grants reading the secret value to some role.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	SecretValueFromJson(key *string) awscdk.SecretValue
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
	EncryptionKey() awskms.IKey
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	SecretName() *string
	// Retrieve the value of the stored secret as a `SecretValue`.
	SecretValue() awscdk.SecretValue
}

// The jsii proxy for ISecret
type jsiiProxy_ISecret struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISecret) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
	var returns RotationSchedule

	_jsii_.Invoke(
		i,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecret) Attach(target ISecretAttachmentTarget) ISecret {
	var returns ISecret

	_jsii_.Invoke(
		i,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		i,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ISecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecret) SecretValueFromJson(key *string) awscdk.SecretValue {
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		i,
		"secretValueFromJson",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

// A secret attachment target.
type ISecretAttachmentTarget interface {
	// Renders the target specifications.
	AsSecretAttachmentTarget() *SecretAttachmentTargetProps
}

// The jsii proxy for ISecretAttachmentTarget
type jsiiProxy_ISecretAttachmentTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ISecretAttachmentTarget) AsSecretAttachmentTarget() *SecretAttachmentTargetProps {
	var returns *SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ISecretTargetAttachment interface {
	ISecret
	// Same as `secretArn`.
	SecretTargetAttachmentSecretArn() *string
}

// The jsii proxy for ISecretTargetAttachment
type jsiiProxy_ISecretTargetAttachment struct {
	jsiiProxy_ISecret
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretTargetAttachmentSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTargetAttachmentSecretArn",
		&returns,
	)
	return returns
}

// Multi user hosted rotation options.
//
// TODO: EXAMPLE
//
type MultiUserHostedRotationOptions struct {
	// A name for the Lambda created to rotate the secret.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// A list of security groups for the Lambda created to rotate the secret.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The VPC where the Lambda rotation function will run.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The master secret for a multi user rotation scheme.
	MasterSecret ISecret `json:"masterSecret" yaml:"masterSecret"`
}

// Secret replica region.
//
// TODO: EXAMPLE
//
type ReplicaRegion struct {
	// The name of the region.
	Region *string `json:"region" yaml:"region"`
	// The customer-managed encryption key to use for encrypting the secret value.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
}

// Resource Policy for SecretsManager Secrets.
//
// Policies define the operations that are allowed on this resource.
//
// You almost never need to define this construct directly.
//
// All AWS resources that support resource policies have a method called
// `addToResourcePolicy()`, which will automatically create a new resource
// policy if one doesn't exist yet, otherwise it will add to the existing
// policy.
//
// Prefer to use `addToResourcePolicy()` instead.
//
// TODO: EXAMPLE
//
type ResourcePolicy interface {
	awscdk.Resource
	Document() awsiam.PolicyDocument
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ResourcePolicy
type jsiiProxy_ResourcePolicy struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ResourcePolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewResourcePolicy(scope constructs.Construct, id *string, props *ResourcePolicyProps) ResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_ResourcePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.ResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewResourcePolicy_Override(r ResourcePolicy, scope constructs.Construct, id *string, props *ResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.ResourcePolicy",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.ResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ResourcePolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.ResourcePolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
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
func (r *jsiiProxy_ResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ResourcePolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ResourcePolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ResourcePolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a ResourcePolicy.
//
// TODO: EXAMPLE
//
type ResourcePolicyProps struct {
	// The secret to attach a resource-based permissions policy.
	Secret ISecret `json:"secret" yaml:"secret"`
}

// A rotation schedule.
//
// TODO: EXAMPLE
//
type RotationSchedule interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for RotationSchedule
type jsiiProxy_RotationSchedule struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_RotationSchedule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RotationSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RotationSchedule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RotationSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRotationSchedule(scope constructs.Construct, id *string, props *RotationScheduleProps) RotationSchedule {
	_init_.Initialize()

	j := jsiiProxy_RotationSchedule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.RotationSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRotationSchedule_Override(r RotationSchedule, scope constructs.Construct, id *string, props *RotationScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.RotationSchedule",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func RotationSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.RotationSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RotationSchedule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.RotationSchedule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
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
func (r *jsiiProxy_RotationSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RotationSchedule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_RotationSchedule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_RotationSchedule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_RotationSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to add a rotation schedule to a secret.
//
// TODO: EXAMPLE
//
type RotationScheduleOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// Hosted rotation.
	HostedRotation HostedRotation `json:"hostedRotation" yaml:"hostedRotation"`
	// A Lambda function that can rotate the secret.
	RotationLambda awslambda.IFunction `json:"rotationLambda" yaml:"rotationLambda"`
}

// Construction properties for a RotationSchedule.
//
// TODO: EXAMPLE
//
type RotationScheduleProps struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// Hosted rotation.
	HostedRotation HostedRotation `json:"hostedRotation" yaml:"hostedRotation"`
	// A Lambda function that can rotate the secret.
	RotationLambda awslambda.IFunction `json:"rotationLambda" yaml:"rotationLambda"`
	// The secret to rotate.
	//
	// If hosted rotation is used, this must be a JSON string with the following format:
	//
	// ```
	// {
	//    "engine": <required: database engine>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port will be used>,
	//    "masterarn": <required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	//
	// This is typically the case for a secret referenced from an `AWS::SecretsManager::SecretTargetAttachment`
	// or an `ISecret` returned by the `attach()` method of `Secret`.
	Secret ISecret `json:"secret" yaml:"secret"`
}

// Creates a new secret in AWS SecretsManager.
//
// TODO: EXAMPLE
//
type Secret interface {
	awscdk.Resource
	ISecret
	ArnForPolicies() *string
	AutoCreatePolicy() *bool
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	SecretArn() *string
	SecretFullArn() *string
	SecretName() *string
	SecretValue() awscdk.SecretValue
	Stack() awscdk.Stack
	AddReplicaRegion(region *string, encryptionKey awskms.IKey)
	AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	Attach(target ISecretAttachmentTarget) ISecret
	DenyAccountRootDelete()
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	ToString() *string
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecret
}

func (j *jsiiProxy_Secret) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSecret(scope constructs.Construct, id *string, props *SecretProps) Secret {
	_init_.Initialize()

	j := jsiiProxy_Secret{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecret_Override(s Secret, scope constructs.Construct, id *string, props *SecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an existing secret into the Stack.
func Secret_FromSecretAttributes(scope constructs.Construct, id *string, attrs *SecretAttributes) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"fromSecretAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a secret by complete ARN.
//
// The complete ARN is the ARN with the Secrets Manager-supplied suffix.
func Secret_FromSecretCompleteArn(scope constructs.Construct, id *string, secretCompleteArn *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"fromSecretCompleteArn",
		[]interface{}{scope, id, secretCompleteArn},
		&returns,
	)

	return returns
}

// Imports a secret by secret name.
//
// A secret with this name must exist in the same account & region.
// Replaces the deprecated `fromSecretName`.
func Secret_FromSecretNameV2(scope constructs.Construct, id *string, secretName *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"fromSecretNameV2",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by partial ARN.
//
// The partial ARN is the ARN without the Secrets Manager-supplied suffix.
func Secret_FromSecretPartialArn(scope constructs.Construct, id *string, secretPartialArn *string) ISecret {
	_init_.Initialize()

	var returns ISecret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"fromSecretPartialArn",
		[]interface{}{scope, id, secretPartialArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Secret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Secret_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.Secret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a replica region for the secret.
func (s *jsiiProxy_Secret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	_jsii_.InvokeVoid(
		s,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

// Adds a rotation schedule to the secret.
func (s *jsiiProxy_Secret) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
	var returns RotationSchedule

	_jsii_.Invoke(
		s,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a statement to the IAM resource policy associated with this secret.
//
// If this secret was created in this stack, a resource policy will be
// automatically created upon the first call to `addToResourcePolicy`. If
// the secret is imported, then this is a no-op.
func (s *jsiiProxy_Secret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
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
func (s *jsiiProxy_Secret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Attach a target to this secret.
//
// Returns: An attached secret
func (s *jsiiProxy_Secret) Attach(target ISecretAttachmentTarget) ISecret {
	var returns ISecret

	_jsii_.Invoke(
		s,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Denies the `DeleteSecret` action to all principals within the current account.
func (s *jsiiProxy_Secret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		s,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Secret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_Secret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_Secret) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants reading the secret value to some role.
func (s *jsiiProxy_Secret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

// Grants writing and updating the secret value to some role.
func (s *jsiiProxy_Secret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
func (s *jsiiProxy_Secret) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		s,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_Secret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attachment target specifications.
//
// TODO: EXAMPLE
//
type SecretAttachmentTargetProps struct {
	// The id of the target to attach the secret to.
	TargetId *string `json:"targetId" yaml:"targetId"`
	// The type of the target to attach the secret to.
	TargetType AttachmentTargetType `json:"targetType" yaml:"targetType"`
}

// Attributes required to import an existing secret into the Stack.
//
// One ARN format (`secretArn`, `secretCompleteArn`, `secretPartialArn`) must be provided.
//
// TODO: EXAMPLE
//
type SecretAttributes struct {
	// The encryption key that is used to encrypt the secret, unless the default SecretsManager key is used.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The complete ARN of the secret in SecretsManager.
	//
	// This is the ARN including the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretPartialArn`.
	SecretCompleteArn *string `json:"secretCompleteArn" yaml:"secretCompleteArn"`
	// The partial ARN of the secret in SecretsManager.
	//
	// This is the ARN without the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretCompleteArn`.
	SecretPartialArn *string `json:"secretPartialArn" yaml:"secretPartialArn"`
}

// The properties required to create a new secret in AWS Secrets Manager.
//
// TODO: EXAMPLE
//
type SecretProps struct {
	// An optional, human-friendly description of the secret.
	Description *string `json:"description" yaml:"description"`
	// The customer-managed encryption key to use for encrypting the secret value.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Configuration for how to generate a secret value.
	//
	// Only one of `secretString` and `generateSecretString` can be provided.
	GenerateSecretString *SecretStringGenerator `json:"generateSecretString" yaml:"generateSecretString"`
	// Policy to apply when the secret is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*ReplicaRegion `json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	//
	// Note that deleting secrets from SecretsManager does not happen immediately, but after a 7 to
	// 30 days blackout period. During that period, it is not possible to create another secret that shares the same name.
	SecretName *string `json:"secretName" yaml:"secretName"`
	// Initial value for the secret.
	//
	// **NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
	// The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
	// and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
	// another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
	// to the CloudFormation template (via the AWS Console, SDKs, or CLI).
	//
	// Specifies text data that you want to encrypt and store in this new version of the secret.
	// May be a simple string value, or a string representation of a JSON structure.
	//
	// Only one of `secretString` and `generateSecretString` can be provided.
	SecretStringBeta1 SecretStringValueBeta1 `json:"secretStringBeta1" yaml:"secretStringBeta1"`
}

// Secret rotation for a service or database.
//
// TODO: EXAMPLE
//
type SecretRotation interface {
	constructs.Construct
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for SecretRotation
type jsiiProxy_SecretRotation struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SecretRotation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSecretRotation(scope constructs.Construct, id *string, props *SecretRotationProps) SecretRotation {
	_init_.Initialize()

	j := jsiiProxy_SecretRotation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecretRotation_Override(s SecretRotation, scope constructs.Construct, id *string, props *SecretRotationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SecretRotation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretRotation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SecretRotation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A secret rotation serverless application.
//
// TODO: EXAMPLE
//
type SecretRotationApplication interface {
	IsMultiUser() *bool
	ApplicationArnForPartition(partition *string) *string
	SemanticVersionForPartition(partition *string) *string
}

// The jsii proxy struct for SecretRotationApplication
type jsiiProxy_SecretRotationApplication struct {
	_ byte // padding
}

func (j *jsiiProxy_SecretRotationApplication) IsMultiUser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isMultiUser",
		&returns,
	)
	return returns
}


func NewSecretRotationApplication(applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) SecretRotationApplication {
	_init_.Initialize()

	j := jsiiProxy_SecretRotationApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		[]interface{}{applicationId, semanticVersion, options},
		&j,
	)

	return &j
}

func NewSecretRotationApplication_Override(s SecretRotationApplication, applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		[]interface{}{applicationId, semanticVersion, options},
		s,
	)
}

func SecretRotationApplication_MARIADB_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MARIADB_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MARIADB_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MARIADB_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MONGODB_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MONGODB_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MONGODB_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MONGODB_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MYSQL_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MYSQL_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MYSQL_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MYSQL_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_ORACLE_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"ORACLE_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_ORACLE_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"ORACLE_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_POSTGRES_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"POSTGRES_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_POSTGRES_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"POSTGRES_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_REDSHIFT_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"REDSHIFT_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_REDSHIFT_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"REDSHIFT_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_SQLSERVER_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"SQLSERVER_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_SQLSERVER_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"SQLSERVER_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

// Returns the application ARN for the current partition.
//
// Can be used in combination with a `CfnMapping` to automatically select the correct ARN based on the current partition.
func (s *jsiiProxy_SecretRotationApplication) ApplicationArnForPartition(partition *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"applicationArnForPartition",
		[]interface{}{partition},
		&returns,
	)

	return returns
}

// The semantic version of the app for the current partition.
//
// Can be used in combination with a `CfnMapping` to automatically select the correct version based on the current partition.
func (s *jsiiProxy_SecretRotationApplication) SemanticVersionForPartition(partition *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"semanticVersionForPartition",
		[]interface{}{partition},
		&returns,
	)

	return returns
}

// Options for a SecretRotationApplication.
//
// TODO: EXAMPLE
//
type SecretRotationApplicationOptions struct {
	// Whether the rotation application uses the mutli user scheme.
	IsMultiUser *bool `json:"isMultiUser" yaml:"isMultiUser"`
}

// Construction properties for a SecretRotation.
//
// TODO: EXAMPLE
//
type SecretRotationProps struct {
	// The serverless application for the rotation.
	Application SecretRotationApplication `json:"application" yaml:"application"`
	// The secret to rotate. It must be a JSON string with the following format:.
	//
	// ```
	// {
	//    "engine": <required: database engine>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port will be used>,
	//    "masterarn": <required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	//
	// This is typically the case for a secret referenced from an `AWS::SecretsManager::SecretTargetAttachment`
	// or an `ISecret` returned by the `attach()` method of `Secret`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
	//
	Secret ISecret `json:"secret" yaml:"secret"`
	// The target service or database.
	Target awsec2.IConnectable `json:"target" yaml:"target"`
	// The VPC where the Lambda rotation function will run.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	Endpoint awsec2.IInterfaceVpcEndpoint `json:"endpoint" yaml:"endpoint"`
	// Characters which should not appear in the generated password.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret for a multi user rotation scheme.
	MasterSecret ISecret `json:"masterSecret" yaml:"masterSecret"`
	// The security group for the Lambda rotation function.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Configuration to generate secrets such as passwords automatically.
//
// TODO: EXAMPLE
//
type SecretStringGenerator struct {
	// A string that includes characters that shouldn't be included in the generated password.
	//
	// The string can be a minimum
	// of ``0`` and a maximum of ``4096`` characters long.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies that the generated password shouldn't include lowercase letters.
	ExcludeLowercase *bool `json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies that the generated password shouldn't include digits.
	ExcludeNumbers *bool `json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies that the generated password shouldn't include punctuation characters.
	ExcludePunctuation *bool `json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies that the generated password shouldn't include uppercase letters.
	ExcludeUppercase *bool `json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name that's used to add the generated password to the JSON structure specified by the ``secretStringTemplate`` parameter.
	//
	// If you specify ``generateStringKey`` then ``secretStringTemplate``
	// must be also be specified.
	GenerateStringKey *string `json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies that the generated password can include the space character.
	IncludeSpace *bool `json:"includeSpace" yaml:"includeSpace"`
	// The desired length of the generated password.
	PasswordLength *float64 `json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether the generated password must include at least one of every allowed character type.
	RequireEachIncludedType *bool `json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A properly structured JSON string that the generated password can be added to.
	//
	// The ``generateStringKey`` is
	// combined with the generated random string and inserted into the JSON structure that's specified by this parameter.
	// The merged JSON string is returned as the completed SecretString of the secret. If you specify ``secretStringTemplate``
	// then ``generateStringKey`` must be also be specified.
	SecretStringTemplate *string `json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

// An experimental class used to specify an initial secret value for a Secret.
//
// The class wraps a simple string (or JSON representation) in order to provide some safety checks and warnings
// about the dangers of using plaintext strings as initial secret seed values via CDK/CloudFormation.
//
// TODO: EXAMPLE
//
type SecretStringValueBeta1 interface {
	SecretValue() *string
}

// The jsii proxy struct for SecretStringValueBeta1
type jsiiProxy_SecretStringValueBeta1 struct {
	_ byte // padding
}

// Creates a `SecretValueValueBeta1` from a string value coming from a Token.
//
// The intent is to enable creating secrets from references (e.g., `Ref`, `Fn::GetAtt`) from other resources.
// This might be the direct output of another Construct, or the output of a Custom Resource.
// This method throws if it determines the input is an unsafe plaintext string.
//
// For example:
// ```ts
//      // Creates a new IAM user, access and secret keys, and stores the secret access key in a Secret.
//      const user = new iam.User(this, 'User');
//      const accessKey = new iam.AccessKey(this, 'AccessKey', { user });
//      const secretValue = secretsmanager.SecretStringValueBeta1.fromToken(accessKey.secretAccessKey.toString());
//      new secretsmanager.Secret(this, 'Secret', {
//        secretStringBeta1: secretValue,
//      });
// ```
//
// The secret may also be embedded in a string representation of a JSON structure:
//      const secretValue = secretsmanager.SecretStringValueBeta1.fromToken(JSON.stringify({
//        username: user.userName,
//        database: 'foo',
//        password: accessKey.secretAccessKey.toString(),
//      }));
//
// Note that the value being a Token does *not* guarantee safety. For example, a Lazy-evaluated string
// (e.g., `Lazy.string({ produce: () => 'myInsecurePassword' }))`) is a Token, but as the output is
// ultimately a plaintext string, and so insecure.
func SecretStringValueBeta1_FromToken(secretValueFromToken *string) SecretStringValueBeta1 {
	_init_.Initialize()

	var returns SecretStringValueBeta1

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1",
		"fromToken",
		[]interface{}{secretValueFromToken},
		&returns,
	)

	return returns
}

// Creates a `SecretStringValueBeta1` from a plaintext value.
//
// This approach is inherently unsafe, as the secret value may be visible in your source control repository
// and will also appear in plaintext in the resulting CloudFormation template, including in the AWS Console or APIs.
// Usage of this method is discouraged, especially for production workloads.
func SecretStringValueBeta1_FromUnsafePlaintext(secretValue *string) SecretStringValueBeta1 {
	_init_.Initialize()

	var returns SecretStringValueBeta1

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1",
		"fromUnsafePlaintext",
		[]interface{}{secretValue},
		&returns,
	)

	return returns
}

// Returns the secret value.
func (s *jsiiProxy_SecretStringValueBeta1) SecretValue() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"secretValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An attached secret.
//
// TODO: EXAMPLE
//
type SecretTargetAttachment interface {
	awscdk.Resource
	ISecret
	ISecretTargetAttachment
	ArnForPolicies() *string
	AutoCreatePolicy() *bool
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	SecretArn() *string
	SecretFullArn() *string
	SecretName() *string
	SecretTargetAttachmentSecretArn() *string
	SecretValue() awscdk.SecretValue
	Stack() awscdk.Stack
	AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	Attach(target ISecretAttachmentTarget) ISecret
	DenyAccountRootDelete()
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	ToString() *string
}

// The jsii proxy struct for SecretTargetAttachment
type jsiiProxy_SecretTargetAttachment struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecret
	jsiiProxy_ISecretTargetAttachment
}

func (j *jsiiProxy_SecretTargetAttachment) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) SecretTargetAttachmentSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTargetAttachmentSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretTargetAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSecretTargetAttachment(scope constructs.Construct, id *string, props *SecretTargetAttachmentProps) SecretTargetAttachment {
	_init_.Initialize()

	j := jsiiProxy_SecretTargetAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretTargetAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecretTargetAttachment_Override(s SecretTargetAttachment, scope constructs.Construct, id *string, props *SecretTargetAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretTargetAttachment",
		[]interface{}{scope, id, props},
		s,
	)
}

func SecretTargetAttachment_FromSecretTargetAttachmentSecretArn(scope constructs.Construct, id *string, secretTargetAttachmentSecretArn *string) ISecretTargetAttachment {
	_init_.Initialize()

	var returns ISecretTargetAttachment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretTargetAttachment",
		"fromSecretTargetAttachmentSecretArn",
		[]interface{}{scope, id, secretTargetAttachmentSecretArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SecretTargetAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretTargetAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SecretTargetAttachment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.SecretTargetAttachment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a rotation schedule to the secret.
func (s *jsiiProxy_SecretTargetAttachment) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
	var returns RotationSchedule

	_jsii_.Invoke(
		s,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a statement to the IAM resource policy associated with this secret.
//
// If this secret was created in this stack, a resource policy will be
// automatically created upon the first call to `addToResourcePolicy`. If
// the secret is imported, then this is a no-op.
func (s *jsiiProxy_SecretTargetAttachment) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
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
func (s *jsiiProxy_SecretTargetAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Attach a target to this secret.
//
// Returns: An attached secret
func (s *jsiiProxy_SecretTargetAttachment) Attach(target ISecretAttachmentTarget) ISecret {
	var returns ISecret

	_jsii_.Invoke(
		s,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Denies the `DeleteSecret` action to all principals within the current account.
func (s *jsiiProxy_SecretTargetAttachment) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		s,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretTargetAttachment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SecretTargetAttachment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SecretTargetAttachment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants reading the secret value to some role.
func (s *jsiiProxy_SecretTargetAttachment) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

// Grants writing and updating the secret value to some role.
func (s *jsiiProxy_SecretTargetAttachment) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
func (s *jsiiProxy_SecretTargetAttachment) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		s,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SecretTargetAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for an AttachedSecret.
//
// TODO: EXAMPLE
//
type SecretTargetAttachmentProps struct {
	// The target to attach the secret to.
	Target ISecretAttachmentTarget `json:"target" yaml:"target"`
	// The secret to attach to the target.
	Secret ISecret `json:"secret" yaml:"secret"`
}

// Single user hosted rotation options.
//
// TODO: EXAMPLE
//
type SingleUserHostedRotationOptions struct {
	// A name for the Lambda created to rotate the secret.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// A list of security groups for the Lambda created to rotate the secret.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The VPC where the Lambda rotation function will run.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

