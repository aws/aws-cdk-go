package previewawsec2mixins


// Information about your commitment for a future-dated Capacity Reservation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commitmentInfoProperty := &CommitmentInfoProperty{
//   	CommitmentEndDate: jsii.String("commitmentEndDate"),
//   	CommittedInstanceCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-commitmentinfo.html
//
type CfnCapacityReservationPropsMixin_CommitmentInfoProperty struct {
	// The date and time at which the commitment duration expires, in the ISO8601 format in the UTC time zone ( `YYYY-MM-DDThh:mm:ss.sssZ` ). You can't decrease the instance count or cancel the Capacity Reservation before this date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-commitmentinfo.html#cfn-ec2-capacityreservation-commitmentinfo-commitmentenddate
	//
	CommitmentEndDate *string `field:"optional" json:"commitmentEndDate" yaml:"commitmentEndDate"`
	// The instance capacity that you committed to when you requested the future-dated Capacity Reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservation-commitmentinfo.html#cfn-ec2-capacityreservation-commitmentinfo-committedinstancecount
	//
	CommittedInstanceCount *float64 `field:"optional" json:"committedInstanceCount" yaml:"committedInstanceCount"`
}

