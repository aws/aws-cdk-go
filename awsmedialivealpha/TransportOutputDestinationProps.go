package awsmedialivealpha


// A destination address (IP or host) and port for a transport-stream output.
//
// Example:
//   var video EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Rtp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("ts"),
//   			Fec: &FecOutputSettings{
//   				Mode: medialive.FecMode_COLUMN_AND_ROW(),
//   				ColumnDepth: jsii.Number(10),
//   				RowLength: jsii.Number(10),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type TransportOutputDestinationProps struct {
	// The destination address — a unicast or multicast IP, or a hostname.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// The destination port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

