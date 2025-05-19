package awsec2


// Describes the placement of an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementProperty := &PlacementProperty{
//   	Affinity: jsii.String("affinity"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	GroupName: jsii.String("groupName"),
//   	HostId: jsii.String("hostId"),
//   	HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   	PartitionNumber: jsii.Number(123),
//   	SpreadDomain: jsii.String("spreadDomain"),
//   	Tenancy: jsii.String("tenancy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html
//
type CfnEC2Fleet_PlacementProperty struct {
	// The affinity setting for the instance on the Dedicated Host.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) or [ImportInstance](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-affinity
	//
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone of the instance.
	//
	// If not specified, an Availability Zone will be automatically chosen for you based on the load balancing criteria for the Region.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the placement group that the instance is in.
	//
	// If you specify `GroupName` , you can't specify `GroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ID of the Dedicated Host on which the instance resides.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) or [ImportInstance](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-hostid
	//
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify this parameter, either omit the *Tenancy* parameter or set it to `host` .
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-hostresourcegrouparn
	//
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The number of the partition that the instance is in.
	//
	// Valid only if the placement group strategy is set to `partition` .
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-partitionnumber
	//
	PartitionNumber *float64 `field:"optional" json:"partitionNumber" yaml:"partitionNumber"`
	// Reserved for future use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-spreaddomain
	//
	SpreadDomain *string `field:"optional" json:"spreadDomain" yaml:"spreadDomain"`
	// The tenancy of the instance. An instance with a tenancy of `dedicated` runs on single-tenant hardware.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) . The `host` tenancy is not supported for [ImportInstance](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html) or for T3 instances that are configured for the `unlimited` CPU credit option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-tenancy
	//
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

