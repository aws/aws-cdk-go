package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reservedCapacityOptionsRequestProperty := &ReservedCapacityOptionsRequestProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	CapacityReservationTarget: &CapacityReservationTargetRequestProperty{
//   		CapacityReservationIds: []*string{
//   			jsii.String("capacityReservationIds"),
//   		},
//   		CapacityReservationResourceGroupArns: []*string{
//   			jsii.String("capacityReservationResourceGroupArns"),
//   		},
//   	},
//   	ReservationTypes: []*string{
//   		jsii.String("reservationTypes"),
//   	},
//   	ReservedCapacityFallbackOptions: &ReservedCapacityFallbackOptionsRequestProperty{
//   		MarketTypes: []*string{
//   			jsii.String("marketTypes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html
//
type CfnEC2FleetPropsMixin_ReservedCapacityOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityoptionsrequest-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityoptionsrequest-capacityreservationtarget
	//
	CapacityReservationTarget interface{} `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityoptionsrequest-reservationtypes
	//
	ReservationTypes *[]*string `field:"optional" json:"reservationTypes" yaml:"reservationTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityoptionsrequest-reservedcapacityfallbackoptions
	//
	ReservedCapacityFallbackOptions interface{} `field:"optional" json:"reservedCapacityFallbackOptions" yaml:"reservedCapacityFallbackOptions"`
}

