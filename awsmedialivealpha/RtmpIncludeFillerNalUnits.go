package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP include filler NAL units.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpIncludeFillerNalUnits := medialive_alpha.RtmpIncludeFillerNalUnits_AUTO()
//
// Experimental.
type RtmpIncludeFillerNalUnits interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpIncludeFillerNalUnits
type jsiiProxy_RtmpIncludeFillerNalUnits struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpIncludeFillerNalUnits) Value() *string {
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
func RtmpIncludeFillerNalUnits_Of(value *string) RtmpIncludeFillerNalUnits {
	_init_.Initialize()

	if err := validateRtmpIncludeFillerNalUnits_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpIncludeFillerNalUnits

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpIncludeFillerNalUnits",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpIncludeFillerNalUnits_AUTO() RtmpIncludeFillerNalUnits {
	_init_.Initialize()
	var returns RtmpIncludeFillerNalUnits
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpIncludeFillerNalUnits",
		"AUTO",
		&returns,
	)
	return returns
}

func RtmpIncludeFillerNalUnits_DROP() RtmpIncludeFillerNalUnits {
	_init_.Initialize()
	var returns RtmpIncludeFillerNalUnits
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpIncludeFillerNalUnits",
		"DROP",
		&returns,
	)
	return returns
}

func RtmpIncludeFillerNalUnits_INCLUDE() RtmpIncludeFillerNalUnits {
	_init_.Initialize()
	var returns RtmpIncludeFillerNalUnits
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpIncludeFillerNalUnits",
		"INCLUDE",
		&returns,
	)
	return returns
}

