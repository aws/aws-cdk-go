package awsglobalaccelerator


// The list of port ranges for the connections from clients to the accelerator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRange := &portRange{
//   	fromPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	toPort: jsii.Number(123),
//   }
//
type PortRange struct {
	// The first port in the range of ports, inclusive.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The last port in the range of ports, inclusive.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

