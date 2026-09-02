package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Enables column-only or column-and-row FEC for a UDP output.
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
type FecMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for FecMode
type jsiiProxy_FecMode struct {
	_ byte // padding
}

func (j *jsiiProxy_FecMode) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func FecMode_Of(value *string) FecMode {
	_init_.Initialize()

	if err := validateFecMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns FecMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FecMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FecMode_COLUMN() FecMode {
	_init_.Initialize()
	var returns FecMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FecMode",
		"COLUMN",
		&returns,
	)
	return returns
}

func FecMode_COLUMN_AND_ROW() FecMode {
	_init_.Initialize()
	var returns FecMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FecMode",
		"COLUMN_AND_ROW",
		&returns,
	)
	return returns
}

