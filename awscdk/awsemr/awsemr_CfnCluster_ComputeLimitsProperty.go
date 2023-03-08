package awsemr


// The EC2 unit limits for a managed scaling policy.
//
// The managed scaling activity of a cluster can not be above or below these limits. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeLimitsProperty := &computeLimitsProperty{
//   	maximumCapacityUnits: jsii.Number(123),
//   	minimumCapacityUnits: jsii.Number(123),
//   	unitType: jsii.String("unitType"),
//
//   	// the properties below are optional
//   	maximumCoreCapacityUnits: jsii.Number(123),
//   	maximumOnDemandCapacityUnits: jsii.Number(123),
//   }
//
type CfnCluster_ComputeLimitsProperty struct {
	// The upper boundary of EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. Managed scaling activities are not allowed beyond this boundary. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	MaximumCapacityUnits *float64 `field:"required" json:"maximumCapacityUnits" yaml:"maximumCapacityUnits"`
	// The lower boundary of EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. Managed scaling activities are not allowed beyond this boundary. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	MinimumCapacityUnits *float64 `field:"required" json:"minimumCapacityUnits" yaml:"minimumCapacityUnits"`
	// The unit type used for specifying a managed scaling policy.
	UnitType *string `field:"required" json:"unitType" yaml:"unitType"`
	// The upper boundary of EC2 units for core node type in a cluster.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. The core units are not allowed to scale beyond this boundary. The parameter is used to split capacity allocation between core and task nodes.
	MaximumCoreCapacityUnits *float64 `field:"optional" json:"maximumCoreCapacityUnits" yaml:"maximumCoreCapacityUnits"`
	// The upper boundary of On-Demand EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. The On-Demand units are not allowed to scale beyond this boundary. The parameter is used to split capacity allocation between On-Demand and Spot Instances.
	MaximumOnDemandCapacityUnits *float64 `field:"optional" json:"maximumOnDemandCapacityUnits" yaml:"maximumOnDemandCapacityUnits"`
}

