package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties of a LaunchTemplate.
//
// Example:
//   bootHookConf := ec2.userData.forLinux()
//   bootHookConf.addCommands(jsii.String("cloud-init-per once docker_options echo 'OPTIONS=\"${OPTIONS} --storage-opt dm.basesize=40G\"' >> /etc/sysconfig/docker"))
//
//   setupCommands := ec2.userData.forLinux()
//   setupCommands.addCommands(jsii.String("sudo yum install awscli && echo Packages installed らと > /var/tmp/setup"))
//
//   multipartUserData := ec2.NewMultipartUserData()
//   // The docker has to be configured at early stage, so content type is overridden to boothook
//   multipartUserData.addPart(ec2.multipartBody.fromUserData(bootHookConf, jsii.String("text/cloud-boothook; charset=\"us-ascii\"")))
//   // Execute the rest of setup
//   multipartUserData.addPart(ec2.multipartBody.fromUserData(setupCommands))
//
//   ec2.NewLaunchTemplate(this, jsii.String(""), &launchTemplateProps{
//   	userData: multipartUserData,
//   	blockDevices: []blockDevice{
//   	},
//   })
//
// Experimental.
type LaunchTemplateProps struct {
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	// Experimental.
	BlockDevices *[]*BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// CPU credit type for burstable EC2 instance types.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html
	//
	// Experimental.
	CpuCredits CpuCredits `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
	// If set to true, then detailed monitoring will be enabled on instances created with this launch template.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html
	//
	// Experimental.
	DetailedMonitoring *bool `field:"optional" json:"detailedMonitoring" yaml:"detailedMonitoring"`
	// If you set this parameter to true, you cannot terminate the instances launched with this launch template using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can.
	// Experimental.
	DisableApiTermination *bool `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instances are optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput
	// to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization
	// isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	// Experimental.
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// If you set this parameter to true, the instance is enabled for hibernation.
	// Experimental.
	HibernationConfigured *bool `field:"optional" json:"hibernationConfigured" yaml:"hibernationConfigured"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
	//
	// Experimental.
	InstanceInitiatedShutdownBehavior InstanceInitiatedShutdownBehavior `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Type of instance to launch.
	// Experimental.
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Name of SSH keypair to grant access to instance.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Name for this launch template.
	// Experimental.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The AMI that will be used by instances.
	// Experimental.
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	// Experimental.
	NitroEnclaveEnabled *bool `field:"optional" json:"nitroEnclaveEnabled" yaml:"nitroEnclaveEnabled"`
	// Whether IMDSv2 should be required on launched instances.
	// Experimental.
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// An IAM role to associate with the instance profile that is used by instances.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &roleProps{
	//   	assumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security group to assign to instances created with the launch template.
	// Experimental.
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If this property is defined, then the Launch Template's InstanceMarketOptions will be set to use Spot instances, and the options for the Spot instances will be as defined.
	// Experimental.
	SpotOptions *LaunchTemplateSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The AMI that will be used by instances.
	// Experimental.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

