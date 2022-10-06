package awsopsworks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::OpsWorks::Instance`.
//
// Creates an instance in a specified stack. For more information, see [Adding an Instance to a Layer](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html) .
//
// *Required Permissions* : To use this action, an IAM user must have a Manage permissions level for the stack, or an attached policy that explicitly grants permissions. For more information on user permissions, see [Managing User Permissions](https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstance := awscdk.Aws_opsworks.NewCfnInstance(this, jsii.String("MyCfnInstance"), &cfnInstanceProps{
//   	instanceType: jsii.String("instanceType"),
//   	layerIds: []*string{
//   		jsii.String("layerIds"),
//   	},
//   	stackId: jsii.String("stackId"),
//
//   	// the properties below are optional
//   	agentVersion: jsii.String("agentVersion"),
//   	amiId: jsii.String("amiId"),
//   	architecture: jsii.String("architecture"),
//   	autoScalingType: jsii.String("autoScalingType"),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	blockDeviceMappings: []interface{}{
//   		&blockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsBlockDeviceProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				snapshotId: jsii.String("snapshotId"),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   	elasticIps: []*string{
//   		jsii.String("elasticIps"),
//   	},
//   	hostname: jsii.String("hostname"),
//   	installUpdatesOnBoot: jsii.Boolean(false),
//   	os: jsii.String("os"),
//   	rootDeviceType: jsii.String("rootDeviceType"),
//   	sshKeyName: jsii.String("sshKeyName"),
//   	subnetId: jsii.String("subnetId"),
//   	tenancy: jsii.String("tenancy"),
//   	timeBasedAutoScaling: &timeBasedAutoScalingProperty{
//   		friday: map[string]*string{
//   			"fridayKey": jsii.String("friday"),
//   		},
//   		monday: map[string]*string{
//   			"mondayKey": jsii.String("monday"),
//   		},
//   		saturday: map[string]*string{
//   			"saturdayKey": jsii.String("saturday"),
//   		},
//   		sunday: map[string]*string{
//   			"sundayKey": jsii.String("sunday"),
//   		},
//   		thursday: map[string]*string{
//   			"thursdayKey": jsii.String("thursday"),
//   		},
//   		tuesday: map[string]*string{
//   			"tuesdayKey": jsii.String("tuesday"),
//   		},
//   		wednesday: map[string]*string{
//   			"wednesdayKey": jsii.String("wednesday"),
//   		},
//   	},
//   	virtualizationType: jsii.String("virtualizationType"),
//   	volumes: []*string{
//   		jsii.String("volumes"),
//   	},
//   })
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The default AWS OpsWorks Stacks agent version. You have the following options:.
	//
	// - `INHERIT` - Use the stack's default agent version setting.
	// - *version_number* - Use the specified agent version. This value overrides the stack's default setting. To update the agent version, edit the instance configuration and specify a new version. AWS OpsWorks Stacks installs that version on the instance.
	//
	// The default setting is `INHERIT` . To specify an agent version, you must use the complete version number, not the abbreviated number shown on the console. For a list of available agent version numbers, call [DescribeAgentVersions](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeAgentVersions) . AgentVersion cannot be set to Chef 12.2.
	AgentVersion() *string
	SetAgentVersion(val *string)
	// A custom AMI ID to be used to create the instance.
	//
	// The AMI should be based on one of the supported operating systems. For more information, see [Using Custom AMIs](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html) .
	//
	// > If you specify a custom AMI, you must set `Os` to `Custom` .
	AmiId() *string
	SetAmiId(val *string)
	// The instance architecture.
	//
	// The default option is `x86_64` . Instance types do not necessarily support both architectures. For a list of the architectures that are supported by the different instance types, see [Instance Families and Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) .
	Architecture() *string
	SetArchitecture(val *string)
	// The Availability Zone of the AWS OpsWorks instance, such as `us-east-2a` .
	AttrAvailabilityZone() *string
	// The private DNS name of the AWS OpsWorks instance.
	AttrPrivateDnsName() *string
	// The private IP address of the AWS OpsWorks instance, such as `192.0.2.0` .
	AttrPrivateIp() *string
	// The public DNS name of the AWS OpsWorks instance.
	AttrPublicDnsName() *string
	// The public IP address of the AWS OpsWorks instance, such as `192.0.2.0` .
	//
	// > Use this attribute only when the AWS OpsWorks instance is in an AWS OpsWorks layer that auto-assigns public IP addresses.
	AttrPublicIp() *string
	// For load-based or time-based instances, the type.
	//
	// Windows stacks can use only time-based instances.
	AutoScalingType() *string
	SetAutoScalingType(val *string)
	// The Availability Zone of the AWS OpsWorks instance, such as `us-east-2a` .
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// An array of `BlockDeviceMapping` objects that specify the instance's block devices.
	//
	// For more information, see [Block Device Mapping](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) . Note that block device mappings are not supported for custom AMIs.
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Whether to create an Amazon EBS-optimized instance.
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	// A list of Elastic IP addresses to associate with the instance.
	ElasticIps() *[]*string
	SetElasticIps(val *[]*string)
	// The instance host name. The following are character limits for instance host names.
	//
	// - Linux-based instances: 63 characters
	// - Windows-based instances: 15 characters.
	Hostname() *string
	SetHostname(val *string)
	// Whether to install operating system and package updates when the instance boots.
	//
	// The default value is `true` . To control when updates are installed, set this value to `false` . You must then update your instances manually by using [CreateDeployment](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateDeployment) to run the `update_dependencies` stack command or by manually running `yum` (Amazon Linux) or `apt-get` (Ubuntu) on the instances.
	//
	// > We strongly recommend using the default value of `true` to ensure that your instances have the latest security updates.
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	// The instance type, such as `t2.micro` . For a list of supported instance types, open the stack in the console, choose *Instances* , and choose *+ Instance* . The *Size* list contains the currently supported types. For more information, see [Instance Families and Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) . The parameter values that you use to specify the various types are in the *API Name* column of the *Available Instance Types* table.
	InstanceType() *string
	SetInstanceType(val *string)
	// An array that contains the instance's layer IDs.
	LayerIds() *[]*string
	SetLayerIds(val *[]*string)
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The instance's operating system, which must be set to one of the following.
	//
	// - A supported Linux operating system: An Amazon Linux version, such as `Amazon Linux 2` , `Amazon Linux 2018.03` , `Amazon Linux 2017.09` , `Amazon Linux 2017.03` , `Amazon Linux 2016.09` , `Amazon Linux 2016.03` , `Amazon Linux 2015.09` , or `Amazon Linux 2015.03` .
	// - A supported Ubuntu operating system, such as `Ubuntu 18.04 LTS` , `Ubuntu 16.04 LTS` , `Ubuntu 14.04 LTS` , or `Ubuntu 12.04 LTS` .
	// - `CentOS Linux 7`
	// - `Red Hat Enterprise Linux 7`
	// - A supported Windows operating system, such as `Microsoft Windows Server 2012 R2 Base` , `Microsoft Windows Server 2012 R2 with SQL Server Express` , `Microsoft Windows Server 2012 R2 with SQL Server Standard` , or `Microsoft Windows Server 2012 R2 with SQL Server Web` .
	// - A custom AMI: `Custom` .
	//
	// Not all operating systems are supported with all versions of Chef. For more information about the supported operating systems, see [AWS OpsWorks Stacks Operating Systems](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html) .
	//
	// The default option is the current Amazon Linux version. If you set this parameter to `Custom` , you must use the [CreateInstance](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateInstance) action's AmiId parameter to specify the custom AMI that you want to use. Block device mappings are not supported if the value is `Custom` . For more information about how to use custom AMIs with AWS OpsWorks Stacks, see [Using Custom AMIs](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html) .
	Os() *string
	SetOs(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The instance root device type.
	//
	// For more information, see [Storage for the Root Device](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device) .
	RootDeviceType() *string
	SetRootDeviceType(val *string)
	// The instance's Amazon EC2 key-pair name.
	SshKeyName() *string
	SetSshKeyName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The stack ID.
	StackId() *string
	SetStackId(val *string)
	// The ID of the instance's subnet.
	//
	// If the stack is running in a VPC, you can use this parameter to override the stack's default subnet ID value and direct AWS OpsWorks Stacks to launch the instance in a different subnet.
	SubnetId() *string
	SetSubnetId(val *string)
	// The instance's tenancy option.
	//
	// The default option is no tenancy, or if the instance is running in a VPC, inherit tenancy settings from the VPC. The following are valid values for this parameter: `dedicated` , `default` , or `host` . Because there are costs associated with changes in tenancy options, we recommend that you research tenancy options before choosing them for your instances. For more information about dedicated hosts, see [Dedicated Hosts Overview](https://docs.aws.amazon.com/ec2/dedicated-hosts/) and [Amazon EC2 Dedicated Hosts](https://docs.aws.amazon.com/ec2/dedicated-hosts/) . For more information about dedicated instances, see [Dedicated Instances](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/dedicated-instance.html) and [Amazon EC2 Dedicated Instances](https://docs.aws.amazon.com/ec2/purchasing-options/dedicated-instances/) .
	Tenancy() *string
	SetTenancy(val *string)
	// The time-based scaling configuration for the instance.
	TimeBasedAutoScaling() interface{}
	SetTimeBasedAutoScaling(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The instance's virtualization type, `paravirtual` or `hvm` .
	VirtualizationType() *string
	SetVirtualizationType(val *string)
	// A list of AWS OpsWorks volume IDs to associate with the instance.
	//
	// For more information, see [`AWS::OpsWorks::Volume`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html) .
	Volumes() *[]*string
	SetVolumes(val *[]*string)
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

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AutoScalingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BlockDeviceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) ElasticIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LayerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) RootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) TimeBasedAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Volumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::OpsWorks::Instance`.
func NewCfnInstance(scope awscdk.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	if err := validateNewCfnInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"monocdk.aws_opsworks.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::OpsWorks::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope awscdk.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_opsworks.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance)SetAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetAmiId(val *string) {
	_jsii_.Set(
		j,
		"amiId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetAutoScalingType(val *string) {
	_jsii_.Set(
		j,
		"autoScalingType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetBlockDeviceMappings(val interface{}) {
	if err := j.validateSetBlockDeviceMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDeviceMappings",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetElasticIps(val *[]*string) {
	_jsii_.Set(
		j,
		"elasticIps",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetInstallUpdatesOnBoot(val interface{}) {
	if err := j.validateSetInstallUpdatesOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetLayerIds(val *[]*string) {
	if err := j.validateSetLayerIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layerIds",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetOs(val *string) {
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetRootDeviceType(val *string) {
	_jsii_.Set(
		j,
		"rootDeviceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSshKeyName(val *string) {
	_jsii_.Set(
		j,
		"sshKeyName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetTenancy(val *string) {
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetTimeBasedAutoScaling(val interface{}) {
	if err := j.validateSetTimeBasedAutoScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeBasedAutoScaling",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetVirtualizationType(val *string) {
	_jsii_.Set(
		j,
		"virtualizationType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetVolumes(val *[]*string) {
	_jsii_.Set(
		j,
		"volumes",
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
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_opsworks.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_opsworks.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnInstance) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

