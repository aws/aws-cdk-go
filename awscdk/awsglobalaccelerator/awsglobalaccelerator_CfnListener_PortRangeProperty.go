package awsglobalaccelerator


// A complex type for a range of ports for a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &portRangeProperty{
//   	fromPort: jsii.Number(123),
//   	toPort: jsii.Number(123),
//   }
//
type CfnListener_PortRangeProperty struct {
	// The first port in the range of ports, inclusive.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The last port in the range of ports, inclusive.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

