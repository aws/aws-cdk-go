package awsiotfleetwise


// Information about the number of nodes and node types in a vehicle network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeCountsProperty := &NodeCountsProperty{
//   	TotalActuators: jsii.Number(123),
//   	TotalAttributes: jsii.Number(123),
//   	TotalBranches: jsii.Number(123),
//   	TotalNodes: jsii.Number(123),
//   	TotalSensors: jsii.Number(123),
//   }
//
type CfnSignalCatalog_NodeCountsProperty struct {
	// (Optional) The total number of nodes in a vehicle network that represent actuators.
	TotalActuators *float64 `field:"optional" json:"totalActuators" yaml:"totalActuators"`
	// (Optional) The total number of nodes in a vehicle network that represent attributes.
	TotalAttributes *float64 `field:"optional" json:"totalAttributes" yaml:"totalAttributes"`
	// (Optional) The total number of nodes in a vehicle network that represent branches.
	TotalBranches *float64 `field:"optional" json:"totalBranches" yaml:"totalBranches"`
	// (Optional) The total number of nodes in a vehicle network.
	TotalNodes *float64 `field:"optional" json:"totalNodes" yaml:"totalNodes"`
	// (Optional) The total number of nodes in a vehicle network that represent sensors.
	TotalSensors *float64 `field:"optional" json:"totalSensors" yaml:"totalSensors"`
}

