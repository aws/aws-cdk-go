package awsautoscaling


// Properties for defining a `CfnLaunchConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchConfigurationProps := &CfnLaunchConfigurationProps{
//   	ImageId: jsii.String("imageId"),
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//
//   			// the properties below are optional
//   			Ebs: &BlockDeviceProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				Throughput: jsii.Number(123),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.Boolean(false),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	ClassicLinkVpcId: jsii.String("classicLinkVpcId"),
//   	ClassicLinkVpcSecurityGroups: []*string{
//   		jsii.String("classicLinkVpcSecurityGroups"),
//   	},
//   	EbsOptimized: jsii.Boolean(false),
//   	IamInstanceProfile: jsii.String("iamInstanceProfile"),
//   	InstanceId: jsii.String("instanceId"),
//   	InstanceMonitoring: jsii.Boolean(false),
//   	KernelId: jsii.String("kernelId"),
//   	KeyName: jsii.String("keyName"),
//   	LaunchConfigurationName: jsii.String("launchConfigurationName"),
//   	MetadataOptions: &MetadataOptionsProperty{
//   		HttpEndpoint: jsii.String("httpEndpoint"),
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   	},
//   	PlacementTenancy: jsii.String("placementTenancy"),
//   	RamDiskId: jsii.String("ramDiskId"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SpotPrice: jsii.String("spotPrice"),
//   	UserData: jsii.String("userData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html
//
type CfnLaunchConfigurationProps struct {
	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	//
	// For more information, see [Finding a Linux AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// If you specify `InstanceId` , an `ImageId` is not required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-imageid
	//
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Specifies the instance type of the EC2 instance.
	//
	// For information about available instance types, see [Available instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// If you specify `InstanceId` , an `InstanceType` is not required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Specifies whether to assign a public IPv4 address to the group's instances.
	//
	// If the instance is launched into a default subnet, the default is to assign a public IPv4 address, unless you disabled the option to assign a public IPv4 address on the subnet. If the instance is launched into a nondefault subnet, the default is not to assign a public IPv4 address, unless you enabled the option to assign a public IPv4 address on the subnet.
	//
	// If you specify `true` , each instance in the Auto Scaling group receives a unique public IPv4 address. For more information, see [Launching Auto Scaling instances in a VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify this property, you must specify at least one subnet for `VPCZoneIdentifier` when you create your group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-associatepublicipaddress
	//
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// The block device mapping entries that define the block devices to attach to the instances at launch.
	//
	// By default, the block devices specified in the block device mapping for the AMI are used. For more information, see [Block device mappings](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-blockdevicemappings
	//
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Available for backward compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-classiclinkvpcid
	//
	ClassicLinkVpcId *string `field:"optional" json:"classicLinkVpcId" yaml:"classicLinkVpcId"`
	// Available for backward compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-classiclinkvpcsecuritygroups
	//
	ClassicLinkVpcSecurityGroups *[]*string `field:"optional" json:"classicLinkVpcSecurityGroups" yaml:"classicLinkVpcSecurityGroups"`
	// Specifies whether the launch configuration is optimized for EBS I/O ( `true` ) or not ( `false` ).
	//
	// The optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization is not available with all instance types. Additional fees are incurred when you enable EBS optimization for an instance type that is not EBS-optimized by default. For more information, see [Amazon EBS-optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.
	//
	// The instance profile contains the IAM role. For more information, see [IAM role for applications that run on Amazon EC2 instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-iaminstanceprofile
	//
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the Amazon EC2 instance to use to create the launch configuration.
	//
	// When you use an instance to create a launch configuration, all properties are derived from the instance with the exception of `BlockDeviceMapping` and `AssociatePublicIpAddress` . You can override any properties from the instance by specifying them in the launch configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Controls whether instances in this group are launched with detailed ( `true` ) or basic ( `false` ) monitoring.
	//
	// The default value is `true` (enabled).
	//
	// > When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes. For more information, see [Configure Monitoring for Auto Scaling Instances](https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-instancemonitoring
	//
	InstanceMonitoring interface{} `field:"optional" json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// The ID of the kernel associated with the AMI.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-kernelid
	//
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair.
	//
	// For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-keyname
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The name of the launch configuration.
	//
	// This name must be unique per Region per account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-launchconfigurationname
	//
	LaunchConfigurationName *string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// The metadata options for the instances.
	//
	// For more information, see [Configuring the Instance Metadata Options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-metadataoptions
	//
	MetadataOptions interface{} `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// The tenancy of the instance, either `default` or `dedicated` .
	//
	// An instance with `dedicated` tenancy runs on isolated, single-tenant hardware and can only be launched into a VPC. To launch dedicated instances into a shared tenancy VPC (a VPC with the instance placement tenancy attribute set to `default` ), you must set the value of this property to `dedicated` . For more information, see [Configuring instance tenancy with Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-dedicated-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// If you specify `PlacementTenancy` , you must specify at least one subnet for `VPCZoneIdentifier` when you create your group.
	//
	// Valid values: `default` | `dedicated`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-placementtenancy
	//
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// The ID of the RAM disk to select.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-ramdiskid
	//
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// A list that contains the security groups to assign to the instances in the Auto Scaling group.
	//
	// The list can contain both the IDs of existing security groups and references to [SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
	//
	// For more information, see [Control traffic to resources using security groups](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The maximum hourly price to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are launched when the price you specify exceeds the current Spot price. For more information, see [Request Spot Instances for fault-tolerant and flexible applications](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-template-spot-instances.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid Range: Minimum value of 0.001
	//
	// > When you change your maximum price by creating a new launch configuration, running instances will continue to run as long as the maximum price for those running instances is higher than the current Spot price.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-spotprice
	//
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The Base64-encoded user data to make available to the launched EC2 instances.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html#cfn-autoscaling-launchconfiguration-userdata
	//
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

