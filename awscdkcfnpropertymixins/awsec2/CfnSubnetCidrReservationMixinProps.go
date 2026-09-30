package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSubnetCidrReservationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSubnetCidrReservationMixinProps := &CfnSubnetCidrReservationMixinProps{
//   	Cidr: jsii.String("cidr"),
//   	Description: jsii.String("description"),
//   	ReservationType: jsii.String("reservationType"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html
//
type CfnSubnetCidrReservationMixinProps struct {
	// The IPv4 or IPv6 CIDR range to reserve.
	//
	// Must lie inside the subnet's CIDR block and must not overlap an existing reservation. Supply an IPv6 range exactly as the service returns it: Amazon EC2 stores an IPv6 CIDR in RFC 5952 compressed form, and a noncanonical spelling of the same range is reported as drift because this property is create-only and read back.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html#cfn-ec2-subnetcidrreservation-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The description of the subnet CIDR reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html#cfn-ec2-subnetcidrreservation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of reservation.
	//
	// A prefix reservation is used for an IPv6 prefix delegated to a network interface; an explicit reservation is used for a range that Amazon EC2 must not assign automatically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html#cfn-ec2-subnetcidrreservation-reservationtype
	//
	ReservationType *string `field:"optional" json:"reservationType" yaml:"reservationType"`
	// The identifier of the subnet the CIDR range is reserved in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html#cfn-ec2-subnetcidrreservation-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags assigned to the subnet CIDR reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrreservation.html#cfn-ec2-subnetcidrreservation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

