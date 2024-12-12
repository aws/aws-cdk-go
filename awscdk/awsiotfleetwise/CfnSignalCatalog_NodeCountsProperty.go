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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html
//
type CfnSignalCatalog_NodeCountsProperty struct {
	// The total number of nodes in a vehicle network that represent actuators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html#cfn-iotfleetwise-signalcatalog-nodecounts-totalactuators
	//
	TotalActuators *float64 `field:"optional" json:"totalActuators" yaml:"totalActuators"`
	// The total number of nodes in a vehicle network that represent attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html#cfn-iotfleetwise-signalcatalog-nodecounts-totalattributes
	//
	TotalAttributes *float64 `field:"optional" json:"totalAttributes" yaml:"totalAttributes"`
	// The total number of nodes in a vehicle network that represent branches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html#cfn-iotfleetwise-signalcatalog-nodecounts-totalbranches
	//
	TotalBranches *float64 `field:"optional" json:"totalBranches" yaml:"totalBranches"`
	// The total number of nodes in a vehicle network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html#cfn-iotfleetwise-signalcatalog-nodecounts-totalnodes
	//
	TotalNodes *float64 `field:"optional" json:"totalNodes" yaml:"totalNodes"`
	// The total number of nodes in a vehicle network that represent sensors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-nodecounts.html#cfn-iotfleetwise-signalcatalog-nodecounts-totalsensors
	//
	TotalSensors *float64 `field:"optional" json:"totalSensors" yaml:"totalSensors"`
}

