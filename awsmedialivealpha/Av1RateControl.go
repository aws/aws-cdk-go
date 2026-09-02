package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 rate control.
//
// AV1 supports QVBR and CBR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitrate Bitrate
//
//   av1RateControl := medialive_alpha.Av1RateControl_Cbr(&CbrRateControlProps{
//   	Bitrate: bitrate,
//   })
//
// Experimental.
type Av1RateControl interface {
}

// The jsii proxy struct for Av1RateControl
type jsiiProxy_Av1RateControl struct {
	_ byte // padding
}

// Constant bitrate.
// Experimental.
func Av1RateControl_Cbr(props *CbrRateControlProps) Av1RateControl {
	_init_.Initialize()

	if err := validateAv1RateControl_CbrParameters(props); err != nil {
		panic(err)
	}
	var returns Av1RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1RateControl",
		"cbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Quality-defined variable bitrate.
// Experimental.
func Av1RateControl_Qvbr(props *QvbrRateControlProps) Av1RateControl {
	_init_.Initialize()

	if err := validateAv1RateControl_QvbrParameters(props); err != nil {
		panic(err)
	}
	var returns Av1RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1RateControl",
		"qvbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

