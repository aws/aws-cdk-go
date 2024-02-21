package awsec2


// Properties of the bastion host.
//
// Example:
//   host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &BastionHostLinuxProps{
//   	Vpc: Vpc,
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("/dev/sdh"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(10), &EbsDeviceOptions{
//   				Encrypted: jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   })
//
type BastionHostLinuxProps struct {
	// VPC to launch the instance in.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// In which AZ to place the instance within the VPC.
	// Default: - Random zone.
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	// Default: - Uses the block device mapping of the AMI.
	//
	BlockDevices *[]*BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Apply the given CloudFormation Init configuration to the instance at startup.
	// Default: - no CloudFormation init.
	//
	Init CloudFormationInit `field:"optional" json:"init" yaml:"init"`
	// Use the given options for applying CloudFormation Init.
	//
	// Describes the configsets to use and the timeout to wait.
	// Default: - default options.
	//
	InitOptions *ApplyCloudFormationInitOptions `field:"optional" json:"initOptions" yaml:"initOptions"`
	// The name of the instance.
	// Default: 'BastionHost'.
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Type of instance to launch.
	// Default: 't3.nano'
	//
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The machine image to use, assumed to have SSM Agent preinstalled.
	// Default: - An Amazon Linux 2 image which is kept up-to-date automatically (the instance
	// may be replaced on every deployment) and already has SSM Agent installed.
	//
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Whether IMDSv2 should be required on this instance.
	// Default: - false.
	//
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// Security Group to assign to this instance.
	// Default: - create new security group with no inbound and all outbound traffic allowed.
	//
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Select the subnets to run the bastion host in.
	//
	// Set this to PUBLIC if you need to connect to this instance via the internet and cannot use SSM.
	// You have to allow port 22 manually by using the connections field.
	// Default: - private subnets of the supplied VPC.
	//
	SubnetSelection *SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

