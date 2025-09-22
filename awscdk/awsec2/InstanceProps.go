package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of an EC2 Instance.
//
// Example:
//   // Creates a distribution from an EC2 instance
//   var vpc vpc
//
//   // Create an EC2 instance in a VPC. 'subnetType' can be private.
//   instance := ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.VpcOrigin_WithEc2Instance(instance),
//   	},
//   })
//
type InstanceProps struct {
	// Type of instance to launch.
	InstanceType InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// AMI to launch.
	MachineImage IMachineImage `field:"required" json:"machineImage" yaml:"machineImage"`
	// VPC to launch the instance in.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether the instance could initiate IPv6 connections to anywhere by default.
	//
	// This property is only used when you do not provide a security group.
	// Default: false.
	//
	AllowAllIpv6Outbound *bool `field:"optional" json:"allowAllIpv6Outbound" yaml:"allowAllIpv6Outbound"`
	// Whether the instance could initiate connections to anywhere by default.
	//
	// This property is only used when you do not provide a security group.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Whether to associate a public IP address to the primary network interface attached to this instance.
	//
	// You cannot specify this property and `ipv6AddressCount` at the same time.
	// Default: - public IP address is automatically assigned based on default behavior.
	//
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
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
	// Specifying the CPU credit type for burstable EC2 instance types (T2, T3, T3a, etc).
	//
	// The unlimited CPU credit option is not supported for T3 instances with a dedicated host.
	// Default: - T2 instances are standard, while T3, T4g, and T3a instances are unlimited.
	//
	CreditSpecification CpuCredits `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Whether "Detailed Monitoring" is enabled for this instance Keep in mind that Detailed Monitoring results in extra charges.
	// See: http://aws.amazon.com/cloudwatch/pricing/
	//
	// Default: - false.
	//
	DetailedMonitoring *bool `field:"optional" json:"detailedMonitoring" yaml:"detailedMonitoring"`
	// If true, the instance will not be able to be terminated using the Amazon EC2 console, CLI, or API.
	//
	// To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html).
	// Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance
	// by running the shutdown command from the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instance.html#cfn-ec2-instance-disableapitermination
	//
	// Default: false.
	//
	DisableApiTermination *bool `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance.
	// This optimization isn't available with all instance types.
	// Additional usage charges apply when using an EBS-optimized instance.
	// Default: false.
	//
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Whether the instance is enabled for AWS Nitro Enclaves.
	//
	// Nitro Enclaves requires a Nitro-based virtualized parent instance with specific Intel/AMD with at least 4 vCPUs
	// or Graviton with at least 2 vCPUs instance types and Linux/Windows host OS,
	// while the enclave itself supports only Linux OS.
	//
	// You can't set both `enclaveEnabled` and `hibernationEnabled` to true on the same instance.
	// See: https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html#nitro-enclave-reqs
	//
	// Default: - false.
	//
	EnclaveEnabled *bool `field:"optional" json:"enclaveEnabled" yaml:"enclaveEnabled"`
	// Whether the instance is enabled for hibernation.
	//
	// You can't set both `enclaveEnabled` and `hibernationEnabled` to true on the same instance.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-hibernationoptions.html
	//
	// Default: - false.
	//
	HibernationEnabled *bool `field:"optional" json:"hibernationEnabled" yaml:"hibernationEnabled"`
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
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
	//
	// Default: InstanceInitiatedShutdownBehavior.STOP
	//
	InstanceInitiatedShutdownBehavior InstanceInitiatedShutdownBehavior `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The name of the instance.
	// Default: - CDK generated name.
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// The instance profile used to pass role information to EC2 instances.
	//
	// Note: You can provide an instanceProfile or a role, but not both.
	// Default: - No instance profile.
	//
	InstanceProfile awsiam.IInstanceProfile `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
	// The number of IPv6 addresses to associate with the primary network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	//
	// You cannot specify this property and `associatePublicIpAddress` at the same time.
	// Default: - For instances associated with an IPv6 subnet, use 1; otherwise, use 0.
	//
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Name of SSH keypair to grant access to instance.
	// Default: - No SSH access will be possible.
	//
	// Deprecated: - Use `keyPair` instead - https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2-readme.html#using-an-existing-ec2-key-pair
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The SSH keypair to grant access to the instance.
	// Default: - No SSH access will be possible.
	//
	KeyPair IKeyPair `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The placement group that you want to launch the instance into.
	// Default: - no placement group will be used for this instance.
	//
	PlacementGroup IPlacementGroupRef `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Defines a private IP address to associate with an instance.
	//
	// Private IP should be available within the VPC that the instance is build within.
	// Default: - no association.
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Propagate the EC2 instance tags to the EBS volumes.
	// Default: - false.
	//
	PropagateTagsToVolumeOnCreation *bool `field:"optional" json:"propagateTagsToVolumeOnCreation" yaml:"propagateTagsToVolumeOnCreation"`
	// Whether IMDSv2 should be required on this instance.
	// Default: - false.
	//
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// The length of time to wait for the resourceSignalCount.
	//
	// The maximum value is 43200 (12 hours).
	// Default: Duration.minutes(5)
	//
	ResourceSignalTimeout awscdk.Duration `field:"optional" json:"resourceSignalTimeout" yaml:"resourceSignalTimeout"`
	// An IAM role to associate with the instance profile assigned to this Auto Scaling Group.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	// Note: You can provide an instanceProfile or a role, but not both.
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	// Default: - A role will automatically be created, it can be accessed via the `role` property.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to assign to this instance.
	// Default: - create new security group.
	//
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	//
	// This controls whether source/destination checking is enabled on the instance.
	// A value of true means that checking is enabled, and false means that checking is disabled.
	// The value must be false for the instance to perform NAT.
	// Default: true.
	//
	SourceDestCheck *bool `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Add SSM session permissions to the instance role.
	//
	// Setting this to `true` adds the necessary permissions to connect
	// to the instance using SSM Session Manager. You can do this
	// from the AWS Console.
	//
	// NOTE: Setting this flag to `true` may not be enough by itself.
	// You must also use an AMI that comes with the SSM Agent, or install
	// the SSM Agent yourself. See
	// [Working with SSM Agent](https://docs.aws.amazon.com/systems-manager/latest/userguide/ssm-agent.html)
	// in the SSM Developer Guide.
	// Default: false.
	//
	SsmSessionPermissions *bool `field:"optional" json:"ssmSessionPermissions" yaml:"ssmSessionPermissions"`
	// Specific UserData to use.
	//
	// The UserData may still be mutated after creation.
	// Default: - A UserData object appropriate for the MachineImage's
	// Operating System is created.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// Changes to the UserData force replacement.
	//
	// Depending the EC2 instance type, changing UserData either
	// restarts the instance or replaces the instance.
	//
	// - Instance store-backed instances are replaced.
	// - EBS-backed instances are restarted.
	//
	// By default, restarting does not execute the new UserData so you
	// will need a different mechanism to ensure the instance is restarted.
	//
	// Setting this to `true` will make the instance's Logical ID depend on the
	// UserData, which will cause CloudFormation to replace it if the UserData
	// changes.
	// Default: - true if `initOptions` is specified, false otherwise.
	//
	UserDataCausesReplacement *bool `field:"optional" json:"userDataCausesReplacement" yaml:"userDataCausesReplacement"`
	// Where to place the instance within the VPC.
	// Default: - Private subnets.
	//
	VpcSubnets *SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

