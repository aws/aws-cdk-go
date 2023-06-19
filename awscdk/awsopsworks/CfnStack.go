package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::OpsWorks::Stack`.
//
// Creates a new stack. For more information, see [Create a New Stack](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-edit.html) .
//
// *Required Permissions* : To use this action, an IAM user must have an attached policy that explicitly grants permissions. For more information about user permissions, see [Managing User Permissions](https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customJson interface{}
//
//   cfnStack := awscdk.Aws_opsworks.NewCfnStack(this, jsii.String("MyCfnStack"), &CfnStackProps{
//   	DefaultInstanceProfileArn: jsii.String("defaultInstanceProfileArn"),
//   	Name: jsii.String("name"),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	AgentVersion: jsii.String("agentVersion"),
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	ChefConfiguration: &ChefConfigurationProperty{
//   		BerkshelfVersion: jsii.String("berkshelfVersion"),
//   		ManageBerkshelf: jsii.Boolean(false),
//   	},
//   	CloneAppIds: []*string{
//   		jsii.String("cloneAppIds"),
//   	},
//   	ClonePermissions: jsii.Boolean(false),
//   	ConfigurationManager: &StackConfigurationManagerProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	CustomCookbooksSource: &SourceProperty{
//   		Password: jsii.String("password"),
//   		Revision: jsii.String("revision"),
//   		SshKey: jsii.String("sshKey"),
//   		Type: jsii.String("type"),
//   		Url: jsii.String("url"),
//   		Username: jsii.String("username"),
//   	},
//   	CustomJson: customJson,
//   	DefaultAvailabilityZone: jsii.String("defaultAvailabilityZone"),
//   	DefaultOs: jsii.String("defaultOs"),
//   	DefaultRootDeviceType: jsii.String("defaultRootDeviceType"),
//   	DefaultSshKeyName: jsii.String("defaultSshKeyName"),
//   	DefaultSubnetId: jsii.String("defaultSubnetId"),
//   	EcsClusterArn: jsii.String("ecsClusterArn"),
//   	ElasticIps: []interface{}{
//   		&ElasticIpProperty{
//   			Ip: jsii.String("ip"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	HostnameTheme: jsii.String("hostnameTheme"),
//   	RdsDbInstances: []interface{}{
//   		&RdsDbInstanceProperty{
//   			DbPassword: jsii.String("dbPassword"),
//   			DbUser: jsii.String("dbUser"),
//   			RdsDbInstanceArn: jsii.String("rdsDbInstanceArn"),
//   		},
//   	},
//   	SourceStackId: jsii.String("sourceStackId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseCustomCookbooks: jsii.Boolean(false),
//   	UseOpsworksSecurityGroups: jsii.Boolean(false),
//   	VpcId: jsii.String("vpcId"),
//   })
//
type CfnStack interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The default AWS OpsWorks Stacks agent version. You have the following options:.
	//
	// - Auto-update - Set this parameter to `LATEST` . AWS OpsWorks Stacks automatically installs new agent versions on the stack's instances as soon as they are available.
	// - Fixed version - Set this parameter to your preferred agent version. To update the agent version, you must edit the stack configuration and specify a new version. AWS OpsWorks Stacks installs that version on the stack's instances.
	//
	// The default setting is the most recent release of the agent. To specify an agent version, you must use the complete version number, not the abbreviated number shown on the console. For a list of available agent version numbers, call [DescribeAgentVersions](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeAgentVersions) . AgentVersion cannot be set to Chef 12.2.
	//
	// > You can also specify an agent version when you create or update an instance, which overrides the stack's default setting.
	AgentVersion() *string
	SetAgentVersion(val *string)
	// One or more user-defined key-value pairs to be added to the stack attributes.
	Attributes() interface{}
	SetAttributes(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A `ChefConfiguration` object that specifies whether to enable Berkshelf and the Berkshelf version on Chef 11.10 stacks. For more information, see [Create a New Stack](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html) .
	ChefConfiguration() interface{}
	SetChefConfiguration(val interface{})
	// If you're cloning an AWS OpsWorks stack, a list of AWS OpsWorks application stack IDs from the source stack to include in the cloned stack.
	CloneAppIds() *[]*string
	SetCloneAppIds(val *[]*string)
	// If you're cloning an AWS OpsWorks stack, indicates whether to clone the source stack's permissions.
	ClonePermissions() interface{}
	SetClonePermissions(val interface{})
	// The configuration manager.
	//
	// When you create a stack we recommend that you use the configuration manager to specify the Chef version: 12, 11.10, or 11.4 for Linux stacks, or 12.2 for Windows stacks. The default value for Linux stacks is currently 12.
	ConfigurationManager() interface{}
	SetConfigurationManager(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Contains the information required to retrieve an app or cookbook from a repository.
	//
	// For more information, see [Adding Apps](https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html) or [Cookbooks and Recipes](https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook.html) .
	CustomCookbooksSource() interface{}
	SetCustomCookbooksSource(val interface{})
	// A string that contains user-defined, custom JSON.
	//
	// It can be used to override the corresponding default stack configuration attribute values or to pass data to recipes. The string should be in the following format:
	//
	// `"{\"key1\": \"value1\", \"key2\": \"value2\",...}"`
	//
	// For more information about custom JSON, see [Use Custom JSON to Modify the Stack Configuration Attributes](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-json.html) .
	CustomJson() interface{}
	SetCustomJson(val interface{})
	// The stack's default Availability Zone, which must be in the specified region.
	//
	// For more information, see [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html) . If you also specify a value for `DefaultSubnetId` , the subnet must be in the same zone. For more information, see the `VpcId` parameter description.
	DefaultAvailabilityZone() *string
	SetDefaultAvailabilityZone(val *string)
	// The Amazon Resource Name (ARN) of an IAM profile that is the default profile for all of the stack's EC2 instances.
	//
	// For more information about IAM ARNs, see [Using Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) .
	DefaultInstanceProfileArn() *string
	SetDefaultInstanceProfileArn(val *string)
	// The stack's default operating system, which is installed on every instance unless you specify a different operating system when you create the instance.
	//
	// You can specify one of the following.
	//
	// - A supported Linux operating system: An Amazon Linux version, such as `Amazon Linux 2` , `Amazon Linux 2018.03` , `Amazon Linux 2017.09` , `Amazon Linux 2017.03` , `Amazon Linux 2016.09` , `Amazon Linux 2016.03` , `Amazon Linux 2015.09` , or `Amazon Linux 2015.03` .
	// - A supported Ubuntu operating system, such as `Ubuntu 18.04 LTS` , `Ubuntu 16.04 LTS` , `Ubuntu 14.04 LTS` , or `Ubuntu 12.04 LTS` .
	// - `CentOS Linux 7`
	// - `Red Hat Enterprise Linux 7`
	// - A supported Windows operating system, such as `Microsoft Windows Server 2012 R2 Base` , `Microsoft Windows Server 2012 R2 with SQL Server Express` , `Microsoft Windows Server 2012 R2 with SQL Server Standard` , or `Microsoft Windows Server 2012 R2 with SQL Server Web` .
	// - A custom AMI: `Custom` . You specify the custom AMI you want to use when you create instances. For more information, see [Using Custom AMIs](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html) .
	//
	// The default option is the current Amazon Linux version. Not all operating systems are supported with all versions of Chef. For more information about supported operating systems, see [AWS OpsWorks Stacks Operating Systems](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html) .
	DefaultOs() *string
	SetDefaultOs(val *string)
	// The default root device type.
	//
	// This value is the default for all instances in the stack, but you can override it when you create an instance. The default option is `instance-store` . For more information, see [Storage for the Root Device](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device) .
	DefaultRootDeviceType() *string
	SetDefaultRootDeviceType(val *string)
	// A default Amazon EC2 key pair name.
	//
	// The default value is none. If you specify a key pair name, AWS OpsWorks installs the public key on the instance and you can use the private key with an SSH client to log in to the instance. For more information, see [Using SSH to Communicate with an Instance](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-ssh.html) and [Managing SSH Access](https://docs.aws.amazon.com/opsworks/latest/userguide/security-ssh-access.html) . You can override this setting by specifying a different key pair, or no key pair, when you [create an instance](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html) .
	DefaultSshKeyName() *string
	SetDefaultSshKeyName(val *string)
	// The stack's default subnet ID.
	//
	// All instances are launched into this subnet unless you specify another subnet ID when you create the instance. This parameter is required if you specify a value for the `VpcId` parameter. If you also specify a value for `DefaultAvailabilityZone` , the subnet must be in that zone.
	DefaultSubnetId() *string
	SetDefaultSubnetId(val *string)
	// The Amazon Resource Name (ARN) of the Amazon Elastic Container Service ( Amazon ECS ) cluster to register with the AWS OpsWorks stack.
	//
	// > If you specify a cluster that's registered with another AWS OpsWorks stack, AWS CloudFormation deregisters the existing association before registering the cluster.
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	// A list of Elastic IP addresses to register with the AWS OpsWorks stack.
	//
	// > If you specify an IP address that's registered with another AWS OpsWorks stack, AWS CloudFormation deregisters the existing association before registering the IP address.
	ElasticIps() interface{}
	SetElasticIps(val interface{})
	// The stack's host name theme, with spaces replaced by underscores.
	//
	// The theme is used to generate host names for the stack's instances. By default, `HostnameTheme` is set to `Layer_Dependent` , which creates host names by appending integers to the layer's short name. The other themes are:
	//
	// - `Baked_Goods`
	// - `Clouds`
	// - `Europe_Cities`
	// - `Fruits`
	// - `Greek_Deities_and_Titans`
	// - `Legendary_creatures_from_Japan`
	// - `Planets_and_Moons`
	// - `Roman_Deities`
	// - `Scottish_Islands`
	// - `US_Cities`
	// - `Wild_Cats`
	//
	// To obtain a generated host name, call `GetHostNameSuggestion` , which returns a host name based on the current theme.
	HostnameTheme() *string
	SetHostnameTheme(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The stack name.
	//
	// Stack names can be a maximum of 64 characters.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Amazon Relational Database Service ( Amazon RDS ) database instance to register with the AWS OpsWorks stack.
	//
	// > If you specify a database instance that's registered with another AWS OpsWorks stack, AWS CloudFormation deregisters the existing association before registering the database instance.
	RdsDbInstances() interface{}
	SetRdsDbInstances(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack's IAM role, which allows AWS OpsWorks Stacks to work with AWS resources on your behalf.
	//
	// You must set this parameter to the Amazon Resource Name (ARN) for an existing IAM role. For more information about IAM ARNs, see [Using Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) .
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// If you're cloning an AWS OpsWorks stack, the stack ID of the source AWS OpsWorks stack to clone.
	SourceStackId() *string
	SetSourceStackId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A map that contains tag keys and tag values that are attached to a stack or layer.
	//
	// - The key cannot be empty.
	// - The key can be a maximum of 127 characters, and can contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : /`
	// - The value can be a maximum 255 characters, and contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : /`
	// - Leading and trailing white spaces are trimmed from both the key and value.
	// - A maximum of 40 tags is allowed for any resource.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Whether the stack uses custom cookbooks.
	UseCustomCookbooks() interface{}
	SetUseCustomCookbooks(val interface{})
	// Whether to associate the AWS OpsWorks Stacks built-in security groups with the stack's layers.
	//
	// AWS OpsWorks Stacks provides a standard set of built-in security groups, one for each layer, which are associated with layers by default. With `UseOpsworksSecurityGroups` you can instead provide your own custom security groups. `UseOpsworksSecurityGroups` has the following settings:
	//
	// - True - AWS OpsWorks Stacks automatically associates the appropriate built-in security group with each layer (default setting). You can associate additional security groups with a layer after you create it, but you cannot delete the built-in security group.
	// - False - AWS OpsWorks Stacks does not associate built-in security groups with layers. You must create appropriate EC2 security groups and associate a security group with each layer that you create. However, you can still manually associate a built-in security group with a layer on creation; custom security groups are required only for those layers that need custom settings.
	//
	// For more information, see [Create a New Stack](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html) .
	UseOpsworksSecurityGroups() interface{}
	SetUseOpsworksSecurityGroups(val interface{})
	// The ID of the VPC that the stack is to be launched into.
	//
	// The VPC must be in the stack's region. All instances are launched into this VPC. You cannot change the ID later.
	//
	// - If your account supports EC2-Classic, the default value is `no VPC` .
	// - If your account does not support EC2-Classic, the default value is the default VPC for the specified region.
	//
	// If the VPC ID corresponds to a default VPC and you have specified either the `DefaultAvailabilityZone` or the `DefaultSubnetId` parameter only, AWS OpsWorks Stacks infers the value of the other parameter. If you specify neither parameter, AWS OpsWorks Stacks sets these parameters to the first valid Availability Zone for the specified region and the corresponding default VPC subnet ID, respectively.
	//
	// If you specify a nondefault VPC ID, note the following:
	//
	// - It must belong to a VPC in your account that is in the specified region.
	// - You must specify a value for `DefaultSubnetId` .
	//
	// For more information about how to use AWS OpsWorks Stacks with a VPC, see [Running a Stack in a VPC](https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-vpc.html) . For more information about default VPC and EC2-Classic, see [Supported Platforms](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) .
	VpcId() *string
	SetVpcId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStack
type jsiiProxy_CfnStack struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStack) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Attributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ChefConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chefConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CloneAppIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloneAppIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ClonePermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clonePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ConfigurationManager() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomCookbooksSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) CustomJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultOs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultRootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) DefaultSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) HostnameTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) RdsDbInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsDbInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) SourceStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceStackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseCustomCookbooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) UseOpsworksSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStack) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack(scope awscdk.Construct, id *string, props *CfnStackProps) CfnStack {
	_init_.Initialize()

	if err := validateNewCfnStackParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStack{}

	_jsii_.Create(
		"monocdk.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Stack`.
func NewCfnStack_Override(c CfnStack, scope awscdk.Construct, id *string, props *CfnStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_opsworks.CfnStack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStack)SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetAttributes(val interface{}) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetChefConfiguration(val interface{}) {
	if err := j.validateSetChefConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chefConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCloneAppIds(val *[]*string) {
	_jsii_.Set(
		j,
		"cloneAppIds",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetClonePermissions(val interface{}) {
	if err := j.validateSetClonePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clonePermissions",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetConfigurationManager(val interface{}) {
	if err := j.validateSetConfigurationManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationManager",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCustomCookbooksSource(val interface{}) {
	if err := j.validateSetCustomCookbooksSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCookbooksSource",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetCustomJson(val interface{}) {
	if err := j.validateSetCustomJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"defaultAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultInstanceProfileArn(val *string) {
	if err := j.validateSetDefaultInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultOs(val *string) {
	_jsii_.Set(
		j,
		"defaultOs",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"defaultRootDeviceType",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"defaultSshKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetDefaultSubnetId(val *string) {
	_jsii_.Set(
		j,
		"defaultSubnetId",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetEcsClusterArn(val *string) {
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetElasticIps(val interface{}) {
	if err := j.validateSetElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetHostnameTheme(val *string) {
	_jsii_.Set(
		j,
		"hostnameTheme",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetRdsDbInstances(val interface{}) {
	if err := j.validateSetRdsDbInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstances",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetSourceStackId(val *string) {
	_jsii_.Set(
		j,
		"sourceStackId",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetUseCustomCookbooks(val interface{}) {
	if err := j.validateSetUseCustomCookbooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomCookbooks",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetUseOpsworksSecurityGroups(val interface{}) {
	if err := j.validateSetUseOpsworksSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOpsworksSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnStack)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
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
func CfnStack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnStack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStack_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnStack",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_opsworks.CfnStack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStack) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStack) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStack) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStack) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStack) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStack) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStack) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnStack) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStack) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStack) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStack) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStack) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnStack) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStack) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

