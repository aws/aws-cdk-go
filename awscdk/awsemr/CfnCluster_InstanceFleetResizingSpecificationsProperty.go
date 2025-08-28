package awsemr


// The resize specification for On-Demand and Spot Instances in the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetResizingSpecificationsProperty := &InstanceFleetResizingSpecificationsProperty{
//   	OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   			CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   			CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   			UsageStrategy: jsii.String("usageStrategy"),
//   		},
//   		TimeoutDurationMinutes: jsii.Number(123),
//   	},
//   	SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		TimeoutDurationMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetresizingspecifications.html
//
type CfnCluster_InstanceFleetResizingSpecificationsProperty struct {
	// The resize specification for On-Demand Instances in the instance fleet, which contains the allocation strategy, capacity reservation options, and the resize timeout period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetresizingspecifications.html#cfn-emr-cluster-instancefleetresizingspecifications-ondemandresizespecification
	//
	OnDemandResizeSpecification interface{} `field:"optional" json:"onDemandResizeSpecification" yaml:"onDemandResizeSpecification"`
	// The resize specification for Spot Instances in the instance fleet, which contains the allocation strategy and the resize timeout period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetresizingspecifications.html#cfn-emr-cluster-instancefleetresizingspecifications-spotresizespecification
	//
	SpotResizeSpecification interface{} `field:"optional" json:"spotResizeSpecification" yaml:"spotResizeSpecification"`
}

