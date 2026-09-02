package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capacityReservationTargetRequestProperty := &CapacityReservationTargetRequestProperty{
//   	CapacityReservationIds: []*string{
//   		jsii.String("capacityReservationIds"),
//   	},
//   	CapacityReservationResourceGroupArns: []*string{
//   		jsii.String("capacityReservationResourceGroupArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-capacityreservationtargetrequest.html
//
type CfnEC2FleetPropsMixin_CapacityReservationTargetRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-capacityreservationtargetrequest.html#cfn-ec2-ec2fleet-capacityreservationtargetrequest-capacityreservationids
	//
	CapacityReservationIds *[]*string `field:"optional" json:"capacityReservationIds" yaml:"capacityReservationIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-capacityreservationtargetrequest.html#cfn-ec2-ec2fleet-capacityreservationtargetrequest-capacityreservationresourcegrouparns
	//
	CapacityReservationResourceGroupArns *[]*string `field:"optional" json:"capacityReservationResourceGroupArns" yaml:"capacityReservationResourceGroupArns"`
}

