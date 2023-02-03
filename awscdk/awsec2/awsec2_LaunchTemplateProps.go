package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of a LaunchTemplate.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   template := ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &launchTemplateProps{
//   	machineImage: ec2.machineImage.latestAmazonLinux(),
//   	securityGroup: ec2.NewSecurityGroup(this, jsii.String("LaunchTemplateSG"), &securityGroupProps{
//   		vpc: vpc,
//   	}),
//   })
//
type LaunchTemplateProps struct {
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// CPU credit type for burstable EC2 instance types.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html
	//
	CpuCredits CpuCredits `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
	// If set to true, then detailed monitoring will be enabled on instances created with this launch template.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html
	//
	DetailedMonitoring *bool `field:"optional" json:"detailedMonitoring" yaml:"detailedMonitoring"`
	// If you set this parameter to true, you cannot terminate the instances launched with this launch template using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can.
	DisableApiTermination *bool `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instances are optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput
	// to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization
	// isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// If you set this parameter to true, the instance is enabled for hibernation.
	HibernationConfigured *bool `field:"optional" json:"hibernationConfigured" yaml:"hibernationConfigured"`
	// Enables or disables the HTTP metadata endpoint on your instances.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpendpoint
	//
	HttpEndpoint *bool `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpprotocolipv6
	//
	HttpProtocolIpv6 *bool `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httpputresponsehoplimit
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// The state of token usage for your instance metadata requests.
	//
	// The default state is `optional` if not specified. However,
	// if requireImdsv2 is true, the state must be `required`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens
	//
	HttpTokens LaunchTemplateHttpTokens `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
	//
	InstanceInitiatedShutdownBehavior InstanceInitiatedShutdownBehavior `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Set to enabled to allow access to instance tags from the instance metadata.
	//
	// Set to disabled to turn off access to instance tags from the instance metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-instancemetadatatags
	//
	InstanceMetadataTags *bool `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
	// Type of instance to launch.
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Name of SSH keypair to grant access to instance.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Name for this launch template.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The AMI that will be used by instances.
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	NitroEnclaveEnabled *bool `field:"optional" json:"nitroEnclaveEnabled" yaml:"nitroEnclaveEnabled"`
	// Whether IMDSv2 should be required on launched instances.
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// An IAM role to associate with the instance profile that is used by instances.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   role := iam.NewRole(this, jsii.String("MyRole"), &roleProps{
	//   	assumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security group to assign to instances created with the launch template.
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If this property is defined, then the Launch Template's InstanceMarketOptions will be set to use Spot instances, and the options for the Spot instances will be as defined.
	SpotOptions *LaunchTemplateSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The AMI that will be used by instances.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

