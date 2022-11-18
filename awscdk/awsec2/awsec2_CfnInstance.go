package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EC2::Instance`.
//
// Specifies an EC2 instance.
//
// If an Elastic IP address is attached to your instance, AWS CloudFormation reattaches the Elastic IP address after it updates the instance. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstance := awscdk.Aws_ec2.NewCfnInstance(this, jsii.String("MyCfnInstance"), &cfnInstanceProps{
//   	additionalInfo: jsii.String("additionalInfo"),
//   	affinity: jsii.String("affinity"),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	blockDeviceMappings: []interface{}{
//   		&blockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//
//   			// the properties below are optional
//   			ebs: &ebsProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: &noDeviceProperty{
//   			},
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	cpuOptions: &cpuOptionsProperty{
//   		coreCount: jsii.Number(123),
//   		threadsPerCore: jsii.Number(123),
//   	},
//   	creditSpecification: &creditSpecificationProperty{
//   		cpuCredits: jsii.String("cpuCredits"),
//   	},
//   	disableApiTermination: jsii.Boolean(false),
//   	ebsOptimized: jsii.Boolean(false),
//   	elasticGpuSpecifications: []interface{}{
//   		&elasticGpuSpecificationProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   	elasticInferenceAccelerators: []interface{}{
//   		&elasticInferenceAcceleratorProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			count: jsii.Number(123),
//   		},
//   	},
//   	enclaveOptions: &enclaveOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	hibernationOptions: &hibernationOptionsProperty{
//   		configured: jsii.Boolean(false),
//   	},
//   	hostId: jsii.String("hostId"),
//   	hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   	iamInstanceProfile: jsii.String("iamInstanceProfile"),
//   	imageId: jsii.String("imageId"),
//   	instanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   	instanceType: jsii.String("instanceType"),
//   	ipv6AddressCount: jsii.Number(123),
//   	ipv6Addresses: []interface{}{
//   		&instanceIpv6AddressProperty{
//   			ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	kernelId: jsii.String("kernelId"),
//   	keyName: jsii.String("keyName"),
//   	launchTemplate: &launchTemplateSpecificationProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	licenseSpecifications: []interface{}{
//   		&licenseSpecificationProperty{
//   			licenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   		},
//   	},
//   	monitoring: jsii.Boolean(false),
//   	networkInterfaces: []interface{}{
//   		&networkInterfaceProperty{
//   			deviceIndex: jsii.String("deviceIndex"),
//
//   			// the properties below are optional
//   			associateCarrierIpAddress: jsii.Boolean(false),
//   			associatePublicIpAddress: jsii.Boolean(false),
//   			deleteOnTermination: jsii.Boolean(false),
//   			description: jsii.String("description"),
//   			groupSet: []*string{
//   				jsii.String("groupSet"),
//   			},
//   			ipv6AddressCount: jsii.Number(123),
//   			ipv6Addresses: []interface{}{
//   				&instanceIpv6AddressProperty{
//   					ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			networkInterfaceId: jsii.String("networkInterfaceId"),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   			privateIpAddresses: []interface{}{
//   				&privateIpAddressSpecificationProperty{
//   					primary: jsii.Boolean(false),
//   					privateIpAddress: jsii.String("privateIpAddress"),
//   				},
//   			},
//   			secondaryPrivateIpAddressCount: jsii.Number(123),
//   			subnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	placementGroupName: jsii.String("placementGroupName"),
//   	privateDnsNameOptions: &privateDnsNameOptionsProperty{
//   		enableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   		enableResourceNameDnsARecord: jsii.Boolean(false),
//   		hostnameType: jsii.String("hostnameType"),
//   	},
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	propagateTagsToVolumeOnCreation: jsii.Boolean(false),
//   	ramdiskId: jsii.String("ramdiskId"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	sourceDestCheck: jsii.Boolean(false),
//   	ssmAssociations: []interface{}{
//   		&ssmAssociationProperty{
//   			documentName: jsii.String("documentName"),
//
//   			// the properties below are optional
//   			associationParameters: []interface{}{
//   				&associationParameterProperty{
//   					key: jsii.String("key"),
//   					value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tenancy: jsii.String("tenancy"),
//   	userData: jsii.String("userData"),
//   	volumes: []interface{}{
//   		&volumeProperty{
//   			device: jsii.String("device"),
//   			volumeId: jsii.String("volumeId"),
//   		},
//   	},
//   })
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// This property is reserved for internal use.
	//
	// If you use it, the stack fails with this error: `Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX)` .
	AdditionalInfo() *string
	SetAdditionalInfo(val *string)
	// Indicates whether the instance is associated with a dedicated host.
	//
	// If you want the instance to always restart on the same host on which it was launched, specify `host` . If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify `default` .
	Affinity() *string
	SetAffinity(val *string)
	// The Availability Zone where the specified instance is launched. For example: `us-east-1b` .
	//
	// You can retrieve a list of all Availability Zones for a Region by using the [Fn::GetAZs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getavailabilityzones.html) intrinsic function.
	AttrAvailabilityZone() *string
	// The private DNS name of the specified instance.
	//
	// For example: `ip-10-24-34-0.ec2.internal` .
	AttrPrivateDnsName() *string
	// The private IP address of the specified instance.
	//
	// For example: `10.24.34.0` .
	AttrPrivateIp() *string
	// The public DNS name of the specified instance.
	//
	// For example: `ec2-107-20-50-45.compute-1.amazonaws.com` .
	AttrPublicDnsName() *string
	// The public IP address of the specified instance.
	//
	// For example: `192.0.2.0` .
	AttrPublicIp() *string
	// The Availability Zone of the instance.
	//
	// If not specified, an Availability Zone will be automatically chosen for you based on the load balancing criteria for the Region.
	//
	// This parameter is not supported by [DescribeImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImageAttribute.html) .
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	//
	// By default, the block devices specified in the block device mapping for the AMI are used. You can override the AMI block device mapping using the instance block device mapping. For the root volume, you can override only the volume size, volume type, volume encryption settings, and the `DeleteOnTermination` setting.
	//
	// > After the instance is running, you can modify only the `DeleteOnTermination` parameter for the attached volumes without interrupting the instance. Modifying any other parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The CPU options for the instance.
	//
	// For more information, see [Optimize CPU options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
	CpuOptions() interface{}
	SetCpuOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The credit option for CPU usage of the burstable performance instance.
	//
	// Valid values are `standard` and `unlimited` . To change this attribute after launch, use [ModifyInstanceCreditSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html) . For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `standard` (T2 instances) or `unlimited` (T3/T3a instances)
	//
	// For T3 instances with `host` tenancy, only `standard` is supported.
	CreditSpecification() interface{}
	SetCreditSpecification(val interface{})
	// If you set this parameter to `true` , you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html) . Alternatively, if you set `InstanceInitiatedShutdownBehavior` to `terminate` , you can terminate the instance by running the shutdown command from the instance.
	//
	// Default: `false`.
	DisableApiTermination() interface{}
	SetDisableApiTermination(val interface{})
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	//
	// Default: `false`.
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	// An elastic GPU to associate with the instance.
	//
	// An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see [Amazon EC2 Elastic GPUs](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the *Amazon EC2 User Guide* .
	ElasticGpuSpecifications() interface{}
	SetElasticGpuSpecifications(val interface{})
	// An elastic inference accelerator to associate with the instance.
	//
	// Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.
	//
	// You cannot specify accelerators from different generations in the same request.
	ElasticInferenceAccelerators() interface{}
	SetElasticInferenceAccelerators(val interface{})
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	EnclaveOptions() interface{}
	SetEnclaveOptions(val interface{})
	// Indicates whether an instance is enabled for hibernation.
	//
	// For more information, see [Hibernate your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon EC2 User Guide* .
	//
	// You can't enable hibernation and AWS Nitro Enclaves on the same instance.
	HibernationOptions() interface{}
	SetHibernationOptions(val interface{})
	// If you specify host for the `Affinity` property, the ID of a dedicated host that the instance is associated with.
	//
	// If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account. This type of launch is called an untargeted launch. Note that for untargeted launches, you must have a compatible, dedicated host available to successfully launch instances.
	HostId() *string
	SetHostId(val *string)
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to `host` .
	HostResourceGroupArn() *string
	SetHostResourceGroupArn(val *string)
	// The name of an IAM instance profile.
	//
	// To create a new IAM instance profile, use the [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html) resource.
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	// The ID of the AMI.
	//
	// An AMI ID is required to launch an instance and must be specified here or in a launch template.
	ImageId() *string
	SetImageId(val *string)
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Default: `stop`.
	InstanceInitiatedShutdownBehavior() *string
	SetInstanceInitiatedShutdownBehavior(val *string)
	// The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `m1.small`
	InstanceType() *string
	SetInstanceType(val *string)
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	//
	// You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	Ipv6Addresses() interface{}
	SetIpv6Addresses(val interface{})
	// The ID of the kernel.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [PV-GRUB](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	KernelId() *string
	SetKernelId(val *string)
	// The name of the key pair. You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html) .
	//
	// > If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.
	KeyName() *string
	SetKeyName(val *string)
	// The launch template to use to launch the instances.
	//
	// Any parameters that you specify in the AWS CloudFormation template override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	// The license configurations.
	LicenseSpecifications() interface{}
	SetLicenseSpecifications(val interface{})
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
	// Specifies whether detailed monitoring is enabled for the instance.
	//
	// Specify `true` to enable detailed monitoring. Otherwise, basic monitoring is enabled. For more information about detailed monitoring, see [Enable or turn off detailed monitoring for your instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html) in the *Amazon EC2 User Guide* .
	Monitoring() interface{}
	SetMonitoring(val interface{})
	// The network interfaces to associate with the instance.
	//
	// > If you use this property to point to a network interface, you must terminate the original interface before attaching a new one to allow the update of the instance to succeed.
	// >
	// > If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the VPC-gateway attachment.
	NetworkInterfaces() interface{}
	SetNetworkInterfaces(val interface{})
	// The tree node.
	Node() constructs.Node
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	PlacementGroupName() *string
	SetPlacementGroupName(val *string)
	// The options for the instance hostname.
	PrivateDnsNameOptions() interface{}
	SetPrivateDnsNameOptions(val interface{})
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	//
	// Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	//
	// If you make an update to an instance that requires replacement, you must assign a new private IP address. During a replacement, AWS CloudFormation creates a new instance but doesn't delete the old instance until the stack has successfully updated. If the stack update fails, AWS CloudFormation uses the old instance to roll back the stack to the previous working state. The old and new instances cannot have the same private IP address.
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch.
	//
	// If you specify `true` and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify `false` , those tags are not assigned to the attached volumes.
	PropagateTagsToVolumeOnCreation() interface{}
	SetPropagateTagsToVolumeOnCreation(val interface{})
	// The ID of the RAM disk to select.
	//
	// Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the AWS Resource Center and search for the kernel ID.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [PV-GRUB](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	RamdiskId() *string
	SetRamdiskId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IDs of the security groups. You can create a security group using [CreateSecurityGroup](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html) .
	//
	// If you specify a network interface, you must specify any security groups as part of the network interface.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// [EC2-Classic, default VPC] The names of the security groups.
	//
	// For a nondefault VPC, you must use security group IDs instead.
	//
	// You cannot specify this option and the network interfaces option in the same request. The list can contain both the name of existing Amazon EC2 security groups or references to AWS::EC2::SecurityGroup resources created in the template.
	//
	// Default: Amazon EC2 uses the default security group.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.
	//
	// If the value is `true` , source/destination checks are enabled; otherwise, they are disabled. The default value is `true` . You must disable source/destination checks if the instance runs services such as network address translation, routing, or firewalls.
	SourceDestCheck() interface{}
	SetSourceDestCheck(val interface{})
	// The SSM [document](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html) and parameter values in AWS Systems Manager to associate with this instance. To use this property, you must specify an IAM instance profile role for the instance. For more information, see [Create an Instance Profile for Systems Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-configuring-access-role.html) in the *AWS Systems Manager User Guide* .
	//
	// > You can currently associate only one document with an instance.
	SsmAssociations() interface{}
	SetSsmAssociations(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	//
	// If you specify a network interface, you must specify any subnets as part of the network interface.
	SubnetId() *string
	SetSubnetId(val *string)
	// The tags to add to the instance.
	//
	// These tags are not applied to the EBS volumes, such as the root volume.
	Tags() awscdk.TagManager
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of `dedicated` runs on single-tenant hardware.
	Tenancy() *string
	SetTenancy(val *string)
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
	// The user data script to make available to the instance.
	//
	// For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and [Run commands on your Windows instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html) . If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.
	UserData() *string
	SetUserData(val *string)
	// The volumes to attach to the instance.
	Volumes() interface{}
	SetVolumes(val interface{})
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

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Affinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinity",
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

func (j *jsiiProxy_CfnInstance) CpuOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuOptions",
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

func (j *jsiiProxy_CfnInstance) CreditSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
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

func (j *jsiiProxy_CfnInstance) ElasticGpuSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticGpuSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) ElasticInferenceAccelerators() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticInferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) EnclaveOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) HibernationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
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

func (j *jsiiProxy_CfnInstance) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ipv6Addresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LaunchTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LicenseSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecifications",
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

func (j *jsiiProxy_CfnInstance) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) NetworkInterfaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) PlacementGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) PrivateDnsNameOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) PropagateTagsToVolumeOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTagsToVolumeOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) RamdiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskId",
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

func (j *jsiiProxy_CfnInstance) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) SsmAssociations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmAssociations",
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

func (j *jsiiProxy_CfnInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
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

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Volumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	if err := validateNewCfnInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance)SetAdditionalInfo(val *string) {
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetAffinity(val *string) {
	_jsii_.Set(
		j,
		"affinity",
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

func (j *jsiiProxy_CfnInstance)SetCpuOptions(val interface{}) {
	if err := j.validateSetCpuOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuOptions",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetCreditSpecification(val interface{}) {
	if err := j.validateSetCreditSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
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

func (j *jsiiProxy_CfnInstance)SetElasticGpuSpecifications(val interface{}) {
	if err := j.validateSetElasticGpuSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticGpuSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetElasticInferenceAccelerators(val interface{}) {
	if err := j.validateSetElasticInferenceAcceleratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticInferenceAccelerators",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetEnclaveOptions(val interface{}) {
	if err := j.validateSetEnclaveOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enclaveOptions",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetHibernationOptions(val interface{}) {
	if err := j.validateSetHibernationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetHostId(val *string) {
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetHostResourceGroupArn(val *string) {
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetIamInstanceProfile(val *string) {
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetInstanceInitiatedShutdownBehavior(val *string) {
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetIpv6AddressCount(val *float64) {
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetIpv6Addresses(val interface{}) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetKernelId(val *string) {
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetKeyName(val *string) {
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetLaunchTemplate(val interface{}) {
	if err := j.validateSetLaunchTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetLicenseSpecifications(val interface{}) {
	if err := j.validateSetLicenseSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetNetworkInterfaces(val interface{}) {
	if err := j.validateSetNetworkInterfacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaces",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetPlacementGroupName(val *string) {
	_jsii_.Set(
		j,
		"placementGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetPrivateDnsNameOptions(val interface{}) {
	if err := j.validateSetPrivateDnsNameOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsNameOptions",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetPrivateIpAddress(val *string) {
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetPropagateTagsToVolumeOnCreation(val interface{}) {
	if err := j.validateSetPropagateTagsToVolumeOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTagsToVolumeOnCreation",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetRamdiskId(val *string) {
	_jsii_.Set(
		j,
		"ramdiskId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetSsmAssociations(val interface{}) {
	if err := j.validateSetSsmAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssmAssociations",
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

func (j *jsiiProxy_CfnInstance)SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_CfnInstance)SetVolumes(val interface{}) {
	if err := j.validateSetVolumesParameters(val); err != nil {
		panic(err)
	}
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
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnInstance",
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
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnInstance",
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
		"aws-cdk-lib.aws_ec2.CfnInstance",
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

