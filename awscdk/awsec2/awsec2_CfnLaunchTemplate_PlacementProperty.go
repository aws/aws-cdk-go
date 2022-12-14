package awsec2


// Specifies the placement of an instance.
//
// `Placement` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementProperty := &placementProperty{
//   	affinity: jsii.String("affinity"),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	groupId: jsii.String("groupId"),
//   	groupName: jsii.String("groupName"),
//   	hostId: jsii.String("hostId"),
//   	hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   	partitionNumber: jsii.Number(123),
//   	spreadDomain: jsii.String("spreadDomain"),
//   	tenancy: jsii.String("tenancy"),
//   }
//
type CfnLaunchTemplate_PlacementProperty struct {
	// The affinity setting for an instance on a Dedicated Host.
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone for the instance.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// `CfnLaunchTemplate.PlacementProperty.GroupId`.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The name of the placement group for the instance.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ID of the Dedicated Host for the instance.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to `host` .
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The number of the partition the instance should launch in.
	//
	// Valid only if the placement group strategy is set to `partition` .
	PartitionNumber *float64 `field:"optional" json:"partitionNumber" yaml:"partitionNumber"`
	// Reserved for future use.
	SpreadDomain *string `field:"optional" json:"spreadDomain" yaml:"spreadDomain"`
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of dedicated runs on single-tenant hardware.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

