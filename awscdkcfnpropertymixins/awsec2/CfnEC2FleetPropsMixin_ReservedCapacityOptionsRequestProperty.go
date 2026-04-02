package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reservedCapacityOptionsRequestProperty := &ReservedCapacityOptionsRequestProperty{
//   	ReservationTypes: []*string{
//   		jsii.String("reservationTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html
//
type CfnEC2FleetPropsMixin_ReservedCapacityOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.html#cfn-ec2-ec2fleet-reservedcapacityoptionsrequest-reservationtypes
	//
	ReservationTypes *[]*string `field:"optional" json:"reservationTypes" yaml:"reservationTypes"`
}

