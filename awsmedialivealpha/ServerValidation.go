package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Server validation mode for HTTPS inputs.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			NetworkInputSettings: &NetworkInputSettings{
//   				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
//   				HlsInputSettings: &HlsInputSettings{
//   					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
//   					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
//   				},
//   			},
//   			LogicalInterfaceNames: []*string{
//   				jsii.String("eth0"),
//   				jsii.String("eth1"),
//   			},
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ServerValidation interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ServerValidation
type jsiiProxy_ServerValidation struct {
	_ byte // padding
}

func (j *jsiiProxy_ServerValidation) Value() *string {
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
func ServerValidation_Of(value *string) ServerValidation {
	_init_.Initialize()

	if err := validateServerValidation_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ServerValidation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ServerValidation",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME() ServerValidation {
	_init_.Initialize()
	var returns ServerValidation
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ServerValidation",
		"CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME",
		&returns,
	)
	return returns
}

func ServerValidation_CHECK_CRYPTOGRAPHY_ONLY() ServerValidation {
	_init_.Initialize()
	var returns ServerValidation
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ServerValidation",
		"CHECK_CRYPTOGRAPHY_ONLY",
		&returns,
	)
	return returns
}

