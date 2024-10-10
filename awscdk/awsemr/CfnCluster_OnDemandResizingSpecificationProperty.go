package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onDemandResizingSpecificationProperty := &OnDemandResizingSpecificationProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   		CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   		UsageStrategy: jsii.String("usageStrategy"),
//   	},
//   	TimeoutDurationMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html
//
type CfnCluster_OnDemandResizingSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-capacityreservationoptions
	//
	CapacityReservationOptions interface{} `field:"optional" json:"capacityReservationOptions" yaml:"capacityReservationOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-timeoutdurationminutes
	//
	TimeoutDurationMinutes *float64 `field:"optional" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
}

