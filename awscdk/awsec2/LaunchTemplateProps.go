package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of a LaunchTemplate.
//
// Example:
//   var vpc vpc
//
//
//   sg1 := ec2.NewSecurityGroup(this, jsii.String("sg1"), &SecurityGroupProps{
//   	Vpc: vpc,
//   })
//   sg2 := ec2.NewSecurityGroup(this, jsii.String("sg2"), &SecurityGroupProps{
//   	Vpc: vpc,
//   })
//
//   launchTemplate := ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &LaunchTemplateProps{
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	SecurityGroup: sg1,
//   })
//
//   launchTemplate.AddSecurityGroup(sg2)
//
type LaunchTemplateProps struct {
	// Whether instances should have a public IP addresses associated with them.
	// Default: - Use subnet settings.
	//
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
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
	// CPU credit type for burstable EC2 instance types.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html
	//
	// Default: - No credit type is specified in the Launch Template.
	//
	CpuCredits CpuCredits `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
	// If set to true, then detailed monitoring will be enabled on instances created with this launch template.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html
	//
	// Default: False - Detailed monitoring is disabled.
	//
	DetailedMonitoring *bool `field:"optional" json:"detailedMonitoring" yaml:"detailedMonitoring"`
	// If you set this parameter to true, you cannot terminate the instances launched with this launch template using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can.
	// Default: - The API termination setting is not specified in the Launch Template.
	//
	DisableApiTermination *bool `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instances are optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput
	// to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization
	// isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	// Default: - EBS optimization is not specified in the launch template.
	//
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// If you set this parameter to true, the instance is enabled for hibernation.
	// Default: - Hibernation configuration is not specified in the launch template; defaulting to false.
	//
	HibernationConfigured *bool `field:"optional" json:"hibernationConfigured" yaml:"hibernationConfigured"`
	// Enables or disables the HTTP metadata endpoint on your instances.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpendpoint
	//
	// Default: true.
	//
	HttpEndpoint *bool `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpprotocolipv6
	//
	// Default: true.
	//
	HttpProtocolIpv6 *bool `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpputresponsehoplimit
	//
	// Default: 1.
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// The state of token usage for your instance metadata requests.
	//
	// The default state is `optional` if not specified. However,
	// if requireImdsv2 is true, the state must be `required`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens
	//
	// Default: LaunchTemplateHttpTokens.OPTIONAL
	//
	HttpTokens LaunchTemplateHttpTokens `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
	//
	// Default: - Shutdown behavior is not specified in the launch template; defaults to STOP.
	//
	InstanceInitiatedShutdownBehavior InstanceInitiatedShutdownBehavior `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Set to enabled to allow access to instance tags from the instance metadata.
	//
	// Set to disabled to turn off access to instance tags from the instance metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-instancemetadatatags
	//
	// Default: false.
	//
	InstanceMetadataTags *bool `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
	// The instance profile used to pass role information to EC2 instances.
	//
	// Note: You can provide an instanceProfile or a role, but not both.
	// Default: - No instance profile.
	//
	InstanceProfile awsiam.IInstanceProfile `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
	// Type of instance to launch.
	// Default: - This Launch Template does not specify a default Instance Type.
	//
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Name of SSH keypair to grant access to instance.
	// Default: - No SSH access will be possible.
	//
	// Deprecated: - Use `keyPair` instead - https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2-readme.html#using-an-existing-ec2-key-pair
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The SSH keypair to grant access to the instance.
	// Default: - No SSH access will be possible.
	//
	KeyPair IKeyPair `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Name for this launch template.
	// Default: Automatically generated name.
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The AMI that will be used by instances.
	// Default: - This Launch Template does not specify a default AMI.
	//
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	// Default: - Enablement of Nitro enclaves is not specified in the launch template; defaulting to false.
	//
	NitroEnclaveEnabled *bool `field:"optional" json:"nitroEnclaveEnabled" yaml:"nitroEnclaveEnabled"`
	// Whether IMDSv2 should be required on launched instances.
	// Default: - false.
	//
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// An IAM role to associate with the instance profile that is used by instances.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`.
	// Note: You can provide an instanceProfile or a role, but not both.
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	// Default: - No new role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security group to assign to instances created with the launch template.
	// Default: No security group is assigned.
	//
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If this property is defined, then the Launch Template's InstanceMarketOptions will be set to use Spot instances, and the options for the Spot instances will be as defined.
	// Default: - Instance launched with this template will not be spot instances.
	//
	SpotOptions *LaunchTemplateSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The AMI that will be used by instances.
	// Default: - This Launch Template creates a UserData based on the type of provided
	// machineImage; no UserData is created if a machineImage is not provided.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

