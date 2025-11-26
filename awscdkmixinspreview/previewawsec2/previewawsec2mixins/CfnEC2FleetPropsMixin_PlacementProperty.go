package previewawsec2mixins


// Describes the placement of an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnEC2FleetPropsMixin_PlacementProperty struct {
	// The affinity setting for the instance on the Dedicated Host.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) or [ImportInstance](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-affinity
	//
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone of the instance.
	//
	// On input, you can specify `AvailabilityZone` or `AvailabilityZoneId` , but not both. If you specify neither one, Amazon EC2 automatically selects an Availability Zone for you.
	//
	// This parameter is not supported for [CreateFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-placement.html#cfn-ec2-ec2fleet-placement-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the placement group that the instance is in.
	//
	// On input, you can specify `GroupId` or `GroupName` , but not both.
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
	// On input, if you specify this parameter, either omit the *Tenancy* parameter or set it to `host` .
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

