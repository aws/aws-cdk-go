package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How DVB Service Description Table (SDT) information is inserted.
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
type DvbSdtOutputMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for DvbSdtOutputMode
type jsiiProxy_DvbSdtOutputMode struct {
	_ byte // padding
}

func (j *jsiiProxy_DvbSdtOutputMode) Value() *string {
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
func DvbSdtOutputMode_Of(value *string) DvbSdtOutputMode {
	_init_.Initialize()

	if err := validateDvbSdtOutputMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DvbSdtOutputMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.DvbSdtOutputMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DvbSdtOutputMode_SDT_FOLLOW() DvbSdtOutputMode {
	_init_.Initialize()
	var returns DvbSdtOutputMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbSdtOutputMode",
		"SDT_FOLLOW",
		&returns,
	)
	return returns
}

func DvbSdtOutputMode_SDT_FOLLOW_IF_PRESENT() DvbSdtOutputMode {
	_init_.Initialize()
	var returns DvbSdtOutputMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbSdtOutputMode",
		"SDT_FOLLOW_IF_PRESENT",
		&returns,
	)
	return returns
}

func DvbSdtOutputMode_SDT_MANUAL() DvbSdtOutputMode {
	_init_.Initialize()
	var returns DvbSdtOutputMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbSdtOutputMode",
		"SDT_MANUAL",
		&returns,
	)
	return returns
}

func DvbSdtOutputMode_SDT_NONE() DvbSdtOutputMode {
	_init_.Initialize()
	var returns DvbSdtOutputMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbSdtOutputMode",
		"SDT_NONE",
		&returns,
	)
	return returns
}

