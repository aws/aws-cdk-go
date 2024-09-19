package awsemr


// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// `InstanceTypeConfig` is a sub-property of `InstanceFleetConfig` . `InstanceTypeConfig` determines the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &InstanceFleetProvisioningSpecificationsProperty{
//   	OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	SpotSpecification: &SpotProvisioningSpecificationProperty{
//   		TimeoutAction: jsii.String("timeoutAction"),
//   		TimeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		BlockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancefleetprovisioningspecifications.html
//
type CfnInstanceFleetConfig_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy and capacity reservation options.
	//
	// > The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR releases 5.12.1 and later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-emr-instancefleetconfig-instancefleetprovisioningspecifications-ondemandspecification
	//
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// The launch specification for Spot instances in the fleet, which determines the allocation strategy, defined duration, and provisioning timeout behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-emr-instancefleetconfig-instancefleetprovisioningspecifications-spotspecification
	//
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

