package previewawsec2mixins


// Describes Spot Instance placement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotPlacementProperty := &SpotPlacementProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	GroupName: jsii.String("groupName"),
//   	Tenancy: jsii.String("tenancy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotplacement.html
//
type CfnSpotFleetPropsMixin_SpotPlacementProperty struct {
	// The Availability Zone.
	//
	// To specify multiple Availability Zones, separate them using commas; for example, "us-west-2a, us-west-2b".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotplacement.html#cfn-ec2-spotfleet-spotplacement-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotplacement.html#cfn-ec2-spotfleet-spotplacement-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The name of the placement group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotplacement.html#cfn-ec2-spotfleet-spotplacement-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of `dedicated` runs on single-tenant hardware. The `host` tenancy is not supported for Spot Instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotplacement.html#cfn-ec2-spotfleet-spotplacement-tenancy
	//
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

