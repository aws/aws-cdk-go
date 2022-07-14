package awsopsworkscm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworkscm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::OpsWorksCM::Server`.
//
// The `AWS::OpsWorksCM::Server` resource creates an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise configuration management server. For more information, see [Create a Chef Automate Server in AWS CloudFormation](https://docs.aws.amazon.com/opsworks/latest/userguide/opscm-create-server-cfn.html) or [Create a Puppet Enterprise Master in AWS CloudFormation](https://docs.aws.amazon.com/opsworks/latest/userguide/opspup-create-server-cfn.html) in the *AWS OpsWorks User Guide* , and [CreateServer](https://docs.aws.amazon.com/opsworks-cm/latest/APIReference/API_CreateServer.html) in the *AWS OpsWorks CM API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServer := awscdk.Aws_opsworkscm.NewCfnServer(this, jsii.String("MyCfnServer"), &cfnServerProps{
//   	instanceProfileArn: jsii.String("instanceProfileArn"),
//   	instanceType: jsii.String("instanceType"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	backupId: jsii.String("backupId"),
//   	backupRetentionCount: jsii.Number(123),
//   	customCertificate: jsii.String("customCertificate"),
//   	customDomain: jsii.String("customDomain"),
//   	customPrivateKey: jsii.String("customPrivateKey"),
//   	disableAutomatedBackup: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineAttributes: []interface{}{
//   		&engineAttributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	engineModel: jsii.String("engineModel"),
//   	engineVersion: jsii.String("engineVersion"),
//   	keyPair: jsii.String("keyPair"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	serverName: jsii.String("serverName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnServer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Associate a public IP address with a server that you are launching.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	// The Amazon Resource Name (ARN) of the server, such as `arn:aws:OpsWorksCM:us-east-1:123456789012:server/server-a1bzhi` .
	AttrArn() *string
	// A DNS name that can be used to access the engine.
	//
	// Example: `myserver-asdfghjkl.us-east-1.opsworks.io` .
	AttrEndpoint() *string
	AttrId() *string
	// If you specify this field, AWS OpsWorks CM creates the server by using the backup represented by BackupId.
	BackupId() *string
	SetBackupId(val *string)
	// The number of automated backups that you want to keep.
	//
	// Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
	BackupRetentionCount() *float64
	SetBackupRetentionCount(val *float64)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Supported on servers running Chef Automate 2.0 only. A PEM-formatted HTTPS certificate. The value can be be a single, self-signed certificate, or a certificate chain. If you specify a custom certificate, you must also specify values for `CustomDomain` and `CustomPrivateKey` . The following are requirements for the `CustomCertificate` value:.
	//
	// - You can provide either a self-signed, custom certificate, or the full certificate chain.
	// - The certificate must be a valid X509 certificate, or a certificate chain in PEM format.
	// - The certificate must be valid at the time of upload. A certificate can't be used before its validity period begins (the certificate's `NotBefore` date), or after it expires (the certificate's `NotAfter` date).
	// - The certificate’s common name or subject alternative names (SANs), if present, must match the value of `CustomDomain` .
	// - The certificate must match the value of `CustomPrivateKey` .
	CustomCertificate() *string
	SetCustomCertificate(val *string)
	// Supported on servers running Chef Automate 2.0 only. An optional public endpoint of a server, such as `https://aws.my-company.com` . To access the server, create a CNAME DNS record in your preferred DNS service that points the custom domain to the endpoint that is generated when the server is created (the value of the CreateServer Endpoint attribute). You cannot access the server by using the generated `Endpoint` value if the server is using a custom domain. If you specify a custom domain, you must also specify values for `CustomCertificate` and `CustomPrivateKey` .
	CustomDomain() *string
	SetCustomDomain(val *string)
	// Supported on servers running Chef Automate 2.0 only. A private key in PEM format for connecting to the server by using HTTPS. The private key must not be encrypted; it cannot be protected by a password or passphrase. If you specify a custom private key, you must also specify values for `CustomDomain` and `CustomCertificate` .
	CustomPrivateKey() *string
	SetCustomPrivateKey(val *string)
	// Enable or disable scheduled backups.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	DisableAutomatedBackup() interface{}
	SetDisableAutomatedBackup(val interface{})
	// The configuration management engine to use.
	//
	// Valid values include `ChefAutomate` and `Puppet` .
	Engine() *string
	SetEngine(val *string)
	// Optional engine attributes on a specified server.
	//
	// **Attributes accepted in a Chef createServer request:** - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. When no CHEF_AUTOMATE_PIVOTAL_KEY is set, a private key is generated and returned in the response. When you are specifying the value of CHEF_AUTOMATE_PIVOTAL_KEY as a parameter in the AWS CloudFormation console, you must add newline ( `\n` ) characters at the end of each line of the pivotal key value.
	// - `CHEF_AUTOMATE_ADMIN_PASSWORD` : The password for the administrative user in the Chef Automate web-based dashboard. The password length is a minimum of eight characters, and a maximum of 32. The password can contain letters, numbers, and special characters (!/@#$%^&+=_). The password must contain at least one lower case letter, one upper case letter, one number, and one special character. When no CHEF_AUTOMATE_ADMIN_PASSWORD is set, one is generated and returned in the response.
	//
	// **Attributes accepted in a Puppet createServer request:** - `PUPPET_ADMIN_PASSWORD` : To work with the Puppet Enterprise console, a password must use ASCII characters.
	// - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
	// - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add PUPPET_R10K_PRIVATE_KEY to specify a PEM-encoded private SSH key.
	EngineAttributes() interface{}
	SetEngineAttributes(val interface{})
	// The engine model of the server.
	//
	// Valid values in this release include `Monolithic` for Puppet and `Single` for Chef.
	EngineModel() *string
	SetEngineModel(val *string)
	// The major release version of the engine that you want to use.
	//
	// For a Chef server, the valid value for EngineVersion is currently `2` . For a Puppet server, valid values are `2019` or `2017` .
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The ARN of the instance profile that your Amazon EC2 instances use.
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	// The Amazon EC2 instance type to use.
	//
	// For example, `m5.large` .
	InstanceType() *string
	SetInstanceType(val *string)
	// The Amazon EC2 key pair to set for the instance.
	//
	// This parameter is optional; if desired, you may specify this parameter to connect to your instances by using SSH.
	KeyPair() *string
	SetKeyPair(val *string)
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
	// The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled.
	//
	// Valid values must be specified in one of the following formats:
	//
	// - `HH:MM` for daily backups
	// - `DDD:HH:MM` for weekly backups
	//
	// `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
	//
	// *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance.
	//
	// Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of security group IDs to attach to the Amazon EC2 instance.
	//
	// If you add this parameter, the specified security groups must be within the VPC that is specified by `SubnetIds` .
	//
	// If you do not specify this parameter, AWS OpsWorks CM creates one new security group that uses TCP ports 22 and 443, open to 0.0.0.0/0 (everyone).
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The name of the server.
	//
	// The server name must be unique within your AWS account, within each region. Server names must start with a letter; then letters, numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	ServerName() *string
	SetServerName(val *string)
	// The service role that the AWS OpsWorks CM service backend uses to work with your account.
	//
	// Although the AWS OpsWorks management console typically creates the service role for you, if you are using the AWS CLI or API commands, run the service-role-creation.yaml AWS CloudFormation template, located at https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml. This template creates a CloudFormation stack that includes the service role and instance profile that you need.
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The IDs of subnets in which to launch the server EC2 instance.
	//
	// Amazon EC2-Classic customers: This field is required. All servers must run within a VPC. The VPC must have "Auto Assign Public IP" enabled.
	//
	// EC2-VPC customers: This field is optional. If you do not specify subnet IDs, your EC2 instances are created in a default subnet that is selected by Amazon EC2. If you specify subnet IDs, the VPC must have "Auto Assign Public IP" enabled.
	//
	// For more information about supported Amazon EC2 platforms, see [Supported Platforms](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) .
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// A map that contains tag keys and tag values to attach to an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
	//
	// - The key cannot be empty.
	// - The key can be a maximum of 127 characters, and can contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - The value can be a maximum 255 characters, and contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - Leading and trailing spaces are trimmed from both the key and value.
	// - A maximum of 50 user-applied tags is allowed for any AWS OpsWorks CM server.
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

func (j *jsiiProxy_CfnServer) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnServer(scope constructs.Construct, id *string, props *CfnServerProps) CfnServer {
	_init_.Initialize()

	j := jsiiProxy_CfnServer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorksCM::Server`.
func NewCfnServer_Override(c CfnServer, scope constructs.Construct, id *string, props *CfnServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
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
func CfnServer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnServer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
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
func CfnServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
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
		"aws-cdk-lib.aws_opsworkscm.CfnServer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnServer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (c *jsiiProxy_CfnServer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `EngineAttribute` property type specifies administrator credentials for an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
//
// `EngineAttribute` is a property of the `AWS::OpsWorksCM::Server` resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   engineAttributeProperty := &engineAttributeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnServer_EngineAttributeProperty struct {
	// The name of the engine attribute.
	//
	// *Attribute name for Chef Automate servers:*
	//
	// - `CHEF_AUTOMATE_ADMIN_PASSWORD`
	//
	// *Attribute names for Puppet Enterprise servers:*
	//
	// - `PUPPET_ADMIN_PASSWORD`
	// - `PUPPET_R10K_REMOTE`
	// - `PUPPET_R10K_PRIVATE_KEY`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the engine attribute.
	//
	// *Attribute value for Chef Automate servers:*
	//
	// - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. You can generate this key by running the following [OpenSSL](https://docs.aws.amazon.com/https://www.openssl.org/) command on Linux-based computers.
	//
	// `openssl genrsa -out *pivotal_key_file_name* .pem 2048`
	//
	// On Windows-based computers, you can use the PuTTYgen utility to generate a base64-encoded RSA private key. For more information, see [PuTTYgen - Key Generator for PuTTY on Windows](https://docs.aws.amazon.com/https://www.ssh.com/ssh/putty/windows/puttygen) on SSH.com.
	//
	// *Attribute values for Puppet Enterprise servers:*
	//
	// - `PUPPET_ADMIN_PASSWORD` : An administrator password that you can use to sign in to the Puppet Enterprise console webpage after the server is online. The password must use between 8 and 32 ASCII characters.
	// - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
	// - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add `PUPPET_R10K_PRIVATE_KEY` to specify a PEM-encoded private SSH key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerProps := &cfnServerProps{
//   	instanceProfileArn: jsii.String("instanceProfileArn"),
//   	instanceType: jsii.String("instanceType"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	backupId: jsii.String("backupId"),
//   	backupRetentionCount: jsii.Number(123),
//   	customCertificate: jsii.String("customCertificate"),
//   	customDomain: jsii.String("customDomain"),
//   	customPrivateKey: jsii.String("customPrivateKey"),
//   	disableAutomatedBackup: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineAttributes: []interface{}{
//   		&engineAttributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	engineModel: jsii.String("engineModel"),
//   	engineVersion: jsii.String("engineVersion"),
//   	keyPair: jsii.String("keyPair"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	serverName: jsii.String("serverName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServerProps struct {
	// The ARN of the instance profile that your Amazon EC2 instances use.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The Amazon EC2 instance type to use.
	//
	// For example, `m5.large` .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The service role that the AWS OpsWorks CM service backend uses to work with your account.
	//
	// Although the AWS OpsWorks management console typically creates the service role for you, if you are using the AWS CLI or API commands, run the service-role-creation.yaml AWS CloudFormation template, located at https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml. This template creates a CloudFormation stack that includes the service role and instance profile that you need.
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Associate a public IP address with a server that you are launching.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// If you specify this field, AWS OpsWorks CM creates the server by using the backup represented by BackupId.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The number of automated backups that you want to keep.
	//
	// Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
	BackupRetentionCount *float64 `field:"optional" json:"backupRetentionCount" yaml:"backupRetentionCount"`
	// Supported on servers running Chef Automate 2.0 only. A PEM-formatted HTTPS certificate. The value can be be a single, self-signed certificate, or a certificate chain. If you specify a custom certificate, you must also specify values for `CustomDomain` and `CustomPrivateKey` . The following are requirements for the `CustomCertificate` value:.
	//
	// - You can provide either a self-signed, custom certificate, or the full certificate chain.
	// - The certificate must be a valid X509 certificate, or a certificate chain in PEM format.
	// - The certificate must be valid at the time of upload. A certificate can't be used before its validity period begins (the certificate's `NotBefore` date), or after it expires (the certificate's `NotAfter` date).
	// - The certificate’s common name or subject alternative names (SANs), if present, must match the value of `CustomDomain` .
	// - The certificate must match the value of `CustomPrivateKey` .
	CustomCertificate *string `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Supported on servers running Chef Automate 2.0 only. An optional public endpoint of a server, such as `https://aws.my-company.com` . To access the server, create a CNAME DNS record in your preferred DNS service that points the custom domain to the endpoint that is generated when the server is created (the value of the CreateServer Endpoint attribute). You cannot access the server by using the generated `Endpoint` value if the server is using a custom domain. If you specify a custom domain, you must also specify values for `CustomCertificate` and `CustomPrivateKey` .
	CustomDomain *string `field:"optional" json:"customDomain" yaml:"customDomain"`
	// Supported on servers running Chef Automate 2.0 only. A private key in PEM format for connecting to the server by using HTTPS. The private key must not be encrypted; it cannot be protected by a password or passphrase. If you specify a custom private key, you must also specify values for `CustomDomain` and `CustomCertificate` .
	CustomPrivateKey *string `field:"optional" json:"customPrivateKey" yaml:"customPrivateKey"`
	// Enable or disable scheduled backups.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	DisableAutomatedBackup interface{} `field:"optional" json:"disableAutomatedBackup" yaml:"disableAutomatedBackup"`
	// The configuration management engine to use.
	//
	// Valid values include `ChefAutomate` and `Puppet` .
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Optional engine attributes on a specified server.
	//
	// **Attributes accepted in a Chef createServer request:** - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. When no CHEF_AUTOMATE_PIVOTAL_KEY is set, a private key is generated and returned in the response. When you are specifying the value of CHEF_AUTOMATE_PIVOTAL_KEY as a parameter in the AWS CloudFormation console, you must add newline ( `\n` ) characters at the end of each line of the pivotal key value.
	// - `CHEF_AUTOMATE_ADMIN_PASSWORD` : The password for the administrative user in the Chef Automate web-based dashboard. The password length is a minimum of eight characters, and a maximum of 32. The password can contain letters, numbers, and special characters (!/@#$%^&+=_). The password must contain at least one lower case letter, one upper case letter, one number, and one special character. When no CHEF_AUTOMATE_ADMIN_PASSWORD is set, one is generated and returned in the response.
	//
	// **Attributes accepted in a Puppet createServer request:** - `PUPPET_ADMIN_PASSWORD` : To work with the Puppet Enterprise console, a password must use ASCII characters.
	// - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
	// - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add PUPPET_R10K_PRIVATE_KEY to specify a PEM-encoded private SSH key.
	EngineAttributes interface{} `field:"optional" json:"engineAttributes" yaml:"engineAttributes"`
	// The engine model of the server.
	//
	// Valid values in this release include `Monolithic` for Puppet and `Single` for Chef.
	EngineModel *string `field:"optional" json:"engineModel" yaml:"engineModel"`
	// The major release version of the engine that you want to use.
	//
	// For a Chef server, the valid value for EngineVersion is currently `2` . For a Puppet server, valid values are `2019` or `2017` .
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The Amazon EC2 key pair to set for the instance.
	//
	// This parameter is optional; if desired, you may specify this parameter to connect to your instances by using SSH.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled.
	//
	// Valid values must be specified in one of the following formats:
	//
	// - `HH:MM` for daily backups
	// - `DDD:HH:MM` for weekly backups
	//
	// `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
	//
	// *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance.
	//
	// Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A list of security group IDs to attach to the Amazon EC2 instance.
	//
	// If you add this parameter, the specified security groups must be within the VPC that is specified by `SubnetIds` .
	//
	// If you do not specify this parameter, AWS OpsWorks CM creates one new security group that uses TCP ports 22 and 443, open to 0.0.0.0/0 (everyone).
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The name of the server.
	//
	// The server name must be unique within your AWS account, within each region. Server names must start with a letter; then letters, numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The IDs of subnets in which to launch the server EC2 instance.
	//
	// Amazon EC2-Classic customers: This field is required. All servers must run within a VPC. The VPC must have "Auto Assign Public IP" enabled.
	//
	// EC2-VPC customers: This field is optional. If you do not specify subnet IDs, your EC2 instances are created in a default subnet that is selected by Amazon EC2. If you specify subnet IDs, the VPC must have "Auto Assign Public IP" enabled.
	//
	// For more information about supported Amazon EC2 platforms, see [Supported Platforms](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A map that contains tag keys and tag values to attach to an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
	//
	// - The key cannot be empty.
	// - The key can be a maximum of 127 characters, and can contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - The value can be a maximum 255 characters, and contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - Leading and trailing spaces are trimmed from both the key and value.
	// - A maximum of 50 user-applied tags is allowed for any AWS OpsWorks CM server.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

