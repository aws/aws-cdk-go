package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoT::AccountAuditConfiguration`.
//
// Use the `AWS::IoT::AccountAuditConfiguration` resource to configure or reconfigure the Device Defender audit settings for your account. Settings include how audit notifications are sent and which audit checks are enabled or disabled. For API reference, see [UpdateAccountAuditConfiguration](https://docs.aws.amazon.com/iot/latest/apireference/API_UpdateAccountAuditConfiguration.html) and for detailed information on all available audit checks, see [Audit checks](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-audit-checks.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnAccountAuditConfiguration := iot.NewCfnAccountAuditConfiguration(this, jsii.String("MyCfnAccountAuditConfiguration"), &cfnAccountAuditConfigurationProps{
//   	accountId: jsii.String("accountId"),
//   	auditCheckConfigurations: &auditCheckConfigurationsProperty{
//   		authenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		conflictingClientIdsCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateSharedCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotPolicyOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasAllowsAccessToUnusedServicesCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		loggingDisabledCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedCaCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedDeviceCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	auditNotificationTargetConfigurations: &auditNotificationTargetConfigurationsProperty{
//   		sns: &auditNotificationTargetProperty{
//   			enabled: jsii.Boolean(false),
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//   		},
//   	},
//   })
//
type CfnAccountAuditConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the account.
	//
	// You can use the expression `!Sub "${AWS::AccountId}"` to use your account ID.
	AccountId() *string
	SetAccountId(val *string)
	// Specifies which audit checks are enabled and disabled for this account.
	//
	// Some data collection might start immediately when certain checks are enabled. When a check is disabled, any data collected so far in relation to the check is deleted. To disable a check, set the value of the `Enabled:` key to `false` .
	//
	// If an enabled check is removed from the template, it will also be disabled.
	//
	// You can't disable a check if it's used by any scheduled audit. You must delete the check from the scheduled audit or delete the scheduled audit itself to disable the check.
	//
	// For more information on avialbe auidt checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)
	AuditCheckConfigurations() interface{}
	SetAuditCheckConfigurations(val interface{})
	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations() interface{}
	SetAuditNotificationTargetConfigurations(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the role that grants permission to AWS IoT to access information about your devices, policies, certificates, and other items as required when performing an audit.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnAccountAuditConfiguration
type jsiiProxy_CfnAccountAuditConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) AuditCheckConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditCheckConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) AuditNotificationTargetConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditNotificationTargetConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::AccountAuditConfiguration`.
func NewCfnAccountAuditConfiguration(scope constructs.Construct, id *string, props *CfnAccountAuditConfigurationProps) CfnAccountAuditConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnAccountAuditConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::AccountAuditConfiguration`.
func NewCfnAccountAuditConfiguration_Override(c CfnAccountAuditConfiguration, scope constructs.Construct, id *string, props *CfnAccountAuditConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) SetAuditCheckConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"auditCheckConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) SetAuditNotificationTargetConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"auditNotificationTargetConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnAccountAuditConfiguration) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAccountAuditConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAccountAuditConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAccountAuditConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccountAuditConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnAccountAuditConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Which audit checks are enabled and disabled for this account.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   auditCheckConfigurationProperty := &auditCheckConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnAccountAuditConfiguration_AuditCheckConfigurationProperty struct {
	// True if this audit check is enabled for this account.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// The types of audit checks that can be performed.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   auditCheckConfigurationsProperty := &auditCheckConfigurationsProperty{
//   	authenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	caCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	caCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	conflictingClientIdsCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	deviceCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	deviceCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	deviceCertificateSharedCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	iotPolicyOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	iotRoleAliasAllowsAccessToUnusedServicesCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	iotRoleAliasOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	loggingDisabledCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	revokedCaCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	revokedDeviceCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	unauthenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnAccountAuditConfiguration_AuditCheckConfigurationsProperty struct {
	// Checks the permissiveness of an authenticated Amazon Cognito identity pool role.
	//
	// For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker during the 31 days before the audit is performed.
	AuthenticatedCognitoRoleOverlyPermissiveCheck interface{} `json:"authenticatedCognitoRoleOverlyPermissiveCheck" yaml:"authenticatedCognitoRoleOverlyPermissiveCheck"`
	// Checks if a CA certificate is expiring.
	//
	// This check applies to CA certificates expiring within 30 days or that have expired.
	CaCertificateExpiringCheck interface{} `json:"caCertificateExpiringCheck" yaml:"caCertificateExpiringCheck"`
	// Checks the quality of the CA certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size. This check applies to CA certificates that are `ACTIVE` or `PENDING_TRANSFER` .
	CaCertificateKeyQualityCheck interface{} `json:"caCertificateKeyQualityCheck" yaml:"caCertificateKeyQualityCheck"`
	// Checks if multiple devices connect using the same client ID.
	ConflictingClientIdsCheck interface{} `json:"conflictingClientIdsCheck" yaml:"conflictingClientIdsCheck"`
	// Checks if a device certificate is expiring.
	//
	// This check applies to device certificates expiring within 30 days or that have expired.
	DeviceCertificateExpiringCheck interface{} `json:"deviceCertificateExpiringCheck" yaml:"deviceCertificateExpiringCheck"`
	// Checks the quality of the device certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority, and if the key meets a minimum required size.
	DeviceCertificateKeyQualityCheck interface{} `json:"deviceCertificateKeyQualityCheck" yaml:"deviceCertificateKeyQualityCheck"`
	// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT .
	DeviceCertificateSharedCheck interface{} `json:"deviceCertificateSharedCheck" yaml:"deviceCertificateSharedCheck"`
	// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
	IotPolicyOverlyPermissiveCheck interface{} `json:"iotPolicyOverlyPermissiveCheck" yaml:"iotPolicyOverlyPermissiveCheck"`
	// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
	IotRoleAliasAllowsAccessToUnusedServicesCheck interface{} `json:"iotRoleAliasAllowsAccessToUnusedServicesCheck" yaml:"iotRoleAliasAllowsAccessToUnusedServicesCheck"`
	// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
	IotRoleAliasOverlyPermissiveCheck interface{} `json:"iotRoleAliasOverlyPermissiveCheck" yaml:"iotRoleAliasOverlyPermissiveCheck"`
	// Checks if AWS IoT logs are disabled.
	LoggingDisabledCheck interface{} `json:"loggingDisabledCheck" yaml:"loggingDisabledCheck"`
	// Checks if a revoked CA certificate is still active.
	RevokedCaCertificateStillActiveCheck interface{} `json:"revokedCaCertificateStillActiveCheck" yaml:"revokedCaCertificateStillActiveCheck"`
	// Checks if a revoked device certificate is still active.
	RevokedDeviceCertificateStillActiveCheck interface{} `json:"revokedDeviceCertificateStillActiveCheck" yaml:"revokedDeviceCertificateStillActiveCheck"`
	// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
	UnauthenticatedCognitoRoleOverlyPermissiveCheck interface{} `json:"unauthenticatedCognitoRoleOverlyPermissiveCheck" yaml:"unauthenticatedCognitoRoleOverlyPermissiveCheck"`
}

// The configuration of the audit notification target.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   auditNotificationTargetConfigurationsProperty := &auditNotificationTargetConfigurationsProperty{
//   	sns: &auditNotificationTargetProperty{
//   		enabled: jsii.Boolean(false),
//   		roleArn: jsii.String("roleArn"),
//   		targetArn: jsii.String("targetArn"),
//   	},
//   }
//
type CfnAccountAuditConfiguration_AuditNotificationTargetConfigurationsProperty struct {
	// The `Sns` notification target.
	Sns interface{} `json:"sns" yaml:"sns"`
}

// Information about the targets to which audit notifications are sent.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   auditNotificationTargetProperty := &auditNotificationTargetProperty{
//   	enabled: jsii.Boolean(false),
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnAccountAuditConfiguration_AuditNotificationTargetProperty struct {
	// True if notifications to the target are enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The ARN of the role that grants permission to send notifications to the target.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The ARN of the target (SNS topic) to which audit notifications are sent.
	TargetArn *string `json:"targetArn" yaml:"targetArn"`
}

// Properties for defining a `CfnAccountAuditConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnAccountAuditConfigurationProps := &cfnAccountAuditConfigurationProps{
//   	accountId: jsii.String("accountId"),
//   	auditCheckConfigurations: &auditCheckConfigurationsProperty{
//   		authenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		caCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		conflictingClientIdsCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateExpiringCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateKeyQualityCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		deviceCertificateSharedCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotPolicyOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasAllowsAccessToUnusedServicesCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		iotRoleAliasOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		loggingDisabledCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedCaCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		revokedDeviceCertificateStillActiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		unauthenticatedCognitoRoleOverlyPermissiveCheck: &auditCheckConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	auditNotificationTargetConfigurations: &auditNotificationTargetConfigurationsProperty{
//   		sns: &auditNotificationTargetProperty{
//   			enabled: jsii.Boolean(false),
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//   		},
//   	},
//   }
//
type CfnAccountAuditConfigurationProps struct {
	// The ID of the account.
	//
	// You can use the expression `!Sub "${AWS::AccountId}"` to use your account ID.
	AccountId *string `json:"accountId" yaml:"accountId"`
	// Specifies which audit checks are enabled and disabled for this account.
	//
	// Some data collection might start immediately when certain checks are enabled. When a check is disabled, any data collected so far in relation to the check is deleted. To disable a check, set the value of the `Enabled:` key to `false` .
	//
	// If an enabled check is removed from the template, it will also be disabled.
	//
	// You can't disable a check if it's used by any scheduled audit. You must delete the check from the scheduled audit or delete the scheduled audit itself to disable the check.
	//
	// For more information on avialbe auidt checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)
	AuditCheckConfigurations interface{} `json:"auditCheckConfigurations" yaml:"auditCheckConfigurations"`
	// The Amazon Resource Name (ARN) of the role that grants permission to AWS IoT to access information about your devices, policies, certificates, and other items as required when performing an audit.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations interface{} `json:"auditNotificationTargetConfigurations" yaml:"auditNotificationTargetConfigurations"`
}

// A CloudFormation `AWS::IoT::Authorizer`.
//
// Specifies an authorizer.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnAuthorizer := iot.NewCfnAuthorizer(this, jsii.String("MyCfnAuthorizer"), &cfnAuthorizerProps{
//   	authorizerFunctionArn: jsii.String("authorizerFunctionArn"),
//
//   	// the properties below are optional
//   	authorizerName: jsii.String("authorizerName"),
//   	enableCachingForHttp: jsii.Boolean(false),
//   	signingDisabled: jsii.Boolean(false),
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tokenKeyName: jsii.String("tokenKeyName"),
//   	tokenSigningPublicKeys: map[string]*string{
//   		"tokenSigningPublicKeysKey": jsii.String("tokenSigningPublicKeys"),
//   	},
//   })
//
type CfnAuthorizer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the authorizer.
	AttrArn() *string
	// The authorizer's Lambda function ARN.
	AuthorizerFunctionArn() *string
	SetAuthorizerFunctionArn(val *string)
	// The authorizer name.
	AuthorizerName() *string
	SetAuthorizerName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::IoT::Authorizer.EnableCachingForHttp`.
	EnableCachingForHttp() interface{}
	SetEnableCachingForHttp(val interface{})
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies whether AWS IoT validates the token signature in an authorization request.
	SigningDisabled() interface{}
	SetSigningDisabled(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the authorizer.
	//
	// Valid values: `ACTIVE` | `INACTIVE`.
	Status() *string
	SetStatus(val *string)
	// Metadata which can be used to manage the custom authorizer.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags() awscdk.TagManager
	// The key used to extract the token from the HTTP headers.
	TokenKeyName() *string
	SetTokenKeyName(val *string)
	// The public keys used to validate the token signature returned by your custom authentication service.
	TokenSigningPublicKeys() interface{}
	SetTokenSigningPublicKeys(val interface{})
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnAuthorizer
type jsiiProxy_CfnAuthorizer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAuthorizer) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) EnableCachingForHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCachingForHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) SigningDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signingDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) TokenKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) TokenSigningPublicKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenSigningPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Authorizer`.
func NewCfnAuthorizer(scope constructs.Construct, id *string, props *CfnAuthorizerProps) CfnAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_CfnAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Authorizer`.
func NewCfnAuthorizer_Override(c CfnAuthorizer, scope constructs.Construct, id *string, props *CfnAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerFunctionArn(val *string) {
	_jsii_.Set(
		j,
		"authorizerFunctionArn",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerName(val *string) {
	_jsii_.Set(
		j,
		"authorizerName",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetEnableCachingForHttp(val interface{}) {
	_jsii_.Set(
		j,
		"enableCachingForHttp",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetSigningDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"signingDisabled",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetTokenKeyName(val *string) {
	_jsii_.Set(
		j,
		"tokenKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetTokenSigningPublicKeys(val interface{}) {
	_jsii_.Set(
		j,
		"tokenSigningPublicKeys",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAuthorizer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAuthorizer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthorizer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnAuthorizer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAuthorizer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAuthorizer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAuthorizer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAuthorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAuthorizer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAuthorizer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnAuthorizerProps := &cfnAuthorizerProps{
//   	authorizerFunctionArn: jsii.String("authorizerFunctionArn"),
//
//   	// the properties below are optional
//   	authorizerName: jsii.String("authorizerName"),
//   	enableCachingForHttp: jsii.Boolean(false),
//   	signingDisabled: jsii.Boolean(false),
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tokenKeyName: jsii.String("tokenKeyName"),
//   	tokenSigningPublicKeys: map[string]*string{
//   		"tokenSigningPublicKeysKey": jsii.String("tokenSigningPublicKeys"),
//   	},
//   }
//
type CfnAuthorizerProps struct {
	// The authorizer's Lambda function ARN.
	AuthorizerFunctionArn *string `json:"authorizerFunctionArn" yaml:"authorizerFunctionArn"`
	// The authorizer name.
	AuthorizerName *string `json:"authorizerName" yaml:"authorizerName"`
	// `AWS::IoT::Authorizer.EnableCachingForHttp`.
	EnableCachingForHttp interface{} `json:"enableCachingForHttp" yaml:"enableCachingForHttp"`
	// Specifies whether AWS IoT validates the token signature in an authorization request.
	SigningDisabled interface{} `json:"signingDisabled" yaml:"signingDisabled"`
	// The status of the authorizer.
	//
	// Valid values: `ACTIVE` | `INACTIVE`.
	Status *string `json:"status" yaml:"status"`
	// Metadata which can be used to manage the custom authorizer.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The key used to extract the token from the HTTP headers.
	TokenKeyName *string `json:"tokenKeyName" yaml:"tokenKeyName"`
	// The public keys used to validate the token signature returned by your custom authentication service.
	TokenSigningPublicKeys interface{} `json:"tokenSigningPublicKeys" yaml:"tokenSigningPublicKeys"`
}

// A CloudFormation `AWS::IoT::Certificate`.
//
// Use the `AWS::IoT::Certificate` resource to declare an AWS IoT X.509 certificate. For information about working with X.509 certificates, see [X.509 Client Certificates](https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnCertificate := iot.NewCfnCertificate(this, jsii.String("MyCfnCertificate"), &cfnCertificateProps{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	caCertificatePem: jsii.String("caCertificatePem"),
//   	certificateMode: jsii.String("certificateMode"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   })
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) for the instance profile. For example:.
	//
	// `{ "Fn::GetAtt": ["MyCertificate", "Arn"] }`
	//
	// A value similar to the following is returned:
	//
	// `arn:aws:iot:ap-southeast-2:123456789012:cert/a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2`.
	AttrArn() *string
	// The certificate ID.
	AttrId() *string
	// The CA certificate used to sign the device certificate being registered, not available when CertificateMode is SNI_ONLY.
	CaCertificatePem() *string
	SetCaCertificatePem(val *string)
	// Specifies which mode of certificate registration to use with this resource.
	//
	// Valid options are DEFAULT with CaCertificatePem and CertificatePem, SNI_ONLY with CertificatePem, and Default with CertificateSigningRequest.
	CertificateMode() *string
	SetCertificateMode(val *string)
	// The certificate data in PEM format.
	//
	// Requires SNI_ONLY for the certificate mode or the accompanying CACertificatePem for registration.
	CertificatePem() *string
	SetCertificatePem(val *string)
	// The certificate signing request (CSR).
	CertificateSigningRequest() *string
	SetCertificateSigningRequest(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the certificate.
	//
	// Valid values are ACTIVE, INACTIVE, REVOKED, PENDING_TRANSFER, and PENDING_ACTIVATION.
	//
	// The status value REGISTER_INACTIVE is deprecated and should not be used.
	Status() *string
	SetStatus(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnCertificate
type jsiiProxy_CfnCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CaCertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Certificate`.
func NewCfnCertificate(scope constructs.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope constructs.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCaCertificatePem(val *string) {
	_jsii_.Set(
		j,
		"caCertificatePem",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateMode(val *string) {
	_jsii_.Set(
		j,
		"certificateMode",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificatePem(val *string) {
	_jsii_.Set(
		j,
		"certificatePem",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateSigningRequest(val *string) {
	_jsii_.Set(
		j,
		"certificateSigningRequest",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCertificate`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnCertificateProps := &cfnCertificateProps{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	caCertificatePem: jsii.String("caCertificatePem"),
//   	certificateMode: jsii.String("certificateMode"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   }
//
type CfnCertificateProps struct {
	// The status of the certificate.
	//
	// Valid values are ACTIVE, INACTIVE, REVOKED, PENDING_TRANSFER, and PENDING_ACTIVATION.
	//
	// The status value REGISTER_INACTIVE is deprecated and should not be used.
	Status *string `json:"status" yaml:"status"`
	// The CA certificate used to sign the device certificate being registered, not available when CertificateMode is SNI_ONLY.
	CaCertificatePem *string `json:"caCertificatePem" yaml:"caCertificatePem"`
	// Specifies which mode of certificate registration to use with this resource.
	//
	// Valid options are DEFAULT with CaCertificatePem and CertificatePem, SNI_ONLY with CertificatePem, and Default with CertificateSigningRequest.
	CertificateMode *string `json:"certificateMode" yaml:"certificateMode"`
	// The certificate data in PEM format.
	//
	// Requires SNI_ONLY for the certificate mode or the accompanying CACertificatePem for registration.
	CertificatePem *string `json:"certificatePem" yaml:"certificatePem"`
	// The certificate signing request (CSR).
	CertificateSigningRequest *string `json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
}

// A CloudFormation `AWS::IoT::CustomMetric`.
//
// Use the `AWS::IoT::CustomMetric` resource to define a custom metric published by your devices to Device Defender. For API reference, see [CreateCustomMetric](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCustomMetric.html) and for general information, see [Custom metrics](https://docs.aws.amazon.com/iot/latest/developerguide/dd-detect-custom-metrics.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnCustomMetric := iot.NewCfnCustomMetric(this, jsii.String("MyCfnCustomMetric"), &cfnCustomMetricProps{
//   	metricType: jsii.String("metricType"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   	metricName: jsii.String("metricName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCustomMetric interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Number (ARN) of the custom metric;
	//
	// for example, `arn: *aws-partition* :iot: *region* : *accountId* :custommetric/ *metricName*` .
	AttrMetricArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The friendly name in the console for the custom metric.
	//
	// This name doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. You can update the friendly name after you define it.
	DisplayName() *string
	SetDisplayName(val *string)
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
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. The name can't begin with `aws:` . You can’t change the name after you define it.
	MetricName() *string
	SetMetricName(val *string)
	// The type of the custom metric. Types include `string-list` , `ip-address-list` , `number-list` , and `number` .
	//
	// > The type `number` only takes a single metric value as an input, but when you submit the metrics value in the DeviceMetrics report, you must pass it as an array with a single value.
	MetricType() *string
	SetMetricType(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the custom metric.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnCustomMetric
type jsiiProxy_CfnCustomMetric struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCustomMetric) AttrMetricArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMetricArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) MetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetric) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::CustomMetric`.
func NewCfnCustomMetric(scope constructs.Construct, id *string, props *CfnCustomMetricProps) CfnCustomMetric {
	_init_.Initialize()

	j := jsiiProxy_CfnCustomMetric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::CustomMetric`.
func NewCfnCustomMetric_Override(c CfnCustomMetric, scope constructs.Construct, id *string, props *CfnCustomMetricProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomMetric) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnCustomMetric) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnCustomMetric) SetMetricType(val *string) {
	_jsii_.Set(
		j,
		"metricType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCustomMetric_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCustomMetric_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCustomMetric_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomMetric_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnCustomMetric",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomMetric) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomMetric) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomMetric) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomMetric) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomMetric) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomMetric) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomMetric) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCustomMetric) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomMetric) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomMetric) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomMetric) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCustomMetric) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomMetric) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomMetric) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomMetric) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCustomMetric`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnCustomMetricProps := &cfnCustomMetricProps{
//   	metricType: jsii.String("metricType"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   	metricName: jsii.String("metricName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomMetricProps struct {
	// The type of the custom metric. Types include `string-list` , `ip-address-list` , `number-list` , and `number` .
	//
	// > The type `number` only takes a single metric value as an input, but when you submit the metrics value in the DeviceMetrics report, you must pass it as an array with a single value.
	MetricType *string `json:"metricType" yaml:"metricType"`
	// The friendly name in the console for the custom metric.
	//
	// This name doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. You can update the friendly name after you define it.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. The name can't begin with `aws:` . You can’t change the name after you define it.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// Metadata that can be used to manage the custom metric.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoT::Dimension`.
//
// Use the `AWS::IoT::Dimension` to limit the scope of a metric used in a security profile for AWS IoT Device Defender . For example, using a `TOPIC_FILTER` dimension, you can narrow down the scope of the metric to only MQTT topics where the name matches the pattern specified in the dimension. For API reference, see [CreateDimension](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateDimension.html) and for general information, see [Scoping metrics in security profiles using dimensions](https://docs.aws.amazon.com/iot/latest/developerguide/scoping-security-behavior.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnDimension := iot.NewCfnDimension(this, jsii.String("MyCfnDimension"), &cfnDimensionProps{
//   	stringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDimension interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the dimension.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// A unique identifier for the dimension.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies the value or list of values for the dimension.
	//
	// For `TOPIC_FILTER` dimensions, this is a pattern used to match the MQTT topic (for example, "admin/#").
	StringValues() *[]*string
	SetStringValues(val *[]*string)
	// Metadata that can be used to manage the dimension.
	Tags() awscdk.TagManager
	// Specifies the type of dimension.
	//
	// Supported types: `TOPIC_FILTER.`
	Type() *string
	SetType(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnDimension
type jsiiProxy_CfnDimension struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDimension) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) StringValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDimension) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Dimension`.
func NewCfnDimension(scope constructs.Construct, id *string, props *CfnDimensionProps) CfnDimension {
	_init_.Initialize()

	j := jsiiProxy_CfnDimension{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnDimension",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Dimension`.
func NewCfnDimension_Override(c CfnDimension, scope constructs.Construct, id *string, props *CfnDimensionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnDimension",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDimension) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDimension) SetStringValues(val *[]*string) {
	_jsii_.Set(
		j,
		"stringValues",
		val,
	)
}

func (j *jsiiProxy_CfnDimension) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDimension_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDimension",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDimension_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDimension",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDimension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDimension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDimension_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnDimension",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDimension) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDimension) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDimension) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDimension) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDimension) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDimension) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDimension) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDimension) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDimension) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDimension) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDimension) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDimension) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDimension) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDimension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDimension) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDimension`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnDimensionProps := &cfnDimensionProps{
//   	stringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDimensionProps struct {
	// Specifies the value or list of values for the dimension.
	//
	// For `TOPIC_FILTER` dimensions, this is a pattern used to match the MQTT topic (for example, "admin/#").
	StringValues *[]*string `json:"stringValues" yaml:"stringValues"`
	// Specifies the type of dimension.
	//
	// Supported types: `TOPIC_FILTER.`
	Type *string `json:"type" yaml:"type"`
	// A unique identifier for the dimension.
	Name *string `json:"name" yaml:"name"`
	// Metadata that can be used to manage the dimension.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoT::DomainConfiguration`.
//
// Specifies a domain configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnDomainConfiguration := iot.NewCfnDomainConfiguration(this, jsii.String("MyCfnDomainConfiguration"), &cfnDomainConfigurationProps{
//   	authorizerConfig: &authorizerConfigProperty{
//   		allowAuthorizerOverride: jsii.Boolean(false),
//   		defaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   	},
//   	domainConfigurationName: jsii.String("domainConfigurationName"),
//   	domainConfigurationStatus: jsii.String("domainConfigurationStatus"),
//   	domainName: jsii.String("domainName"),
//   	serverCertificateArns: []*string{
//   		jsii.String("serverCertificateArns"),
//   	},
//   	serviceType: jsii.String("serviceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	validationCertificateArn: jsii.String("validationCertificateArn"),
//   })
//
type CfnDomainConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the domain configuration.
	AttrArn() *string
	// The type of service delivered by the domain.
	AttrDomainType() *string
	// The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake.
	//
	// Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
	AttrServerCertificates() awscdk.IResolvable
	// An object that specifies the authorization service for a domain.
	AuthorizerConfig() interface{}
	SetAuthorizerConfig(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the domain configuration.
	//
	// This value must be unique to a region.
	DomainConfigurationName() *string
	SetDomainConfigurationName(val *string)
	// The status to which the domain configuration should be updated.
	//
	// Valid values: `ENABLED` | `DISABLED`.
	DomainConfigurationStatus() *string
	SetDomainConfigurationStatus(val *string)
	// The name of the domain.
	DomainName() *string
	SetDomainName(val *string)
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake.
	//
	// Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
	ServerCertificateArns() *[]*string
	SetServerCertificateArns(val *[]*string)
	// The type of service delivered by the endpoint.
	//
	// > AWS IoT Core currently supports only the `DATA` service type.
	ServiceType() *string
	SetServiceType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the domain configuration.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The certificate used to validate the server certificate and prove domain name ownership.
	//
	// This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
	ValidationCertificateArn() *string
	SetValidationCertificateArn(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnDomainConfiguration
type jsiiProxy_CfnDomainConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomainConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) AttrDomainType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) AttrServerCertificates() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrServerCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) AuthorizerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) DomainConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) DomainConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) ServerCertificateArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainConfiguration) ValidationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationCertificateArn",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::DomainConfiguration`.
func NewCfnDomainConfiguration(scope constructs.Construct, id *string, props *CfnDomainConfigurationProps) CfnDomainConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnDomainConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::DomainConfiguration`.
func NewCfnDomainConfiguration_Override(c CfnDomainConfiguration, scope constructs.Construct, id *string, props *CfnDomainConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetAuthorizerConfig(val interface{}) {
	_jsii_.Set(
		j,
		"authorizerConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetDomainConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"domainConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetDomainConfigurationStatus(val *string) {
	_jsii_.Set(
		j,
		"domainConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetServerCertificateArns(val *[]*string) {
	_jsii_.Set(
		j,
		"serverCertificateArns",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetServiceType(val *string) {
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_CfnDomainConfiguration) SetValidationCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"validationCertificateArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomainConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomainConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDomainConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnDomainConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomainConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that specifies the authorization service for a domain.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   authorizerConfigProperty := &authorizerConfigProperty{
//   	allowAuthorizerOverride: jsii.Boolean(false),
//   	defaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   }
//
type CfnDomainConfiguration_AuthorizerConfigProperty struct {
	// A Boolean that specifies whether the domain configuration's authorization service can be overridden.
	AllowAuthorizerOverride interface{} `json:"allowAuthorizerOverride" yaml:"allowAuthorizerOverride"`
	// The name of the authorization service for a domain configuration.
	DefaultAuthorizerName *string `json:"defaultAuthorizerName" yaml:"defaultAuthorizerName"`
}

// An object that contains information about a server certificate.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   serverCertificateSummaryProperty := &serverCertificateSummaryProperty{
//   	serverCertificateArn: jsii.String("serverCertificateArn"),
//   	serverCertificateStatus: jsii.String("serverCertificateStatus"),
//   	serverCertificateStatusDetail: jsii.String("serverCertificateStatusDetail"),
//   }
//
type CfnDomainConfiguration_ServerCertificateSummaryProperty struct {
	// The ARN of the server certificate.
	ServerCertificateArn *string `json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The status of the server certificate.
	ServerCertificateStatus *string `json:"serverCertificateStatus" yaml:"serverCertificateStatus"`
	// Details that explain the status of the server certificate.
	ServerCertificateStatusDetail *string `json:"serverCertificateStatusDetail" yaml:"serverCertificateStatusDetail"`
}

// Properties for defining a `CfnDomainConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnDomainConfigurationProps := &cfnDomainConfigurationProps{
//   	authorizerConfig: &authorizerConfigProperty{
//   		allowAuthorizerOverride: jsii.Boolean(false),
//   		defaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   	},
//   	domainConfigurationName: jsii.String("domainConfigurationName"),
//   	domainConfigurationStatus: jsii.String("domainConfigurationStatus"),
//   	domainName: jsii.String("domainName"),
//   	serverCertificateArns: []*string{
//   		jsii.String("serverCertificateArns"),
//   	},
//   	serviceType: jsii.String("serviceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	validationCertificateArn: jsii.String("validationCertificateArn"),
//   }
//
type CfnDomainConfigurationProps struct {
	// An object that specifies the authorization service for a domain.
	AuthorizerConfig interface{} `json:"authorizerConfig" yaml:"authorizerConfig"`
	// The name of the domain configuration.
	//
	// This value must be unique to a region.
	DomainConfigurationName *string `json:"domainConfigurationName" yaml:"domainConfigurationName"`
	// The status to which the domain configuration should be updated.
	//
	// Valid values: `ENABLED` | `DISABLED`.
	DomainConfigurationStatus *string `json:"domainConfigurationStatus" yaml:"domainConfigurationStatus"`
	// The name of the domain.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake.
	//
	// Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
	ServerCertificateArns *[]*string `json:"serverCertificateArns" yaml:"serverCertificateArns"`
	// The type of service delivered by the endpoint.
	//
	// > AWS IoT Core currently supports only the `DATA` service type.
	ServiceType *string `json:"serviceType" yaml:"serviceType"`
	// Metadata which can be used to manage the domain configuration.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The certificate used to validate the server certificate and prove domain name ownership.
	//
	// This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
	ValidationCertificateArn *string `json:"validationCertificateArn" yaml:"validationCertificateArn"`
}

// A CloudFormation `AWS::IoT::FleetMetric`.
//
// Use the `AWS::IoT::FleetMetric` resource to declare a fleet metric.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnFleetMetric := iot.NewCfnFleetMetric(this, jsii.String("MyCfnFleetMetric"), &cfnFleetMetricProps{
//   	metricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	aggregationField: jsii.String("aggregationField"),
//   	aggregationType: &aggregationTypeProperty{
//   		name: jsii.String("name"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	indexName: jsii.String("indexName"),
//   	period: jsii.Number(123),
//   	queryString: jsii.String("queryString"),
//   	queryVersion: jsii.String("queryVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   })
//
type CfnFleetMetric interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The field to aggregate.
	AggregationField() *string
	SetAggregationField(val *string)
	// The type of the aggregation query.
	AggregationType() interface{}
	SetAggregationType(val interface{})
	// The time the fleet metric was created.
	AttrCreationDate() awscdk.IResolvable
	// The time the fleet metric was last modified.
	AttrLastModifiedDate() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the fleet metric.
	AttrMetricArn() *string
	// The fleet metric version.
	AttrVersion() awscdk.IResolvable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The fleet metric description.
	Description() *string
	SetDescription(val *string)
	// The name of the index to search.
	IndexName() *string
	SetIndexName(val *string)
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
	// The name of the fleet metric to create.
	MetricName() *string
	SetMetricName(val *string)
	// The tree node.
	Node() constructs.Node
	// The time in seconds between fleet metric emissions.
	//
	// Range [60(1 min), 86400(1 day)] and must be multiple of 60.
	Period() *float64
	SetPeriod(val *float64)
	// The search query string.
	QueryString() *string
	SetQueryString(val *string)
	// The query version.
	QueryVersion() *string
	SetQueryVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the fleet metric.
	Tags() awscdk.TagManager
	// Used to support unit transformation such as milliseconds to seconds.
	//
	// Must be a unit supported by CW metric. Default to null.
	Unit() *string
	SetUnit(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnFleetMetric
type jsiiProxy_CfnFleetMetric struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFleetMetric) AggregationField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) AggregationType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) AttrCreationDate() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) AttrLastModifiedDate() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) AttrMetricArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMetricArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) AttrVersion() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) QueryVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetMetric) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::FleetMetric`.
func NewCfnFleetMetric(scope constructs.Construct, id *string, props *CfnFleetMetricProps) CfnFleetMetric {
	_init_.Initialize()

	j := jsiiProxy_CfnFleetMetric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::FleetMetric`.
func NewCfnFleetMetric_Override(c CfnFleetMetric, scope constructs.Construct, id *string, props *CfnFleetMetricProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetAggregationField(val *string) {
	_jsii_.Set(
		j,
		"aggregationField",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetAggregationType(val interface{}) {
	_jsii_.Set(
		j,
		"aggregationType",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetIndexName(val *string) {
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetPeriod(val *float64) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetQueryString(val *string) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetQueryVersion(val *string) {
	_jsii_.Set(
		j,
		"queryVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFleetMetric) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFleetMetric_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFleetMetric_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnFleetMetric_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFleetMetric_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnFleetMetric",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFleetMetric) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFleetMetric) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFleetMetric) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFleetMetric) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFleetMetric) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFleetMetric) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFleetMetric) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFleetMetric) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleetMetric) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleetMetric) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFleetMetric) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFleetMetric) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleetMetric) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleetMetric) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleetMetric) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The type of aggregation queries.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   aggregationTypeProperty := &aggregationTypeProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnFleetMetric_AggregationTypeProperty struct {
	// The name of the aggregation type.
	Name *string `json:"name" yaml:"name"`
	// A list of the values of aggregation types.
	Values *[]*string `json:"values" yaml:"values"`
}

// Properties for defining a `CfnFleetMetric`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnFleetMetricProps := &cfnFleetMetricProps{
//   	metricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	aggregationField: jsii.String("aggregationField"),
//   	aggregationType: &aggregationTypeProperty{
//   		name: jsii.String("name"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	indexName: jsii.String("indexName"),
//   	period: jsii.Number(123),
//   	queryString: jsii.String("queryString"),
//   	queryVersion: jsii.String("queryVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   }
//
type CfnFleetMetricProps struct {
	// The name of the fleet metric to create.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The field to aggregate.
	AggregationField *string `json:"aggregationField" yaml:"aggregationField"`
	// The type of the aggregation query.
	AggregationType interface{} `json:"aggregationType" yaml:"aggregationType"`
	// The fleet metric description.
	Description *string `json:"description" yaml:"description"`
	// The name of the index to search.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The time in seconds between fleet metric emissions.
	//
	// Range [60(1 min), 86400(1 day)] and must be multiple of 60.
	Period *float64 `json:"period" yaml:"period"`
	// The search query string.
	QueryString *string `json:"queryString" yaml:"queryString"`
	// The query version.
	QueryVersion *string `json:"queryVersion" yaml:"queryVersion"`
	// Metadata which can be used to manage the fleet metric.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// Used to support unit transformation such as milliseconds to seconds.
	//
	// Must be a unit supported by CW metric. Default to null.
	Unit *string `json:"unit" yaml:"unit"`
}

// A CloudFormation `AWS::IoT::JobTemplate`.
//
// Represents a job template.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//
//   var abortConfig interface{}
//   var jobExecutionsRetryConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//   cfnJobTemplate := iot.NewCfnJobTemplate(this, jsii.String("MyCfnJobTemplate"), &cfnJobTemplateProps{
//   	description: jsii.String("description"),
//   	jobTemplateId: jsii.String("jobTemplateId"),
//
//   	// the properties below are optional
//   	abortConfig: abortConfig,
//   	document: jsii.String("document"),
//   	documentSource: jsii.String("documentSource"),
//   	jobArn: jsii.String("jobArn"),
//   	jobExecutionsRetryConfig: jobExecutionsRetryConfig,
//   	jobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	presignedUrlConfig: presignedUrlConfig,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutConfig: timeoutConfig,
//   })
//
type CfnJobTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The criteria that determine when and how a job abort takes place.
	AbortConfig() interface{}
	SetAbortConfig(val interface{})
	// The ARN of the job to use as the basis for the job template.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the job template.
	Description() *string
	SetDescription(val *string)
	// The job document.
	//
	// Required if you don't specify a value for `documentSource` .
	Document() *string
	SetDocument(val *string)
	// An S3 link to the job document to use in the template.
	//
	// Required if you don't specify a value for `document` .
	//
	// > If the job document resides in an S3 bucket, you must use a placeholder link when specifying the document.
	// >
	// > The placeholder link is of the following form:
	// >
	// > `${aws:iot:s3-presigned-url:https://s3.amazonaws.com/ *bucket* / *key* }`
	// >
	// > where *bucket* is your bucket name and *key* is the object in the bucket to which you are linking.
	DocumentSource() *string
	SetDocumentSource(val *string)
	// The ARN of the job to use as the basis for the job template.
	JobArn() *string
	SetJobArn(val *string)
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig() interface{}
	SetJobExecutionsRetryConfig(val interface{})
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig() interface{}
	SetJobExecutionsRolloutConfig(val interface{})
	// A unique identifier for the job template.
	//
	// We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId() *string
	SetJobTemplateId(val *string)
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
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig() interface{}
	SetPresignedUrlConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the job template.
	Tags() awscdk.TagManager
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// A timer is started when the job execution status is set to `IN_PROGRESS` . If the job execution status is not set to another terminal state before the timer expires, it will be automatically set to `TIMED_OUT` .
	TimeoutConfig() interface{}
	SetTimeoutConfig(val interface{})
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnJobTemplate
type jsiiProxy_CfnJobTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobTemplate) AbortConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Document() *string {
	var returns *string
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) DocumentSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) JobExecutionsRetryConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRetryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) JobExecutionsRolloutConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) JobTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) PresignedUrlConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presignedUrlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) TimeoutConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::JobTemplate`.
func NewCfnJobTemplate(scope constructs.Construct, id *string, props *CfnJobTemplateProps) CfnJobTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnJobTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::JobTemplate`.
func NewCfnJobTemplate_Override(c CfnJobTemplate, scope constructs.Construct, id *string, props *CfnJobTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetAbortConfig(val interface{}) {
	_jsii_.Set(
		j,
		"abortConfig",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetDocument(val *string) {
	_jsii_.Set(
		j,
		"document",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetDocumentSource(val *string) {
	_jsii_.Set(
		j,
		"documentSource",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetJobArn(val *string) {
	_jsii_.Set(
		j,
		"jobArn",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetJobExecutionsRetryConfig(val interface{}) {
	_jsii_.Set(
		j,
		"jobExecutionsRetryConfig",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetJobExecutionsRolloutConfig(val interface{}) {
	_jsii_.Set(
		j,
		"jobExecutionsRolloutConfig",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetJobTemplateId(val *string) {
	_jsii_.Set(
		j,
		"jobTemplateId",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetPresignedUrlConfig(val interface{}) {
	_jsii_.Set(
		j,
		"presignedUrlConfig",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate) SetTimeoutConfig(val interface{}) {
	_jsii_.Set(
		j,
		"timeoutConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJobTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJobTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnJobTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnJobTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJobTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJobTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJobTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnJobTemplate`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//
//   var abortConfig interface{}
//   var jobExecutionsRetryConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//   cfnJobTemplateProps := &cfnJobTemplateProps{
//   	description: jsii.String("description"),
//   	jobTemplateId: jsii.String("jobTemplateId"),
//
//   	// the properties below are optional
//   	abortConfig: abortConfig,
//   	document: jsii.String("document"),
//   	documentSource: jsii.String("documentSource"),
//   	jobArn: jsii.String("jobArn"),
//   	jobExecutionsRetryConfig: jobExecutionsRetryConfig,
//   	jobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	presignedUrlConfig: presignedUrlConfig,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutConfig: timeoutConfig,
//   }
//
type CfnJobTemplateProps struct {
	// A description of the job template.
	Description *string `json:"description" yaml:"description"`
	// A unique identifier for the job template.
	//
	// We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId *string `json:"jobTemplateId" yaml:"jobTemplateId"`
	// The criteria that determine when and how a job abort takes place.
	AbortConfig interface{} `json:"abortConfig" yaml:"abortConfig"`
	// The job document.
	//
	// Required if you don't specify a value for `documentSource` .
	Document *string `json:"document" yaml:"document"`
	// An S3 link to the job document to use in the template.
	//
	// Required if you don't specify a value for `document` .
	//
	// > If the job document resides in an S3 bucket, you must use a placeholder link when specifying the document.
	// >
	// > The placeholder link is of the following form:
	// >
	// > `${aws:iot:s3-presigned-url:https://s3.amazonaws.com/ *bucket* / *key* }`
	// >
	// > where *bucket* is your bucket name and *key* is the object in the bucket to which you are linking.
	DocumentSource *string `json:"documentSource" yaml:"documentSource"`
	// The ARN of the job to use as the basis for the job template.
	JobArn *string `json:"jobArn" yaml:"jobArn"`
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig interface{} `json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig interface{} `json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig interface{} `json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Metadata that can be used to manage the job template.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// A timer is started when the job execution status is set to `IN_PROGRESS` . If the job execution status is not set to another terminal state before the timer expires, it will be automatically set to `TIMED_OUT` .
	TimeoutConfig interface{} `json:"timeoutConfig" yaml:"timeoutConfig"`
}

// A CloudFormation `AWS::IoT::Logging`.
//
// Configure logging.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnLogging := iot.NewCfnLogging(this, jsii.String("MyCfnLogging"), &cfnLoggingProps{
//   	accountId: jsii.String("accountId"),
//   	defaultLogLevel: jsii.String("defaultLogLevel"),
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnLogging interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The account ID.
	AccountId() *string
	SetAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	DefaultLogLevel() *string
	SetDefaultLogLevel(val *string)
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The role ARN used for the log.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnLogging
type jsiiProxy_CfnLogging struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLogging) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) DefaultLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogging) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Logging`.
func NewCfnLogging(scope constructs.Construct, id *string, props *CfnLoggingProps) CfnLogging {
	_init_.Initialize()

	j := jsiiProxy_CfnLogging{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnLogging",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Logging`.
func NewCfnLogging_Override(c CfnLogging, scope constructs.Construct, id *string, props *CfnLoggingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnLogging",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLogging) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CfnLogging) SetDefaultLogLevel(val *string) {
	_jsii_.Set(
		j,
		"defaultLogLevel",
		val,
	)
}

func (j *jsiiProxy_CfnLogging) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLogging_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnLogging",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLogging_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnLogging",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLogging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnLogging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogging_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnLogging",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogging) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLogging) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogging) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLogging) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLogging) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLogging) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLogging) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLogging) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogging) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogging) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLogging) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLogging) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogging) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogging) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLogging`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnLoggingProps := &cfnLoggingProps{
//   	accountId: jsii.String("accountId"),
//   	defaultLogLevel: jsii.String("defaultLogLevel"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnLoggingProps struct {
	// The account ID.
	AccountId *string `json:"accountId" yaml:"accountId"`
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	DefaultLogLevel *string `json:"defaultLogLevel" yaml:"defaultLogLevel"`
	// The role ARN used for the log.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::IoT::MitigationAction`.
//
// Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask. For API reference, see [CreateMitigationAction](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateMitigationAction.html) and for general information, see [Mitigation actions](https://docs.aws.amazon.com/iot/latest/developerguide/dd-mitigation-actions.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnMitigationAction := iot.NewCfnMitigationAction(this, jsii.String("MyCfnMitigationAction"), &cfnMitigationActionProps{
//   	actionParams: &actionParamsProperty{
//   		addThingsToThingGroupParams: &addThingsToThingGroupParamsProperty{
//   			thingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
//
//   			// the properties below are optional
//   			overrideDynamicGroups: jsii.Boolean(false),
//   		},
//   		enableIoTLoggingParams: &enableIoTLoggingParamsProperty{
//   			logLevel: jsii.String("logLevel"),
//   			roleArnForLogging: jsii.String("roleArnForLogging"),
//   		},
//   		publishFindingToSnsParams: &publishFindingToSnsParamsProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   		replaceDefaultPolicyVersionParams: &replaceDefaultPolicyVersionParamsProperty{
//   			templateName: jsii.String("templateName"),
//   		},
//   		updateCaCertificateParams: &updateCACertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   		updateDeviceCertificateParams: &updateDeviceCertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnMitigationAction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The friendly name of the mitigation action.
	ActionName() *string
	SetActionName(val *string)
	// The set of parameters for this mitigation action.
	//
	// The parameters vary, depending on the kind of action you apply.
	ActionParams() interface{}
	SetActionParams(val interface{})
	// The Amazon Resource Name (ARN) of the mitigation action.
	AttrMitigationActionArn() *string
	// The ID of the mitigation action.
	AttrMitigationActionId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IAM role ARN used to apply this mitigation action.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the mitigation action.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnMitigationAction
type jsiiProxy_CfnMitigationAction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMitigationAction) ActionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) ActionParams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) AttrMitigationActionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMitigationActionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) AttrMitigationActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMitigationActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationAction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::MitigationAction`.
func NewCfnMitigationAction(scope constructs.Construct, id *string, props *CfnMitigationActionProps) CfnMitigationAction {
	_init_.Initialize()

	j := jsiiProxy_CfnMitigationAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::MitigationAction`.
func NewCfnMitigationAction_Override(c CfnMitigationAction, scope constructs.Construct, id *string, props *CfnMitigationActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMitigationAction) SetActionName(val *string) {
	_jsii_.Set(
		j,
		"actionName",
		val,
	)
}

func (j *jsiiProxy_CfnMitigationAction) SetActionParams(val interface{}) {
	_jsii_.Set(
		j,
		"actionParams",
		val,
	)
}

func (j *jsiiProxy_CfnMitigationAction) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMitigationAction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMitigationAction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnMitigationAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMitigationAction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnMitigationAction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMitigationAction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMitigationAction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMitigationAction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMitigationAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMitigationAction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMitigationAction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMitigationAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMitigationAction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMitigationAction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMitigationAction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMitigationAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMitigationAction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMitigationAction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMitigationAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMitigationAction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Defines the type of action and the parameters for that action.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   actionParamsProperty := &actionParamsProperty{
//   	addThingsToThingGroupParams: &addThingsToThingGroupParamsProperty{
//   		thingGroupNames: []*string{
//   			jsii.String("thingGroupNames"),
//   		},
//
//   		// the properties below are optional
//   		overrideDynamicGroups: jsii.Boolean(false),
//   	},
//   	enableIoTLoggingParams: &enableIoTLoggingParamsProperty{
//   		logLevel: jsii.String("logLevel"),
//   		roleArnForLogging: jsii.String("roleArnForLogging"),
//   	},
//   	publishFindingToSnsParams: &publishFindingToSnsParamsProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   	replaceDefaultPolicyVersionParams: &replaceDefaultPolicyVersionParamsProperty{
//   		templateName: jsii.String("templateName"),
//   	},
//   	updateCaCertificateParams: &updateCACertificateParamsProperty{
//   		action: jsii.String("action"),
//   	},
//   	updateDeviceCertificateParams: &updateDeviceCertificateParamsProperty{
//   		action: jsii.String("action"),
//   	},
//   }
//
type CfnMitigationAction_ActionParamsProperty struct {
	// Specifies the group to which you want to add the devices.
	AddThingsToThingGroupParams interface{} `json:"addThingsToThingGroupParams" yaml:"addThingsToThingGroupParams"`
	// Specifies the logging level and the role with permissions for logging.
	//
	// You cannot specify a logging level of `DISABLED` .
	EnableIoTLoggingParams interface{} `json:"enableIoTLoggingParams" yaml:"enableIoTLoggingParams"`
	// Specifies the topic to which the finding should be published.
	PublishFindingToSnsParams interface{} `json:"publishFindingToSnsParams" yaml:"publishFindingToSnsParams"`
	// Replaces the policy version with a default or blank policy.
	//
	// You specify the template name. Only a value of `BLANK_POLICY` is currently supported.
	ReplaceDefaultPolicyVersionParams interface{} `json:"replaceDefaultPolicyVersionParams" yaml:"replaceDefaultPolicyVersionParams"`
	// Specifies the new state for the CA certificate.
	//
	// Only a value of `DEACTIVATE` is currently supported.
	UpdateCaCertificateParams interface{} `json:"updateCaCertificateParams" yaml:"updateCaCertificateParams"`
	// Specifies the new state for a device certificate.
	//
	// Only a value of `DEACTIVATE` is currently supported.
	UpdateDeviceCertificateParams interface{} `json:"updateDeviceCertificateParams" yaml:"updateDeviceCertificateParams"`
}

// Parameters used when defining a mitigation action that move a set of things to a thing group.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   addThingsToThingGroupParamsProperty := &addThingsToThingGroupParamsProperty{
//   	thingGroupNames: []*string{
//   		jsii.String("thingGroupNames"),
//   	},
//
//   	// the properties below are optional
//   	overrideDynamicGroups: jsii.Boolean(false),
//   }
//
type CfnMitigationAction_AddThingsToThingGroupParamsProperty struct {
	// The list of groups to which you want to add the things that triggered the mitigation action.
	//
	// You can add a thing to a maximum of 10 groups, but you can't add a thing to more than one group in the same hierarchy.
	ThingGroupNames *[]*string `json:"thingGroupNames" yaml:"thingGroupNames"`
	// Specifies if this mitigation action can move the things that triggered the mitigation action even if they are part of one or more dynamic thing groups.
	OverrideDynamicGroups interface{} `json:"overrideDynamicGroups" yaml:"overrideDynamicGroups"`
}

// Parameters used when defining a mitigation action that enable AWS IoT Core logging.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   enableIoTLoggingParamsProperty := &enableIoTLoggingParamsProperty{
//   	logLevel: jsii.String("logLevel"),
//   	roleArnForLogging: jsii.String("roleArnForLogging"),
//   }
//
type CfnMitigationAction_EnableIoTLoggingParamsProperty struct {
	// Specifies the type of information to be logged.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
	// The Amazon Resource Name (ARN) of the IAM role used for logging.
	RoleArnForLogging *string `json:"roleArnForLogging" yaml:"roleArnForLogging"`
}

// Parameters to define a mitigation action that publishes findings to Amazon SNS.
//
// You can implement your own custom actions in response to the Amazon SNS messages.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   publishFindingToSnsParamsProperty := &publishFindingToSnsParamsProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnMitigationAction_PublishFindingToSnsParamsProperty struct {
	// The ARN of the topic to which you want to publish the findings.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// Parameters to define a mitigation action that adds a blank policy to restrict permissions.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   replaceDefaultPolicyVersionParamsProperty := &replaceDefaultPolicyVersionParamsProperty{
//   	templateName: jsii.String("templateName"),
//   }
//
type CfnMitigationAction_ReplaceDefaultPolicyVersionParamsProperty struct {
	// The name of the template to be applied.
	//
	// The only supported value is `BLANK_POLICY` .
	TemplateName *string `json:"templateName" yaml:"templateName"`
}

// Parameters to define a mitigation action that changes the state of the CA certificate to inactive.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   updateCACertificateParamsProperty := &updateCACertificateParamsProperty{
//   	action: jsii.String("action"),
//   }
//
type CfnMitigationAction_UpdateCACertificateParamsProperty struct {
	// The action that you want to apply to the CA certificate.
	//
	// The only supported value is `DEACTIVATE` .
	Action *string `json:"action" yaml:"action"`
}

// Parameters to define a mitigation action that changes the state of the device certificate to inactive.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   updateDeviceCertificateParamsProperty := &updateDeviceCertificateParamsProperty{
//   	action: jsii.String("action"),
//   }
//
type CfnMitigationAction_UpdateDeviceCertificateParamsProperty struct {
	// The action that you want to apply to the device certificate.
	//
	// The only supported value is `DEACTIVATE` .
	Action *string `json:"action" yaml:"action"`
}

// Properties for defining a `CfnMitigationAction`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnMitigationActionProps := &cfnMitigationActionProps{
//   	actionParams: &actionParamsProperty{
//   		addThingsToThingGroupParams: &addThingsToThingGroupParamsProperty{
//   			thingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
//
//   			// the properties below are optional
//   			overrideDynamicGroups: jsii.Boolean(false),
//   		},
//   		enableIoTLoggingParams: &enableIoTLoggingParamsProperty{
//   			logLevel: jsii.String("logLevel"),
//   			roleArnForLogging: jsii.String("roleArnForLogging"),
//   		},
//   		publishFindingToSnsParams: &publishFindingToSnsParamsProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   		replaceDefaultPolicyVersionParams: &replaceDefaultPolicyVersionParamsProperty{
//   			templateName: jsii.String("templateName"),
//   		},
//   		updateCaCertificateParams: &updateCACertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   		updateDeviceCertificateParams: &updateDeviceCertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMitigationActionProps struct {
	// The set of parameters for this mitigation action.
	//
	// The parameters vary, depending on the kind of action you apply.
	ActionParams interface{} `json:"actionParams" yaml:"actionParams"`
	// The IAM role ARN used to apply this mitigation action.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The friendly name of the mitigation action.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Metadata that can be used to manage the mitigation action.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoT::Policy`.
//
// Use the `AWS::IoT::Policy` resource to declare an AWS IoT policy. For more information about working with AWS IoT policies, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//
//   var policyDocument interface{}
//   cfnPolicy := iot.NewCfnPolicy(this, jsii.String("MyCfnPolicy"), &cfnPolicyProps{
//   	policyDocument: policyDocument,
//
//   	// the properties below are optional
//   	policyName: jsii.String("policyName"),
//   })
//
type CfnPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the AWS IoT policy, such as `arn:aws:iot:us-east-2:123456789012:policy/MyPolicy` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The JSON document that describes the policy.
	PolicyDocument() interface{}
	SetPolicyDocument(val interface{})
	// The policy name.
	PolicyName() *string
	SetPolicyName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPolicy
type jsiiProxy_CfnPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) PolicyDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Policy`.
func NewCfnPolicy(scope constructs.Construct, id *string, props *CfnPolicyProps) CfnPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Policy`.
func NewCfnPolicy_Override(c CfnPolicy, scope constructs.Construct, id *string, props *CfnPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPolicy) SetPolicyDocument(val interface{}) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::IoT::PolicyPrincipalAttachment`.
//
// Use the `AWS::IoT::PolicyPrincipalAttachment` resource to attach an AWS IoT policy to a principal (an X.509 certificate or other credential).
//
// For information about working with AWS IoT policies and principals, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnPolicyPrincipalAttachment := iot.NewCfnPolicyPrincipalAttachment(this, jsii.String("MyCfnPolicyPrincipalAttachment"), &cfnPolicyPrincipalAttachmentProps{
//   	policyName: jsii.String("policyName"),
//   	principal: jsii.String("principal"),
//   })
//
type CfnPolicyPrincipalAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The name of the AWS IoT policy.
	PolicyName() *string
	SetPolicyName(val *string)
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	Principal() *string
	SetPrincipal(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPolicyPrincipalAttachment
type jsiiProxy_CfnPolicyPrincipalAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::PolicyPrincipalAttachment`.
func NewCfnPolicyPrincipalAttachment(scope constructs.Construct, id *string, props *CfnPolicyPrincipalAttachmentProps) CfnPolicyPrincipalAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnPolicyPrincipalAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::PolicyPrincipalAttachment`.
func NewCfnPolicyPrincipalAttachment_Override(c CfnPolicyPrincipalAttachment, scope constructs.Construct, id *string, props *CfnPolicyPrincipalAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachment) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPolicyPrincipalAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPolicyPrincipalAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPolicyPrincipalAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyPrincipalAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnPolicyPrincipalAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPolicyPrincipalAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnPolicyPrincipalAttachmentProps := &cfnPolicyPrincipalAttachmentProps{
//   	policyName: jsii.String("policyName"),
//   	principal: jsii.String("principal"),
//   }
//
type CfnPolicyPrincipalAttachmentProps struct {
	// The name of the AWS IoT policy.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	Principal *string `json:"principal" yaml:"principal"`
}

// Properties for defining a `CfnPolicy`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//
//   var policyDocument interface{}
//   cfnPolicyProps := &cfnPolicyProps{
//   	policyDocument: policyDocument,
//
//   	// the properties below are optional
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnPolicyProps struct {
	// The JSON document that describes the policy.
	PolicyDocument interface{} `json:"policyDocument" yaml:"policyDocument"`
	// The policy name.
	PolicyName *string `json:"policyName" yaml:"policyName"`
}

// A CloudFormation `AWS::IoT::ProvisioningTemplate`.
//
// Creates a fleet provisioning template.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnProvisioningTemplate := iot.NewCfnProvisioningTemplate(this, jsii.String("MyCfnProvisioningTemplate"), &cfnProvisioningTemplateProps{
//   	provisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	templateBody: jsii.String("templateBody"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	preProvisioningHook: &provisioningHookProperty{
//   		payloadVersion: jsii.String("payloadVersion"),
//   		targetArn: jsii.String("targetArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateName: jsii.String("templateName"),
//   })
//
type CfnProvisioningTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN that identifies the provisioning template.
	AttrTemplateArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the fleet provisioning template.
	Description() *string
	SetDescription(val *string)
	// True to enable the fleet provisioning template, otherwise false.
	Enabled() interface{}
	SetEnabled(val interface{})
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
	// Creates a pre-provisioning hook template.
	PreProvisioningHook() interface{}
	SetPreProvisioningHook(val interface{})
	// The role ARN for the role associated with the fleet provisioning template.
	//
	// This IoT role grants permission to provision a device.
	ProvisioningRoleArn() *string
	SetProvisioningRoleArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the fleet provisioning template.
	Tags() awscdk.TagManager
	// The JSON formatted contents of the fleet provisioning template version.
	TemplateBody() *string
	SetTemplateBody(val *string)
	// The name of the fleet provisioning template.
	TemplateName() *string
	SetTemplateName(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnProvisioningTemplate
type jsiiProxy_CfnProvisioningTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProvisioningTemplate) AttrTemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTemplateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) PreProvisioningHook() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preProvisioningHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) ProvisioningRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProvisioningTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::ProvisioningTemplate`.
func NewCfnProvisioningTemplate(scope constructs.Construct, id *string, props *CfnProvisioningTemplateProps) CfnProvisioningTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnProvisioningTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::ProvisioningTemplate`.
func NewCfnProvisioningTemplate_Override(c CfnProvisioningTemplate, scope constructs.Construct, id *string, props *CfnProvisioningTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetPreProvisioningHook(val interface{}) {
	_jsii_.Set(
		j,
		"preProvisioningHook",
		val,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetProvisioningRoleArn(val *string) {
	_jsii_.Set(
		j,
		"provisioningRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CfnProvisioningTemplate) SetTemplateName(val *string) {
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnProvisioningTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnProvisioningTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnProvisioningTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProvisioningTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnProvisioningTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProvisioningTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProvisioningTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Structure that contains payloadVersion and targetArn.
//
// Provisioning hooks can be used when fleet provisioning to validate device parameters before allowing the device to be provisioned.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   provisioningHookProperty := &provisioningHookProperty{
//   	payloadVersion: jsii.String("payloadVersion"),
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnProvisioningTemplate_ProvisioningHookProperty struct {
	// The payload that was sent to the target function.
	//
	// The valid payload is `"2020-04-01"` .
	PayloadVersion *string `json:"payloadVersion" yaml:"payloadVersion"`
	// The ARN of the target function.
	TargetArn *string `json:"targetArn" yaml:"targetArn"`
}

// Properties for defining a `CfnProvisioningTemplate`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnProvisioningTemplateProps := &cfnProvisioningTemplateProps{
//   	provisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	templateBody: jsii.String("templateBody"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	preProvisioningHook: &provisioningHookProperty{
//   		payloadVersion: jsii.String("payloadVersion"),
//   		targetArn: jsii.String("targetArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateName: jsii.String("templateName"),
//   }
//
type CfnProvisioningTemplateProps struct {
	// The role ARN for the role associated with the fleet provisioning template.
	//
	// This IoT role grants permission to provision a device.
	ProvisioningRoleArn *string `json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// The JSON formatted contents of the fleet provisioning template version.
	TemplateBody *string `json:"templateBody" yaml:"templateBody"`
	// The description of the fleet provisioning template.
	Description *string `json:"description" yaml:"description"`
	// True to enable the fleet provisioning template, otherwise false.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Creates a pre-provisioning hook template.
	PreProvisioningHook interface{} `json:"preProvisioningHook" yaml:"preProvisioningHook"`
	// Metadata that can be used to manage the fleet provisioning template.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The name of the fleet provisioning template.
	TemplateName *string `json:"templateName" yaml:"templateName"`
}

// A CloudFormation `AWS::IoT::ResourceSpecificLogging`.
//
// Configure resource-specific logging.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnResourceSpecificLogging := iot.NewCfnResourceSpecificLogging(this, jsii.String("MyCfnResourceSpecificLogging"), &cfnResourceSpecificLoggingProps{
//   	logLevel: jsii.String("logLevel"),
//   	targetName: jsii.String("targetName"),
//   	targetType: jsii.String("targetType"),
//   })
//
type CfnResourceSpecificLogging interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The target Id.
	AttrTargetId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	LogLevel() *string
	SetLogLevel(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The target name.
	TargetName() *string
	SetTargetName(val *string)
	// The target type.
	//
	// Valid Values: `DEFAULT | THING_GROUP`.
	TargetType() *string
	SetTargetType(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnResourceSpecificLogging
type jsiiProxy_CfnResourceSpecificLogging struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceSpecificLogging) AttrTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) TargetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSpecificLogging) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::ResourceSpecificLogging`.
func NewCfnResourceSpecificLogging(scope constructs.Construct, id *string, props *CfnResourceSpecificLoggingProps) CfnResourceSpecificLogging {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceSpecificLogging{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::ResourceSpecificLogging`.
func NewCfnResourceSpecificLogging_Override(c CfnResourceSpecificLogging, scope constructs.Construct, id *string, props *CfnResourceSpecificLoggingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceSpecificLogging) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CfnResourceSpecificLogging) SetTargetName(val *string) {
	_jsii_.Set(
		j,
		"targetName",
		val,
	)
}

func (j *jsiiProxy_CfnResourceSpecificLogging) SetTargetType(val *string) {
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
func CfnResourceSpecificLogging_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceSpecificLogging_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnResourceSpecificLogging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceSpecificLogging_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnResourceSpecificLogging",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceSpecificLogging) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceSpecificLogging) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourceSpecificLogging`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnResourceSpecificLoggingProps := &cfnResourceSpecificLoggingProps{
//   	logLevel: jsii.String("logLevel"),
//   	targetName: jsii.String("targetName"),
//   	targetType: jsii.String("targetType"),
//   }
//
type CfnResourceSpecificLoggingProps struct {
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
	// The target name.
	TargetName *string `json:"targetName" yaml:"targetName"`
	// The target type.
	//
	// Valid Values: `DEFAULT | THING_GROUP`.
	TargetType *string `json:"targetType" yaml:"targetType"`
}

// A CloudFormation `AWS::IoT::ScheduledAudit`.
//
// Use the `AWS::IoT::ScheduledAudit` resource to create a scheduled audit that is run at a specified time interval. For API reference, see [CreateScheduleAudit](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateScheduledAudit.html) and for general information, see [Audit](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-audit.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnScheduledAudit := iot.NewCfnScheduledAudit(this, jsii.String("MyCfnScheduledAudit"), &cfnScheduledAuditProps{
//   	frequency: jsii.String("frequency"),
//   	targetCheckNames: []*string{
//   		jsii.String("targetCheckNames"),
//   	},
//
//   	// the properties below are optional
//   	dayOfMonth: jsii.String("dayOfMonth"),
//   	dayOfWeek: jsii.String("dayOfWeek"),
//   	scheduledAuditName: jsii.String("scheduledAuditName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnScheduledAudit interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the scheduled audit.
	AttrScheduledAuditArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The day of the month on which the scheduled audit is run (if the `frequency` is "MONTHLY").
	//
	// If days 29-31 are specified, and the month does not have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth() *string
	SetDayOfMonth(val *string)
	// The day of the week on which the scheduled audit is run (if the `frequency` is "WEEKLY" or "BIWEEKLY").
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	// How often the scheduled audit occurs.
	Frequency() *string
	SetFrequency(val *string)
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name of the scheduled audit.
	ScheduledAuditName() *string
	SetScheduledAuditName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the scheduled audit.
	Tags() awscdk.TagManager
	// Which checks are performed during the scheduled audit.
	//
	// Checks must be enabled for your account. (Use `DescribeAccountAuditConfiguration` to see the list of all checks, including those that are enabled or use `UpdateAccountAuditConfiguration` to select which checks are enabled.)
	//
	// The following checks are currently aviable:
	//
	// - `AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`
	// - `CA_CERTIFICATE_EXPIRING_CHECK`
	// - `CA_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `CONFLICTING_CLIENT_IDS_CHECK`
	// - `DEVICE_CERTIFICATE_EXPIRING_CHECK`
	// - `DEVICE_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `DEVICE_CERTIFICATE_SHARED_CHECK`
	// - `IOT_POLICY_OVERLY_PERMISSIVE_CHECK`
	// - `IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK`
	// - `IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK`
	// - `LOGGING_DISABLED_CHECK`
	// - `REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`.
	TargetCheckNames() *[]*string
	SetTargetCheckNames(val *[]*string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnScheduledAudit
type jsiiProxy_CfnScheduledAudit struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScheduledAudit) AttrScheduledAuditArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScheduledAuditArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) DayOfMonth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) ScheduledAuditName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledAuditName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) TargetCheckNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetCheckNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAudit) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::ScheduledAudit`.
func NewCfnScheduledAudit(scope constructs.Construct, id *string, props *CfnScheduledAuditProps) CfnScheduledAudit {
	_init_.Initialize()

	j := jsiiProxy_CfnScheduledAudit{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::ScheduledAudit`.
func NewCfnScheduledAudit_Override(c CfnScheduledAudit, scope constructs.Construct, id *string, props *CfnScheduledAuditProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScheduledAudit) SetDayOfMonth(val *string) {
	_jsii_.Set(
		j,
		"dayOfMonth",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAudit) SetDayOfWeek(val *string) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAudit) SetFrequency(val *string) {
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAudit) SetScheduledAuditName(val *string) {
	_jsii_.Set(
		j,
		"scheduledAuditName",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAudit) SetTargetCheckNames(val *[]*string) {
	_jsii_.Set(
		j,
		"targetCheckNames",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScheduledAudit_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScheduledAudit_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnScheduledAudit_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledAudit_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnScheduledAudit",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScheduledAudit) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAudit) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnScheduledAudit`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnScheduledAuditProps := &cfnScheduledAuditProps{
//   	frequency: jsii.String("frequency"),
//   	targetCheckNames: []*string{
//   		jsii.String("targetCheckNames"),
//   	},
//
//   	// the properties below are optional
//   	dayOfMonth: jsii.String("dayOfMonth"),
//   	dayOfWeek: jsii.String("dayOfWeek"),
//   	scheduledAuditName: jsii.String("scheduledAuditName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnScheduledAuditProps struct {
	// How often the scheduled audit occurs.
	Frequency *string `json:"frequency" yaml:"frequency"`
	// Which checks are performed during the scheduled audit.
	//
	// Checks must be enabled for your account. (Use `DescribeAccountAuditConfiguration` to see the list of all checks, including those that are enabled or use `UpdateAccountAuditConfiguration` to select which checks are enabled.)
	//
	// The following checks are currently aviable:
	//
	// - `AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`
	// - `CA_CERTIFICATE_EXPIRING_CHECK`
	// - `CA_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `CONFLICTING_CLIENT_IDS_CHECK`
	// - `DEVICE_CERTIFICATE_EXPIRING_CHECK`
	// - `DEVICE_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `DEVICE_CERTIFICATE_SHARED_CHECK`
	// - `IOT_POLICY_OVERLY_PERMISSIVE_CHECK`
	// - `IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK`
	// - `IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK`
	// - `LOGGING_DISABLED_CHECK`
	// - `REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`.
	TargetCheckNames *[]*string `json:"targetCheckNames" yaml:"targetCheckNames"`
	// The day of the month on which the scheduled audit is run (if the `frequency` is "MONTHLY").
	//
	// If days 29-31 are specified, and the month does not have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week on which the scheduled audit is run (if the `frequency` is "WEEKLY" or "BIWEEKLY").
	DayOfWeek *string `json:"dayOfWeek" yaml:"dayOfWeek"`
	// The name of the scheduled audit.
	ScheduledAuditName *string `json:"scheduledAuditName" yaml:"scheduledAuditName"`
	// Metadata that can be used to manage the scheduled audit.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoT::SecurityProfile`.
//
// Use the `AWS::IoT::SecurityProfile` resource to create a Device Defender security profile. For API reference, see [CreateSecurityProfile](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateSecurityProfile.html) and for general information, see [Detect](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-detect.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnSecurityProfile := iot.NewCfnSecurityProfile(this, jsii.String("MyCfnSecurityProfile"), &cfnSecurityProfileProps{
//   	additionalMetricsToRetainV2: []interface{}{
//   		&metricToRetainProperty{
//   			metric: jsii.String("metric"),
//
//   			// the properties below are optional
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	alertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	behaviors: []interface{}{
//   		&behaviorProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			criteria: &behaviorCriteriaProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				consecutiveDatapointsToAlarm: jsii.Number(123),
//   				consecutiveDatapointsToClear: jsii.Number(123),
//   				durationSeconds: jsii.Number(123),
//   				mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   					confidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				statisticalThreshold: &statisticalThresholdProperty{
//   					statistic: jsii.String("statistic"),
//   				},
//   				value: &metricValueProperty{
//   					cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					count: jsii.String("count"),
//   					number: jsii.Number(123),
//   					numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			metric: jsii.String("metric"),
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   			suppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	securityProfileDescription: jsii.String("securityProfileDescription"),
//   	securityProfileName: jsii.String("securityProfileName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//   })
//
type CfnSecurityProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of metrics whose data is retained (stored).
	//
	// By default, data is retained for any metric used in the profile's `behaviors` , but it's also retained for any metric specified here. Can be used with custom metrics; can't be used with dimensions.
	AdditionalMetricsToRetainV2() interface{}
	SetAdditionalMetricsToRetainV2(val interface{})
	// Specifies the destinations to which alerts are sent.
	//
	// (Alerts are always sent to the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets() interface{}
	SetAlertTargets(val interface{})
	// The Amazon Resource Name (ARN) of the security profile.
	AttrSecurityProfileArn() *string
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors() interface{}
	SetBehaviors(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A description of the security profile.
	SecurityProfileDescription() *string
	SetSecurityProfileDescription(val *string)
	// The name you gave to the security profile.
	SecurityProfileName() *string
	SetSecurityProfileName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that can be used to manage the security profile.
	Tags() awscdk.TagManager
	// The ARN of the target (thing group) to which the security profile is attached.
	TargetArns() *[]*string
	SetTargetArns(val *[]*string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnSecurityProfile
type jsiiProxy_CfnSecurityProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecurityProfile) AdditionalMetricsToRetainV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalMetricsToRetainV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) AlertTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) AttrSecurityProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSecurityProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Behaviors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) SecurityProfileDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) SecurityProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::SecurityProfile`.
func NewCfnSecurityProfile(scope constructs.Construct, id *string, props *CfnSecurityProfileProps) CfnSecurityProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnSecurityProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::SecurityProfile`.
func NewCfnSecurityProfile_Override(c CfnSecurityProfile, scope constructs.Construct, id *string, props *CfnSecurityProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetAdditionalMetricsToRetainV2(val interface{}) {
	_jsii_.Set(
		j,
		"additionalMetricsToRetainV2",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetAlertTargets(val interface{}) {
	_jsii_.Set(
		j,
		"alertTargets",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetBehaviors(val interface{}) {
	_jsii_.Set(
		j,
		"behaviors",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetSecurityProfileDescription(val *string) {
	_jsii_.Set(
		j,
		"securityProfileDescription",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetSecurityProfileName(val *string) {
	_jsii_.Set(
		j,
		"securityProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile) SetTargetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetArns",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSecurityProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSecurityProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnSecurityProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A structure containing the alert target ARN and the role ARN.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   alertTargetProperty := &alertTargetProperty{
//   	alertTargetArn: jsii.String("alertTargetArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnSecurityProfile_AlertTargetProperty struct {
	// The Amazon Resource Name (ARN) of the notification target to which alerts are sent.
	AlertTargetArn *string `json:"alertTargetArn" yaml:"alertTargetArn"`
	// The ARN of the role that grants permission to send alerts to the notification target.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// The criteria by which the behavior is determined to be normal.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   behaviorCriteriaProperty := &behaviorCriteriaProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	consecutiveDatapointsToAlarm: jsii.Number(123),
//   	consecutiveDatapointsToClear: jsii.Number(123),
//   	durationSeconds: jsii.Number(123),
//   	mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   		confidenceLevel: jsii.String("confidenceLevel"),
//   	},
//   	statisticalThreshold: &statisticalThresholdProperty{
//   		statistic: jsii.String("statistic"),
//   	},
//   	value: &metricValueProperty{
//   		cidrs: []*string{
//   			jsii.String("cidrs"),
//   		},
//   		count: jsii.String("count"),
//   		number: jsii.Number(123),
//   		numbers: []interface{}{
//   			jsii.Number(123),
//   		},
//   		ports: []interface{}{
//   			jsii.Number(123),
//   		},
//   		strings: []*string{
//   			jsii.String("strings"),
//   		},
//   	},
//   }
//
type CfnSecurityProfile_BehaviorCriteriaProperty struct {
	// The operator that relates the thing measured ( `metric` ) to the criteria (containing a `value` or `statisticalThreshold` ).
	//
	// Valid operators include:
	//
	// - `string-list` : `in-set` and `not-in-set`
	// - `number-list` : `in-set` and `not-in-set`
	// - `ip-address-list` : `in-cidr-set` and `not-in-cidr-set`
	// - `number` : `less-than` , `less-than-equals` , `greater-than` , and `greater-than-equals`.
	ComparisonOperator *string `json:"comparisonOperator" yaml:"comparisonOperator"`
	// If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs.
	//
	// If not specified, the default is 1.
	ConsecutiveDatapointsToAlarm *float64 `json:"consecutiveDatapointsToAlarm" yaml:"consecutiveDatapointsToAlarm"`
	// If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared.
	//
	// If not specified, the default is 1.
	ConsecutiveDatapointsToClear *float64 `json:"consecutiveDatapointsToClear" yaml:"consecutiveDatapointsToClear"`
	// Use this to specify the time duration over which the behavior is evaluated, for those criteria that have a time dimension (for example, `NUM_MESSAGES_SENT` ).
	//
	// For a `statisticalThreshhold` metric comparison, measurements from all devices are accumulated over this time duration before being used to calculate percentiles, and later, measurements from an individual device are also accumulated over this time duration before being given a percentile rank. Cannot be used with list-based metric datatypes.
	DurationSeconds *float64 `json:"durationSeconds" yaml:"durationSeconds"`
	// The confidence level of the detection model.
	MlDetectionConfig interface{} `json:"mlDetectionConfig" yaml:"mlDetectionConfig"`
	// A statistical ranking (percentile)that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
	StatisticalThreshold interface{} `json:"statisticalThreshold" yaml:"statisticalThreshold"`
	// The value to be compared with the `metric` .
	Value interface{} `json:"value" yaml:"value"`
}

// A Device Defender security profile behavior.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   behaviorProperty := &behaviorProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	criteria: &behaviorCriteriaProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		consecutiveDatapointsToAlarm: jsii.Number(123),
//   		consecutiveDatapointsToClear: jsii.Number(123),
//   		durationSeconds: jsii.Number(123),
//   		mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   			confidenceLevel: jsii.String("confidenceLevel"),
//   		},
//   		statisticalThreshold: &statisticalThresholdProperty{
//   			statistic: jsii.String("statistic"),
//   		},
//   		value: &metricValueProperty{
//   			cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			count: jsii.String("count"),
//   			number: jsii.Number(123),
//   			numbers: []interface{}{
//   				jsii.Number(123),
//   			},
//   			ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   			strings: []*string{
//   				jsii.String("strings"),
//   			},
//   		},
//   	},
//   	metric: jsii.String("metric"),
//   	metricDimension: &metricDimensionProperty{
//   		dimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		operator: jsii.String("operator"),
//   	},
//   	suppressAlerts: jsii.Boolean(false),
//   }
//
type CfnSecurityProfile_BehaviorProperty struct {
	// The name you've given to the behavior.
	Name *string `json:"name" yaml:"name"`
	// The criteria that determine if a device is behaving normally in regard to the `metric` .
	Criteria interface{} `json:"criteria" yaml:"criteria"`
	// What is measured by the behavior.
	Metric *string `json:"metric" yaml:"metric"`
	// The dimension of the metric.
	MetricDimension interface{} `json:"metricDimension" yaml:"metricDimension"`
	// The alert status.
	//
	// If you set the value to `true` , alerts will be suppressed.
	SuppressAlerts interface{} `json:"suppressAlerts" yaml:"suppressAlerts"`
}

// The `MachineLearningDetectionConfig` property type controls confidence of the machine learning model.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   machineLearningDetectionConfigProperty := &machineLearningDetectionConfigProperty{
//   	confidenceLevel: jsii.String("confidenceLevel"),
//   }
//
type CfnSecurityProfile_MachineLearningDetectionConfigProperty struct {
	// The model confidence level.
	//
	// There are three levels of confidence, `"high"` , `"medium"` , and `"low"` .
	//
	// The higher the confidence level, the lower the sensitivity, and the lower the alarm frequency will be.
	ConfidenceLevel *string `json:"confidenceLevel" yaml:"confidenceLevel"`
}

// The dimension of the metric.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   metricDimensionProperty := &metricDimensionProperty{
//   	dimensionName: jsii.String("dimensionName"),
//
//   	// the properties below are optional
//   	operator: jsii.String("operator"),
//   }
//
type CfnSecurityProfile_MetricDimensionProperty struct {
	// The name of the dimension.
	DimensionName *string `json:"dimensionName" yaml:"dimensionName"`
	// Operators are constructs that perform logical operations.
	//
	// Valid values are `IN` and `NOT_IN` .
	Operator *string `json:"operator" yaml:"operator"`
}

// The metric you want to retain.
//
// Dimensions are optional.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   metricToRetainProperty := &metricToRetainProperty{
//   	metric: jsii.String("metric"),
//
//   	// the properties below are optional
//   	metricDimension: &metricDimensionProperty{
//   		dimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		operator: jsii.String("operator"),
//   	},
//   }
//
type CfnSecurityProfile_MetricToRetainProperty struct {
	// A standard of measurement.
	Metric *string `json:"metric" yaml:"metric"`
	// The dimension of the metric.
	MetricDimension interface{} `json:"metricDimension" yaml:"metricDimension"`
}

// The value to be compared with the `metric` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   metricValueProperty := &metricValueProperty{
//   	cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	count: jsii.String("count"),
//   	number: jsii.Number(123),
//   	numbers: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ports: []interface{}{
//   		jsii.Number(123),
//   	},
//   	strings: []*string{
//   		jsii.String("strings"),
//   	},
//   }
//
type CfnSecurityProfile_MetricValueProperty struct {
	// If the `comparisonOperator` calls for a set of CIDRs, use this to specify that set to be compared with the `metric` .
	Cidrs *[]*string `json:"cidrs" yaml:"cidrs"`
	// If the `comparisonOperator` calls for a numeric value, use this to specify that numeric value to be compared with the `metric` .
	Count *string `json:"count" yaml:"count"`
	// The numeric values of a metric.
	Number *float64 `json:"number" yaml:"number"`
	// The numeric value of a metric.
	Numbers interface{} `json:"numbers" yaml:"numbers"`
	// If the `comparisonOperator` calls for a set of ports, use this to specify that set to be compared with the `metric` .
	Ports interface{} `json:"ports" yaml:"ports"`
	// The string values of a metric.
	Strings *[]*string `json:"strings" yaml:"strings"`
}

// A statistical ranking (percentile) that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   statisticalThresholdProperty := &statisticalThresholdProperty{
//   	statistic: jsii.String("statistic"),
//   }
//
type CfnSecurityProfile_StatisticalThresholdProperty struct {
	// The percentile that resolves to a threshold value by which compliance with a behavior is determined.
	//
	// Metrics are collected over the specified period ( `durationSeconds` ) from all reporting devices in your account and statistical ranks are calculated. Then, the measurements from a device are collected over the same period. If the accumulated measurements from the device fall above or below ( `comparisonOperator` ) the value associated with the percentile specified, then the device is considered to be in compliance with the behavior, otherwise a violation occurs.
	Statistic *string `json:"statistic" yaml:"statistic"`
}

// Properties for defining a `CfnSecurityProfile`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnSecurityProfileProps := &cfnSecurityProfileProps{
//   	additionalMetricsToRetainV2: []interface{}{
//   		&metricToRetainProperty{
//   			metric: jsii.String("metric"),
//
//   			// the properties below are optional
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	alertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	behaviors: []interface{}{
//   		&behaviorProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			criteria: &behaviorCriteriaProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				consecutiveDatapointsToAlarm: jsii.Number(123),
//   				consecutiveDatapointsToClear: jsii.Number(123),
//   				durationSeconds: jsii.Number(123),
//   				mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   					confidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				statisticalThreshold: &statisticalThresholdProperty{
//   					statistic: jsii.String("statistic"),
//   				},
//   				value: &metricValueProperty{
//   					cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					count: jsii.String("count"),
//   					number: jsii.Number(123),
//   					numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			metric: jsii.String("metric"),
//   			metricDimension: &metricDimensionProperty{
//   				dimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				operator: jsii.String("operator"),
//   			},
//   			suppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	securityProfileDescription: jsii.String("securityProfileDescription"),
//   	securityProfileName: jsii.String("securityProfileName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//   }
//
type CfnSecurityProfileProps struct {
	// A list of metrics whose data is retained (stored).
	//
	// By default, data is retained for any metric used in the profile's `behaviors` , but it's also retained for any metric specified here. Can be used with custom metrics; can't be used with dimensions.
	AdditionalMetricsToRetainV2 interface{} `json:"additionalMetricsToRetainV2" yaml:"additionalMetricsToRetainV2"`
	// Specifies the destinations to which alerts are sent.
	//
	// (Alerts are always sent to the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets interface{} `json:"alertTargets" yaml:"alertTargets"`
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors interface{} `json:"behaviors" yaml:"behaviors"`
	// A description of the security profile.
	SecurityProfileDescription *string `json:"securityProfileDescription" yaml:"securityProfileDescription"`
	// The name you gave to the security profile.
	SecurityProfileName *string `json:"securityProfileName" yaml:"securityProfileName"`
	// Metadata that can be used to manage the security profile.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The ARN of the target (thing group) to which the security profile is attached.
	TargetArns *[]*string `json:"targetArns" yaml:"targetArns"`
}

// A CloudFormation `AWS::IoT::Thing`.
//
// Use the `AWS::IoT::Thing` resource to declare an AWS IoT thing.
//
// For information about working with things, see [How AWS IoT Works](https://docs.aws.amazon.com/iot/latest/developerguide/aws-iot-how-it-works.html) and [Device Registry for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/thing-registry.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnThing := iot.NewCfnThing(this, jsii.String("MyCfnThing"), &cfnThingProps{
//   	attributePayload: &attributePayloadProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   	},
//   	thingName: jsii.String("thingName"),
//   })
//
type CfnThing interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A string that contains up to three key value pairs.
	//
	// Maximum length of 800. Duplicates not allowed.
	AttributePayload() interface{}
	SetAttributePayload(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the thing to update.
	//
	// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
	ThingName() *string
	SetThingName(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnThing
type jsiiProxy_CfnThing struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnThing) AttributePayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributePayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) ThingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThing) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::Thing`.
func NewCfnThing(scope constructs.Construct, id *string, props *CfnThingProps) CfnThing {
	_init_.Initialize()

	j := jsiiProxy_CfnThing{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnThing",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::Thing`.
func NewCfnThing_Override(c CfnThing, scope constructs.Construct, id *string, props *CfnThingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnThing",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnThing) SetAttributePayload(val interface{}) {
	_jsii_.Set(
		j,
		"attributePayload",
		val,
	)
}

func (j *jsiiProxy_CfnThing) SetThingName(val *string) {
	_jsii_.Set(
		j,
		"thingName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnThing_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThing",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnThing_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThing",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnThing_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThing",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnThing_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnThing",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnThing) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnThing) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnThing) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnThing) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnThing) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnThing) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnThing) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnThing) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThing) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThing) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnThing) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnThing) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThing) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThing) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThing) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The AttributePayload property specifies up to three attributes for an AWS IoT as key-value pairs.
//
// AttributePayload is a property of the [AWS::IoT::Thing](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html) resource.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   attributePayloadProperty := &attributePayloadProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   }
//
type CfnThing_AttributePayloadProperty struct {
	// A JSON string containing up to three key-value pair in JSON format. For example:.
	//
	// `{\"attributes\":{\"string1\":\"string2\"}}`.
	Attributes interface{} `json:"attributes" yaml:"attributes"`
}

// A CloudFormation `AWS::IoT::ThingPrincipalAttachment`.
//
// Use the `AWS::IoT::ThingPrincipalAttachment` resource to attach a principal (an X.509 certificate or another credential) to a thing.
//
// For more information about working with AWS IoT things and principals, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnThingPrincipalAttachment := iot.NewCfnThingPrincipalAttachment(this, jsii.String("MyCfnThingPrincipalAttachment"), &cfnThingPrincipalAttachmentProps{
//   	principal: jsii.String("principal"),
//   	thingName: jsii.String("thingName"),
//   })
//
type CfnThingPrincipalAttachment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	Principal() *string
	SetPrincipal(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the AWS IoT thing.
	ThingName() *string
	SetThingName(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnThingPrincipalAttachment
type jsiiProxy_CfnThingPrincipalAttachment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) ThingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::ThingPrincipalAttachment`.
func NewCfnThingPrincipalAttachment(scope constructs.Construct, id *string, props *CfnThingPrincipalAttachmentProps) CfnThingPrincipalAttachment {
	_init_.Initialize()

	j := jsiiProxy_CfnThingPrincipalAttachment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::ThingPrincipalAttachment`.
func NewCfnThingPrincipalAttachment_Override(c CfnThingPrincipalAttachment, scope constructs.Construct, id *string, props *CfnThingPrincipalAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnThingPrincipalAttachment) SetThingName(val *string) {
	_jsii_.Set(
		j,
		"thingName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnThingPrincipalAttachment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnThingPrincipalAttachment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnThingPrincipalAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnThingPrincipalAttachment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnThingPrincipalAttachment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnThingPrincipalAttachment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnThingPrincipalAttachmentProps := &cfnThingPrincipalAttachmentProps{
//   	principal: jsii.String("principal"),
//   	thingName: jsii.String("thingName"),
//   }
//
type CfnThingPrincipalAttachmentProps struct {
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	Principal *string `json:"principal" yaml:"principal"`
	// The name of the AWS IoT thing.
	ThingName *string `json:"thingName" yaml:"thingName"`
}

// Properties for defining a `CfnThing`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnThingProps := &cfnThingProps{
//   	attributePayload: &attributePayloadProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   	},
//   	thingName: jsii.String("thingName"),
//   }
//
type CfnThingProps struct {
	// A string that contains up to three key value pairs.
	//
	// Maximum length of 800. Duplicates not allowed.
	AttributePayload interface{} `json:"attributePayload" yaml:"attributePayload"`
	// The name of the thing to update.
	//
	// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
	ThingName *string `json:"thingName" yaml:"thingName"`
}

// A CloudFormation `AWS::IoT::TopicRule`.
//
// Use the `AWS::IoT::TopicRule` resource to declare an AWS IoT rule. For information about working with AWS IoT rules, see [Rules for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnTopicRule := iot.NewCfnTopicRule(this, jsii.String("MyCfnTopicRule"), &cfnTopicRuleProps{
//   	topicRulePayload: &topicRulePayloadProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   					alarmName: jsii.String("alarmName"),
//   					roleArn: jsii.String("roleArn"),
//   					stateReason: jsii.String("stateReason"),
//   					stateValue: jsii.String("stateValue"),
//   				},
//   				cloudwatchLogs: &cloudwatchLogsActionProperty{
//   					logGroupName: jsii.String("logGroupName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchMetric: &cloudwatchMetricActionProperty{
//   					metricName: jsii.String("metricName"),
//   					metricNamespace: jsii.String("metricNamespace"),
//   					metricUnit: jsii.String("metricUnit"),
//   					metricValue: jsii.String("metricValue"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					metricTimestamp: jsii.String("metricTimestamp"),
//   				},
//   				dynamoDb: &dynamoDBActionProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2ActionProperty{
//   					putItem: &putItemInputProperty{
//   						tableName: jsii.String("tableName"),
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				elasticsearch: &elasticsearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				firehose: &firehoseActionProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					separator: jsii.String("separator"),
//   				},
//   				http: &httpActionProperty{
//   					url: jsii.String("url"),
//
//   					// the properties below are optional
//   					auth: &httpAuthorizationProperty{
//   						sigv4: &sigV4AuthorizationProperty{
//   							roleArn: jsii.String("roleArn"),
//   							serviceName: jsii.String("serviceName"),
//   							signingRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					confirmationUrl: jsii.String("confirmationUrl"),
//   					headers: []interface{}{
//   						&httpActionHeaderProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				iotAnalytics: &iotAnalyticsActionProperty{
//   					channelName: jsii.String("channelName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   				},
//   				iotEvents: &iotEventsActionProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					messageId: jsii.String("messageId"),
//   				},
//   				iotSiteWise: &iotSiteWiseActionProperty{
//   					putAssetPropertyValueEntries: []interface{}{
//   						&putAssetPropertyValueEntryProperty{
//   							propertyValues: []interface{}{
//   								&assetPropertyValueProperty{
//   									timestamp: &assetPropertyTimestampProperty{
//   										timeInSeconds: jsii.String("timeInSeconds"),
//
//   										// the properties below are optional
//   										offsetInNanos: jsii.String("offsetInNanos"),
//   									},
//   									value: &assetPropertyVariantProperty{
//   										booleanValue: jsii.String("booleanValue"),
//   										doubleValue: jsii.String("doubleValue"),
//   										integerValue: jsii.String("integerValue"),
//   										stringValue: jsii.String("stringValue"),
//   									},
//
//   									// the properties below are optional
//   									quality: jsii.String("quality"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				kafka: &kafkaActionProperty{
//   					clientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					destinationArn: jsii.String("destinationArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					key: jsii.String("key"),
//   					partition: jsii.String("partition"),
//   				},
//   				kinesis: &kinesisActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					streamName: jsii.String("streamName"),
//
//   					// the properties below are optional
//   					partitionKey: jsii.String("partitionKey"),
//   				},
//   				lambda: &lambdaActionProperty{
//   					functionArn: jsii.String("functionArn"),
//   				},
//   				openSearch: &openSearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				republish: &republishActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					qos: jsii.Number(123),
//   				},
//   				s3: &s3ActionProperty{
//   					bucketName: jsii.String("bucketName"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   				sns: &snsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					messageFormat: jsii.String("messageFormat"),
//   				},
//   				sqs: &sqsActionProperty{
//   					queueUrl: jsii.String("queueUrl"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					useBase64: jsii.Boolean(false),
//   				},
//   				stepFunctions: &stepFunctionsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					stateMachineName: jsii.String("stateMachineName"),
//
//   					// the properties below are optional
//   					executionNamePrefix: jsii.String("executionNamePrefix"),
//   				},
//   				timestream: &timestreamActionProperty{
//   					databaseName: jsii.String("databaseName"),
//   					dimensions: []interface{}{
//   						&timestreamDimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					timestamp: &timestreamTimestampProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		sql: jsii.String("sql"),
//
//   		// the properties below are optional
//   		awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   		description: jsii.String("description"),
//   		errorAction: &actionProperty{
//   			cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   				alarmName: jsii.String("alarmName"),
//   				roleArn: jsii.String("roleArn"),
//   				stateReason: jsii.String("stateReason"),
//   				stateValue: jsii.String("stateValue"),
//   			},
//   			cloudwatchLogs: &cloudwatchLogsActionProperty{
//   				logGroupName: jsii.String("logGroupName"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			cloudwatchMetric: &cloudwatchMetricActionProperty{
//   				metricName: jsii.String("metricName"),
//   				metricNamespace: jsii.String("metricNamespace"),
//   				metricUnit: jsii.String("metricUnit"),
//   				metricValue: jsii.String("metricValue"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				metricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			dynamoDb: &dynamoDBActionProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2ActionProperty{
//   				putItem: &putItemInputProperty{
//   					tableName: jsii.String("tableName"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			elasticsearch: &elasticsearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			firehose: &firehoseActionProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				separator: jsii.String("separator"),
//   			},
//   			http: &httpActionProperty{
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				auth: &httpAuthorizationProperty{
//   					sigv4: &sigV4AuthorizationProperty{
//   						roleArn: jsii.String("roleArn"),
//   						serviceName: jsii.String("serviceName"),
//   						signingRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				confirmationUrl: jsii.String("confirmationUrl"),
//   				headers: []interface{}{
//   					&httpActionHeaderProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			iotAnalytics: &iotAnalyticsActionProperty{
//   				channelName: jsii.String("channelName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			iotEvents: &iotEventsActionProperty{
//   				inputName: jsii.String("inputName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				messageId: jsii.String("messageId"),
//   			},
//   			iotSiteWise: &iotSiteWiseActionProperty{
//   				putAssetPropertyValueEntries: []interface{}{
//   					&putAssetPropertyValueEntryProperty{
//   						propertyValues: []interface{}{
//   							&assetPropertyValueProperty{
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kafka: &kafkaActionProperty{
//   				clientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				destinationArn: jsii.String("destinationArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				partition: jsii.String("partition"),
//   			},
//   			kinesis: &kinesisActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				partitionKey: jsii.String("partitionKey"),
//   			},
//   			lambda: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//   			},
//   			openSearch: &openSearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			republish: &republishActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				qos: jsii.Number(123),
//   			},
//   			s3: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//   				key: jsii.String("key"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			sns: &snsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				messageFormat: jsii.String("messageFormat"),
//   			},
//   			sqs: &sqsActionProperty{
//   				queueUrl: jsii.String("queueUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				useBase64: jsii.Boolean(false),
//   			},
//   			stepFunctions: &stepFunctionsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				stateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				executionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			timestream: &timestreamActionProperty{
//   				databaseName: jsii.String("databaseName"),
//   				dimensions: []interface{}{
//   					&timestreamDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				timestamp: &timestreamTimestampProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		ruleDisabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	ruleName: jsii.String("ruleName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnTopicRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the AWS IoT rule, such as `arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name of the rule.
	RuleName() *string
	SetRuleName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags() awscdk.TagManager
	// The rule payload.
	TopicRulePayload() interface{}
	SetTopicRulePayload(val interface{})
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnTopicRule
type jsiiProxy_CfnTopicRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopicRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) TopicRulePayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicRulePayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule(scope constructs.Construct, id *string, props *CfnTopicRuleProps) CfnTopicRule {
	_init_.Initialize()

	j := jsiiProxy_CfnTopicRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::TopicRule`.
func NewCfnTopicRule_Override(c CfnTopicRule, scope constructs.Construct, id *string, props *CfnTopicRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopicRule) SetRuleName(val *string) {
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_CfnTopicRule) SetTopicRulePayload(val interface{}) {
	_jsii_.Set(
		j,
		"topicRulePayload",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTopicRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTopicRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnTopicRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopicRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopicRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopicRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopicRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopicRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the actions associated with a rule.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   actionProperty := &actionProperty{
//   	cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   		alarmName: jsii.String("alarmName"),
//   		roleArn: jsii.String("roleArn"),
//   		stateReason: jsii.String("stateReason"),
//   		stateValue: jsii.String("stateValue"),
//   	},
//   	cloudwatchLogs: &cloudwatchLogsActionProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	cloudwatchMetric: &cloudwatchMetricActionProperty{
//   		metricName: jsii.String("metricName"),
//   		metricNamespace: jsii.String("metricNamespace"),
//   		metricUnit: jsii.String("metricUnit"),
//   		metricValue: jsii.String("metricValue"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		metricTimestamp: jsii.String("metricTimestamp"),
//   	},
//   	dynamoDb: &dynamoDBActionProperty{
//   		hashKeyField: jsii.String("hashKeyField"),
//   		hashKeyValue: jsii.String("hashKeyValue"),
//   		roleArn: jsii.String("roleArn"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		hashKeyType: jsii.String("hashKeyType"),
//   		payloadField: jsii.String("payloadField"),
//   		rangeKeyField: jsii.String("rangeKeyField"),
//   		rangeKeyType: jsii.String("rangeKeyType"),
//   		rangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	dynamoDBv2: &dynamoDBv2ActionProperty{
//   		putItem: &putItemInputProperty{
//   			tableName: jsii.String("tableName"),
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	elasticsearch: &elasticsearchActionProperty{
//   		endpoint: jsii.String("endpoint"),
//   		id: jsii.String("id"),
//   		index: jsii.String("index"),
//   		roleArn: jsii.String("roleArn"),
//   		type: jsii.String("type"),
//   	},
//   	firehose: &firehoseActionProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   		separator: jsii.String("separator"),
//   	},
//   	http: &httpActionProperty{
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		auth: &httpAuthorizationProperty{
//   			sigv4: &sigV4AuthorizationProperty{
//   				roleArn: jsii.String("roleArn"),
//   				serviceName: jsii.String("serviceName"),
//   				signingRegion: jsii.String("signingRegion"),
//   			},
//   		},
//   		confirmationUrl: jsii.String("confirmationUrl"),
//   		headers: []interface{}{
//   			&httpActionHeaderProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	iotAnalytics: &iotAnalyticsActionProperty{
//   		channelName: jsii.String("channelName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   	},
//   	iotEvents: &iotEventsActionProperty{
//   		inputName: jsii.String("inputName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   		messageId: jsii.String("messageId"),
//   	},
//   	iotSiteWise: &iotSiteWiseActionProperty{
//   		putAssetPropertyValueEntries: []interface{}{
//   			&putAssetPropertyValueEntryProperty{
//   				propertyValues: []interface{}{
//   					&assetPropertyValueProperty{
//   						timestamp: &assetPropertyTimestampProperty{
//   							timeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							offsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   						value: &assetPropertyVariantProperty{
//   							booleanValue: jsii.String("booleanValue"),
//   							doubleValue: jsii.String("doubleValue"),
//   							integerValue: jsii.String("integerValue"),
//   							stringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						quality: jsii.String("quality"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	kafka: &kafkaActionProperty{
//   		clientProperties: map[string]*string{
//   			"clientPropertiesKey": jsii.String("clientProperties"),
//   		},
//   		destinationArn: jsii.String("destinationArn"),
//   		topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   		partition: jsii.String("partition"),
//   	},
//   	kinesis: &kinesisActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		streamName: jsii.String("streamName"),
//
//   		// the properties below are optional
//   		partitionKey: jsii.String("partitionKey"),
//   	},
//   	lambda: &lambdaActionProperty{
//   		functionArn: jsii.String("functionArn"),
//   	},
//   	openSearch: &openSearchActionProperty{
//   		endpoint: jsii.String("endpoint"),
//   		id: jsii.String("id"),
//   		index: jsii.String("index"),
//   		roleArn: jsii.String("roleArn"),
//   		type: jsii.String("type"),
//   	},
//   	republish: &republishActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		qos: jsii.Number(123),
//   	},
//   	s3: &s3ActionProperty{
//   		bucketName: jsii.String("bucketName"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   	sns: &snsActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		targetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		messageFormat: jsii.String("messageFormat"),
//   	},
//   	sqs: &sqsActionProperty{
//   		queueUrl: jsii.String("queueUrl"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		useBase64: jsii.Boolean(false),
//   	},
//   	stepFunctions: &stepFunctionsActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		stateMachineName: jsii.String("stateMachineName"),
//
//   		// the properties below are optional
//   		executionNamePrefix: jsii.String("executionNamePrefix"),
//   	},
//   	timestream: &timestreamActionProperty{
//   		databaseName: jsii.String("databaseName"),
//   		dimensions: []interface{}{
//   			&timestreamDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   		timestamp: &timestreamTimestampProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_ActionProperty struct {
	// Change the state of a CloudWatch alarm.
	CloudwatchAlarm interface{} `json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// Sends data to CloudWatch.
	CloudwatchLogs interface{} `json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Capture a CloudWatch metric.
	CloudwatchMetric interface{} `json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Write to a DynamoDB table.
	DynamoDb interface{} `json:"dynamoDb" yaml:"dynamoDb"`
	// Write to a DynamoDB table.
	//
	// This is a new version of the DynamoDB action. It allows you to write each attribute in an MQTT message payload into a separate DynamoDB column.
	DynamoDBv2 interface{} `json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Write data to an Amazon OpenSearch Service domain.
	//
	// > The `Elasticsearch` action can only be used by existing rule actions. To create a new rule action or to update an existing rule action, use the `OpenSearch` rule action instead. For more information, see [OpenSearchAction](https://docs.aws.amazon.com//iot/latest/apireference/API_OpenSearchAction.html) .
	Elasticsearch interface{} `json:"elasticsearch" yaml:"elasticsearch"`
	// Write to an Amazon Kinesis Firehose stream.
	Firehose interface{} `json:"firehose" yaml:"firehose"`
	// Send data to an HTTPS endpoint.
	Http interface{} `json:"http" yaml:"http"`
	// Sends message data to an AWS IoT Analytics channel.
	IotAnalytics interface{} `json:"iotAnalytics" yaml:"iotAnalytics"`
	// Sends an input to an AWS IoT Events detector.
	IotEvents interface{} `json:"iotEvents" yaml:"iotEvents"`
	// Sends data from the MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
	IotSiteWise interface{} `json:"iotSiteWise" yaml:"iotSiteWise"`
	// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
	Kafka interface{} `json:"kafka" yaml:"kafka"`
	// Write data to an Amazon Kinesis stream.
	Kinesis interface{} `json:"kinesis" yaml:"kinesis"`
	// Invoke a Lambda function.
	Lambda interface{} `json:"lambda" yaml:"lambda"`
	// Write data to an Amazon OpenSearch Service domain.
	OpenSearch interface{} `json:"openSearch" yaml:"openSearch"`
	// Publish to another MQTT topic.
	Republish interface{} `json:"republish" yaml:"republish"`
	// Write to an Amazon S3 bucket.
	S3 interface{} `json:"s3" yaml:"s3"`
	// Publish to an Amazon SNS topic.
	Sns interface{} `json:"sns" yaml:"sns"`
	// Publish to an Amazon SQS queue.
	Sqs interface{} `json:"sqs" yaml:"sqs"`
	// Starts execution of a Step Functions state machine.
	StepFunctions interface{} `json:"stepFunctions" yaml:"stepFunctions"`
	// Writes attributes from an MQTT message.
	Timestream interface{} `json:"timestream" yaml:"timestream"`
}

// An asset property timestamp entry containing the following information.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   assetPropertyTimestampProperty := &assetPropertyTimestampProperty{
//   	timeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	offsetInNanos: jsii.String("offsetInNanos"),
//   }
//
type CfnTopicRule_AssetPropertyTimestampProperty struct {
	// A string that contains the time in seconds since epoch.
	//
	// Accepts substitution templates.
	TimeInSeconds *string `json:"timeInSeconds" yaml:"timeInSeconds"`
	// Optional.
	//
	// A string that contains the nanosecond time offset. Accepts substitution templates.
	OffsetInNanos *string `json:"offsetInNanos" yaml:"offsetInNanos"`
}

// An asset property value entry containing the following information.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   assetPropertyValueProperty := &assetPropertyValueProperty{
//   	timestamp: &assetPropertyTimestampProperty{
//   		timeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		offsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   	value: &assetPropertyVariantProperty{
//   		booleanValue: jsii.String("booleanValue"),
//   		doubleValue: jsii.String("doubleValue"),
//   		integerValue: jsii.String("integerValue"),
//   		stringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	quality: jsii.String("quality"),
//   }
//
type CfnTopicRule_AssetPropertyValueProperty struct {
	// The asset property value timestamp.
	Timestamp interface{} `json:"timestamp" yaml:"timestamp"`
	// The value of the asset property.
	Value interface{} `json:"value" yaml:"value"`
	// Optional.
	//
	// A string that describes the quality of the value. Accepts substitution templates. Must be `GOOD` , `BAD` , or `UNCERTAIN` .
	Quality *string `json:"quality" yaml:"quality"`
}

// Contains an asset property value (of a single type).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   assetPropertyVariantProperty := &assetPropertyVariantProperty{
//   	booleanValue: jsii.String("booleanValue"),
//   	doubleValue: jsii.String("doubleValue"),
//   	integerValue: jsii.String("integerValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnTopicRule_AssetPropertyVariantProperty struct {
	// Optional.
	//
	// A string that contains the boolean value ( `true` or `false` ) of the value entry. Accepts substitution templates.
	BooleanValue *string `json:"booleanValue" yaml:"booleanValue"`
	// Optional.
	//
	// A string that contains the double value of the value entry. Accepts substitution templates.
	DoubleValue *string `json:"doubleValue" yaml:"doubleValue"`
	// Optional.
	//
	// A string that contains the integer value of the value entry. Accepts substitution templates.
	IntegerValue *string `json:"integerValue" yaml:"integerValue"`
	// Optional.
	//
	// The string value of the value entry. Accepts substitution templates.
	StringValue *string `json:"stringValue" yaml:"stringValue"`
}

// Describes an action that updates a CloudWatch alarm.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cloudwatchAlarmActionProperty := &cloudwatchAlarmActionProperty{
//   	alarmName: jsii.String("alarmName"),
//   	roleArn: jsii.String("roleArn"),
//   	stateReason: jsii.String("stateReason"),
//   	stateValue: jsii.String("stateValue"),
//   }
//
type CfnTopicRule_CloudwatchAlarmActionProperty struct {
	// The CloudWatch alarm name.
	AlarmName *string `json:"alarmName" yaml:"alarmName"`
	// The IAM role that allows access to the CloudWatch alarm.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The reason for the alarm change.
	StateReason *string `json:"stateReason" yaml:"stateReason"`
	// The value of the alarm state.
	//
	// Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	StateValue *string `json:"stateValue" yaml:"stateValue"`
}

// Describes an action that updates a CloudWatch log.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cloudwatchLogsActionProperty := &cloudwatchLogsActionProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnTopicRule_CloudwatchLogsActionProperty struct {
	// The CloudWatch log name.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The IAM role that allows access to the CloudWatch log.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Describes an action that captures a CloudWatch metric.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cloudwatchMetricActionProperty := &cloudwatchMetricActionProperty{
//   	metricName: jsii.String("metricName"),
//   	metricNamespace: jsii.String("metricNamespace"),
//   	metricUnit: jsii.String("metricUnit"),
//   	metricValue: jsii.String("metricValue"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	metricTimestamp: jsii.String("metricTimestamp"),
//   }
//
type CfnTopicRule_CloudwatchMetricActionProperty struct {
	// The CloudWatch metric name.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The CloudWatch metric namespace name.
	MetricNamespace *string `json:"metricNamespace" yaml:"metricNamespace"`
	// The [metric unit](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.
	MetricUnit *string `json:"metricUnit" yaml:"metricUnit"`
	// The CloudWatch metric value.
	MetricValue *string `json:"metricValue" yaml:"metricValue"`
	// The IAM role that allows access to the CloudWatch metric.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// An optional [Unix timestamp](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp) .
	MetricTimestamp *string `json:"metricTimestamp" yaml:"metricTimestamp"`
}

// Describes an action to write to a DynamoDB table.
//
// The `tableName` , `hashKeyField` , and `rangeKeyField` values must match the values used when you created the table.
//
// The `hashKeyValue` and `rangeKeyvalue` fields use a substitution template syntax. These templates provide data at runtime. The syntax is as follows: ${ *sql-expression* }.
//
// You can specify any valid expression in a WHERE or SELECT clause, including JSON properties, comparisons, calculations, and functions. For example, the following field uses the third level of the topic:
//
// `"hashKeyValue": "${topic(3)}"`
//
// The following field uses the timestamp:
//
// `"rangeKeyValue": "${timestamp()}"`
//
// For more information, see [DynamoDBv2 Action](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rule-actions.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   dynamoDBActionProperty := &dynamoDBActionProperty{
//   	hashKeyField: jsii.String("hashKeyField"),
//   	hashKeyValue: jsii.String("hashKeyValue"),
//   	roleArn: jsii.String("roleArn"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	hashKeyType: jsii.String("hashKeyType"),
//   	payloadField: jsii.String("payloadField"),
//   	rangeKeyField: jsii.String("rangeKeyField"),
//   	rangeKeyType: jsii.String("rangeKeyType"),
//   	rangeKeyValue: jsii.String("rangeKeyValue"),
//   }
//
type CfnTopicRule_DynamoDBActionProperty struct {
	// The hash key name.
	HashKeyField *string `json:"hashKeyField" yaml:"hashKeyField"`
	// The hash key value.
	HashKeyValue *string `json:"hashKeyValue" yaml:"hashKeyValue"`
	// The ARN of the IAM role that grants access to the DynamoDB table.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The name of the DynamoDB table.
	TableName *string `json:"tableName" yaml:"tableName"`
	// The hash key type.
	//
	// Valid values are "STRING" or "NUMBER".
	HashKeyType *string `json:"hashKeyType" yaml:"hashKeyType"`
	// The action payload.
	//
	// This name can be customized.
	PayloadField *string `json:"payloadField" yaml:"payloadField"`
	// The range key name.
	RangeKeyField *string `json:"rangeKeyField" yaml:"rangeKeyField"`
	// The range key type.
	//
	// Valid values are "STRING" or "NUMBER".
	RangeKeyType *string `json:"rangeKeyType" yaml:"rangeKeyType"`
	// The range key value.
	RangeKeyValue *string `json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

// Describes an action to write to a DynamoDB table.
//
// This DynamoDB action writes each attribute in the message payload into it's own column in the DynamoDB table.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   dynamoDBv2ActionProperty := &dynamoDBv2ActionProperty{
//   	putItem: &putItemInputProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnTopicRule_DynamoDBv2ActionProperty struct {
	// Specifies the DynamoDB table to which the message data will be written. For example:.
	//
	// `{ "dynamoDBv2": { "roleArn": "aws:iam:12341251:my-role" "putItem": { "tableName": "my-table" } } }`
	//
	// Each attribute in the message payload will be written to a separate column in the DynamoDB database.
	PutItem interface{} `json:"putItem" yaml:"putItem"`
	// The ARN of the IAM role that grants access to the DynamoDB table.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Describes an action that writes data to an Amazon OpenSearch Service domain.
//
// > The `Elasticsearch` action can only be used by existing rule actions. To create a new rule action or to update an existing rule action, use the `OpenSearch` rule action instead. For more information, see [OpenSearchAction](https://docs.aws.amazon.com//iot/latest/apireference/API_OpenSearchAction.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   elasticsearchActionProperty := &elasticsearchActionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	id: jsii.String("id"),
//   	index: jsii.String("index"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnTopicRule_ElasticsearchActionProperty struct {
	// The endpoint of your OpenSearch domain.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The unique identifier for the document you are storing.
	Id *string `json:"id" yaml:"id"`
	// The index where you want to store your data.
	Index *string `json:"index" yaml:"index"`
	// The IAM role ARN that has access to OpenSearch.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The type of document you are storing.
	Type *string `json:"type" yaml:"type"`
}

// Describes an action that writes data to an Amazon Kinesis Firehose stream.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   firehoseActionProperty := &firehoseActionProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	batchMode: jsii.Boolean(false),
//   	separator: jsii.String("separator"),
//   }
//
type CfnTopicRule_FirehoseActionProperty struct {
	// The delivery stream name.
	DeliveryStreamName *string `json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The IAM role that grants access to the Amazon Kinesis Firehose stream.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using [`PutRecordBatch`](https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html) . The default value is `false` .
	//
	// When `batchMode` is `true` and the rule's SQL statement evaluates to an Array, each Array element forms one record in the [`PutRecordBatch`](https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html) request. The resulting array can't have more than 500 records.
	BatchMode interface{} `json:"batchMode" yaml:"batchMode"`
	// A character separator that will be used to separate records written to the Firehose stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator *string `json:"separator" yaml:"separator"`
}

// The HTTP action header.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   httpActionHeaderProperty := &httpActionHeaderProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_HttpActionHeaderProperty struct {
	// The HTTP header key.
	Key *string `json:"key" yaml:"key"`
	// The HTTP header value.
	//
	// Substitution templates are supported.
	Value *string `json:"value" yaml:"value"`
}

// Send data to an HTTPS endpoint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   httpActionProperty := &httpActionProperty{
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	auth: &httpAuthorizationProperty{
//   		sigv4: &sigV4AuthorizationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			serviceName: jsii.String("serviceName"),
//   			signingRegion: jsii.String("signingRegion"),
//   		},
//   	},
//   	confirmationUrl: jsii.String("confirmationUrl"),
//   	headers: []interface{}{
//   		&httpActionHeaderProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_HttpActionProperty struct {
	// The endpoint URL.
	//
	// If substitution templates are used in the URL, you must also specify a `confirmationUrl` . If this is a new destination, a new `TopicRuleDestination` is created if possible.
	Url *string `json:"url" yaml:"url"`
	// The authentication method to use when sending data to an HTTPS endpoint.
	Auth interface{} `json:"auth" yaml:"auth"`
	// The URL to which AWS IoT sends a confirmation message.
	//
	// The value of the confirmation URL must be a prefix of the endpoint URL. If you do not specify a confirmation URL AWS IoT uses the endpoint URL as the confirmation URL. If you use substitution templates in the confirmationUrl, you must create and enable topic rule destinations that match each possible value of the substitution template before traffic is allowed to your endpoint URL.
	ConfirmationUrl *string `json:"confirmationUrl" yaml:"confirmationUrl"`
	// The HTTP headers to send with the message data.
	Headers interface{} `json:"headers" yaml:"headers"`
}

// The authorization method used to send messages.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   httpAuthorizationProperty := &httpAuthorizationProperty{
//   	sigv4: &sigV4AuthorizationProperty{
//   		roleArn: jsii.String("roleArn"),
//   		serviceName: jsii.String("serviceName"),
//   		signingRegion: jsii.String("signingRegion"),
//   	},
//   }
//
type CfnTopicRule_HttpAuthorizationProperty struct {
	// Use Sig V4 authorization.
	//
	// For more information, see [Signature Version 4 Signing Process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
	Sigv4 interface{} `json:"sigv4" yaml:"sigv4"`
}

// Sends message data to an AWS IoT Analytics channel.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   iotAnalyticsActionProperty := &iotAnalyticsActionProperty{
//   	channelName: jsii.String("channelName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	batchMode: jsii.Boolean(false),
//   }
//
type CfnTopicRule_IotAnalyticsActionProperty struct {
	// The name of the IoT Analytics channel to which message data will be sent.
	ChannelName *string `json:"channelName" yaml:"channelName"`
	// The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Whether to process the action as a batch. The default value is `false` .
	//
	// When `batchMode` is `true` and the rule SQL statement evaluates to an Array, each Array element is delivered as a separate message when passed by [`BatchPutMessage`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_BatchPutMessage.html) The resulting array can't have more than 100 messages.
	BatchMode interface{} `json:"batchMode" yaml:"batchMode"`
}

// Sends an input to an AWS IoT Events detector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   iotEventsActionProperty := &iotEventsActionProperty{
//   	inputName: jsii.String("inputName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	batchMode: jsii.Boolean(false),
//   	messageId: jsii.String("messageId"),
//   }
//
type CfnTopicRule_IotEventsActionProperty struct {
	// The name of the AWS IoT Events input.
	InputName *string `json:"inputName" yaml:"inputName"`
	// The ARN of the role that grants AWS IoT permission to send an input to an AWS IoT Events detector.
	//
	// ("Action":"iotevents:BatchPutMessage").
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Whether to process the event actions as a batch. The default value is `false` .
	//
	// When `batchMode` is `true` , you can't specify a `messageId` .
	//
	// When `batchMode` is `true` and the rule SQL statement evaluates to an Array, each Array element is treated as a separate message when Events by calling [`BatchPutMessage`](https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchPutMessage.html) . The resulting array can't have more than 10 messages.
	BatchMode interface{} `json:"batchMode" yaml:"batchMode"`
	// The ID of the message. The default `messageId` is a new UUID value.
	//
	// When `batchMode` is `true` , you can't specify a `messageId` --a new UUID value will be assigned.
	//
	// Assign a value to this property to ensure that only one input (message) with a given `messageId` will be processed by an AWS IoT Events detector.
	MessageId *string `json:"messageId" yaml:"messageId"`
}

// Describes an action to send data from an MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   iotSiteWiseActionProperty := &iotSiteWiseActionProperty{
//   	putAssetPropertyValueEntries: []interface{}{
//   		&putAssetPropertyValueEntryProperty{
//   			propertyValues: []interface{}{
//   				&assetPropertyValueProperty{
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			assetId: jsii.String("assetId"),
//   			entryId: jsii.String("entryId"),
//   			propertyAlias: jsii.String("propertyAlias"),
//   			propertyId: jsii.String("propertyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnTopicRule_IotSiteWiseActionProperty struct {
	// A list of asset property value entries.
	PutAssetPropertyValueEntries interface{} `json:"putAssetPropertyValueEntries" yaml:"putAssetPropertyValueEntries"`
	// The ARN of the role that grants AWS IoT permission to send an asset property value to AWS IoT SiteWise.
	//
	// ( `"Action": "iotsitewise:BatchPutAssetPropertyValue"` ). The trust policy can restrict access to specific asset hierarchy paths.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   kafkaActionProperty := &kafkaActionProperty{
//   	clientProperties: map[string]*string{
//   		"clientPropertiesKey": jsii.String("clientProperties"),
//   	},
//   	destinationArn: jsii.String("destinationArn"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   	partition: jsii.String("partition"),
//   }
//
type CfnTopicRule_KafkaActionProperty struct {
	// Properties of the Apache Kafka producer client.
	ClientProperties interface{} `json:"clientProperties" yaml:"clientProperties"`
	// The ARN of Kafka action's VPC `TopicRuleDestination` .
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// The Kafka topic for messages to be sent to the Kafka broker.
	Topic *string `json:"topic" yaml:"topic"`
	// The Kafka message key.
	Key *string `json:"key" yaml:"key"`
	// The Kafka message partition.
	Partition *string `json:"partition" yaml:"partition"`
}

// Describes an action to write data to an Amazon Kinesis stream.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   kinesisActionProperty := &kinesisActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	streamName: jsii.String("streamName"),
//
//   	// the properties below are optional
//   	partitionKey: jsii.String("partitionKey"),
//   }
//
type CfnTopicRule_KinesisActionProperty struct {
	// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The name of the Amazon Kinesis stream.
	StreamName *string `json:"streamName" yaml:"streamName"`
	// The partition key.
	PartitionKey *string `json:"partitionKey" yaml:"partitionKey"`
}

// Describes an action to invoke a Lambda function.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   lambdaActionProperty := &lambdaActionProperty{
//   	functionArn: jsii.String("functionArn"),
//   }
//
type CfnTopicRule_LambdaActionProperty struct {
	// The ARN of the Lambda function.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
}

// Describes an action that writes data to an Amazon OpenSearch Service domain.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   openSearchActionProperty := &openSearchActionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	id: jsii.String("id"),
//   	index: jsii.String("index"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnTopicRule_OpenSearchActionProperty struct {
	// The endpoint of your OpenSearch domain.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The unique identifier for the document you are storing.
	Id *string `json:"id" yaml:"id"`
	// The OpenSearch index where you want to store your data.
	Index *string `json:"index" yaml:"index"`
	// The IAM role ARN that has access to OpenSearch.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The type of document you are storing.
	Type *string `json:"type" yaml:"type"`
}

// An asset property value entry containing the following information.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   putAssetPropertyValueEntryProperty := &putAssetPropertyValueEntryProperty{
//   	propertyValues: []interface{}{
//   		&assetPropertyValueProperty{
//   			timestamp: &assetPropertyTimestampProperty{
//   				timeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				offsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   			value: &assetPropertyVariantProperty{
//   				booleanValue: jsii.String("booleanValue"),
//   				doubleValue: jsii.String("doubleValue"),
//   				integerValue: jsii.String("integerValue"),
//   				stringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			quality: jsii.String("quality"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	assetId: jsii.String("assetId"),
//   	entryId: jsii.String("entryId"),
//   	propertyAlias: jsii.String("propertyAlias"),
//   	propertyId: jsii.String("propertyId"),
//   }
//
type CfnTopicRule_PutAssetPropertyValueEntryProperty struct {
	// A list of property values to insert that each contain timestamp, quality, and value (TQV) information.
	PropertyValues interface{} `json:"propertyValues" yaml:"propertyValues"`
	// The ID of the AWS IoT SiteWise asset.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	AssetId *string `json:"assetId" yaml:"assetId"`
	// Optional.
	//
	// A unique identifier for this entry that you can define to better track which message caused an error in case of failure. Accepts substitution templates. Defaults to a new UUID.
	EntryId *string `json:"entryId" yaml:"entryId"`
	// The name of the property alias associated with your asset property.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	PropertyAlias *string `json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset's property.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	PropertyId *string `json:"propertyId" yaml:"propertyId"`
}

// The input for the DynamoActionVS action that specifies the DynamoDB table to which the message data will be written.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   putItemInputProperty := &putItemInputProperty{
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnTopicRule_PutItemInputProperty struct {
	// The table where the message data will be written.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// Describes an action to republish to another topic.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   republishActionProperty := &republishActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	qos: jsii.Number(123),
//   }
//
type CfnTopicRule_RepublishActionProperty struct {
	// The ARN of the IAM role that grants access.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The name of the MQTT topic.
	Topic *string `json:"topic" yaml:"topic"`
	// The Quality of Service (QoS) level to use when republishing messages.
	//
	// The default value is 0.
	Qos *float64 `json:"qos" yaml:"qos"`
}

// Describes an action to write data to an Amazon S3 bucket.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   s3ActionProperty := &s3ActionProperty{
//   	bucketName: jsii.String("bucketName"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnTopicRule_S3ActionProperty struct {
	// The Amazon S3 bucket.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// The object key.
	//
	// For more information, see [Actions, resources, and condition keys for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html) .
	Key *string `json:"key" yaml:"key"`
	// The ARN of the IAM role that grants access.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The Amazon S3 canned ACL that controls access to the object identified by the object key.
	//
	// For more information, see [S3 canned ACLs](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) .
	CannedAcl *string `json:"cannedAcl" yaml:"cannedAcl"`
}

// For more information, see [Signature Version 4 signing process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   sigV4AuthorizationProperty := &sigV4AuthorizationProperty{
//   	roleArn: jsii.String("roleArn"),
//   	serviceName: jsii.String("serviceName"),
//   	signingRegion: jsii.String("signingRegion"),
//   }
//
type CfnTopicRule_SigV4AuthorizationProperty struct {
	// The ARN of the signing role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The service name to use while signing with Sig V4.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
	// The signing region.
	SigningRegion *string `json:"signingRegion" yaml:"signingRegion"`
}

// Describes an action to publish to an Amazon SNS topic.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   snsActionProperty := &snsActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	messageFormat: jsii.String("messageFormat"),
//   }
//
type CfnTopicRule_SnsActionProperty struct {
	// The ARN of the IAM role that grants access.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The ARN of the SNS topic.
	TargetArn *string `json:"targetArn" yaml:"targetArn"`
	// (Optional) The message format of the message to publish.
	//
	// Accepted values are "JSON" and "RAW". The default value of the attribute is "RAW". SNS uses this setting to determine if the payload should be parsed and relevant platform-specific bits of the payload should be extracted. For more information, see [Amazon SNS Message and JSON Formats](https://docs.aws.amazon.com/sns/latest/dg/json-formats.html) in the *Amazon Simple Notification Service Developer Guide* .
	MessageFormat *string `json:"messageFormat" yaml:"messageFormat"`
}

// Describes an action to publish data to an Amazon SQS queue.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   sqsActionProperty := &sqsActionProperty{
//   	queueUrl: jsii.String("queueUrl"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	useBase64: jsii.Boolean(false),
//   }
//
type CfnTopicRule_SqsActionProperty struct {
	// The URL of the Amazon SQS queue.
	QueueUrl *string `json:"queueUrl" yaml:"queueUrl"`
	// The ARN of the IAM role that grants access.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Specifies whether to use Base64 encoding.
	UseBase64 interface{} `json:"useBase64" yaml:"useBase64"`
}

// Starts execution of a Step Functions state machine.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   stepFunctionsActionProperty := &stepFunctionsActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	stateMachineName: jsii.String("stateMachineName"),
//
//   	// the properties below are optional
//   	executionNamePrefix: jsii.String("executionNamePrefix"),
//   }
//
type CfnTopicRule_StepFunctionsActionProperty struct {
	// The ARN of the role that grants IoT permission to start execution of a state machine ("Action":"states:StartExecution").
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The name of the Step Functions state machine whose execution will be started.
	StateMachineName *string `json:"stateMachineName" yaml:"stateMachineName"`
	// (Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID.
	//
	// Step Functions automatically creates a unique name for each state machine execution if one is not provided.
	ExecutionNamePrefix *string `json:"executionNamePrefix" yaml:"executionNamePrefix"`
}

// Describes an action that writes records into an Amazon Timestream table.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   timestreamActionProperty := &timestreamActionProperty{
//   	databaseName: jsii.String("databaseName"),
//   	dimensions: []interface{}{
//   		&timestreamDimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	batchMode: jsii.Boolean(false),
//   	timestamp: &timestreamTimestampProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnTopicRule_TimestreamActionProperty struct {
	// The name of an Amazon Timestream database that has the table to write records into.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Metadata attributes of the time series that are written in each measure record.
	Dimensions interface{} `json:"dimensions" yaml:"dimensions"`
	// The Amazon Resource Name (ARN) of the role that grants AWS IoT permission to write to the Timestream database table.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The table where the message data will be written.
	TableName *string `json:"tableName" yaml:"tableName"`
	// Whether to process the action as a batch.
	BatchMode interface{} `json:"batchMode" yaml:"batchMode"`
	// The value to use for the entry's timestamp.
	//
	// If blank, the time that the entry was processed is used.
	Timestamp interface{} `json:"timestamp" yaml:"timestamp"`
}

// Metadata attributes of the time series that are written in each measure record.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   timestreamDimensionProperty := &timestreamDimensionProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_TimestreamDimensionProperty struct {
	// The metadata dimension name.
	//
	// This is the name of the column in the Amazon Timestream database table record.
	Name *string `json:"name" yaml:"name"`
	// The value to write in this column of the database record.
	Value *string `json:"value" yaml:"value"`
}

// The value to use for the entry's timestamp.
//
// If blank, the time that the entry was processed is used.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   timestreamTimestampProperty := &timestreamTimestampProperty{
//   	unit: jsii.String("unit"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_TimestreamTimestampProperty struct {
	// The precision of the timestamp value that results from the expression described in `value` .
	Unit *string `json:"unit" yaml:"unit"`
	// An expression that returns a long epoch time value.
	Value *string `json:"value" yaml:"value"`
}

// Describes a rule.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   topicRulePayloadProperty := &topicRulePayloadProperty{
//   	actions: []interface{}{
//   		&actionProperty{
//   			cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   				alarmName: jsii.String("alarmName"),
//   				roleArn: jsii.String("roleArn"),
//   				stateReason: jsii.String("stateReason"),
//   				stateValue: jsii.String("stateValue"),
//   			},
//   			cloudwatchLogs: &cloudwatchLogsActionProperty{
//   				logGroupName: jsii.String("logGroupName"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			cloudwatchMetric: &cloudwatchMetricActionProperty{
//   				metricName: jsii.String("metricName"),
//   				metricNamespace: jsii.String("metricNamespace"),
//   				metricUnit: jsii.String("metricUnit"),
//   				metricValue: jsii.String("metricValue"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				metricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			dynamoDb: &dynamoDBActionProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2ActionProperty{
//   				putItem: &putItemInputProperty{
//   					tableName: jsii.String("tableName"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			elasticsearch: &elasticsearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			firehose: &firehoseActionProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				separator: jsii.String("separator"),
//   			},
//   			http: &httpActionProperty{
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				auth: &httpAuthorizationProperty{
//   					sigv4: &sigV4AuthorizationProperty{
//   						roleArn: jsii.String("roleArn"),
//   						serviceName: jsii.String("serviceName"),
//   						signingRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				confirmationUrl: jsii.String("confirmationUrl"),
//   				headers: []interface{}{
//   					&httpActionHeaderProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			iotAnalytics: &iotAnalyticsActionProperty{
//   				channelName: jsii.String("channelName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			iotEvents: &iotEventsActionProperty{
//   				inputName: jsii.String("inputName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				messageId: jsii.String("messageId"),
//   			},
//   			iotSiteWise: &iotSiteWiseActionProperty{
//   				putAssetPropertyValueEntries: []interface{}{
//   					&putAssetPropertyValueEntryProperty{
//   						propertyValues: []interface{}{
//   							&assetPropertyValueProperty{
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kafka: &kafkaActionProperty{
//   				clientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				destinationArn: jsii.String("destinationArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				partition: jsii.String("partition"),
//   			},
//   			kinesis: &kinesisActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				partitionKey: jsii.String("partitionKey"),
//   			},
//   			lambda: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//   			},
//   			openSearch: &openSearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			republish: &republishActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				qos: jsii.Number(123),
//   			},
//   			s3: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//   				key: jsii.String("key"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			sns: &snsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				messageFormat: jsii.String("messageFormat"),
//   			},
//   			sqs: &sqsActionProperty{
//   				queueUrl: jsii.String("queueUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				useBase64: jsii.Boolean(false),
//   			},
//   			stepFunctions: &stepFunctionsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				stateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				executionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			timestream: &timestreamActionProperty{
//   				databaseName: jsii.String("databaseName"),
//   				dimensions: []interface{}{
//   					&timestreamDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				timestamp: &timestreamTimestampProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	description: jsii.String("description"),
//   	errorAction: &actionProperty{
//   		cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   			alarmName: jsii.String("alarmName"),
//   			roleArn: jsii.String("roleArn"),
//   			stateReason: jsii.String("stateReason"),
//   			stateValue: jsii.String("stateValue"),
//   		},
//   		cloudwatchLogs: &cloudwatchLogsActionProperty{
//   			logGroupName: jsii.String("logGroupName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		cloudwatchMetric: &cloudwatchMetricActionProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricUnit: jsii.String("metricUnit"),
//   			metricValue: jsii.String("metricValue"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			metricTimestamp: jsii.String("metricTimestamp"),
//   		},
//   		dynamoDb: &dynamoDBActionProperty{
//   			hashKeyField: jsii.String("hashKeyField"),
//   			hashKeyValue: jsii.String("hashKeyValue"),
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			hashKeyType: jsii.String("hashKeyType"),
//   			payloadField: jsii.String("payloadField"),
//   			rangeKeyField: jsii.String("rangeKeyField"),
//   			rangeKeyType: jsii.String("rangeKeyType"),
//   			rangeKeyValue: jsii.String("rangeKeyValue"),
//   		},
//   		dynamoDBv2: &dynamoDBv2ActionProperty{
//   			putItem: &putItemInputProperty{
//   				tableName: jsii.String("tableName"),
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		elasticsearch: &elasticsearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		firehose: &firehoseActionProperty{
//   			deliveryStreamName: jsii.String("deliveryStreamName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			separator: jsii.String("separator"),
//   		},
//   		http: &httpActionProperty{
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			auth: &httpAuthorizationProperty{
//   				sigv4: &sigV4AuthorizationProperty{
//   					roleArn: jsii.String("roleArn"),
//   					serviceName: jsii.String("serviceName"),
//   					signingRegion: jsii.String("signingRegion"),
//   				},
//   			},
//   			confirmationUrl: jsii.String("confirmationUrl"),
//   			headers: []interface{}{
//   				&httpActionHeaderProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		iotAnalytics: &iotAnalyticsActionProperty{
//   			channelName: jsii.String("channelName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   		},
//   		iotEvents: &iotEventsActionProperty{
//   			inputName: jsii.String("inputName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			messageId: jsii.String("messageId"),
//   		},
//   		iotSiteWise: &iotSiteWiseActionProperty{
//   			putAssetPropertyValueEntries: []interface{}{
//   				&putAssetPropertyValueEntryProperty{
//   					propertyValues: []interface{}{
//   						&assetPropertyValueProperty{
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kafka: &kafkaActionProperty{
//   			clientProperties: map[string]*string{
//   				"clientPropertiesKey": jsii.String("clientProperties"),
//   			},
//   			destinationArn: jsii.String("destinationArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   			partition: jsii.String("partition"),
//   		},
//   		kinesis: &kinesisActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			streamName: jsii.String("streamName"),
//
//   			// the properties below are optional
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		lambda: &lambdaActionProperty{
//   			functionArn: jsii.String("functionArn"),
//   		},
//   		openSearch: &openSearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		republish: &republishActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			qos: jsii.Number(123),
//   		},
//   		s3: &s3ActionProperty{
//   			bucketName: jsii.String("bucketName"),
//   			key: jsii.String("key"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   		sns: &snsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//
//   			// the properties below are optional
//   			messageFormat: jsii.String("messageFormat"),
//   		},
//   		sqs: &sqsActionProperty{
//   			queueUrl: jsii.String("queueUrl"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			useBase64: jsii.Boolean(false),
//   		},
//   		stepFunctions: &stepFunctionsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			stateMachineName: jsii.String("stateMachineName"),
//
//   			// the properties below are optional
//   			executionNamePrefix: jsii.String("executionNamePrefix"),
//   		},
//   		timestream: &timestreamActionProperty{
//   			databaseName: jsii.String("databaseName"),
//   			dimensions: []interface{}{
//   				&timestreamDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			timestamp: &timestreamTimestampProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ruleDisabled: jsii.Boolean(false),
//   }
//
type CfnTopicRule_TopicRulePayloadProperty struct {
	// The actions associated with the rule.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The SQL statement used to query the topic.
	//
	// For more information, see [AWS IoT SQL Reference](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) in the *AWS IoT Developer Guide* .
	Sql *string `json:"sql" yaml:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	//
	// The default value is 2015-10-08.
	AwsIotSqlVersion *string `json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The description of the rule.
	Description *string `json:"description" yaml:"description"`
	// The action to take when an error occurs.
	ErrorAction interface{} `json:"errorAction" yaml:"errorAction"`
	// Specifies whether the rule is disabled.
	RuleDisabled interface{} `json:"ruleDisabled" yaml:"ruleDisabled"`
}

// A CloudFormation `AWS::IoT::TopicRuleDestination`.
//
// A topic rule destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnTopicRuleDestination := iot.NewCfnTopicRuleDestination(this, jsii.String("MyCfnTopicRuleDestination"), &cfnTopicRuleDestinationProps{
//   	httpUrlProperties: &httpUrlDestinationSummaryProperty{
//   		confirmationUrl: jsii.String("confirmationUrl"),
//   	},
//   	status: jsii.String("status"),
//   	vpcProperties: &vpcDestinationPropertiesProperty{
//   		roleArn: jsii.String("roleArn"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcId: jsii.String("vpcId"),
//   	},
//   })
//
type CfnTopicRuleDestination interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The topic rule destination URL.
	AttrArn() *string
	// Additional details or reason why the topic rule destination is in the current status.
	AttrStatusReason() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Properties of the HTTP URL.
	HttpUrlProperties() interface{}
	SetHttpUrlProperties(val interface{})
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// - **IN_PROGRESS** - A topic rule destination was created but has not been confirmed.
	//
	// You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	// - **ENABLED** - Confirmation was completed, and traffic to this destination is allowed. You can set status to `DISABLED` by calling `UpdateTopicRuleDestination` .
	// - **DISABLED** - Confirmation was completed, and traffic to this destination is not allowed. You can set status to `ENABLED` by calling `UpdateTopicRuleDestination` .
	// - **ERROR** - Confirmation could not be completed; for example, if the confirmation timed out. You can call `GetTopicRuleDestination` for details about the error. You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	Status() *string
	SetStatus(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Properties of the virtual private cloud (VPC) connection.
	VpcProperties() interface{}
	SetVpcProperties(val interface{})
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnTopicRuleDestination
type jsiiProxy_CfnTopicRuleDestination struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopicRuleDestination) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) AttrStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) HttpUrlProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpUrlProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRuleDestination) VpcProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoT::TopicRuleDestination`.
func NewCfnTopicRuleDestination(scope constructs.Construct, id *string, props *CfnTopicRuleDestinationProps) CfnTopicRuleDestination {
	_init_.Initialize()

	j := jsiiProxy_CfnTopicRuleDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoT::TopicRuleDestination`.
func NewCfnTopicRuleDestination_Override(c CfnTopicRuleDestination, scope constructs.Construct, id *string, props *CfnTopicRuleDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopicRuleDestination) SetHttpUrlProperties(val interface{}) {
	_jsii_.Set(
		j,
		"httpUrlProperties",
		val,
	)
}

func (j *jsiiProxy_CfnTopicRuleDestination) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnTopicRuleDestination) SetVpcProperties(val interface{}) {
	_jsii_.Set(
		j,
		"vpcProperties",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTopicRuleDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTopicRuleDestination_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnTopicRuleDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicRuleDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnTopicRuleDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopicRuleDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopicRuleDestination) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// HTTP URL destination properties.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   httpUrlDestinationSummaryProperty := &httpUrlDestinationSummaryProperty{
//   	confirmationUrl: jsii.String("confirmationUrl"),
//   }
//
type CfnTopicRuleDestination_HttpUrlDestinationSummaryProperty struct {
	// The URL used to confirm the HTTP topic rule destination URL.
	ConfirmationUrl *string `json:"confirmationUrl" yaml:"confirmationUrl"`
}

// The properties of a virtual private cloud (VPC) destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   vpcDestinationPropertiesProperty := &vpcDestinationPropertiesProperty{
//   	roleArn: jsii.String("roleArn"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnTopicRuleDestination_VpcDestinationPropertiesProperty struct {
	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The security groups of the VPC destination.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// The subnet IDs of the VPC destination.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Properties for defining a `CfnTopicRuleDestination`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnTopicRuleDestinationProps := &cfnTopicRuleDestinationProps{
//   	httpUrlProperties: &httpUrlDestinationSummaryProperty{
//   		confirmationUrl: jsii.String("confirmationUrl"),
//   	},
//   	status: jsii.String("status"),
//   	vpcProperties: &vpcDestinationPropertiesProperty{
//   		roleArn: jsii.String("roleArn"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnTopicRuleDestinationProps struct {
	// Properties of the HTTP URL.
	HttpUrlProperties interface{} `json:"httpUrlProperties" yaml:"httpUrlProperties"`
	// - **IN_PROGRESS** - A topic rule destination was created but has not been confirmed.
	//
	// You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	// - **ENABLED** - Confirmation was completed, and traffic to this destination is allowed. You can set status to `DISABLED` by calling `UpdateTopicRuleDestination` .
	// - **DISABLED** - Confirmation was completed, and traffic to this destination is not allowed. You can set status to `ENABLED` by calling `UpdateTopicRuleDestination` .
	// - **ERROR** - Confirmation could not be completed; for example, if the confirmation timed out. You can call `GetTopicRuleDestination` for details about the error. You can set status to `IN_PROGRESS` by calling `UpdateTopicRuleDestination` . Calling `UpdateTopicRuleDestination` causes a new confirmation challenge to be sent to your confirmation endpoint.
	Status *string `json:"status" yaml:"status"`
	// Properties of the virtual private cloud (VPC) connection.
	VpcProperties interface{} `json:"vpcProperties" yaml:"vpcProperties"`
}

// Properties for defining a `CfnTopicRule`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iot "github.com/aws/aws-cdk-go/awscdk/aws_iot"
//   cfnTopicRuleProps := &cfnTopicRuleProps{
//   	topicRulePayload: &topicRulePayloadProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   					alarmName: jsii.String("alarmName"),
//   					roleArn: jsii.String("roleArn"),
//   					stateReason: jsii.String("stateReason"),
//   					stateValue: jsii.String("stateValue"),
//   				},
//   				cloudwatchLogs: &cloudwatchLogsActionProperty{
//   					logGroupName: jsii.String("logGroupName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchMetric: &cloudwatchMetricActionProperty{
//   					metricName: jsii.String("metricName"),
//   					metricNamespace: jsii.String("metricNamespace"),
//   					metricUnit: jsii.String("metricUnit"),
//   					metricValue: jsii.String("metricValue"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					metricTimestamp: jsii.String("metricTimestamp"),
//   				},
//   				dynamoDb: &dynamoDBActionProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2ActionProperty{
//   					putItem: &putItemInputProperty{
//   						tableName: jsii.String("tableName"),
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				elasticsearch: &elasticsearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				firehose: &firehoseActionProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					separator: jsii.String("separator"),
//   				},
//   				http: &httpActionProperty{
//   					url: jsii.String("url"),
//
//   					// the properties below are optional
//   					auth: &httpAuthorizationProperty{
//   						sigv4: &sigV4AuthorizationProperty{
//   							roleArn: jsii.String("roleArn"),
//   							serviceName: jsii.String("serviceName"),
//   							signingRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					confirmationUrl: jsii.String("confirmationUrl"),
//   					headers: []interface{}{
//   						&httpActionHeaderProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				iotAnalytics: &iotAnalyticsActionProperty{
//   					channelName: jsii.String("channelName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   				},
//   				iotEvents: &iotEventsActionProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					messageId: jsii.String("messageId"),
//   				},
//   				iotSiteWise: &iotSiteWiseActionProperty{
//   					putAssetPropertyValueEntries: []interface{}{
//   						&putAssetPropertyValueEntryProperty{
//   							propertyValues: []interface{}{
//   								&assetPropertyValueProperty{
//   									timestamp: &assetPropertyTimestampProperty{
//   										timeInSeconds: jsii.String("timeInSeconds"),
//
//   										// the properties below are optional
//   										offsetInNanos: jsii.String("offsetInNanos"),
//   									},
//   									value: &assetPropertyVariantProperty{
//   										booleanValue: jsii.String("booleanValue"),
//   										doubleValue: jsii.String("doubleValue"),
//   										integerValue: jsii.String("integerValue"),
//   										stringValue: jsii.String("stringValue"),
//   									},
//
//   									// the properties below are optional
//   									quality: jsii.String("quality"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				kafka: &kafkaActionProperty{
//   					clientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					destinationArn: jsii.String("destinationArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					key: jsii.String("key"),
//   					partition: jsii.String("partition"),
//   				},
//   				kinesis: &kinesisActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					streamName: jsii.String("streamName"),
//
//   					// the properties below are optional
//   					partitionKey: jsii.String("partitionKey"),
//   				},
//   				lambda: &lambdaActionProperty{
//   					functionArn: jsii.String("functionArn"),
//   				},
//   				openSearch: &openSearchActionProperty{
//   					endpoint: jsii.String("endpoint"),
//   					id: jsii.String("id"),
//   					index: jsii.String("index"),
//   					roleArn: jsii.String("roleArn"),
//   					type: jsii.String("type"),
//   				},
//   				republish: &republishActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					qos: jsii.Number(123),
//   				},
//   				s3: &s3ActionProperty{
//   					bucketName: jsii.String("bucketName"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					cannedAcl: jsii.String("cannedAcl"),
//   				},
//   				sns: &snsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					messageFormat: jsii.String("messageFormat"),
//   				},
//   				sqs: &sqsActionProperty{
//   					queueUrl: jsii.String("queueUrl"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					useBase64: jsii.Boolean(false),
//   				},
//   				stepFunctions: &stepFunctionsActionProperty{
//   					roleArn: jsii.String("roleArn"),
//   					stateMachineName: jsii.String("stateMachineName"),
//
//   					// the properties below are optional
//   					executionNamePrefix: jsii.String("executionNamePrefix"),
//   				},
//   				timestream: &timestreamActionProperty{
//   					databaseName: jsii.String("databaseName"),
//   					dimensions: []interface{}{
//   						&timestreamDimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					batchMode: jsii.Boolean(false),
//   					timestamp: &timestreamTimestampProperty{
//   						unit: jsii.String("unit"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		sql: jsii.String("sql"),
//
//   		// the properties below are optional
//   		awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   		description: jsii.String("description"),
//   		errorAction: &actionProperty{
//   			cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   				alarmName: jsii.String("alarmName"),
//   				roleArn: jsii.String("roleArn"),
//   				stateReason: jsii.String("stateReason"),
//   				stateValue: jsii.String("stateValue"),
//   			},
//   			cloudwatchLogs: &cloudwatchLogsActionProperty{
//   				logGroupName: jsii.String("logGroupName"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			cloudwatchMetric: &cloudwatchMetricActionProperty{
//   				metricName: jsii.String("metricName"),
//   				metricNamespace: jsii.String("metricNamespace"),
//   				metricUnit: jsii.String("metricUnit"),
//   				metricValue: jsii.String("metricValue"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				metricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			dynamoDb: &dynamoDBActionProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2ActionProperty{
//   				putItem: &putItemInputProperty{
//   					tableName: jsii.String("tableName"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			elasticsearch: &elasticsearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			firehose: &firehoseActionProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				separator: jsii.String("separator"),
//   			},
//   			http: &httpActionProperty{
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				auth: &httpAuthorizationProperty{
//   					sigv4: &sigV4AuthorizationProperty{
//   						roleArn: jsii.String("roleArn"),
//   						serviceName: jsii.String("serviceName"),
//   						signingRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				confirmationUrl: jsii.String("confirmationUrl"),
//   				headers: []interface{}{
//   					&httpActionHeaderProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			iotAnalytics: &iotAnalyticsActionProperty{
//   				channelName: jsii.String("channelName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			iotEvents: &iotEventsActionProperty{
//   				inputName: jsii.String("inputName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				messageId: jsii.String("messageId"),
//   			},
//   			iotSiteWise: &iotSiteWiseActionProperty{
//   				putAssetPropertyValueEntries: []interface{}{
//   					&putAssetPropertyValueEntryProperty{
//   						propertyValues: []interface{}{
//   							&assetPropertyValueProperty{
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kafka: &kafkaActionProperty{
//   				clientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				destinationArn: jsii.String("destinationArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				partition: jsii.String("partition"),
//   			},
//   			kinesis: &kinesisActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				partitionKey: jsii.String("partitionKey"),
//   			},
//   			lambda: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//   			},
//   			openSearch: &openSearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			republish: &republishActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				qos: jsii.Number(123),
//   			},
//   			s3: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//   				key: jsii.String("key"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			sns: &snsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				messageFormat: jsii.String("messageFormat"),
//   			},
//   			sqs: &sqsActionProperty{
//   				queueUrl: jsii.String("queueUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				useBase64: jsii.Boolean(false),
//   			},
//   			stepFunctions: &stepFunctionsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				stateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				executionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			timestream: &timestreamActionProperty{
//   				databaseName: jsii.String("databaseName"),
//   				dimensions: []interface{}{
//   					&timestreamDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				timestamp: &timestreamTimestampProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		ruleDisabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	ruleName: jsii.String("ruleName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRuleProps struct {
	// The rule payload.
	TopicRulePayload interface{} `json:"topicRulePayload" yaml:"topicRulePayload"`
	// The name of the rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

