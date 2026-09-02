package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SCTE-35 passthrough behavior.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp_out"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("ts"),
//   			M2tsSettings: medialive.M2tsSettings_Of(&M2tsSettingsProps{
//   				Bitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   				RateMode: medialive.M2tsRateMode_VBR(),
//   				ProgramNum: jsii.Number(1),
//   				PatInterval: awscdk.Duration_Millis(jsii.Number(100)),
//   				PmtInterval: awscdk.Duration_*Millis(jsii.Number(100)),
//   				Scte35Control: medialive.M2tsScte35Control_PASSTHROUGH(),
//   				DvbSdtSettings: &DvbSdtSettings{
//   					OutputSdt: medialive.DvbSdtOutputMode_SDT_MANUAL(),
//   					ServiceName: jsii.String("My Service"),
//   					RepInterval: awscdk.Duration_*Millis(jsii.Number(2000)),
//   				},
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type M2tsScte35Control interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsScte35Control
type jsiiProxy_M2tsScte35Control struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsScte35Control) Value() *string {
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
func M2tsScte35Control_Of(value *string) M2tsScte35Control {
	_init_.Initialize()

	if err := validateM2tsScte35Control_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsScte35Control

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsScte35Control",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsScte35Control_NONE() M2tsScte35Control {
	_init_.Initialize()
	var returns M2tsScte35Control
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsScte35Control",
		"NONE",
		&returns,
	)
	return returns
}

func M2tsScte35Control_PASSTHROUGH() M2tsScte35Control {
	_init_.Initialize()
	var returns M2tsScte35Control
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsScte35Control",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

