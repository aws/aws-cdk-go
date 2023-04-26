package awsglobalaccelerator


// The list of port ranges for the connections from clients to the accelerator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRange := &PortRange{
//   	FromPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	ToPort: jsii.Number(123),
//   }
//
// Experimental.
type PortRange struct {
	// The first port in the range of ports, inclusive.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The last port in the range of ports, inclusive.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

