package awsec2


// Properties to create a port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portProps := &portProps{
//   	protocol: awscdk.Aws_ec2.protocol_ALL,
//   	stringRepresentation: jsii.String("stringRepresentation"),
//
//   	// the properties below are optional
//   	fromPort: jsii.Number(123),
//   	toPort: jsii.Number(123),
//   }
//
type PortProps struct {
	// The protocol for the range.
	Protocol Protocol `field:"required" json:"protocol" yaml:"protocol"`
	// String representation for this object.
	StringRepresentation *string `field:"required" json:"stringRepresentation" yaml:"stringRepresentation"`
	// The starting port for the range.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The ending port for the range.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

