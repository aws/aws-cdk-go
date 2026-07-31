package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol Options for Sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   sourceProtocol := mediaconnect_alpha.SourceProtocol_CDI()
//
// Experimental.
type SourceProtocol interface {
	// The source protocol string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SourceProtocol
type jsiiProxy_SourceProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceProtocol) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom source protocol value.
// Experimental.
func SourceProtocol_Of(value *string) SourceProtocol {
	_init_.Initialize()

	if err := validateSourceProtocol_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SourceProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SourceProtocol_CDI() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"CDI",
		&returns,
	)
	return returns
}

func SourceProtocol_JPEGXS() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"JPEGXS",
		&returns,
	)
	return returns
}

func SourceProtocol_NDI_SPEED_HQ() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"NDI_SPEED_HQ",
		&returns,
	)
	return returns
}

func SourceProtocol_RIST() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"RIST",
		&returns,
	)
	return returns
}

func SourceProtocol_RTP() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"RTP",
		&returns,
	)
	return returns
}

func SourceProtocol_RTP_FEC() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"RTP_FEC",
		&returns,
	)
	return returns
}

func SourceProtocol_SRT_CALLER() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"SRT_CALLER",
		&returns,
	)
	return returns
}

func SourceProtocol_SRT_LISTENER() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"SRT_LISTENER",
		&returns,
	)
	return returns
}

func SourceProtocol_ZIXI_PUSH() SourceProtocol {
	_init_.Initialize()
	var returns SourceProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.SourceProtocol",
		"ZIXI_PUSH",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SourceProtocol) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

