package awsec2


// Properties to create PortRange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclPortRange := &aclPortRange{
//   	from: jsii.Number(123),
//   	to: jsii.Number(123),
//   }
//
type AclPortRange struct {
	// The first port in the range.
	//
	// Required if you specify 6 (TCP) or 17 (UDP) for the protocol parameter.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The last port in the range.
	//
	// Required if you specify 6 (TCP) or 17 (UDP) for the protocol parameter.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

