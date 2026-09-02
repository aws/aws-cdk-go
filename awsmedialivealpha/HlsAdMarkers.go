package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Ad marker type for an HLS output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsAdMarkers := medialive_alpha.HlsAdMarkers_ADOBE()
//
// Experimental.
type HlsAdMarkers interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsAdMarkers
type jsiiProxy_HlsAdMarkers struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsAdMarkers) Value() *string {
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
func HlsAdMarkers_Of(value *string) HlsAdMarkers {
	_init_.Initialize()

	if err := validateHlsAdMarkers_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsAdMarkers

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsAdMarkers",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsAdMarkers_ADOBE() HlsAdMarkers {
	_init_.Initialize()
	var returns HlsAdMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAdMarkers",
		"ADOBE",
		&returns,
	)
	return returns
}

func HlsAdMarkers_ELEMENTAL() HlsAdMarkers {
	_init_.Initialize()
	var returns HlsAdMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAdMarkers",
		"ELEMENTAL",
		&returns,
	)
	return returns
}

func HlsAdMarkers_ELEMENTAL_SCTE35() HlsAdMarkers {
	_init_.Initialize()
	var returns HlsAdMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAdMarkers",
		"ELEMENTAL_SCTE35",
		&returns,
	)
	return returns
}

