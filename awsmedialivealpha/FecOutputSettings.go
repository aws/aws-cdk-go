package awsmedialivealpha


// Forward Error Correction (FEC) settings for a UDP output (SMPTE 2022-1).
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
type FecOutputSettings struct {
	// Parameter D from SMPTE 2022-1 — the height of the FEC protection matrix (number of transport stream packets per column error-correction packet).
	//
	// Must be 4..20.
	// Default: - service default.
	//
	// Experimental.
	ColumnDepth *float64 `field:"optional" json:"columnDepth" yaml:"columnDepth"`
	// Whether to enable column-only or column-and-row FEC.
	// Default: - service default.
	//
	// Experimental.
	Mode FecMode `field:"optional" json:"mode" yaml:"mode"`
	// Parameter L from SMPTE 2022-1 — the width of the FEC protection matrix.
	//
	// Must be 1..20
	// for column-only FEC, or 4..20 for column-and-row FEC.
	// Default: - service default.
	//
	// Experimental.
	RowLength *float64 `field:"optional" json:"rowLength" yaml:"rowLength"`
}

