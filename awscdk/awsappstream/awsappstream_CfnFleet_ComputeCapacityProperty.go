package awsappstream


// The desired capacity for a fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeCapacityProperty := &computeCapacityProperty{
//   	desiredInstances: jsii.Number(123),
//   }
//
type CfnFleet_ComputeCapacityProperty struct {
	// The desired number of streaming instances.
	DesiredInstances *float64 `field:"required" json:"desiredInstances" yaml:"desiredInstances"`
}

