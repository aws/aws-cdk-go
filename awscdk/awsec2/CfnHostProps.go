package awsec2


// Properties for defining a `CfnHost`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHostProps := &CfnHostProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//
//   	// the properties below are optional
//   	AssetId: jsii.String("assetId"),
//   	AutoPlacement: jsii.String("autoPlacement"),
//   	HostMaintenance: jsii.String("hostMaintenance"),
//   	HostRecovery: jsii.String("hostRecovery"),
//   	InstanceFamily: jsii.String("instanceFamily"),
//   	InstanceType: jsii.String("instanceType"),
//   	OutpostArn: jsii.String("outpostArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
//
type CfnHostProps struct {
	// The Availability Zone in which to allocate the Dedicated Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	//
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Outpost hardware asset on which the Dedicated Host is allocated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-assetid
	//
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	//
	// For more information, see [Understanding auto-placement and affinity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-dedicated-hosts-work.html#dedicated-hosts-understanding) in the *Amazon EC2 User Guide* .
	//
	// Default: `on`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	//
	AutoPlacement *string `field:"optional" json:"autoPlacement" yaml:"autoPlacement"`
	// Indicates whether host maintenance is enabled or disabled for the Dedicated Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-hostmaintenance
	//
	HostMaintenance *string `field:"optional" json:"hostMaintenance" yaml:"hostMaintenance"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host.
	//
	// Host recovery is disabled by default. For more information, see [Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `off`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-hostrecovery
	//
	HostRecovery *string `field:"optional" json:"hostRecovery" yaml:"hostRecovery"`
	// The instance family supported by the Dedicated Host.
	//
	// For example, `m5` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancefamily
	//
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
	// Specifies the instance type to be supported by the Dedicated Hosts.
	//
	// If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The Amazon Resource Name (ARN) of the AWS Outpost on which the Dedicated Host is allocated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-outpostarn
	//
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
}

