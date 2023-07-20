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
//   placementProperty := &PlacementProperty{
//   	Affinity: jsii.String("affinity"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	GroupId: jsii.String("groupId"),
//   	GroupName: jsii.String("groupName"),
//   	HostId: jsii.String("hostId"),
//   	HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   	PartitionNumber: jsii.Number(123),
//   	SpreadDomain: jsii.String("spreadDomain"),
//   	Tenancy: jsii.String("tenancy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html
//
type CfnLaunchTemplate_PlacementProperty struct {
	// The affinity setting for an instance on a Dedicated Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-affinity
	//
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The Group Id of a placement group.
	//
	// You must specify the Placement Group *Group Id* to launch an instance in a shared placement group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The name of the placement group for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ID of the Dedicated Host for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-hostid
	//
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the *Tenancy* parameter or set it to `host` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-hostresourcegrouparn
	//
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The number of the partition the instance should launch in.
	//
	// Valid only if the placement group strategy is set to `partition` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-partitionnumber
	//
	PartitionNumber *float64 `field:"optional" json:"partitionNumber" yaml:"partitionNumber"`
	// Reserved for future use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-spreaddomain
	//
	SpreadDomain *string `field:"optional" json:"spreadDomain" yaml:"spreadDomain"`
	// The tenancy of the instance.
	//
	// An instance with a tenancy of dedicated runs on single-tenant hardware.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-placement.html#cfn-ec2-launchtemplate-placement-tenancy
	//
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

