package awsmedialivealpha


// A source for a multicast input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var multicastProtocol MulticastProtocol
//
//   multicastInputSource := &MulticastInputSource{
//   	Address: jsii.String("address"),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	Protocol: multicastProtocol,
//   	SourceIp: jsii.String("sourceIp"),
//   }
//
// Experimental.
type MulticastInputSource struct {
	// The multicast group address.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// The multicast port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The transport protocol.
	// Default: MulticastProtocol.UDP
	//
	// Experimental.
	Protocol MulticastProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Filter to a specific source IP (source-specific multicast).
	// Default: - accept any source.
	//
	// Experimental.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

