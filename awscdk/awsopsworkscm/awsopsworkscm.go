package awsopsworkscm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsopsworkscm/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::OpsWorksCM::Server`.
type CfnServer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AttrArn() *string
	AttrEndpoint() *string
	AttrId() *string
	BackupId() *string
	SetBackupId(val *string)
	BackupRetentionCount() *float64
	SetBackupRetentionCount(val *float64)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CustomCertificate() *string
	SetCustomCertificate(val *string)
	CustomDomain() *string
	SetCustomDomain(val *string)
	CustomPrivateKey() *string
	SetCustomPrivateKey(val *string)
	DisableAutomatedBackup() interface{}
	SetDisableAutomatedBackup(val interface{})
	Engine() *string
	SetEngine(val *string)
	EngineAttributes() interface{}
	SetEngineAttributes(val interface{})
	EngineModel() *string
	SetEngineModel(val *string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceType() *string
	SetInstanceType(val *string)
	KeyPair() *string
	SetKeyPair(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	Ref() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	ServerName() *string
	SetServerName(val *string)
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	Stack() awscdk.Stack
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnServer
type jsiiProxy_CfnServer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnServer) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) BackupRetentionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CustomCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CustomPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) DisableAutomatedBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) EngineAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"engineAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) EngineModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorksCM::Server`.
func NewCfnServer(scope awscdk.Construct, id *string, props *CfnServerProps) CfnServer {
	_init_.Initialize()

	j := jsiiProxy_CfnServer{}

	_jsii_.Create(
		"monocdk.aws_opsworkscm.CfnServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorksCM::Server`.
func NewCfnServer_Override(c CfnServer, scope awscdk.Construct, id *string, props *CfnServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_opsworkscm.CfnServer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServer) SetAssociatePublicIpAddress(val interface{}) {
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetBackupRetentionCount(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionCount",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetCustomCertificate(val *string) {
	_jsii_.Set(
		j,
		"customCertificate",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetCustomDomain(val *string) {
	_jsii_.Set(
		j,
		"customDomain",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetCustomPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"customPrivateKey",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetDisableAutomatedBackup(val interface{}) {
	_jsii_.Set(
		j,
		"disableAutomatedBackup",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEngineAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"engineAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEngineModel(val *string) {
	_jsii_.Set(
		j,
		"engineModel",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetInstanceProfileArn(val *string) {
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetKeyPair(val *string) {
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
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
func CfnServer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworkscm.CfnServer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnServer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworkscm.CfnServer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworkscm.CfnServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_opsworkscm.CfnServer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnServer) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnServer) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnServer) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnServer) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnServer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnServer) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnServer) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnServer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnServer) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnServer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnServer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnServer) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnServer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnServer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnServer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnServer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnServer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnServer_EngineAttributeProperty struct {
	// `CfnServer.EngineAttributeProperty.Name`.
	Name *string `json:"name"`
	// `CfnServer.EngineAttributeProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::OpsWorksCM::Server`.
type CfnServerProps struct {
	// `AWS::OpsWorksCM::Server.InstanceProfileArn`.
	InstanceProfileArn *string `json:"instanceProfileArn"`
	// `AWS::OpsWorksCM::Server.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `AWS::OpsWorksCM::Server.ServiceRoleArn`.
	ServiceRoleArn *string `json:"serviceRoleArn"`
	// `AWS::OpsWorksCM::Server.AssociatePublicIpAddress`.
	AssociatePublicIpAddress interface{} `json:"associatePublicIpAddress"`
	// `AWS::OpsWorksCM::Server.BackupId`.
	BackupId *string `json:"backupId"`
	// `AWS::OpsWorksCM::Server.BackupRetentionCount`.
	BackupRetentionCount *float64 `json:"backupRetentionCount"`
	// `AWS::OpsWorksCM::Server.CustomCertificate`.
	CustomCertificate *string `json:"customCertificate"`
	// `AWS::OpsWorksCM::Server.CustomDomain`.
	CustomDomain *string `json:"customDomain"`
	// `AWS::OpsWorksCM::Server.CustomPrivateKey`.
	CustomPrivateKey *string `json:"customPrivateKey"`
	// `AWS::OpsWorksCM::Server.DisableAutomatedBackup`.
	DisableAutomatedBackup interface{} `json:"disableAutomatedBackup"`
	// `AWS::OpsWorksCM::Server.Engine`.
	Engine *string `json:"engine"`
	// `AWS::OpsWorksCM::Server.EngineAttributes`.
	EngineAttributes interface{} `json:"engineAttributes"`
	// `AWS::OpsWorksCM::Server.EngineModel`.
	EngineModel *string `json:"engineModel"`
	// `AWS::OpsWorksCM::Server.EngineVersion`.
	EngineVersion *string `json:"engineVersion"`
	// `AWS::OpsWorksCM::Server.KeyPair`.
	KeyPair *string `json:"keyPair"`
	// `AWS::OpsWorksCM::Server.PreferredBackupWindow`.
	PreferredBackupWindow *string `json:"preferredBackupWindow"`
	// `AWS::OpsWorksCM::Server.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// `AWS::OpsWorksCM::Server.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `AWS::OpsWorksCM::Server.ServerName`.
	ServerName *string `json:"serverName"`
	// `AWS::OpsWorksCM::Server.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::OpsWorksCM::Server.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

