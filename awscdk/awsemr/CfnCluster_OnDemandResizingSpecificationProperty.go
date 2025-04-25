package awsemr


// The resize specification for On-Demand Instances in the instance fleet, which contains the resize timeout period.
//
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
	// Specifies the allocation strategy to use to launch On-Demand instances during a resize.
	//
	// The default is `lowest-price` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-capacityreservationoptions
	//
	CapacityReservationOptions interface{} `field:"optional" json:"capacityReservationOptions" yaml:"capacityReservationOptions"`
	// On-Demand resize timeout in minutes.
	//
	// If On-Demand Instances are not provisioned within this time, the resize workflow stops. The minimum value is 5 minutes, and the maximum value is 10,080 minutes (7 days). The timeout applies to all resize workflows on the Instance Fleet. The resize could be triggered by Amazon EMR Managed Scaling or by the customer (via Amazon EMR Console, Amazon EMR CLI modify-instance-fleet or Amazon EMR SDK ModifyInstanceFleet API) or by Amazon EMR due to Amazon EC2 Spot Reclamation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ondemandresizingspecification.html#cfn-emr-cluster-ondemandresizingspecification-timeoutdurationminutes
	//
	TimeoutDurationMinutes *float64 `field:"optional" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
}

