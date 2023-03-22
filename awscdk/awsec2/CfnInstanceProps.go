package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &CfnInstanceProps{
//   	AdditionalInfo: jsii.String("additionalInfo"),
//   	Affinity: jsii.String("affinity"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//
//   			// the properties below are optional
//   			Ebs: &EbsProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				SnapshotId: jsii.String("snapshotId"),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: &NoDeviceProperty{
//   			},
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	CpuOptions: &CpuOptionsProperty{
//   		CoreCount: jsii.Number(123),
//   		ThreadsPerCore: jsii.Number(123),
//   	},
//   	CreditSpecification: &CreditSpecificationProperty{
//   		CpuCredits: jsii.String("cpuCredits"),
//   	},
//   	DisableApiTermination: jsii.Boolean(false),
//   	EbsOptimized: jsii.Boolean(false),
//   	ElasticGpuSpecifications: []interface{}{
//   		&ElasticGpuSpecificationProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ElasticInferenceAccelerators: []interface{}{
//   		&ElasticInferenceAcceleratorProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Count: jsii.Number(123),
//   		},
//   	},
//   	EnclaveOptions: &EnclaveOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	HibernationOptions: &HibernationOptionsProperty{
//   		Configured: jsii.Boolean(false),
//   	},
//   	HostId: jsii.String("hostId"),
//   	HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   	IamInstanceProfile: jsii.String("iamInstanceProfile"),
//   	ImageId: jsii.String("imageId"),
//   	InstanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   	InstanceType: jsii.String("instanceType"),
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&InstanceIpv6AddressProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	KernelId: jsii.String("kernelId"),
//   	KeyName: jsii.String("keyName"),
//   	LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	LicenseSpecifications: []interface{}{
//   		&LicenseSpecificationProperty{
//   			LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   		},
//   	},
//   	Monitoring: jsii.Boolean(false),
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfaceProperty{
//   			DeviceIndex: jsii.String("deviceIndex"),
//
//   			// the properties below are optional
//   			AssociateCarrierIpAddress: jsii.Boolean(false),
//   			AssociatePublicIpAddress: jsii.Boolean(false),
//   			DeleteOnTermination: jsii.Boolean(false),
//   			Description: jsii.String("description"),
//   			GroupSet: []*string{
//   				jsii.String("groupSet"),
//   			},
//   			Ipv6AddressCount: jsii.Number(123),
//   			Ipv6Addresses: []interface{}{
//   				&InstanceIpv6AddressProperty{
//   					Ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   			PrivateIpAddresses: []interface{}{
//   				&PrivateIpAddressSpecificationProperty{
//   					Primary: jsii.Boolean(false),
//   					PrivateIpAddress: jsii.String("privateIpAddress"),
//   				},
//   			},
//   			SecondaryPrivateIpAddressCount: jsii.Number(123),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	PlacementGroupName: jsii.String("placementGroupName"),
//   	PrivateDnsNameOptions: &PrivateDnsNameOptionsProperty{
//   		EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   		EnableResourceNameDnsARecord: jsii.Boolean(false),
//   		HostnameType: jsii.String("hostnameType"),
//   	},
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PropagateTagsToVolumeOnCreation: jsii.Boolean(false),
//   	RamdiskId: jsii.String("ramdiskId"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SourceDestCheck: jsii.Boolean(false),
//   	SsmAssociations: []interface{}{
//   		&SsmAssociationProperty{
//   			DocumentName: jsii.String("documentName"),
//
//   			// the properties below are optional
//   			AssociationParameters: []interface{}{
//   				&AssociationParameterProperty{
//   					Key: jsii.String("key"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tenancy: jsii.String("tenancy"),
//   	UserData: jsii.String("userData"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			Device: jsii.String("device"),
//   			VolumeId: jsii.String("volumeId"),
//   		},
//   	},
//   }
//
type CfnInstanceProps struct {
	// This property is reserved for internal use.
	//
	// If you use it, the stack fails with this error: `Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX)` .
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Indicates whether the instance is associated with a dedicated host.
	//
	// If you want the instance to always restart on the same host on which it was launched, specify `host` . If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify `default` .
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone of the instance.
	//
	// If not specified, an Availability Zone will be automatically chosen for you based on the load balancing criteria for the Region.
	//
	// This parameter is not supported by [DescribeImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImageAttribute.html) .
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	//
	// By default, the block devices specified in the block device mapping for the AMI are used. You can override the AMI block device mapping using the instance block device mapping. For the root volume, you can override only the volume size, volume type, volume encryption settings, and the `DeleteOnTermination` setting.
	//
	// > After the instance is running, you can modify only the `DeleteOnTermination` parameter for the attached volumes without interrupting the instance. Modifying any other parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The CPU options for the instance.
	//
	// For more information, see [Optimize CPU options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
	CpuOptions interface{} `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The credit option for CPU usage of the burstable performance instance.
	//
	// Valid values are `standard` and `unlimited` . To change this attribute after launch, use [ModifyInstanceCreditSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html) . For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `standard` (T2 instances) or `unlimited` (T3/T3a/T4g instances)
	//
	// For T3 instances with `host` tenancy, only `standard` is supported.
	CreditSpecification interface{} `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// If you set this parameter to `true` , you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html) . Alternatively, if you set `InstanceInitiatedShutdownBehavior` to `terminate` , you can terminate the instance by running the shutdown command from the instance.
	//
	// Default: `false`.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	//
	// Default: `false`.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// An elastic GPU to associate with the instance.
	//
	// An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see [Amazon EC2 Elastic GPUs](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the *Amazon EC2 User Guide* .
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance.
	//
	// Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.
	//
	// You cannot specify accelerators from different generations in the same request.
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	EnclaveOptions interface{} `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	//
	// For more information, see [Hibernate your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon EC2 User Guide* .
	//
	// You can't enable hibernation and AWS Nitro Enclaves on the same instance.
	HibernationOptions interface{} `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// If you specify host for the `Affinity` property, the ID of a dedicated host that the instance is associated with.
	//
	// If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account. This type of launch is called an untargeted launch. Note that for untargeted launches, you must have a compatible, dedicated host available to successfully launch instances.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to `host` .
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The name of an IAM instance profile.
	//
	// To create a new IAM instance profile, use the [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html) resource.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	//
	// An AMI ID is required to launch an instance and must be specified here or in a launch template.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Default: `stop`.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `m1.small`
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	//
	// You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The ID of the kernel.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [PV-GRUB](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair. You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html) .
	//
	// > If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The launch template to use to launch the instances.
	//
	// Any parameters that you specify in the AWS CloudFormation template override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The license configurations.
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// Specifies whether detailed monitoring is enabled for the instance.
	//
	// Specify `true` to enable detailed monitoring. Otherwise, basic monitoring is enabled. For more information about detailed monitoring, see [Enable or turn off detailed monitoring for your instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html) in the *Amazon EC2 User Guide* .
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The network interfaces to associate with the instance.
	//
	// > If you use this property to point to a network interface, you must terminate the original interface before attaching a new one to allow the update of the instance to succeed.
	// >
	// > If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the VPC-gateway attachment.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	PlacementGroupName *string `field:"optional" json:"placementGroupName" yaml:"placementGroupName"`
	// The options for the instance hostname.
	PrivateDnsNameOptions interface{} `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	//
	// Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request.
	//
	// You cannot specify this option and the network interfaces option in the same request.
	//
	// If you make an update to an instance that requires replacement, you must assign a new private IP address. During a replacement, AWS CloudFormation creates a new instance but doesn't delete the old instance until the stack has successfully updated. If the stack update fails, AWS CloudFormation uses the old instance to roll back the stack to the previous working state. The old and new instances cannot have the same private IP address.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch.
	//
	// If you specify `true` and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify `false` , those tags are not assigned to the attached volumes.
	PropagateTagsToVolumeOnCreation interface{} `field:"optional" json:"propagateTagsToVolumeOnCreation" yaml:"propagateTagsToVolumeOnCreation"`
	// The ID of the RAM disk to select.
	//
	// Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the AWS Resource Center and search for the kernel ID.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [PV-GRUB](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// The IDs of the security groups. You can create a security group using [CreateSecurityGroup](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html) .
	//
	// If you specify a network interface, you must specify any security groups as part of the network interface.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// [EC2-Classic, default VPC] The names of the security groups.
	//
	// For a nondefault VPC, you must use security group IDs instead.
	//
	// You cannot specify this option and the network interfaces option in the same request. The list can contain both the name of existing Amazon EC2 security groups or references to AWS::EC2::SecurityGroup resources created in the template.
	//
	// Default: Amazon EC2 uses the default security group.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.
	//
	// If the value is `true` , source/destination checks are enabled; otherwise, they are disabled. The default value is `true` . You must disable source/destination checks if the instance runs services such as network address translation, routing, or firewalls.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// The SSM [document](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html) and parameter values in AWS Systems Manager to associate with this instance. To use this property, you must specify an IAM instance profile role for the instance. For more information, see [Create an IAM instance profile for Systems Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-configuring-access-role.html) in the *AWS Systems Manager User Guide* .
	//
	// > You can currently associate only one document with an instance.
	SsmAssociations interface{} `field:"optional" json:"ssmAssociations" yaml:"ssmAssociations"`
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	//
	// If you specify a network interface, you must specify any subnets as part of the network interface.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags to add to the instance.
	//
	// These tags are not applied to the EBS volumes, such as the root volume, unless [PropagateTagsToVolumeOnCreation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-propagatetagstovolumeoncreation) is `true` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of `dedicated` runs on single-tenant hardware.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The user data script to make available to the instance.
	//
	// User data is limited to 16 KB. You must provide base64-encoded text. For more information, see [Fn::Base64](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-base64.html) .
	//
	// User data runs only at instance launch. For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and [Run commands on your Windows instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html) .
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// The volumes to attach to the instance.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

