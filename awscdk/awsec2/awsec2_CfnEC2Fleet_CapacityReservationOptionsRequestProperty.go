package awsec2


// Describes the strategy for using unused Capacity Reservations for fulfilling On-Demand capacity.
//
// > This strategy can only be used if the EC2 Fleet is of type `instant` .
//
// For more information about Capacity Reservations, see [On-Demand Capacity Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html) in the *Amazon EC2 User Guide* . For examples of using Capacity Reservations in an EC2 Fleet, see [EC2 Fleet example configurations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-examples.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationOptionsRequestProperty := &capacityReservationOptionsRequestProperty{
//   	usageStrategy: jsii.String("usageStrategy"),
//   }
//
type CfnEC2Fleet_CapacityReservationOptionsRequestProperty struct {
	// Indicates whether to use unused Capacity Reservations for fulfilling On-Demand capacity.
	//
	// If you specify `use-capacity-reservations-first` , the fleet uses unused Capacity Reservations to fulfill On-Demand capacity up to the target On-Demand capacity. If multiple instance pools have unused Capacity Reservations, the On-Demand allocation strategy ( `lowest-price` or `prioritized` ) is applied. If the number of unused Capacity Reservations is less than the On-Demand target capacity, the remaining On-Demand target capacity is launched according to the On-Demand allocation strategy ( `lowest-price` or `prioritized` ).
	//
	// If you do not specify a value, the fleet fulfils the On-Demand capacity according to the chosen On-Demand allocation strategy.
	UsageStrategy *string `field:"optional" json:"usageStrategy" yaml:"usageStrategy"`
}

