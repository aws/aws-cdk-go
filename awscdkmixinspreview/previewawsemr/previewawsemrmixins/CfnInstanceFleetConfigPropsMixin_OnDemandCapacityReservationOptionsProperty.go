package previewawsemrmixins


// Describes the strategy for using unused Capacity Reservations for fulfilling On-Demand capacity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   onDemandCapacityReservationOptionsProperty := &OnDemandCapacityReservationOptionsProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   	CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   	UsageStrategy: jsii.String("usageStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ondemandcapacityreservationoptions.html
//
type CfnInstanceFleetConfigPropsMixin_OnDemandCapacityReservationOptionsProperty struct {
	// Indicates the instance's Capacity Reservation preferences. Possible preferences include:.
	//
	// - `open` - The instance can run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	// - `none` - The instance avoids running in a Capacity Reservation even if one is available. The instance runs as an On-Demand Instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ondemandcapacityreservationoptions.html#cfn-emr-instancefleetconfig-ondemandcapacityreservationoptions-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// The ARN of the Capacity Reservation resource group in which to run the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ondemandcapacityreservationoptions.html#cfn-emr-instancefleetconfig-ondemandcapacityreservationoptions-capacityreservationresourcegrouparn
	//
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
	// Indicates whether to use unused Capacity Reservations for fulfilling On-Demand capacity.
	//
	// If you specify `use-capacity-reservations-first` , the fleet uses unused Capacity Reservations to fulfill On-Demand capacity up to the target On-Demand capacity. If multiple instance pools have unused Capacity Reservations, the On-Demand allocation strategy ( `lowest-price` ) is applied. If the number of unused Capacity Reservations is less than the On-Demand target capacity, the remaining On-Demand target capacity is launched according to the On-Demand allocation strategy ( `lowest-price` ).
	//
	// If you do not specify a value, the fleet fulfills the On-Demand capacity according to the chosen On-Demand allocation strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ondemandcapacityreservationoptions.html#cfn-emr-instancefleetconfig-ondemandcapacityreservationoptions-usagestrategy
	//
	UsageStrategy *string `field:"optional" json:"usageStrategy" yaml:"usageStrategy"`
}

