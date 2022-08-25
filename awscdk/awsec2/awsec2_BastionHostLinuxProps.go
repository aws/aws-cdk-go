package awsec2


// Properties of the bastion host.
//
// Example:
//   host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &bastionHostLinuxProps{
//   	vpc: vpc,
//   	blockDevices: []blockDevice{
//   		&blockDevice{
//   			deviceName: jsii.String("EBSBastionHost"),
//   			volume: ec2.blockDeviceVolume.ebs(jsii.Number(10), &ebsDeviceOptions{
//   				encrypted: jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   })
//
type BastionHostLinuxProps struct {
	// VPC to launch the instance in.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// In which AZ to place the instance within the VPC.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Apply the given CloudFormation Init configuration to the instance at startup.
	Init CloudFormationInit `field:"optional" json:"init" yaml:"init"`
	// Use the given options for applying CloudFormation Init.
	//
	// Describes the configsets to use and the timeout to wait.
	InitOptions *ApplyCloudFormationInitOptions `field:"optional" json:"initOptions" yaml:"initOptions"`
	// The name of the instance.
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Type of instance to launch.
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The machine image to use, assumed to have SSM Agent preinstalled.
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Whether IMDSv2 should be required on this instance.
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// Security Group to assign to this instance.
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Select the subnets to run the bastion host in.
	//
	// Set this to PUBLIC if you need to connect to this instance via the internet and cannot use SSM.
	// You have to allow port 22 manually by using the connections field.
	SubnetSelection *SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

