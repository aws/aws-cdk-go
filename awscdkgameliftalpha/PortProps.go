package awscdkgameliftalpha


// Properties to create a port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   portProps := &PortProps{
//   	FromPort: jsii.Number(123),
//   	Protocol: gamelift_alpha.Protocol_TCP,
//
//   	// the properties below are optional
//   	ToPort: jsii.Number(123),
//   }
//
// Experimental.
type PortProps struct {
	// A starting value for a range of allowed port numbers.
	//
	// For fleets using Windows and Linux builds, only ports 1026-60000 are valid.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The protocol for the range.
	// Experimental.
	Protocol Protocol `field:"required" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive.
	// This value must be higher than `fromPort`.
	//
	// For fleets using Windows and Linux builds, only ports 1026-60000 are valid.
	// Default: the `fromPort` value.
	//
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

