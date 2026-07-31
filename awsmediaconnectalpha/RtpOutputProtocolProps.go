package awsmediaconnectalpha


// Properties for RTP protocol configuration for outputs.
//
// Example:
//   var networkInterface RouterNetworkInterface
//
//
//   awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   	NetworkInterface: NetworkInterface,
//   	Protocol: awsmediaconnectalpha.RouterOutputProtocol_Rtp(&RtpOutputProtocolProps{
//   		DestinationAddress: jsii.String("10.0.0.1"),
//   		Port: jsii.Number(5000),
//   	}),
//   })
//
// Experimental.
type RtpOutputProtocolProps struct {
	// Destination IP address for RTP traffic.
	// Experimental.
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// Port number for RTP traffic.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Forward Error Correction setting.
	// Default: - undefined; when omitted, MediaConnect applies ForwardErrorCorrection.DISABLED at deploy time
	//
	// Experimental.
	ForwardErrorCorrection ForwardErrorCorrection `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
}

