package awsopsworks


// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &cfnInstanceProps{
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
//   }
//
type CfnInstanceProps struct {
	// The instance type, such as `t2.micro` . For a list of supported instance types, open the stack in the console, choose *Instances* , and choose *+ Instance* . The *Size* list contains the currently supported types. For more information, see [Instance Families and Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) . The parameter values that you use to specify the various types are in the *API Name* column of the *Available Instance Types* table.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// An array that contains the instance's layer IDs.
	LayerIds *[]*string `field:"required" json:"layerIds" yaml:"layerIds"`
	// The stack ID.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The default AWS OpsWorks Stacks agent version. You have the following options:.
	//
	// - `INHERIT` - Use the stack's default agent version setting.
	// - *version_number* - Use the specified agent version. This value overrides the stack's default setting. To update the agent version, edit the instance configuration and specify a new version. AWS OpsWorks Stacks installs that version on the instance.
	//
	// The default setting is `INHERIT` . To specify an agent version, you must use the complete version number, not the abbreviated number shown on the console. For a list of available agent version numbers, call [DescribeAgentVersions](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeAgentVersions) . AgentVersion cannot be set to Chef 12.2.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// A custom AMI ID to be used to create the instance.
	//
	// The AMI should be based on one of the supported operating systems. For more information, see [Using Custom AMIs](https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html) .
	//
	// > If you specify a custom AMI, you must set `Os` to `Custom` .
	AmiId *string `field:"optional" json:"amiId" yaml:"amiId"`
	// The instance architecture.
	//
	// The default option is `x86_64` . Instance types do not necessarily support both architectures. For a list of the architectures that are supported by the different instance types, see [Instance Families and Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) .
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// For load-based or time-based instances, the type.
	//
	// Windows stacks can use only time-based instances.
	AutoScalingType *string `field:"optional" json:"autoScalingType" yaml:"autoScalingType"`
	// The Availability Zone of the AWS OpsWorks instance, such as `us-east-2a` .
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// An array of `BlockDeviceMapping` objects that specify the instance's block devices.
	//
	// For more information, see [Block Device Mapping](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) . Note that block device mappings are not supported for custom AMIs.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Whether to create an Amazon EBS-optimized instance.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// A list of Elastic IP addresses to associate with the instance.
	ElasticIps *[]*string `field:"optional" json:"elasticIps" yaml:"elasticIps"`
	// The instance host name. The following are character limits for instance host names.
	//
	// - Linux-based instances: 63 characters
	// - Windows-based instances: 15 characters.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Whether to install operating system and package updates when the instance boots.
	//
	// The default value is `true` . To control when updates are installed, set this value to `false` . You must then update your instances manually by using [CreateDeployment](https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateDeployment) to run the `update_dependencies` stack command or by manually running `yum` (Amazon Linux) or `apt-get` (Ubuntu) on the instances.
	//
	// > We strongly recommend using the default value of `true` to ensure that your instances have the latest security updates.
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
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
	Os *string `field:"optional" json:"os" yaml:"os"`
	// The instance root device type.
	//
	// For more information, see [Storage for the Root Device](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device) .
	RootDeviceType *string `field:"optional" json:"rootDeviceType" yaml:"rootDeviceType"`
	// The instance's Amazon EC2 key-pair name.
	SshKeyName *string `field:"optional" json:"sshKeyName" yaml:"sshKeyName"`
	// The ID of the instance's subnet.
	//
	// If the stack is running in a VPC, you can use this parameter to override the stack's default subnet ID value and direct AWS OpsWorks Stacks to launch the instance in a different subnet.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The instance's tenancy option.
	//
	// The default option is no tenancy, or if the instance is running in a VPC, inherit tenancy settings from the VPC. The following are valid values for this parameter: `dedicated` , `default` , or `host` . Because there are costs associated with changes in tenancy options, we recommend that you research tenancy options before choosing them for your instances. For more information about dedicated hosts, see [Dedicated Hosts Overview](https://docs.aws.amazon.com/ec2/dedicated-hosts/) and [Amazon EC2 Dedicated Hosts](https://docs.aws.amazon.com/ec2/dedicated-hosts/) . For more information about dedicated instances, see [Dedicated Instances](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/dedicated-instance.html) and [Amazon EC2 Dedicated Instances](https://docs.aws.amazon.com/ec2/purchasing-options/dedicated-instances/) .
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The time-based scaling configuration for the instance.
	TimeBasedAutoScaling interface{} `field:"optional" json:"timeBasedAutoScaling" yaml:"timeBasedAutoScaling"`
	// The instance's virtualization type, `paravirtual` or `hvm` .
	VirtualizationType *string `field:"optional" json:"virtualizationType" yaml:"virtualizationType"`
	// A list of AWS OpsWorks volume IDs to associate with the instance.
	//
	// For more information, see [`AWS::OpsWorks::Volume`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html) .
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

