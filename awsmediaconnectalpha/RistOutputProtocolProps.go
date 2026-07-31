package awsmediaconnectalpha


// Properties for RIST protocol configuration for outputs.
//
// Example:
//   awsmediaconnectalpha.RouterOutputProtocol_Rist(&RistOutputProtocolProps{
//   	DestinationAddress: jsii.String("10.0.0.1"),
//   	Port: jsii.Number(5000),
//   })
//
// Experimental.
type RistOutputProtocolProps struct {
	// Destination IP address for RIST traffic.
	// Experimental.
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// Port number for RIST traffic.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

