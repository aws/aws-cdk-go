package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeCountsProperty := &nodeCountsProperty{
//   	totalActuators: jsii.Number(123),
//   	totalAttributes: jsii.Number(123),
//   	totalBranches: jsii.Number(123),
//   	totalNodes: jsii.Number(123),
//   	totalSensors: jsii.Number(123),
//   }
//
type CfnSignalCatalog_NodeCountsProperty struct {
	// `CfnSignalCatalog.NodeCountsProperty.TotalActuators`.
	TotalActuators *float64 `field:"optional" json:"totalActuators" yaml:"totalActuators"`
	// `CfnSignalCatalog.NodeCountsProperty.TotalAttributes`.
	TotalAttributes *float64 `field:"optional" json:"totalAttributes" yaml:"totalAttributes"`
	// `CfnSignalCatalog.NodeCountsProperty.TotalBranches`.
	TotalBranches *float64 `field:"optional" json:"totalBranches" yaml:"totalBranches"`
	// `CfnSignalCatalog.NodeCountsProperty.TotalNodes`.
	TotalNodes *float64 `field:"optional" json:"totalNodes" yaml:"totalNodes"`
	// `CfnSignalCatalog.NodeCountsProperty.TotalSensors`.
	TotalSensors *float64 `field:"optional" json:"totalSensors" yaml:"totalSensors"`
}

