package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AutoScaling::LaunchConfiguration`.
//
// The `AWS::AutoScaling::LaunchConfiguration` resource specifies the launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.
//
// When you update the launch configuration for an Auto Scaling group, CloudFormation deletes that resource and creates a new launch configuration with the updated properties and a new name. Existing instances are not affected. To update existing instances when you update the `AWS::AutoScaling::LaunchConfiguration` resource, you can specify an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) for the group. You can find sample update policies for rolling updates in [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) .
//
// > Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use launch configurations. They do not provide full functionality for Amazon EC2 Auto Scaling or Amazon EC2. For more information, see [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchConfiguration := awscdk.Aws_autoscaling.NewCfnLaunchConfiguration(this, jsii.String("MyCfnLaunchConfiguration"), &cfnLaunchConfigurationProps{
//   	imageId: jsii.String("imageId"),
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	blockDeviceMappings: []interface{}{
//   		&blockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//
//   			// the properties below are optional
//   			ebs: &blockDeviceProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.Boolean(false),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	classicLinkVpcId: jsii.String("classicLinkVpcId"),
//   	classicLinkVpcSecurityGroups: []*string{
//   		jsii.String("classicLinkVpcSecurityGroups"),
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   	iamInstanceProfile: jsii.String("iamInstanceProfile"),
//   	instanceId: jsii.String("instanceId"),
//   	instanceMonitoring: jsii.Boolean(false),
//   	kernelId: jsii.String("kernelId"),
//   	keyName: jsii.String("keyName"),
//   	launchConfigurationName: jsii.String("launchConfigurationName"),
//   	metadataOptions: &metadataOptionsProperty{
//   		httpEndpoint: jsii.String("httpEndpoint"),
//   		httpPutResponseHopLimit: jsii.Number(123),
//   		httpTokens: jsii.String("httpTokens"),
//   	},
//   	placementTenancy: jsii.String("placementTenancy"),
//   	ramDiskId: jsii.String("ramDiskId"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	spotPrice: jsii.String("spotPrice"),
//   	userData: jsii.String("userData"),
//   })
//
type CfnLaunchConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies whether to assign a public IPv4 address to the group's instances.
	//
	// If the instance is launched into a default subnet, the default is to assign a public IPv4 address, unless you disabled the option to assign a public IPv4 address on the subnet. If the instance is launched into a nondefault subnet, the default is not to assign a public IPv4 address, unless you enabled the option to assign a public IPv4 address on the subnet.
	//
	// If you specify `true` , each instance in the Auto Scaling group receives a unique public IPv4 address. For more information, see [Launching Auto Scaling instances in a VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify this property, you must specify at least one subnet for `VPCZoneIdentifier` when you create your group.
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	// The block device mapping entries that define the block devices to attach to the instances at launch.
	//
	// By default, the block devices specified in the block device mapping for the AMI are used. For more information, see [Block device mappings](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the *Amazon EC2 User Guide for Linux Instances* .
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Available for backward compatibility.
	ClassicLinkVpcId() *string
	SetClassicLinkVpcId(val *string)
	// Available for backward compatibility.
	ClassicLinkVpcSecurityGroups() *[]*string
	SetClassicLinkVpcSecurityGroups(val *[]*string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specifies whether the launch configuration is optimized for EBS I/O ( `true` ) or not ( `false` ).
	//
	// The optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization is not available with all instance types. Additional fees are incurred when you enable EBS optimization for an instance type that is not EBS-optimized by default. For more information, see [Amazon EBS-optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// The default value is `false` .
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	// The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.
	//
	// The instance profile contains the IAM role. For more information, see [IAM role for applications that run on Amazon EC2 instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	//
	// For more information, see [Finding a Linux AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// If you specify `InstanceId` , an `ImageId` is not required.
	ImageId() *string
	SetImageId(val *string)
	// The ID of the Amazon EC2 instance to use to create the launch configuration.
	//
	// When you use an instance to create a launch configuration, all properties are derived from the instance with the exception of `BlockDeviceMapping` and `AssociatePublicIpAddress` . You can override any properties from the instance by specifying them in the launch configuration.
	InstanceId() *string
	SetInstanceId(val *string)
	// Controls whether instances in this group are launched with detailed ( `true` ) or basic ( `false` ) monitoring.
	//
	// The default value is `true` (enabled).
	//
	// > When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes. For more information, see [Configure Monitoring for Auto Scaling Instances](https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html) in the *Amazon EC2 Auto Scaling User Guide* .
	InstanceMonitoring() interface{}
	SetInstanceMonitoring(val interface{})
	// Specifies the instance type of the EC2 instance.
	//
	// For information about available instance types, see [Available instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// If you specify `InstanceId` , an `InstanceType` is not required.
	InstanceType() *string
	SetInstanceType(val *string)
	// The ID of the kernel associated with the AMI.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	KernelId() *string
	SetKernelId(val *string)
	// The name of the key pair.
	//
	// For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon EC2 User Guide for Linux Instances* .
	KeyName() *string
	SetKeyName(val *string)
	// The name of the launch configuration.
	//
	// This name must be unique per Region per account.
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
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
	// The metadata options for the instances.
	//
	// For more information, see [Configuring the Instance Metadata Options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the *Amazon EC2 Auto Scaling User Guide* .
	MetadataOptions() interface{}
	SetMetadataOptions(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The tenancy of the instance, either `default` or `dedicated` .
	//
	// An instance with `dedicated` tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC. To launch dedicated instances into a shared tenancy VPC (a VPC with the instance placement tenancy attribute set to `default` ), you must set the value of this property to `dedicated` . For more information, see [Configuring instance tenancy with Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify `PlacementTenancy` , you must specify at least one subnet for `VPCZoneIdentifier` when you create your group.
	//
	// Valid values: `default` | `dedicated`.
	PlacementTenancy() *string
	SetPlacementTenancy(val *string)
	// The ID of the RAM disk to select.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	RamDiskId() *string
	SetRamDiskId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list that contains the security groups to assign to the instances in the Auto Scaling group.
	//
	// The list can contain both the IDs of existing security groups and references to [SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
	//
	// For more information, see [Control traffic to resources using security groups](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The maximum hourly price to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are launched when the price you specify exceeds the current Spot price. For more information, see [Request Spot Instances for fault-tolerant and flexible applications](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-template-spot-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid Range: Minimum value of 0.001
	//
	// > When you change your maximum price by creating a new launch configuration, running instances will continue to run as long as the maximum price for those running instances is higher than the current Spot price.
	SpotPrice() *string
	SetSpotPrice(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The Base64-encoded user data to make available to the launched EC2 instances.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon EC2 User Guide for Linux Instances* .
	UserData() *string
	SetUserData(val *string)
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

// The jsii proxy struct for CfnLaunchConfiguration
type jsiiProxy_CfnLaunchConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchConfiguration) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) BlockDeviceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ClassicLinkVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicLinkVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ClassicLinkVpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classicLinkVpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) MetadataOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchConfiguration) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScaling::LaunchConfiguration`.
func NewCfnLaunchConfiguration(scope awscdk.Construct, id *string, props *CfnLaunchConfigurationProps) CfnLaunchConfiguration {
	_init_.Initialize()

	if err := validateNewCfnLaunchConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchConfiguration{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScaling::LaunchConfiguration`.
func NewCfnLaunchConfiguration_Override(c CfnLaunchConfiguration, scope awscdk.Construct, id *string, props *CfnLaunchConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetBlockDeviceMappings(val interface{}) {
	if err := j.validateSetBlockDeviceMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDeviceMappings",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetClassicLinkVpcId(val *string) {
	_jsii_.Set(
		j,
		"classicLinkVpcId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetClassicLinkVpcSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"classicLinkVpcSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetIamInstanceProfile(val *string) {
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetInstanceMonitoring(val interface{}) {
	if err := j.validateSetInstanceMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetKernelId(val *string) {
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetKeyName(val *string) {
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetLaunchConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetMetadataOptions(val interface{}) {
	if err := j.validateSetMetadataOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataOptions",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetPlacementTenancy(val *string) {
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetRamDiskId(val *string) {
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetSpotPrice(val *string) {
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchConfiguration)SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
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
func CfnLaunchConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchConfiguration_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLaunchConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchConfiguration_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLaunchConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_autoscaling.CfnLaunchConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnLaunchConfiguration) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnLaunchConfiguration) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnLaunchConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLaunchConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchConfiguration) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

